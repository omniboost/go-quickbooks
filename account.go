package quickbooks

import (
	"encoding/json"
	"fmt"
)

// AccountObject the complete quickbooks account object type
type AccountObject struct {
	Account Account `json:"Account"`
	Time    string  `json:"time"`
}

// Account quickbooks account type
type Account struct {
	Name                          string       `json:"Name"`
	SubAccount                    bool         `json:"SubAccount,omitempty"`
	FullyQualifiedName            string       `json:"FullyQualifiedName,omitempty"`
	Active                        bool         `json:"Active,omitempty"`
	Classification                string       `json:"Classification,omitempty"`
	AccountType                   string       `json:"AccountType"`
	AccountSubType                string       `json:"AccountSubType,omitempty"`
	AcctNum                       string       `json:"AcctNum,omitempty"`
	CurrentBalance                float64      `json:"CurrentBalance,omitempty"`
	CurrentBalanceWithSubAccounts float64      `json:"CurrentBalanceWithSubAccounts,omitempty"`
	CurrencyRef                   *CurrencyRef `json:"CurrencyRef,omitempty"`
	Domain                        string       `json:"domain,omitempty"`
	Sparse                        bool         `json:"sparse,omitempty"`
	ID                            string       `json:"Id,omitempty"`
	SyncToken                     string       `json:"SyncToken,omitempty"`
	MetaData                      *struct {
		CreateTime      string `json:"CreateTime,omitempty"`
		LastUpdatedTime string `json:"LastUpdatedTime,omitempty"`
	} `json:"MetaData,omitempty"`
}

// AccountRef chart of account reference
type AccountRef struct {
	Value string `json:"value"`
	Name  string `json:"name,omitempty"`
}

// CurrencyRef chart of account currency reference
type CurrencyRef struct {
	Value string `json:"value"`
	Name  string `json:"name,omitempty"`
}

// CreateAccount creates a chart of account on quickbooks
func (q *Quickbooks) CreateAccount(account Account) (*AccountObject, error) {
	endpoint := fmt.Sprintf("/company/%s/account", q.RealmID)

	res, err := q.makePostRequest(endpoint, account)
	if err != nil {
		return nil, err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	newAccount := AccountObject{}
	err = json.NewDecoder(res.Body).Decode(&newAccount)
	if err != nil {
		return nil, err
	}

	return &newAccount, nil
}
