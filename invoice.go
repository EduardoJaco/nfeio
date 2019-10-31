package nfeio

import (
	"encoding/json"
	"fmt"
	"time"
)

func (c *Client) GetInvoice(idCompany, idInvoice string) (response Invoice, err error) {
	err = c.Get(fmt.Sprintf("companies/%s/serviceinvoices/%s", idCompany, idInvoice), nil, nil, &response)
	return
}

func (c *Client) GetInvoices(idCompany string, pageCount, pageIndex int64) (response InvoicesResponse, err error) {
	err = c.Get(fmt.Sprintf("companies/%s/serviceinvoices", idCompany), Params{"pageCount": pageCount, "pageIndex": pageIndex}, nil, &response)
	return
}

func (c *Client) IssueInvoice(idCompany string, invoice *Invoice) (response Invoice, err error) {
	err = c.Post(fmt.Sprintf("companies/%s/serviceinvoices", idCompany), invoice, nil, &response)
	return
}

func (c *Client) CancelInvoice(idCompany, idInvoice string) (response Invoice, err error) {
	err = c.Delete(fmt.Sprintf("companies/%s/serviceinvoices/%s", idCompany, idInvoice), nil, nil, &response)
	return
}

func (c *Client) DownloadPDF(idCompany, idInvoice string) (pdf []byte, err error) {
	pdf = make([]byte, 0)
	err = c.Get(fmt.Sprintf("companies/%s/serviceinvoices/%s/pdf", idCompany, idInvoice), nil, nil, &pdf)
	return
}

func (c *Client) DownloadXML(idCompany, idInvoice string) (xml string, err error) {
	err = c.Get(fmt.Sprintf("companies/%s/serviceinvoices/%s/xml", idCompany, idInvoice), nil, nil, &xml)
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
	Id                          string          `json:"id,omitempty"`
	Environment                 string          `json:"environment,omitempty"`
	FlowStatus                  string          `json:"flowStatus,omitempty"`
	FlowMessage                 string          `json:"flowMessage,omitempty"`
	BatchNumber                 json.Number     `json:"batchNumber,omitempty"`
	BatchCheckNumber            json.Number     `json:"batchCheckNumber,omitempty"`
	Number                      json.Number     `json:"number,omitempty"`
	CheckCode                   string          `json:"checkCode,omitempty"`
	Status                      string          `json:"status,omitempty"`
	RpsType                     string          `json:"rpsType,omitempty"`
	RpsStatus                   string          `json:"rpsStatus,omitempty"`
	TaxationType                string          `json:"taxationType,omitempty"`
	IssuedOn                    *time.Time      `json:"issuedOn,omitempty"`
	CancelledOn                 *time.Time      `json:"cancelledOn,omitempty"`
	RpsSerialNumber             json.Number     `json:"rpsSerialNumber,omitempty"`
	RpsNumber                   json.Number     `json:"rpsNumber,omitempty"`
	CityServiceCode             string          `json:"cityServiceCode,omitempty"`
	FederalServiceCode          string          `json:"federalServiceCode,omitempty"`
	Description                 string          `json:"description,omitempty"`
	ServicesAmount              float64         `json:"servicesAmount,omitempty"`
	DeductionsAmount            float64         `json:"deductionsAmount,omitempty"`
	DiscountUnconditionedAmount float64         `json:"discountUnconditionedAmount,omitempty"`
	DiscountConditionedAmount   float64         `json:"discountConditionedAmount,omitempty"`
	BaseTaxAmount               float64         `json:"baseTaxAmount,omitempty"`
	IssRate                     json.Number     `json:"issRate,omitempty"`
	IssTaxAmount                float64         `json:"issTaxAmount,omitempty"`
	IrAmountWithheld            float64         `json:"irAmountWithheld,omitempty"`
	PisAmountWithheld           float64         `json:"pisAmountWithheld,omitempty"`
	CofinsAmountWithheld        float64         `json:"cofinsAmountWithheld,omitempty"`
	CsllAmountWithheld          float64         `json:"csllAmountWithheld,omitempty"`
	InssAmountWithheld          float64         `json:"inssAmountWithheld,omitempty"`
	IssAmountWithheld           float64         `json:"issAmountWithheld,omitempty"`
	OthersAmountWithheld        float64         `json:"othersAmountWithheld,omitempty"`
	AmountWithheld              float64         `json:"amountWithheld,omitempty"`
	AmountNet                   float64         `json:"amountNet,omitempty"`
	Provider                    *Provider       `json:"provider,omitempty"`
	Borrower                    *Borrower       `json:"borrower,omitempty"`
	ApproximateTax              *ApproximateTax `json:"approximateTax,omitempty"`
	AdditionalInformation       string          `json:"additionalInformation,omitempty"`
	CreatedOn                   *time.Time      `json:"createdOn,omitempty"`
	ModifiedOn                  *time.Time      `json:"modifiedOn,omitempty"`
}

type ApproximateTax struct {
	Source    string      `json:"source,omitempty"`
	Version   string      `json:"version,omitempty"`
	TotalRate json.Number `json:"totalRate,omitempty"`
}
