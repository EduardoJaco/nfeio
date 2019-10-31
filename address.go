package nfeio

func (c *Client) SearchAddressByPostalCode(postalCode string) (response Address, err error) {
	response.State = &State{}
	err = c.Get("addresses/"+postalCode, nil, nil, &response)
	return
}

type Address struct {
	StreetSuffix          string      `json:"streetSuffix,omitempty"`
	Street                string      `json:"street,omitempty"`
	Number                string      `json:"number,omitempty"`
	NumberMin             string      `json:"numberMin,omitempty"`
	NumberMax             string      `json:"numberMax,omitempty"`
	AdditionalInformation string      `json:"additionalInformation,omitempty"`
	District              string      `json:"district,omitempty"`
	State                 interface{} `json:"state,omitempty"`
	City                  *City       `json:"city,omitempty"`
	PostalCode            string      `json:"postalCode,omitempty"`
	Country               string      `json:"country,omitempty"`
}

type City struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type State struct {
	Code         string `json:"code,omitempty"`
	Name         string `json:"name,omitempty"`
	Abbreviation string `json:"abbreviation"`
}
