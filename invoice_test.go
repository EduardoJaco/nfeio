package nfeio

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInvoice(t *testing.T) {
	response, err := client.GetInvoice("", "")
	assert.NoError(t, err)
	assert.Equal(t, "", response.Id)
}

func TestGetInvoices(t *testing.T) {
	response, err := client.GetInvoices("", 1, 10)
	assert.NoError(t, err)
	assert.NotZero(t, len(response.Data))
}

func TestIssueInvoice(t *testing.T) {
	invoice := &Invoice{
		ServicesAmount:  100,
		RpsNumber:       "3",
		Description:     "Testando emissao de NFSE",
		CityServiceCode: "1058",
		Borrower: &Borrower{
			Email: "john@doe.com",
		},
	}
	response, err := client.IssueInvoice("", invoice)
	assert.NoError(t, err)
	assert.Equal(t, invoice.RpsNumber, response.RpsNumber)
}

func TestCancelInvoice(t *testing.T) {
	invoice, err := client.CancelInvoice("", "")
	assert.NoError(t, err)
	assert.Equal(t, "", invoice.Id)
}

func TestDownloadPDF(t *testing.T) {
	pdf, err := client.DownloadPDF("", "")
	assert.NoError(t, err)
	assert.NotZero(t, len(pdf))
}

func TestDownloadXML(t *testing.T) {
	xml, err := client.DownloadXML("", "")
	assert.NoError(t, err)
	assert.NotZero(t, len(xml))
}
