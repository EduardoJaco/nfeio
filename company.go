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

func (c *Client) GetCompany(idOrTaxNumber string) (response CompanyResponse, err error) {
	err = c.Get("companies/"+idOrTaxNumber, nil, nil, &response)
	return
}

func (c *Client) GetCompanies(pageCount, pageIndex int64) (response CompaniesResponse, err error) {
	err = c.Get("companies", Params{"pageCount": pageCount, "pageIndex": pageIndex}, nil, &response)
	return
}

func (c *Client) AddCompany(company *Company) (response CompanyResponse, err error) {

	var params interface{}

	if c.apiUrl == ServiceUrl {
		params = company
	} else {
		params = map[string]interface{}{
			"company": company,
		}
	}

	err = c.Post("companies", params, nil, &response)

	return

}

func (c *Client) SetCompany(company *Company) (response CompanyResponse, err error) {

	var params interface{}

	if c.apiUrl == ServiceUrl {
		params = company
	} else {
		params = map[string]interface{}{
			"company": company,
		}
	}

	err = c.Put("companies/"+company.Id, params, nil, &response)

	return

}

func (c *Client) DeleteCompany(id string) (err error) {
	err = c.Delete("companies/"+id, nil, nil, nil)
	return
}

//func (c *Client) UploadCertificate(id, path, password string) (err error) {
//
//	var (
//		file *os.File
//		body bytes.Buffer
//		response *http.Response
//		erm  ErrMessage
//	)
//
//	// read file
//	if file, err = os.Open(path); err != nil {
//		return
//	}
//
//	// close file
//	defer file.Close()
//
//	// create form writer
//	w := multipart.NewWriter(&body)
//
//	// set password
//	p, _ := w.CreateFormField("password")
//	if _, err = p.Write([]byte(password)); err != nil {
//		return
//	}
//
//	// set file
//	f, _ := w.CreateFormFile("file", filepath.Base(file.Name()))
//	if _, err = io.Copy(f, file); err != nil {
//		return
//	}
//
//	// close writer
//	defer w.Close()
//
//	spew.Dump(body.Len())
//
//	// set request
//	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%scompanies/%s/certificate", c.apiUrl, id), &body)
//	request.Header.Add("Authorization", c.apiKey)
//	request.Header.Add("Content-Type", "multipart/form-data")
//	request.Header.Add("Content-Length", fmt.Sprint(body.Len()))
//	request.Header.Add("Accept", "*/*")
//	request.Header.Add("Accept-Encoding", "gzip, deflate")
//	request.Header.Add("Connection", "keep-alive")
//	request.Header.Add("Cache-Control", "no-cache")
//	request.Header.Add("Host", "api.nfe.io")
//
//	service := &http.Client{
//		Transport: &http.Transport{
//			TLSClientConfig: &tls.Config{
//				InsecureSkipVerify: true,
//			},
//		},
//	}
//
//	if response, err = service.Do(request); err != nil {
//		return
//	}
//
//	defer response.Body.Close()
//
//	// read response
//	data, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		return err
//	}
//
//	// check for error message
//	if err = json.Unmarshal(data, erm); err == nil && erm.Message != "" {
//		return erm
//	}
//
//	// verify status code
//	if NotIn(response.StatusCode, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
//		return errors.New(response.Status)
//	}
//
//	return nil
//
//}

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
	if len(header) < 1 || !strings.Contains(header[0], "200") {

		// get status code
		data := strings.Split(strings.Trim(header[0], " "), " ")

		// parse to int
		code, _ := strconv.Atoi(data[1])

		// return error code
		return errors.New(http.StatusText(code))

	}

	return nil

}

type CompanyResponse struct {
	Data      Company
	Company   Company `json:"company"`
	Companies Company `json:"companies"`
}

type CompaniesResponse struct {
	Data []Company `json:"companies"`
	pagination
}

type Company struct {
	Id                    string             `json:"id,omitempty"`
	Name                  string             `json:"name,omitempty"`
	TradeName             string             `json:"tradeName,omitempty"`
	Email                 string             `json:"email,omitempty"`
	FederalTaxNumber      json.Number        `json:"federalTaxNumber,omitempty"`
	RegionalTaxNumber     json.Number        `json:"regionalTaxNumber,omitempty"`
	MunicipalTaxNumber    json.Number        `json:"municipalTaxNumber,omitempty"`
	CompanyRegistryNumber json.Number        `json:"companyRegistryNumber,omitempty"`
	TaxRegime             interface{}        `json:"taxRegime,omitempty"`
	SpecialTaxRegime      interface{}        `json:"specialTaxRegime,omitempty"`
	LegalNature           interface{}        `json:"legalNature,omitempty"`
	RpsSerialNumber       interface{}        `json:"rpsSerialNumber,omitempty"`
	RpsNumber             interface{}        `json:"rpsNumber,omitempty"`
	IssRate               float64            `json:"issRate,omitempty"`
	Environment           string             `json:"environment,omitempty"`
	FiscalStatus          string             `json:"fiscalStatus,omitempty"`
	Address               *Address           `json:"address,omitempty"`
	Certificate           *Certificate       `json:"certificate,omitempty"`
	EconomicActivities    []EconomicActivity `json:"economicActivities,omitempty"`
	OpenningDate          *time.Time         `json:"openningDate,omitempty"`
	CreatedOn             *time.Time         `json:"createdOn,omitempty"`
	ModifiedOn            *time.Time         `json:"modifiedOn,omitempty"`
}

type EconomicActivity struct {
	Type string      `json:"type,omitempty"`
	Code json.Number `json:"code,omitempty"`
}

type Certificate struct {
	Thumbprint string    `json:"thumbprint,omitempty"`
	Status     string    `json:"status,omitempty"`
	ModifiedOn time.Time `json:"modifiedOn,omitempty"`
	ExpiresOn  time.Time `json:"expiresOn,omitempty"`
}

// only to avoid recursion in UnmarshalJSON below
type cr CompanyResponse

// override default json.Unmarshal
func (c *CompanyResponse) UnmarshalJSON(b []byte) (err error) {

	aux := cr{}

	if err = json.Unmarshal(b, &aux); err == nil {

		*c = CompanyResponse(aux)

		if c.Company.Id != "" {
			c.Data = c.Company
		}

		if c.Companies.Id != "" {
			c.Data = c.Companies
		}

	}

	return nil
}
