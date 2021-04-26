package nfeio

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func (c *Client) GetCertificates(id string) (response CertificatesResponse, err error) {
	err = c.Get(fmt.Sprintf("companies/%s/%s", id, Ternary(c.apiUrl == ServiceUrl, "certificate", "certificates")), nil, nil, &response)
	return
}

func (c *Client) UploadCertificate(id, path, password string) (err error) {

	var (
		file *os.File
		part io.Writer
		resp *http.Response
		body []byte
	)

	// init request body
	payload := &bytes.Buffer{}

	// create form writer
	writer := multipart.NewWriter(payload)

	// open file
	if file, err = os.Open(path); err != nil {
		return
	}

	// close file
	defer file.Close()

	// set file
	if part, err = writer.CreateFormFile("file", filepath.Base(file.Name())); err != nil {
		return
	}

	// write file
	if _, err = io.Copy(part, file); err != nil {
		return
	}

	// set password
	if err = writer.WriteField("password", password); err != nil {
		return
	}

	// close writer
	if err = writer.Close(); err != nil {
		return
	}

	// set request
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%scompanies/%s/%s", c.apiUrl, id, Ternary(c.apiUrl == ServiceUrl, "certificate", "certificates")), payload)
	req.Header.Add("Authorization", c.apiKey)
	req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")

	// do request
	if resp, err = http.DefaultClient.Do(req); err != nil {
		return
	}

	// read response
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	// close response
	defer resp.Body.Close()

	// init error response
	erm := &ErrMessage{}

	// check for error message
	if err = json.Unmarshal(body, erm); err == nil && erm.Message != "" {
		return erm
	}

	// init array of errors
	errs := &ErrArray{}

	// check for multiple errors
	if err = json.Unmarshal(body, errs); err == nil && errs.Count() > 0 {
		return errs
	}

	// verify status code
	if NotIn(resp.StatusCode, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent, http.StatusContinue) {

		if len(body) > 0 {
			return errors.New(string(body))
		}

		return errors.New(resp.Status)

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
