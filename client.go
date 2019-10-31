package nfeio

import "strings"

type Client struct {
	ApiKey  string
	OpenKey string
}

func NewClient(apiKey string, openKey string) *Client {
	return &Client{ApiKey: apiKey, OpenKey: openKey}
}

func (c *Client) GetAuthorization(path string) string {

	if strings.Contains(path, "addresses") {
		return c.OpenKey
	}

	return c.ApiKey

}

func (c *Client) GetEndpoint(path string) string {

	if strings.Contains(path, "addresses") {
		return addressUrl
	}

	return baseUrl

}
