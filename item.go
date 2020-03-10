package nfeio

import (
	"encoding/json"
	"time"
)

type Item struct {
	Code                  string              `json:"code,omitempty"`
	Description           string              `json:"description,omitempty"`
	NCM                   string              `json:"ncm,omitempty"`
	CFOP                  json.Number         `json:"cfop,omitempty"`
	CodeGTIN              string              `json:"codeGTIN,omitempty"`
	CodeTaxGTIN           string              `json:"codeTaxGTIN,omitempty"`
	Unit                  string              `json:"unit,omitempty"`
	Quantity              json.Number         `json:"quantity,omitempty"`
	UnitAmount            json.Number         `json:"unitAmount,omitempty"`
	TotalAmount           json.Number         `json:"totalAmount,omitempty"`
	UnitTax               string              `json:"unitTax,omitempty"`
	QuantityTax           json.Number         `json:"quantityTax,omitempty"`
	TaxUnitAmount         json.Number         `json:"taxUnitAmount,omitempty"`
	FreightAmount         json.Number         `json:"freightAmount,omitempty"`
	InsuranceAmount       json.Number         `json:"insuranceAmount,omitempty"`
	DiscountAmount        json.Number         `json:"discountAmount,omitempty"`
	OthersAmount          json.Number         `json:"othersAmount,omitempty"`
	TotalIndicator        bool                `json:"totalIndicator,omitempty"`
	Cest                  string              `json:"cest,omitempty,omitempty"`
	AdditionalInformation string              `json:"additionalInformation,omitempty"`
	NumberOrderBuy        string              `json:"numberOrderBuy,omitempty"`
	ItemNumberOrderBuy    json.Number         `json:"itemNumberOrderBuy,omitempty"`
	Benefit               string              `json:"benefit,omitempty"`
	Tax                   *Tax                `json:"tax,omitempty"`
	FuelDetail            *FuelDetail         `json:"fuelDetail,omitempty"`
	ImportDeclarations    []ImportDeclaration `json:"importDeclarations,omitempty"`
}

type Tax struct {
	TotalTax        json.Number      `json:"totalTax,omitempty"`
	ICMS            *ICMS            `json:"icms,omitempty"`
	IPI             *IPI             `json:"ipi,omitempty"`
	II              *II              `json:"ii,omitempty"`
	PIS             *PIS             `json:"pis,omitempty"`
	COFINS          *COFINS          `json:"cofins,omitempty"`
	ICMSDestination *ICMSDestination `json:"icmsDestination,omitempty"`
}

type ICMS struct {
	Origin                     string      `json:"origin,omitempty"`
	CST                        string      `json:"cst,omitempty"`
	BaseTaxModality            string      `json:"baseTaxModality,omitempty"`
	BaseTax                    json.Number `json:"baseTax,omitempty"`
	BaseTaxSTModality          string      `json:"baseTaxSTModality,omitempty"`
	BaseTaxSTReduction         string      `json:"baseTaxSTReduction,omitempty"`
	BaseTaxST                  json.Number `json:"baseTaxST,omitempty"`
	BaseTaxReduction           json.Number `json:"baseTaxReduction,omitempty"`
	StRate                     json.Number `json:"stRate,omitempty"`
	StAmount                   json.Number `json:"stAmount,omitempty"`
	StMarginAmount             json.Number `json:"stMarginAmount,omitempty"`
	CSOSN                      string      `json:"csosn,omitempty"`
	Rate                       json.Number `json:"rate,omitempty"`
	Amount                     json.Number `json:"amount,omitempty"`
	Percentual                 json.Number `json:"percentual,omitempty"`
	SnCreditRate               json.Number `json:"snCreditRate,omitempty"`
	SnCreditAmount             json.Number `json:"snCreditAmount,omitempty"`
	StMarginAddedAmount        string      `json:"stMarginAddedAmount,omitempty"`
	StRetentionAmount          string      `json:"stRetentionAmount,omitempty"`
	BaseSTRetentionAmount      string      `json:"baseSTRetentionAmount,omitempty"`
	BaseTaxOperationPercentual string      `json:"baseTaxOperationPercentual,omitempty"`
	UFST                       string      `json:"ufst,omitempty"`
	AmountSTReason             string      `json:"amountSTReason,omitempty"`
	BaseSNRetentionAmount      string      `json:"baseSNRetentionAmount,omitempty"`
	SnRetentionAmount          string      `json:"snRetentionAmount,omitempty"`
	AmountOperation            string      `json:"amountOperation,omitempty"`
	PercentualDeferment        string      `json:"percentualDeferment,omitempty"`
	BaseDeferred               string      `json:"baseDeferred,omitempty"`
	ExemptAmount               json.Number `json:"exemptAmount,omitempty"`
	ExemptReason               string      `json:"exemptReason,omitempty"`
	FcpRate                    json.Number `json:"fcpRate,omitempty"`
	FcpAmount                  json.Number `json:"fcpAmount,omitempty"`
	FcpstRate                  json.Number `json:"fcpstRate,omitempty"`
	FcpstAmount                json.Number `json:"fcpstAmount,omitempty"`
	FcpstRetRate               json.Number `json:"fcpstRetRate,omitempty"`
	FcpstRetAmount             json.Number `json:"fcpstRetAmount,omitempty"`
	BaseTaxFCPSTAmount         json.Number `json:"baseTaxFCPSTAmount,omitempty"`
	SubstituteAmount           json.Number `json:"substituteAmount,omitempty"`
}

type IPI struct {
	Classification     string      `json:"classification,omitempty"`
	ProducerCNPJ       string      `json:"producerCNPJ,omitempty"`
	StampCode          string      `json:"stampCode,omitempty"`
	StampQuantity      json.Number `json:"stampQuantity,omitempty"`
	ClassificationCode string      `json:"classificationCode,omitempty"`
	CST                string      `json:"cst,omitempty"`
	Base               json.Number `json:"base,omitempty"`
	Rate               json.Number `json:"rate,omitempty"`
	UnitQuantity       json.Number `json:"unitQuantity,omitempty"`
	UnitAmount         json.Number `json:"unitAmount,omitempty"`
	Amount             json.Number `json:"amount,omitempty"`
}

type II struct {
	BaseTax                  string      `json:"baseTax,omitempty"`
	CustomsExpenditureAmount string      `json:"customsExpenditureAmount,omitempty"`
	Amount                   json.Number `json:"amount,omitempty"`
	IofAmount                json.Number `json:"iofAmount,omitempty"`
}

type PIS struct {
	CST                    string      `json:"cst,omitempty"`
	BaseTax                json.Number `json:"baseTax,omitempty"`
	Rate                   json.Number `json:"rate,omitempty"`
	Amount                 json.Number `json:"amount,omitempty"`
	BaseTaxProductQuantity json.Number `json:"baseTaxProductQuantity,omitempty"`
	ProductRate            json.Number `json:"productRate,omitempty"`
}

type COFINS struct {
	CST                    string      `json:"cst,omitempty"`
	BaseTax                json.Number `json:"baseTax,omitempty"`
	Rate                   json.Number `json:"rate,omitempty"`
	Amount                 json.Number `json:"amount,omitempty"`
	BaseTaxProductQuantity json.Number `json:"baseTaxProductQuantity,omitempty"`
	ProductRate            json.Number `json:"productRate,omitempty"`
}

type ICMSDestination struct {
	VBCUFDest      json.Number `json:"vBCUFDest,omitempty"`
	PFCPUFDest     json.Number `json:"pFCPUFDest,omitempty"`
	PICMSUFDest    json.Number `json:"pICMSUFDest,omitempty"`
	PICMSInter     json.Number `json:"pICMSInter,omitempty"`
	PICMSInterPart json.Number `json:"pICMSInterPart,omitempty"`
	VFCPUFDest     json.Number `json:"vFCPUFDest,omitempty"`
	VICMSUFDest    json.Number `json:"vICMSUFDest,omitempty"`
	VICMSUFRemet   json.Number `json:"vICMSUFRemet,omitempty"`
	VBCFCPUFDest   json.Number `json:"vBCFCPUFDest,omitempty"`
}

type FuelDetail struct {
	CodeANP        string      `json:"codeANP,omitempty"`
	PercentageNG   json.Number `json:"percentageNG,omitempty"`
	DescriptionANP string      `json:"descriptionANP,omitempty"`
	PercentageGLP  json.Number `json:"percentageGLP,omitempty"`
	PercentageNGn  json.Number `json:"percentageNGn,omitempty"`
	PercentageGNi  json.Number `json:"percentageGNi,omitempty"`
	StartingAmount json.Number `json:"startingAmount,omitempty"`
	Codif          string      `json:"codif,omitempty"`
	AmountTemp     json.Number `json:"amountTemp,omitempty"`
	StateBuyer     string      `json:"stateBuyer,omitempty"`
	Cide           *Cide       `json:"cide,omitempty"`
	Pump           *Pump       `json:"pump,omitempty"`
}

type Cide struct {
	Bc         json.Number `json:"bc,omitempty"`
	Rate       json.Number `json:"rate,omitempty"`
	CideAmount json.Number `json:"cideAmount,omitempty"`
}

type Pump struct {
	SpoutNumber     json.Number `json:"spoutNumber,omitempty"`
	Number          json.Number `json:"number,omitempty"`
	TankNumber      json.Number `json:"tankNumber,omitempty"`
	BeginningAmount json.Number `json:"beginningAmount,omitempty"`
	EndAmount       json.Number `json:"endAmount,omitempty"`
}

type ImportDeclaration struct {
	Code                   string     `json:"code,omitempty"`
	CustomsClearanceName   string     `json:"customsClearanceName,omitempty"`
	CustomsClearanceState  string     `json:"customsClearanceState,omitempty"`
	Exporter               string     `json:"exporter,omitempty"`
	InternationalTransport string     `json:"internationalTransport,omitempty"`
	Intermediation         string     `json:"intermediation,omitempty"`
	RegisteredOn           *time.Time `json:"registeredOn,omitempty"`
	CustomsClearancedOn    *time.Time `json:"customsClearancedOn,omitempty"`
	Additions              []Addition `json:"additions,omitempty"`
}

type Addition struct {
	Code         json.Number `json:"code,omitempty"`
	Manufacturer string      `json:"manufacturer,omitempty"`
	Amount       json.Number `json:"amount,omitempty"`
	Drawback     json.Number `json:"drawback,omitempty"`
}
