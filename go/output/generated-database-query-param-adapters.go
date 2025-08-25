package db

import "github.com/google/uuid"
import "time"
import "fmt"

func (saecr *SelectAllEmbroideryCustomersRow) GetID() uuid.UUID {
	if saecr.ID.Valid {
		return saecr.ID.Bytes
	}
	return uuid.UUID{}
}
func (saecr *SelectAllEmbroideryCustomersRow) GetIDText() string {
	if saecr.ID.Valid {
		return fmt.Sprint(saecr.ID.Bytes)
	}
	return ""
}
func (saecr *SelectAllEmbroideryCustomersRow) GetName() string {
	if saecr.Name.Valid {
		return saecr.Name.String
	}
	return ""
}
func (saecr *SelectAllEmbroideryCustomersRow) GetDescription() string {
	if saecr.Description.Valid {
		return saecr.Description.String
	}
	return ""
}
func (saecr *SelectAllEmbroideryCustomersRow) GetOrderin() time.Time {
	if saecr.Orderin.Valid {
		return saecr.Orderin.Time
	}
	return time.Time{}
}
func (saecr *SelectAllEmbroideryCustomersRow) GetOrderinText() string { return "" }
func (saecr *SelectAllEmbroideryCustomersRow) GetOrderdue() time.Time {
	if saecr.Orderdue.Valid {
		return saecr.Orderdue.Time
	}
	return time.Time{}
}
func (saecr *SelectAllEmbroideryCustomersRow) GetOrderdueText() string { return "" }
func (saecr *SelectAllEmbroideryCustomersRow) GetRushed() bool {
	if saecr.Rushed.Valid {
		return saecr.Rushed.Bool
	}
	return false
}
func (saecr *SelectAllEmbroideryCustomersRow) GetRushedText() string {
	if saecr.Rushed.Valid {
		return fmt.Sprint(saecr.Rushed.Bool)
	}
	return ""
}
func (saecr *SelectAllEmbroideryCustomersRow) GetForm() string {
	if saecr.Form.Valid {
		return saecr.Form.String
	}
	return ""
}
func (saecr *SelectAllEmbroideryCustomersRow) GetVisible() int {
	if saecr.Visible.Valid {
		return int(saecr.Visible.Int32)
	}
	return -1
}
func (saecr *SelectAllEmbroideryCustomersRow) GetVisibleText() string {
	if saecr.Visible.Valid {
		return fmt.Sprint(saecr.Visible.Int32)
	}
	return ""
}
func (saecr *SelectAllEmbroideryCustomersRow) GetEmail() string {
	if saecr.Email.Valid {
		return saecr.Email.String
	}
	return ""
}
func (saecr *SelectAllEmbroideryCustomersRow) GetPhone() string {
	if saecr.Phone.Valid {
		return saecr.Phone.String
	}
	return ""
}
func (saecr *SelectAllEmbroideryCustomersRow) GetCompleted() bool {
	if saecr.Completed.Valid {
		return saecr.Completed.Bool
	}
	return false
}
func (saecr *SelectAllEmbroideryCustomersRow) GetCompletedText() string {
	if saecr.Completed.Valid {
		return fmt.Sprint(saecr.Completed.Bool)
	}
	return ""
}
func (saecr *SelectAllEmbroideryCustomersRow) GetRemoved() bool {
	if saecr.Removed.Valid {
		return saecr.Removed.Bool
	}
	return false
}
func (saecr *SelectAllEmbroideryCustomersRow) GetRemovedText() string {
	if saecr.Removed.Valid {
		return fmt.Sprint(saecr.Removed.Bool)
	}
	return ""
}
func (saecr *SelectAllEmbroideryCustomersRow) GetCompletedat() time.Time {
	if saecr.Completedat.Valid {
		return saecr.Completedat.Time
	}
	return time.Time{}
}
func (saecr *SelectAllEmbroideryCustomersRow) GetCompletedatText() string { return "" }
func (saspcr *SelectAllScreenPrintCustomersRow) GetID() uuid.UUID {
	if saspcr.ID.Valid {
		return saspcr.ID.Bytes
	}
	return uuid.UUID{}
}
func (saspcr *SelectAllScreenPrintCustomersRow) GetIDText() string {
	if saspcr.ID.Valid {
		return fmt.Sprint(saspcr.ID.Bytes)
	}
	return ""
}
func (saspcr *SelectAllScreenPrintCustomersRow) GetName() string {
	if saspcr.Name.Valid {
		return saspcr.Name.String
	}
	return ""
}
func (saspcr *SelectAllScreenPrintCustomersRow) GetDescription() string {
	if saspcr.Description.Valid {
		return saspcr.Description.String
	}
	return ""
}
func (saspcr *SelectAllScreenPrintCustomersRow) GetOrderin() time.Time {
	if saspcr.Orderin.Valid {
		return saspcr.Orderin.Time
	}
	return time.Time{}
}
func (saspcr *SelectAllScreenPrintCustomersRow) GetOrderinText() string { return "" }
func (saspcr *SelectAllScreenPrintCustomersRow) GetOrderdue() time.Time {
	if saspcr.Orderdue.Valid {
		return saspcr.Orderdue.Time
	}
	return time.Time{}
}
func (saspcr *SelectAllScreenPrintCustomersRow) GetOrderdueText() string { return "" }
func (saspcr *SelectAllScreenPrintCustomersRow) GetRushed() bool {
	if saspcr.Rushed.Valid {
		return saspcr.Rushed.Bool
	}
	return false
}
func (saspcr *SelectAllScreenPrintCustomersRow) GetRushedText() string {
	if saspcr.Rushed.Valid {
		return fmt.Sprint(saspcr.Rushed.Bool)
	}
	return ""
}
func (saspcr *SelectAllScreenPrintCustomersRow) GetForm() string {
	if saspcr.Form.Valid {
		return saspcr.Form.String
	}
	return ""
}
func (saspcr *SelectAllScreenPrintCustomersRow) GetVisible() int {
	if saspcr.Visible.Valid {
		return int(saspcr.Visible.Int32)
	}
	return -1
}
func (saspcr *SelectAllScreenPrintCustomersRow) GetVisibleText() string {
	if saspcr.Visible.Valid {
		return fmt.Sprint(saspcr.Visible.Int32)
	}
	return ""
}
func (saspcr *SelectAllScreenPrintCustomersRow) GetEmail() string {
	if saspcr.Email.Valid {
		return saspcr.Email.String
	}
	return ""
}
func (saspcr *SelectAllScreenPrintCustomersRow) GetPhone() string {
	if saspcr.Phone.Valid {
		return saspcr.Phone.String
	}
	return ""
}
func (saspcr *SelectAllScreenPrintCustomersRow) GetCompleted() bool {
	if saspcr.Completed.Valid {
		return saspcr.Completed.Bool
	}
	return false
}
func (saspcr *SelectAllScreenPrintCustomersRow) GetCompletedText() string {
	if saspcr.Completed.Valid {
		return fmt.Sprint(saspcr.Completed.Bool)
	}
	return ""
}
func (saspcr *SelectAllScreenPrintCustomersRow) GetRemoved() bool {
	if saspcr.Removed.Valid {
		return saspcr.Removed.Bool
	}
	return false
}
func (saspcr *SelectAllScreenPrintCustomersRow) GetRemovedText() string {
	if saspcr.Removed.Valid {
		return fmt.Sprint(saspcr.Removed.Bool)
	}
	return ""
}
func (saspcr *SelectAllScreenPrintCustomersRow) GetCompletedat() time.Time {
	if saspcr.Completedat.Valid {
		return saspcr.Completedat.Time
	}
	return time.Time{}
}
func (saspcr *SelectAllScreenPrintCustomersRow) GetCompletedatText() string { return "" }
func (satcr *SelectAllTransferCustomersRow) GetID() uuid.UUID {
	if satcr.ID.Valid {
		return satcr.ID.Bytes
	}
	return uuid.UUID{}
}
func (satcr *SelectAllTransferCustomersRow) GetIDText() string {
	if satcr.ID.Valid {
		return fmt.Sprint(satcr.ID.Bytes)
	}
	return ""
}
func (satcr *SelectAllTransferCustomersRow) GetName() string {
	if satcr.Name.Valid {
		return satcr.Name.String
	}
	return ""
}
func (satcr *SelectAllTransferCustomersRow) GetDescription() string {
	if satcr.Description.Valid {
		return satcr.Description.String
	}
	return ""
}
func (satcr *SelectAllTransferCustomersRow) GetOrderin() time.Time {
	if satcr.Orderin.Valid {
		return satcr.Orderin.Time
	}
	return time.Time{}
}
func (satcr *SelectAllTransferCustomersRow) GetOrderinText() string { return "" }
func (satcr *SelectAllTransferCustomersRow) GetOrderdue() time.Time {
	if satcr.Orderdue.Valid {
		return satcr.Orderdue.Time
	}
	return time.Time{}
}
func (satcr *SelectAllTransferCustomersRow) GetOrderdueText() string { return "" }
func (satcr *SelectAllTransferCustomersRow) GetRushed() bool {
	if satcr.Rushed.Valid {
		return satcr.Rushed.Bool
	}
	return false
}
func (satcr *SelectAllTransferCustomersRow) GetRushedText() string {
	if satcr.Rushed.Valid {
		return fmt.Sprint(satcr.Rushed.Bool)
	}
	return ""
}
func (satcr *SelectAllTransferCustomersRow) GetForm() string {
	if satcr.Form.Valid {
		return satcr.Form.String
	}
	return ""
}
func (satcr *SelectAllTransferCustomersRow) GetVisible() int {
	if satcr.Visible.Valid {
		return int(satcr.Visible.Int32)
	}
	return -1
}
func (satcr *SelectAllTransferCustomersRow) GetVisibleText() string {
	if satcr.Visible.Valid {
		return fmt.Sprint(satcr.Visible.Int32)
	}
	return ""
}
func (satcr *SelectAllTransferCustomersRow) GetEmail() string {
	if satcr.Email.Valid {
		return satcr.Email.String
	}
	return ""
}
func (satcr *SelectAllTransferCustomersRow) GetPhone() string {
	if satcr.Phone.Valid {
		return satcr.Phone.String
	}
	return ""
}
func (satcr *SelectAllTransferCustomersRow) GetCompleted() bool {
	if satcr.Completed.Valid {
		return satcr.Completed.Bool
	}
	return false
}
func (satcr *SelectAllTransferCustomersRow) GetCompletedText() string {
	if satcr.Completed.Valid {
		return fmt.Sprint(satcr.Completed.Bool)
	}
	return ""
}
func (satcr *SelectAllTransferCustomersRow) GetRemoved() bool {
	if satcr.Removed.Valid {
		return satcr.Removed.Bool
	}
	return false
}
func (satcr *SelectAllTransferCustomersRow) GetRemovedText() string {
	if satcr.Removed.Valid {
		return fmt.Sprint(satcr.Removed.Bool)
	}
	return ""
}
func (satcr *SelectAllTransferCustomersRow) GetCompletedat() time.Time {
	if satcr.Completedat.Valid {
		return satcr.Completedat.Time
	}
	return time.Time{}
}
func (satcr *SelectAllTransferCustomersRow) GetCompletedatText() string { return "" }
func (slccr *SelectLatestCompletedCustomersRow) GetID() uuid.UUID {
	if slccr.ID.Valid {
		return slccr.ID.Bytes
	}
	return uuid.UUID{}
}
func (slccr *SelectLatestCompletedCustomersRow) GetIDText() string {
	if slccr.ID.Valid {
		return fmt.Sprint(slccr.ID.Bytes)
	}
	return ""
}
func (slccr *SelectLatestCompletedCustomersRow) GetName() string {
	if slccr.Name.Valid {
		return slccr.Name.String
	}
	return ""
}
func (slccr *SelectLatestCompletedCustomersRow) GetDescription() string {
	if slccr.Description.Valid {
		return slccr.Description.String
	}
	return ""
}
func (slccr *SelectLatestCompletedCustomersRow) GetOrderin() time.Time {
	if slccr.Orderin.Valid {
		return slccr.Orderin.Time
	}
	return time.Time{}
}
func (slccr *SelectLatestCompletedCustomersRow) GetOrderinText() string { return "" }
func (slccr *SelectLatestCompletedCustomersRow) GetOrderdue() time.Time {
	if slccr.Orderdue.Valid {
		return slccr.Orderdue.Time
	}
	return time.Time{}
}
func (slccr *SelectLatestCompletedCustomersRow) GetOrderdueText() string { return "" }
func (slccr *SelectLatestCompletedCustomersRow) GetRushed() bool {
	if slccr.Rushed.Valid {
		return slccr.Rushed.Bool
	}
	return false
}
func (slccr *SelectLatestCompletedCustomersRow) GetRushedText() string {
	if slccr.Rushed.Valid {
		return fmt.Sprint(slccr.Rushed.Bool)
	}
	return ""
}
func (slccr *SelectLatestCompletedCustomersRow) GetForm() string {
	if slccr.Form.Valid {
		return slccr.Form.String
	}
	return ""
}
func (slccr *SelectLatestCompletedCustomersRow) GetVisible() int {
	if slccr.Visible.Valid {
		return int(slccr.Visible.Int32)
	}
	return -1
}
func (slccr *SelectLatestCompletedCustomersRow) GetVisibleText() string {
	if slccr.Visible.Valid {
		return fmt.Sprint(slccr.Visible.Int32)
	}
	return ""
}
func (slccr *SelectLatestCompletedCustomersRow) GetEmail() string {
	if slccr.Email.Valid {
		return slccr.Email.String
	}
	return ""
}
func (slccr *SelectLatestCompletedCustomersRow) GetPhone() string {
	if slccr.Phone.Valid {
		return slccr.Phone.String
	}
	return ""
}
func (slccr *SelectLatestCompletedCustomersRow) GetCompleted() bool {
	if slccr.Completed.Valid {
		return slccr.Completed.Bool
	}
	return false
}
func (slccr *SelectLatestCompletedCustomersRow) GetCompletedText() string {
	if slccr.Completed.Valid {
		return fmt.Sprint(slccr.Completed.Bool)
	}
	return ""
}
func (slccr *SelectLatestCompletedCustomersRow) GetRemoved() bool {
	if slccr.Removed.Valid {
		return slccr.Removed.Bool
	}
	return false
}
func (slccr *SelectLatestCompletedCustomersRow) GetRemovedText() string {
	if slccr.Removed.Valid {
		return fmt.Sprint(slccr.Removed.Bool)
	}
	return ""
}
func (slccr *SelectLatestCompletedCustomersRow) GetCompletedat() time.Time {
	if slccr.Completedat.Valid {
		return slccr.Completedat.Time
	}
	return time.Time{}
}
func (slccr *SelectLatestCompletedCustomersRow) GetCompletedatText() string { return "" }
func (icp *InsertCustomerParams) GetName() string {
	if icp.Name.Valid {
		return icp.Name.String
	}
	return ""
}
func (icp *InsertCustomerParams) GetPhone() string {
	if icp.Phone.Valid {
		return icp.Phone.String
	}
	return ""
}
func (icp *InsertCustomerParams) GetEmail() string {
	if icp.Email.Valid {
		return icp.Email.String
	}
	return ""
}
func (icp *InsertCustomerParams) GetOrderin() time.Time {
	if icp.Orderin.Valid {
		return icp.Orderin.Time
	}
	return time.Time{}
}
func (icp *InsertCustomerParams) GetOrderinText() string { return "" }
func (icp *InsertCustomerParams) GetOrderdue() time.Time {
	if icp.Orderdue.Valid {
		return icp.Orderdue.Time
	}
	return time.Time{}
}
func (icp *InsertCustomerParams) GetOrderdueText() string { return "" }
func (icp *InsertCustomerParams) GetRushed() bool {
	if icp.Rushed.Valid {
		return icp.Rushed.Bool
	}
	return false
}
func (icp *InsertCustomerParams) GetRushedText() string {
	if icp.Rushed.Valid {
		return fmt.Sprint(icp.Rushed.Bool)
	}
	return ""
}
func (icp *InsertCustomerParams) GetForm() string {
	if icp.Form.Valid {
		return icp.Form.String
	}
	return ""
}
func (icp *InsertCustomerParams) GetDescription() string {
	if icp.Description.Valid {
		return icp.Description.String
	}
	return ""
}
func (ucp *UpdateCustomerParams) GetName() string {
	if ucp.Name.Valid {
		return ucp.Name.String
	}
	return ""
}
func (ucp *UpdateCustomerParams) GetPhone() string {
	if ucp.Phone.Valid {
		return ucp.Phone.String
	}
	return ""
}
func (ucp *UpdateCustomerParams) GetEmail() string {
	if ucp.Email.Valid {
		return ucp.Email.String
	}
	return ""
}
func (ucp *UpdateCustomerParams) GetOrderin() time.Time {
	if ucp.Orderin.Valid {
		return ucp.Orderin.Time
	}
	return time.Time{}
}
func (ucp *UpdateCustomerParams) GetOrderinText() string { return "" }
func (ucp *UpdateCustomerParams) GetOrderdue() time.Time {
	if ucp.Orderdue.Valid {
		return ucp.Orderdue.Time
	}
	return time.Time{}
}
func (ucp *UpdateCustomerParams) GetOrderdueText() string { return "" }
func (ucp *UpdateCustomerParams) GetRushed() bool {
	if ucp.Rushed.Valid {
		return ucp.Rushed.Bool
	}
	return false
}
func (ucp *UpdateCustomerParams) GetRushedText() string {
	if ucp.Rushed.Valid {
		return fmt.Sprint(ucp.Rushed.Bool)
	}
	return ""
}
func (ucp *UpdateCustomerParams) GetType() string {
	if ucp.Type.Valid {
		return ucp.Type.String
	}
	return ""
}
func (ucp *UpdateCustomerParams) GetDescription() string {
	if ucp.Description.Valid {
		return ucp.Description.String
	}
	return ""
}
func (ucp *UpdateCustomerParams) GetID() uuid.UUID {
	if ucp.ID.Valid {
		return ucp.ID.Bytes
	}
	return uuid.UUID{}
}
func (ucp *UpdateCustomerParams) GetIDText() string {
	if ucp.ID.Valid {
		return fmt.Sprint(ucp.ID.Bytes)
	}
	return ""
}
