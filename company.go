package nfeio

import (
	"encoding/json"
	"time"
)

func (c *Client) GetCompany(idOrTaxNumber string) (response CompanyResponse, err error) {
	err = c.Get("companies/"+idOrTaxNumber, nil, nil, &response)
	return
}

func (c *Client) GetCompanies(pageCount, pageIndex int64) (response CompaniesResponse, err error) {
	err = c.Get("companies", Params{"pageCount": pageCount, "pageIndex": pageIndex}, nil, &response)
	return
}

func (c *Client) AddCompany(company *Company) (response CompanyResponse, err error) {

	var params interface{}

	if c.apiUrl == ServiceUrl {
		params = company
	} else {
		params = map[string]interface{}{
			"company": company,
		}
	}

	err = c.Post("companies", params, nil, &response)

	return

}

func (c *Client) SetCompany(company *Company) (response CompanyResponse, err error) {

	var params interface{}

	if c.apiUrl == ServiceUrl {
		params = company
	} else {
		params = map[string]interface{}{
			"company": company,
		}
	}

	err = c.Put("companies/"+company.Id, params, nil, &response)

	return

}

func (c *Client) DeleteCompany(id string) (err error) {
	err = c.Delete("companies/"+id, nil, nil, nil)
	return
}

type CompanyResponse struct {
	Data      Company
	Company   Company `json:"company"`
	Companies Company `json:"companies"`
}

type CompaniesResponse struct {
	Data []Company `json:"companies"`
	pagination
}

type Company struct {
	Id                    string             `json:"id,omitempty"`
	Name                  string             `json:"name,omitempty"`
	TradeName             string             `json:"tradeName,omitempty"`
	Email                 string             `json:"email,omitempty"`
	FederalTaxNumber      json.Number        `json:"federalTaxNumber,omitempty"`
	RegionalTaxNumber     json.Number        `json:"regionalTaxNumber,omitempty"`
	MunicipalTaxNumber    string             `json:"municipalTaxNumber,omitempty"`
	CompanyRegistryNumber json.Number        `json:"companyRegistryNumber,omitempty"`
	TaxRegime             interface{}        `json:"taxRegime,omitempty"`
	SpecialTaxRegime      interface{}        `json:"specialTaxRegime,omitempty"`
	LegalNature           interface{}        `json:"legalNature,omitempty"`
	RpsSerialNumber       interface{}        `json:"rpsSerialNumber,omitempty"`
	RpsNumber             interface{}        `json:"rpsNumber,omitempty"`
	IssRate               float64            `json:"issRate,omitempty"`
	Environment           string             `json:"environment,omitempty"`
	FiscalStatus          string             `json:"fiscalStatus,omitempty"`
	Address               *Address           `json:"address,omitempty"`
	Certificate           *Certificate       `json:"certificate,omitempty"`
	EconomicActivities    []EconomicActivity `json:"economicActivities,omitempty"`
	OpenningDate          *time.Time         `json:"openningDate,omitempty"`
	CreatedOn             *time.Time         `json:"createdOn,omitempty"`
	ModifiedOn            *time.Time         `json:"modifiedOn,omitempty"`
}

type EconomicActivity struct {
	Type string      `json:"type,omitempty"`
	Code json.Number `json:"code,omitempty"`
}

// only to avoid recursion in UnmarshalJSON below
type cr CompanyResponse

// override default json.Unmarshal
func (c *CompanyResponse) UnmarshalJSON(b []byte) (err error) {

	aux := cr{}

	if err = json.Unmarshal(b, &aux); err == nil {

		*c = CompanyResponse(aux)

		if c.Company.Id != "" {
			c.Data = c.Company
		}

		if c.Companies.Id != "" {
			c.Data = c.Companies
		}

	}

	return nil
}
