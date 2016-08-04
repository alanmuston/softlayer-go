package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"text/template"

	"golang.org/x/tools/imports"
)

type Type struct {
	Name       string              `json:"name"`
	Base       string              `json:"base"`
	TypeDoc    string              `json:"typeDoc"`
	Properties map[string]Property `json:"properties"`
	ServiceDoc string              `json:"serviceDoc"`
	Methods    map[string]Method   `json:"methods"`
	NoService  bool                `json:"noservice"`
}

type Property struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	TypeArray bool   `json:"typeArray"`
	Form      string `json:"form"`
	Doc       string `json:"doc"`
}

type Method struct {
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	TypeArray  bool        `json:"typeArray"`
	Doc        string      `json:"doc"`
	Static     bool        `json:"static"`
	NoAuth     bool        `json:"noauth"`
	Limitable  bool        `json:"limitable"`
	Filterable bool        `json:"filterable"`
	Maskable   bool        `json:"maskable"`
	Parameters []Parameter `json:"parameters"`
}

type Parameter struct {
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	TypeArray    bool        `json:"typeArray"`
	Doc          string      `json:"doc"`
	DefaultValue interface{} `json:"defaultValue"`
}

// Define custom template functions
var fMap = template.FuncMap{
	"convertType":       ConvertType,           // Converts SoftLayer types to Go types
	"prefixWithPackage": PrefixWithPackageName, // Prepend a type with the given package name
	"removePrefix":      RemovePrefix,          // Remove 'SoftLayer_' prefix. if it exists
	"removeReserved":    RemoveReservedWords,   // Substitute language-reserved identifiers
	"titleCase":         strings.Title,         // TitleCase the argument
	"desnake":           Desnake,               // Remove '_' from Snake_Case
}

const datatype = `package datatypes

{{range .}}type {{.Name|removePrefix}} struct {
    {{.Base|removePrefix}}

    {{range .Properties}}{{.Name|titleCase}} {{if .TypeArray}}[]{{else}}*{{end}}{{.Type|convertType|removePrefix}}` +
	"`json:\"{{.Name}},omitempty\"`" + `
    {{end}}
}

{{end}}
`
const service = `package service

{{range .}}{{$base := .Name|removePrefix}}
    type {{$base}} struct {
        Session *Session
        Options
    }

    func (r *Session) Get{{$base | desnake}}Service() {{$base}} {
        return {{$base}}{Session: r}
    }

    {{range .Methods}}func (r *{{$base}}) {{.Name|titleCase}}({{range .Parameters}}{{.Name|removeReserved}} {{if not .TypeArray}}*{{else}}[]{{end}}{{.Type|convertType|prefixWithPackage "datatypes"}}, {{end}}) ({{if .Type|ne "void"}}resp {{if .TypeArray}}[]{{end}}{{.Type|convertType|prefixWithPackage "datatypes"}}, {{end}}err error) {
        {{if .Type|eq "void"}}var resp datatypes.Void
        {{end}}{{if len .Parameters | lt 0}}params := []interface{}{
            {{range .Parameters}}{{.Name|removeReserved}},
            {{end}}
        }
        {{end}}err = invokeMethod({{if len .Parameters | lt 0}}params{{else}}nil{{end}}, r.Session, &r.Options, &resp)
	return
    }
    {{end}}

    {{range .Properties}}{{if .Form|eq "relational"}}func (r *{{$base}}) Get{{.Name|titleCase}}() (resp {{if .TypeArray}}[]{{end}}{{.Type|convertType|prefixWithPackage "datatypes"}}, err error) {
        err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
    }
    {{end}}{{end}}

{{end}}
`

func main() {
	var meta map[string]Type

	outputPath := flag.String("o", ".", "the root of the go project to be refreshed")
	flag.Parse()

	jsonResp, code, err := makeHttpRequest("https://api.softlayer.com/metadata/v3.1", "GET", new(bytes.Buffer))

	if err != nil {
		fmt.Printf("Error retrieving metadata API: %s", err)
		os.Exit(1)
	}

	if code != 200 {
		fmt.Printf("Unexpected HTTP status code received while retrieving metadata API: %d", code)
		os.Exit(1)
	}

	err = json.Unmarshal(jsonResp, &meta)
	if err != nil {
		fmt.Printf("Error unmarshaling json response: %s", err)
		os.Exit(1)
	}

	// Build an array of Types, sorted by name
	// This will ensure consistency in the order that code is later emitted
	keys := getSortedKeys(meta)

	sortedMeta := make([]Type, 0, len(keys))

	for _, name := range keys {
		sortedMeta = append(sortedMeta, meta[name])
	}

	err = writePackage(*outputPath, "datatypes", sortedMeta, datatype)
	if err != nil {
		fmt.Printf("Error writing to file: %s", err)
	}

	err = writePackage(*outputPath, "service", sortedMeta, service)
	if err != nil {
		fmt.Printf("Error writing to file: %s", err)
	}
}

// Exported template functions

func RemovePrefix(args ...interface{}) string {
	s := args[0].(string)

	if strings.HasPrefix(s, "SoftLayer_") {
		return s[10:]
	}

	return s
}

func PrefixWithPackageName(args ...interface{}) string {
	p := args[0].(string)
	s := args[1].(string)

	if !strings.HasPrefix(s, "SoftLayer_") {
		return s
	}

	return p + "." + RemovePrefix(s)
}

func ConvertType(args ...interface{}) string {
	t := args[0].(string)

	// Convert softlayer types to golang types
	switch t {
	case "unsignedLong", "unsignedInt":
		return "uint"
	case "boolean":
		return "bool"
	case "dateTime":
		return "time.Time"
	case "decimal", "float":
		return "float64"
	case "base64Binary":
		return "[]byte"
	case "json", "enum":
		return "string"
	}

	return t
}

func RemoveReservedWords(args ...interface{}) string {
	n := args[0].(string)

	// Replace language reserved identifiers with alternatives
	switch n {
	case "type":
		return "typ"
	}

	return n
}

// Remove '_' from Snake_Case values
func Desnake(args ...interface{}) string {
	s := args[0].(string)
	return strings.Replace(s, "_", "", -1)
}

// private

func getSortedKeys(m map[string]Type) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return keys
}

func writePackage(base string, pkg string, meta []Type, ts string) error {
	var currPrefix string
	var start int

	for i, t := range meta {
		components := strings.Split(RemovePrefix(t.Name), "_")

		if i == 0 {
			currPrefix = components[0]
			continue
		}

		if components[0] != currPrefix {
			err := writeGoFile(base, pkg, currPrefix, meta[start:i], ts)
			if err != nil {
				return err
			}

			currPrefix = components[0]
			start = i
		}
	}

	writeGoFile(base, pkg, currPrefix, meta[start:], ts)

	return nil
}

// Executes a template against the metadata structure, and generates a go source file with the result
func writeGoFile(base string, pkg string, name string, meta []Type, ts string) error {
	filename := base + "/" + pkg + "/" + strings.ToLower(name) + ".go"

	// Generate the source
	var buf bytes.Buffer
	t := template.New(pkg).Funcs(fMap)
	template.Must(t.Parse(ts)).Execute(&buf, meta)

	/*if pkg == "service" && name == "Account"{
		fmt.Println(string(buf.String()))
		os.Exit(0)
	}*/

	// Add the imports
	src, err := imports.Process(filename, buf.Bytes(), &imports.Options{})
	if err != nil {
		fmt.Printf("Error processing imports: %s", err)
	}

	// Format
	pretty, err := format.Source(src)
	if err != nil {
		return fmt.Errorf("Error while formatting source: %s", err)
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Error creating file: %s", err)
	}
	defer f.Close()
	fmt.Fprintf(f, "%s", pretty)

	return nil
}

func makeHttpRequest(url string, requestType string, requestBody *bytes.Buffer) ([]byte, int, error) {
	client := http.DefaultClient

	req, err := http.NewRequest(requestType, url, requestBody)
	if err != nil {
		return nil, 0, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, 520, err
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return responseBody, resp.StatusCode, nil
}
