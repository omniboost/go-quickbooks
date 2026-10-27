package quickbooks

import (
	"encoding/json"
	"fmt"
)

// JournalEntryObject the complete quickbooks journal entry object type
type JournalEntryObject struct {
	JournalEntry JournalEntry `json:"JournalEntry"`
	Time         string       `json:"time"`
}

// Journal Entry quickbooks Journal Entry type
type JournalEntry struct {
	ID           string        `json:"Id,omitempty"`
	Adjustment   bool          `json:"Adjustment,omitempty"`
	Domain       string        `json:"domain,omitempty"`
	Sparse       bool          `json:"sparse,omitempty"`
	SyncToken    string        `json:"SyncToken,omitempty"`
	DocNumber    string        `json:"DocNumber,omitempty"`
	TxnDate      string        `json:"TxnDate,omitempty"`
	Line         []Line        `json:"Line"`
	TxnTaxDetail *TxnTaxDetail `json:"TxnTaxDetail,omitempty"`
	MetaData     *MetaData     `json:"MetaData,omitempty"`
}

// Line type - part of Journal Entry
type Line struct {
	LineID                 string                  `json:"Id,omitempty"`
	Description            string                  `json:"Description,omitempty"`
	Amount                 float64                 `json:"Amount,omitempty"`
	DetailType             string                  `json:"DetailType,omitempty"`
	JournalEntryLineDetail *JournalEntryLineDetail `json:"JournalEntryLineDetail,omitempty"`
	LineNum                *int                    `json:"LineNum,omitempty"`
}

// JournalEntryLineDetail - part of Journal Entry
type JournalEntryLineDetail struct {
	AccountRef      JournalEntryRef `json:"AccountRef,omitempty"`
	PostingType     string          `json:"PostingType,omitempty"`
	TaxCodeRef      *TaxCodeRef     `json:"TaxCodeRef,omitempty"`
	TaxApplicableOn *string         `json:"TaxApplicableOn,omitempty"`
	TaxAmount       *float64        `json:"TaxAmount,omitempty"`
	ClassRef        *ClassRef       `json:"ClassRef,omitempty"`
	DepartmentRef   *DepartmentRef  `json:"DepartmentRef,omitempty"`
	Entity          *Entity         `json:"Entity,omitempty"`
	JournalCodeRef  JournalCodeRef  `json:"JournalCodeRef,omitempty"`
}

// Metadata - info about when the journal entry was created/updated.
type MetaData struct {
	CreateTime      string `json:"CreateTime,omitempty"`
	LastUpdatedTime string `json:"LastUpdatedTime,omitempty"`
}

type JournalCodeRef struct {
	Value string `json:"value,omitempty"`
	Name  string `json:"name,omitempty"`
}

type JournalEntryRef struct {
	Value string `json:"value,omitempty"`
	Name  string `json:"name,omitempty"`
}

type ClassRef struct {
	Value string `json:"value,omitempty"`
	Name  string `json:"name,omitempty"`
}

type DepartmentRef struct {
	Value string `json:"value,omitempty"`
	Name  string `json:"name,omitempty"`
}

type Entity struct {
	Type      string    `json:"Type"`
	EntityRef EntityRef `json:"EntityRef"`
}

type EntityRef struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TxnTaxDetail struct {
	TotalTax *float64  `json:"TotalTax,omitempty"`
	TaxLine  []TaxLine `json:"TaxLine,omitempty"`
}

// CreateJE creates a journal entry on quickbooks
func (q *Quickbooks) CreateJournalEntry(journalentry JournalEntry) (*JournalEntryObject, error) {
	endpoint := fmt.Sprintf("/company/%s/journalentry", q.RealmID)

	response, err := q.makePostRequest(endpoint, journalentry)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	newJE := JournalEntryObject{}
	err = json.NewDecoder(response.Body).Decode(&newJE)
	if err != nil {
		return nil, err
	}

	return &newJE, nil
}
