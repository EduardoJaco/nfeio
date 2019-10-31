package nfeio

var client *Client

func init() {
	client = NewClient(
		"",
		"",
	)
}
