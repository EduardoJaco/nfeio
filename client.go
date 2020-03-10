package nfeio

import "strings"

type Client struct {
	apiKey  string
	openKey string
	apiUrl  string
}

func NewClient(apiKey, openKey, apiUrl string) *Client {
	return &Client{
		apiKey:  apiKey,
		openKey: openKey,
		apiUrl:  apiUrl,
	}
}

func (c *Client) GetAuthorization(path string) string {

	if strings.Contains(path, "addresses") {
		return c.openKey
	}

	return c.apiKey

}

func (c *Client) GetEndpoint(path string) string {

	if strings.Contains(path, "addresses") {
		return AddressUrl
	}

	return c.apiUrl

}
