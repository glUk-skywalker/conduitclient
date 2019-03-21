package objects

// Transaction is the struct for maniphest transaction data
type Transaction struct {
	TaskID          string `json:"taskID"`
	TransactionID   string `json:"transactionID"`
	TransactionPHID string `json:"transactionPHID"`
	TransactionType string `json:"transactionType"`
	DateCreated     string `json:"dateCreated"`
}
