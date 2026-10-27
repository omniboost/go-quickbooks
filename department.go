package quickbooks

// Department is what the QuickBooks UI calls a "Location" (Settings > Advanced >
// Categories > Track locations). The label is configurable per company, but the
// API entity is always Department.
type Department struct {
	Name               string `json:"Name"`
	SubDepartment      bool   `json:"SubDepartment"`
	FullyQualifiedName string `json:"FullyQualifiedName"`
	Active             bool   `json:"Active"`
	Domain             string `json:"domain"`
	Sparse             bool   `json:"sparse"`
	ID                 string `json:"Id"`
	SyncToken          string `json:"SyncToken"`
	MetaData           struct {
		CreateTime      string `json:"CreateTime"`
		LastUpdatedTime string `json:"LastUpdatedTime"`
	} `json:"MetaData"`
}
