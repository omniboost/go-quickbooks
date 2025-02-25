package quickbooks

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// QueryResponse quickbooks search response type
type QueryResponse map[string]interface{}

// Search quickbooks document with a query string
func (q *Quickbooks) Search(query string) (*QueryResponse, error) {
	endpoint := fmt.Sprintf("/company/%s/query?query=%s", q.RealmID, url.QueryEscape(query))

	res, err := q.makePostRequest(endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	queryResponse := QueryResponse{}
	err = json.NewDecoder(res.Body).Decode(&queryResponse)
	if err != nil {
		return nil, err
	}

	return &queryResponse, nil
}

// ToAccount converts a search QueryRespose to Account array type
func (qr QueryResponse) ToAccount() ([]Account, error) {
	queryResponse := qr["QueryResponse"]
	document := queryResponse.(map[string]interface{})

	b, err := json.Marshal(document["Account"])
	if err != nil {
		return nil, err
	}

	accounts := []Account{}
	err = json.Unmarshal(b, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// ToItem converts a search QueryRespose to Item array type
func (qr QueryResponse) ToItem() ([]Item, error) {
	queryResponse := qr["QueryResponse"]
	document := queryResponse.(map[string]interface{})

	b, err := json.Marshal(document["Item"])
	if err != nil {
		return nil, err
	}

	items := []Item{}
	err = json.Unmarshal(b, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

// ToCustomer converts a search QueryRespose to Customer array type
func (qr QueryResponse) ToCustomer() ([]Customer, error) {
	queryResponse := qr["QueryResponse"]
	document := queryResponse.(map[string]interface{})

	b, err := json.Marshal(document["Customer"])
	if err != nil {
		return nil, err
	}

	customers := []Customer{}
	err = json.Unmarshal(b, &customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

// ToInvoice converts a search QueryRespose to Invoice array type
func (qr QueryResponse) ToInvoice() ([]Invoice, error) {
	queryResponse := qr["QueryResponse"]
	document := queryResponse.(map[string]interface{})

	b, err := json.Marshal(document["Invoice"])
	if err != nil {
		return nil, err
	}

	invoices := []Invoice{}
	err = json.Unmarshal(b, &invoices)
	if err != nil {
		return nil, err
	}

	return invoices, nil
}

// ToPayment converts a search QueryRespose to Payment array type
func (qr QueryResponse) ToPayment() ([]Payment, error) {
	queryResponse := qr["QueryResponse"]
	document := queryResponse.(map[string]interface{})

	b, err := json.Marshal(document["Payment"])
	if err != nil {
		return nil, err
	}

	payments := []Payment{}
	err = json.Unmarshal(b, &payments)
	if err != nil {
		return nil, err
	}

	return payments, nil
}

// ToTaxCode converts a search QueryRespose to TaxCode array type
func (qr QueryResponse) ToTaxCode() ([]TaxCode, error) {
	queryResponse := qr["QueryResponse"]
	document := queryResponse.(map[string]interface{})

	b, err := json.Marshal(document["TaxCode"])
	if err != nil {
		return nil, err
	}

	taxCodes := []TaxCode{}
	err = json.Unmarshal(b, &taxCodes)
	if err != nil {
		return nil, err
	}

	return taxCodes, nil
}

// ToClass converts a search QueryRespose to Class array type
func (qr QueryResponse) ToClass() ([]Class, error) {
	queryResponse := qr["QueryResponse"]
	document := queryResponse.(map[string]interface{})

	b, err := json.Marshal(document["Class"])
	if err != nil {
		return nil, err
	}

	classes := []Class{}
	err = json.Unmarshal(b, &classes)
	if err != nil {
		return nil, err
	}

	return classes, nil
}

// ToJournalCode converts a search QueryRespose to JournalCode array type
func (qr QueryResponse) ToJournalCode() ([]JournalCode, error) {
	queryResponse := qr["QueryResponse"]
	document := queryResponse.(map[string]interface{})

	b, err := json.Marshal(document["JournalCode"])
	if err != nil {
		return nil, err
	}

	journalCodes := []JournalCode{}
	err = json.Unmarshal(b, &journalCodes)
	if err != nil {
		return nil, err
	}

	return journalCodes, nil
}

// ToCompanyInfo converts a search QueryRespose to JournalCode array type
func (qr QueryResponse) ToCompanyInfo() ([]CompanyInfo, error) {
	queryResponse := qr["QueryResponse"]
	document := queryResponse.(map[string]interface{})

	b, err := json.Marshal(document["CompanyInfo"])
	if err != nil {
		return nil, err
	}

	companyInfos := []CompanyInfo{}
	err = json.Unmarshal(b, &companyInfos)
	if err != nil {
		return nil, err
	}

	return companyInfos, nil
}

// ToVendor converts a search QueryRespose to Vendor array type
func (qr QueryResponse) ToVendor() ([]Vendor, error) {
	queryResponse := qr["QueryResponse"]
	document := queryResponse.(map[string]interface{})

	b, err := json.Marshal(document["Vendor"])
	if err != nil {
		return nil, err
	}

	vendors := []Vendor{}
	err = json.Unmarshal(b, &vendors)
	if err != nil {
		return nil, err
	}

	return vendors, nil
}

// ToCreditMemo converts a search QueryRespose to CreditMemo array type
func (qr QueryResponse) ToCreditMemo() ([]CreditMemo, error) {
	queryResponse := qr["QueryResponse"]
	document := queryResponse.(map[string]interface{})

	b, err := json.Marshal(document["CreditMemo"])
	if err != nil {
		return nil, err
	}

	invoices := []CreditMemo{}
	err = json.Unmarshal(b, &invoices)
	if err != nil {
		return nil, err
	}

	return invoices, nil
}

// ToJournalEntry converts a search QueryRespose to Invoice array type
func (qr QueryResponse) ToJournalEntry() ([]JournalEntry, error) {
	queryResponse := qr["QueryResponse"]
	document := queryResponse.(map[string]interface{})

	b, err := json.Marshal(document["JournalEntry"])
	if err != nil {
		return nil, err
	}

	journalEntries := []JournalEntry{}
	err = json.Unmarshal(b, &journalEntries)
	if err != nil {
		return nil, err
	}

	return journalEntries, nil
}
