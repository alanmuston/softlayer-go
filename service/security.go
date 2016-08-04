package service

import "github.ibm.com/riethm/gopherlayer/datatypes"

type Security_Certificate struct {
	Session *Session
	Options
}

func (r *Session) GetSecurityCertificateService() Security_Certificate {
	return Security_Certificate{Session: r}
}

func (r *Security_Certificate) CreateObject(templateObject *datatypes.Security_Certificate) (resp datatypes.Security_Certificate, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate) EditObject(templateObject *datatypes.Security_Certificate) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate) FindByCommonName(commonName *string) (resp []datatypes.Security_Certificate, err error) {
	params := []interface{}{
		commonName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate) GetObject() (resp datatypes.Security_Certificate, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate) GetPemFormat() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Security_Certificate) GetAssociatedServiceCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate) GetLoadBalancerVirtualIpAddresses() (resp []datatypes.Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Security_Certificate_Entry struct {
	Session *Session
	Options
}

func (r *Session) GetSecurityCertificateEntryService() Security_Certificate_Entry {
	return Security_Certificate_Entry{Session: r}
}

type Security_Certificate_Request struct {
	Session *Session
	Options
}

func (r *Session) GetSecurityCertificateRequestService() Security_Certificate_Request {
	return Security_Certificate_Request{Session: r}
}

func (r *Security_Certificate_Request) CancelSslOrder() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) GetAdministratorEmailDomains(commonName *string) (resp []string, err error) {
	params := []interface{}{
		commonName,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) GetAdministratorEmailPrefixes() (resp []string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) GetObject() (resp datatypes.Security_Certificate_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) GetPreviousOrderData() (resp datatypes.Container_Product_Order_Security_Certificate, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) GetSslCertificateRequests(accountId *int) (resp []datatypes.Security_Certificate_Request, err error) {
	params := []interface{}{
		accountId,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) ResendEmail(emailType *string) (resp bool, err error) {
	params := []interface{}{
		emailType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) ValidateCsr(csr *string, validityMonths *int, itemId *int, serverType *string) (resp bool, err error) {
	params := []interface{}{
		csr,
		validityMonths,
		itemId,
		serverType,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

func (r *Security_Certificate_Request) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) GetCertificateAuthorityName() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) GetOrder() (resp datatypes.Billing_Order, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) GetOrderItem() (resp datatypes.Billing_Order_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request) GetStatus() (resp datatypes.Security_Certificate_Request_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Security_Certificate_Request_ServerType struct {
	Session *Session
	Options
}

func (r *Session) GetSecurityCertificateRequestServerTypeService() Security_Certificate_Request_ServerType {
	return Security_Certificate_Request_ServerType{Session: r}
}

func (r *Security_Certificate_Request_ServerType) GetAllObjects() (resp []datatypes.Security_Certificate_Request, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request_ServerType) GetObject() (resp datatypes.Security_Certificate_Request_ServerType, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Security_Certificate_Request_Status struct {
	Session *Session
	Options
}

func (r *Session) GetSecurityCertificateRequestStatusService() Security_Certificate_Request_Status {
	return Security_Certificate_Request_Status{Session: r}
}

func (r *Security_Certificate_Request_Status) GetObject() (resp datatypes.Security_Certificate_Request_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Certificate_Request_Status) GetSslRequestStatuses() (resp []datatypes.Security_Certificate_Request_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Security_Directory_Service_Host_Xref_Hardware struct {
	Session *Session
	Options
}

func (r *Session) GetSecurityDirectoryServiceHostXrefHardwareService() Security_Directory_Service_Host_Xref_Hardware {
	return Security_Directory_Service_Host_Xref_Hardware{Session: r}
}

func (r *Security_Directory_Service_Host_Xref_Hardware) GetHost() (resp datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

type Security_SecureTransportCipher struct {
	Session *Session
	Options
}

func (r *Session) GetSecuritySecureTransportCipherService() Security_SecureTransportCipher {
	return Security_SecureTransportCipher{Session: r}
}

type Security_SecureTransportProtocol struct {
	Session *Session
	Options
}

func (r *Session) GetSecuritySecureTransportProtocolService() Security_SecureTransportProtocol {
	return Security_SecureTransportProtocol{Session: r}
}

type Security_Ssh_Key struct {
	Session *Session
	Options
}

func (r *Session) GetSecuritySshKeyService() Security_Ssh_Key {
	return Security_Ssh_Key{Session: r}
}

func (r *Security_Ssh_Key) CreateObject(templateObject *datatypes.Security_Ssh_Key) (resp datatypes.Security_Ssh_Key, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Ssh_Key) DeleteObject() (resp bool, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Ssh_Key) EditObject(templateObject *datatypes.Security_Ssh_Key) (resp bool, err error) {
	params := []interface{}{
		templateObject,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Ssh_Key) GetObject() (resp datatypes.Security_Ssh_Key, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

func (r *Security_Ssh_Key) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Ssh_Key) GetBlockDeviceTemplateGroups() (resp []datatypes.Virtual_Guest_Block_Device_Template_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
func (r *Security_Ssh_Key) GetSoftwarePasswords() (resp []datatypes.Software_Component_Password, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
