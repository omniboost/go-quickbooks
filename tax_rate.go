package quickbooks

type TaxRate struct {
	ID          string  `json:"Id"`
	SyncToken   string  `json:"SyncToken"`
	Name        string  `json:"Name"`
	Description string  `json:"Description"`
	RateValue   float64 `json:"RateValue"`
	Active      bool    `json:"Active"`
	Domain      string  `json:"domain"`
	Sparse      bool    `json:"sparse"`
	MetaData    struct {
		CreateTime      string `json:"CreateTime"`
		LastUpdatedTime string `json:"LastUpdatedTime"`
	} `json:"MetaData"`
	AgencyRef        *struct {
		Value string `json:"value"`
		Name  string `json:"name,omitempty"`
	} `json:"AgencyRef,omitempty"`
	TaxReturnLineRef *struct {
		Value string `json:"value"`
		Name  string `json:"name,omitempty"`
	} `json:"TaxReturnLineRef,omitempty"`
	SpecialTaxType    string              `json:"SpecialTaxType,omitempty"`
	DisplayType       string              `json:"DisplayType,omitempty"`
	EffectiveTaxRate  []EffectiveTaxRate  `json:"EffectiveTaxRate,omitempty"`
}

type EffectiveTaxRate struct {
	RateValue        float64 `json:"RateValue"`
	EffectiveDate    string  `json:"EffectiveDate,omitempty"`
	EndDate          string  `json:"EndDate,omitempty"`
}
