package nfeio

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/seriallink/go-curl"
)

func (c *Client) GetCertificates(id string) (response CertificatesResponse, err error) {
	err = c.Get(fmt.Sprintf("companies/%s/%s", id, Ternary(c.apiUrl == ServiceUrl, "certificate", "certificates")), nil, nil, &response)
	return
}

func (c *Client) UploadCertificate(id, file, password string) error {

	var (
		header []string
		body   bytes.Buffer
		erm    ErrMessage
	)

	// init curl
	easy := curl.EasyInit()

	// close curl
	defer easy.Cleanup()

	// set curl opt
	_ = easy.Setopt(curl.OPT_URL, fmt.Sprintf("%scompanies/%s/%s", c.apiUrl, id, Ternary(c.apiUrl == ServiceUrl, "certificate", "certificates")))
	_ = easy.Setopt(curl.OPT_SSL_VERIFYPEER, false)

	// enable on debug mode
	if os.Getenv("debug") == "true" {
		_ = easy.Setopt(curl.OPT_VERBOSE, true)
	}

	// set header
	_ = easy.Setopt(curl.OPT_HTTPHEADER, []string{
		"Authorization: " + c.apiKey,
		"Content-Type: multipart/form-data",
		"Accept: application/json",
	})

	// init form
	form := curl.NewForm()

	// add form fields
	_ = form.AddFile("file", file)
	_ = form.Add("password", password)

	// define http method
	_ = easy.Setopt(curl.OPT_HTTPPOST, form)

	// print upload progress
	//_ = easy.Setopt(curl.OPT_NOPROGRESS, false)
	//_ = easy.Setopt(curl.OPT_PROGRESSFUNCTION, func(dltotal, dlnow, ultotal, ulnow float64, _ interface{}) bool {
	//	fmt.Printf("Download %3.2f%%, Uploading %3.2f%%\r", dlnow/dltotal*100, ulnow/ultotal*100)
	//	return true
	//})

	// get response body (callback)
	_ = easy.Setopt(curl.OPT_WRITEFUNCTION, func(buf []byte, data interface{}) bool {
		body.Write(buf)
		return true
	})

	// get response header (callback)
	_ = easy.Setopt(curl.OPT_HEADERFUNCTION, func(buf []byte, data interface{}) bool {
		header = append(header, strings.Replace(string(buf), "\r\n", "", -1))
		return true
	})

	// post certificate
	if err := easy.Perform(); err != nil {
		return err
	}

	// check error message
	if err := json.NewDecoder(&body).Decode(&erm); err == nil && erm.Message != "" {
		return erm
	}

	// check response status code
	if len(header) < 1 || NotLikeAny(header[0], "100", "200") {

		// get status code
		data := strings.Split(strings.Trim(header[0], " "), " ")

		// parse to int
		code, _ := strconv.Atoi(data[1])

		// return error code
		return errors.New(http.StatusText(code))

	}

	return nil

}

type CertificatesResponse struct {
	Data []Certificate `json:"certificates"`
	pagination
}

type Certificate struct {
	FederalTaxNumber string    `json:"federalTaxNumber,omitempty"`
	Thumbprint       string    `json:"thumbprint,omitempty"`
	Status           string    `json:"status,omitempty"`
	Subject          string    `json:"subject,omitempty"`
	ModifiedOn       time.Time `json:"modifiedOn,omitempty"`
	ExpiresOn        time.Time `json:"expiresOn,omitempty"`
	ValidUntil       time.Time `json:"validUntil,omitempty"`
}
