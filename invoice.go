package nfeio

import (
	"encoding/json"
	"fmt"
	"time"
)

func (c *Client) getInvoicePath() string {
	if c.apiUrl == ServiceUrl {
		return "serviceinvoices"
	}
	return "productinvoices"
}

func (c *Client) GetInvoice(idCompany, idInvoice string) (response Invoice, err error) {
	err = c.Get(fmt.Sprintf("companies/%s/%s/%s", idCompany, c.getInvoicePath(), idInvoice), nil, nil, &response)
	return
}

func (c *Client) GetInvoices(idCompany string, pageCount, pageIndex int64) (response InvoicesResponse, err error) {
	err = c.Get(fmt.Sprintf("companies/%s/%s", idCompany, c.getInvoicePath()), Params{"pageCount": pageCount, "pageIndex": pageIndex}, nil, &response)
	return
}

func (c *Client) IssueInvoice(idCompany string, invoice *Invoice) (response Invoice, err error) {
	err = c.Post(fmt.Sprintf("companies/%s/%s", idCompany, c.getInvoicePath()), invoice, nil, &response)
	return
}

func (c *Client) CancelInvoice(idCompany, idInvoice string) (response Invoice, err error) {
	err = c.Delete(fmt.Sprintf("companies/%s/%s/%s", idCompany, c.getInvoicePath(), idInvoice), nil, nil, &response)
	return
}

func (c *Client) DownloadPDF(idCompany, idInvoice string) (pdf []byte, err error) {
	pdf = make([]byte, 0)
	err = c.Get(fmt.Sprintf("companies/%s/%s/%s/pdf", idCompany, c.getInvoicePath(), idInvoice), nil, nil, &pdf)
	return
}

func (c *Client) DownloadXML(idCompany, idInvoice string) (xml string, err error) {
	err = c.Get(fmt.Sprintf("companies/%s/%s/%s/xml", idCompany, c.getInvoicePath(), idInvoice), nil, nil, &xml)
	return
}

func (c *Invoice) TotalAmount() (total float64) {

	for _, item := range c.Items {

		quantity, _ := item.Quantity.Float64()
		unitamount, _ := item.UnitAmount.Float64()

		total += quantity * unitamount

	}

	return

}

type InvoiceResponse struct {
	Data Invoice `json:"serviceInvoices"`
}

type InvoicesResponse struct {
	Data []Invoice `json:"serviceInvoices"`
	pagination
}

type Invoice struct {
	Id                          string                 `json:"id,omitempty"`
	Environment                 string                 `json:"environment,omitempty"`
	FlowStatus                  string                 `json:"flowStatus,omitempty"`
	FlowMessage                 string                 `json:"flowMessage,omitempty"`
	BatchNumber                 json.Number            `json:"batchNumber,omitempty"`
	BatchCheckNumber            json.Number            `json:"batchCheckNumber,omitempty"`
	Number                      json.Number            `json:"number,omitempty"`
	CheckCode                   string                 `json:"checkCode,omitempty"`
	Status                      string                 `json:"status,omitempty"`
	RpsType                     string                 `json:"rpsType,omitempty"`
	RpsStatus                   string                 `json:"rpsStatus,omitempty"`
	TaxationType                string                 `json:"taxationType,omitempty"`
	IssuedOn                    *time.Time             `json:"issuedOn,omitempty"`
	CancelledOn                 *time.Time             `json:"cancelledOn,omitempty"`
	RpsNumber                   json.Number            `json:"rpsNumber,omitempty"`
	RpsSerialNumber             string                 `json:"rpsSerialNumber,omitempty"`
	CityServiceCode             string                 `json:"cityServiceCode,omitempty"`
	FederalServiceCode          string                 `json:"federalServiceCode,omitempty"`
	Description                 string                 `json:"description,omitempty"`
	ServicesAmount              float64                `json:"servicesAmount,omitempty"`
	DeductionsAmount            float64                `json:"deductionsAmount,omitempty"`
	DiscountUnconditionedAmount float64                `json:"discountUnconditionedAmount,omitempty"`
	DiscountConditionedAmount   float64                `json:"discountConditionedAmount,omitempty"`
	BaseTaxAmount               float64                `json:"baseTaxAmount,omitempty"`
	IssRate                     json.Number            `json:"issRate,omitempty"`
	IssTaxAmount                float64                `json:"issTaxAmount,omitempty"`
	IrAmountWithheld            float64                `json:"irAmountWithheld,omitempty"`
	PisAmountWithheld           float64                `json:"pisAmountWithheld,omitempty"`
	CofinsAmountWithheld        float64                `json:"cofinsAmountWithheld,omitempty"`
	CsllAmountWithheld          float64                `json:"csllAmountWithheld,omitempty"`
	InssAmountWithheld          float64                `json:"inssAmountWithheld,omitempty"`
	IssAmountWithheld           float64                `json:"issAmountWithheld,omitempty"`
	OthersAmountWithheld        float64                `json:"othersAmountWithheld,omitempty"`
	AmountWithheld              float64                `json:"amountWithheld,omitempty"`
	AmountNet                   float64                `json:"amountNet,omitempty"`
	Provider                    *Provider              `json:"provider,omitempty"`
	Borrower                    *Borrower              `json:"borrower,omitempty"`
	Buyer                       *Buyer                 `json:"buyer,omitempty"`
	Items                       []Item                 `json:"items,omitempty"`
	ApproximateTax              *ApproximateTax        `json:"approximateTax,omitempty"`
	AdditionalInformation       *AdditionalInformation `json:"additionalInformation,omitempty"`
	CreatedOn                   *time.Time             `json:"createdOn,omitempty"`
	ModifiedOn                  *time.Time             `json:"modifiedOn,omitempty"`
}

type ApproximateTax struct {
	Source    string      `json:"source,omitempty"`
	Version   string      `json:"version,omitempty"`
	TotalRate json.Number `json:"totalRate,omitempty"`
}

type AdditionalInformation struct {
	Fisco                 string                  `json:"fisco,omitempty"`
	Taxpayer              string                  `json:"taxpayer,omitempty"`
	XmlAuthorized         []json.Number           `json:"xmlAuthorized,omitempty"`
	Effort                string                  `json:"effort,omitempty"`
	Order                 string                  `json:"order,omitempty"`
	Contract              string                  `json:"contract,omitempty"`
	TaxDocumentsReference []TaxDocumentsReference `json:"taxDocumentsReference,omitempty"`
	TaxpayerComments      []TaxpayerComments      `json:"taxpayerComments,omitempty"`
	ReferencedProcess     []ReferencedProcess     `json:"referencedProcess,omitempty"`
}

type TaxDocumentsReference struct {
	TaxCouponInformation      *TaxCouponInformation      `json:"taxCouponInformation,omitempty"`
	DocumentInvoiceReference  *DocumentInvoiceReference  `json:"documentInvoiceReference,omitempty"`
	DocumentElectronicInvoice *DocumentElectronicInvoice `json:"documentElectronicInvoice,omitempty"`
}

type TaxCouponInformation struct {
	ModelDocumentFiscal string `json:"modelDocumentFiscal,omitempty"`
	OrderECF            string `json:"orderECF,omitempty"`
	OrderCountOperation string `json:"orderCountOperation,omitempty"`
}

type DocumentInvoiceReference struct {
	State            json.Number `json:"state,omitempty"`
	YearMonth        string      `json:"yearMonth,omitempty"`
	FederalTaxNumber string      `json:"federalTaxNumber,omitempty"`
	Model            string      `json:"model,omitempty"`
	Series           string      `json:"series,omitempty"`
	Number           string      `json:"number,omitempty"`
}

type DocumentElectronicInvoice struct {
	AccessKey string `json:"accessKey,omitempty"`
}

type TaxpayerComments struct {
	Field string `json:"field,omitempty"`
	Text  string `json:"text,omitempty"`
}

type ReferencedProcess struct {
	IdentifierConcessory string      `json:"identifierConcessory,omitempty"`
	IdentifierOrigin     json.Number `json:"identifierOrigin,omitempty"`
}
