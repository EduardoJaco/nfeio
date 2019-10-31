package nfeio

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchAddressByPostalCode(t *testing.T) {
	address, err := client.SearchAddressByPostalCode("03181010")
	assert.NoError(t, err)
	assert.Equal(t, address.PostalCode, "03181010")
}
