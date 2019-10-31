package nfeio

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCompany(t *testing.T) {
	response, err := client.GetCompany("")
	assert.NoError(t, err)
	assert.Equal(t, "", response.Data.Id)
}

func TestGetCompanies(t *testing.T) {
	response, err := client.GetCompanies(1, 10)
	assert.NoError(t, err)
	assert.NotZero(t, len(response.Data))
}

func TestAddCompany(t *testing.T) {
	company := &Company{
		Name:               "",
		FederalTaxNumber:   "",
		MunicipalTaxNumber: "",
		Address: &Address{
			Street:                "",
			Number:                "",
			AdditionalInformation: "",
			PostalCode:            "",
			City: &City{
				Name: "São Paulo",
				Code: "3550308",
			},
			State: "SP",
		},
	}
	response, err := client.AddCompany(company)
	assert.NoError(t, err)
	assert.Equal(t, "", response.Data.Id)
}

func TestSetCompany(t *testing.T) {
	company := &Company{
		Id:               "",
		FederalTaxNumber: "",
		Name:             "My New Company",
		Email:            "my@company.com",
	}
	response, err := client.SetCompany(company)
	assert.NoError(t, err)
	assert.Equal(t, company.Id, response.Data.Id)
}

func TestDeleteCompany(t *testing.T) {
	err := client.DeleteCompany("")
	assert.NoError(t, err)
}

//func TestUploadCertificate(t *testing.T) {
//	err := client.UploadCertificate("", "", "")
//	assert.NoError(t, err)
//}
