package nfeio

import (
	"encoding/json"
	"fmt"
	"time"
)

func (c *Client) AddStateTax(companyId string, stateTax *StateTax) (response StateTaxResponse, err error) {
	err = c.Post(fmt.Sprintf("companies/%s/statetaxes", companyId), Params{"stateTax": stateTax}, nil, &response)
	return
}

func (c *Client) DeleteStateTax(companyId, stateTaxId string) (err error) {
	err = c.Delete(fmt.Sprintf("companies/%s/statetaxes/%s", companyId, stateTaxId), nil, nil, nil)
	return
}

type StateTaxResponse struct {
	Data StateTax `json:"stateTax"`
}

type StateTax struct {
	//request
	Code               string              `json:"code,omitempty"`
	EnvironmentType    string              `json:"environmentType,omitempty"`
	TaxNumber          string              `json:"taxNumber,omitempty"`
	SpecialTaxRegime   string              `json:"specialTaxRegime,omitempty"`
	Serie              json.Number         `json:"serie,omitempty"`
	Number             json.Number         `json:"number,omitempty"`
	Type               string              `json:"type,omitempty"`
	SecurityCredential *SecurityCredential `json:"securityCredential,omitempty"`
	// response
	Id         string        `json:"id,omitempty"`
	CompanyId  string        `json:"companyId,omitempty"`
	AccountId  string        `json:"accountId,omitempty"`
	BatchId    json.Number   `json:"batchId,omitempty"`
	Status     string        `json:"status,omitempty"`
	Series     []json.Number `json:"series,omitempty"`
	CreatedOn  *time.Time    `json:"createdOn,omitempty"`
	ModifiedOn *time.Time    `json:"modifiedOn,omitempty"`
}

type SecurityCredential struct {
	Id   json.Number `json:"id,omitempty"`
	Code string      `json:"code,omitempty"`
}
