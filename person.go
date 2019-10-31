package nfeio

import (
	"encoding/json"
	"time"
)

type Provider struct {
	Id                    string             `json:"id,omitempty"`
	ParentId              string             `json:"parentId,omitempty"`
	Name                  string             `json:"name,omitempty"`
	Email                 string             `json:"email,omitempty"`
	TradeName             string             `json:"tradeName,omitempty"`
	TaxRegime             json.Number        `json:"taxRegime,omitempty"`
	SpecialTaxRegime      json.Number        `json:"specialTaxRegime,omitempty"`
	LegalNature           json.Number        `json:"legalNature,omitempty"`
	CompanyRegistryNumber json.Number        `json:"companyRegistryNumber,omitempty"`
	RegionalTaxNumber     json.Number        `json:"regionalTaxNumber,omitempty"`
	MunicipalTaxNumber    json.Number        `json:"municipalTaxNumber,omitempty"`
	FederalTaxNumber      json.Number        `json:"federalTaxNumber,omitempty"`
	IssRate               json.Number        `json:"issRate,omitempty"`
	Status                string             `json:"status,omitempty"`
	Type                  string             `json:"type,omitempty"`
	Address               *Address           `json:"address,omitempty"`
	EconomicActivities    []EconomicActivity `json:"economicActivities,omitempty"`
	OpenningDate          *time.Time         `json:"openningDate,omitempty"`
	CreatedOn             *time.Time         `json:"createdOn,omitempty"`
	ModifiedOn            *time.Time         `json:"modifiedOn,omitempty"`
}

type Borrower struct {
	Id                 string      `json:"id,omitempty"`
	ParentId           string      `json:"parentId,omitempty"`
	Name               string      `json:"name,omitempty"`
	Email              string      `json:"email,omitempty"`
	FederalTaxNumber   json.Number `json:"federalTaxNumber,omitempty"`
	MunicipalTaxNumber json.Number `json:"municipalTaxNumber,omitempty"`
	Status             string      `json:"status,omitempty"`
	Type               string      `json:"type,omitempty"`
	Address            *Address    `json:"address,omitempty"`
	CreatedOn          *time.Time  `json:"createdOn,omitempty"`
	ModifiedOn         *time.Time  `json:"modifiedOn,omitempty"`
}
