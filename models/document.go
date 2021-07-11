package models

type Document struct {
	DocumentId      int    `json:"document_id"`
	DocumentName    string `json:"document_name"`
	DocumentNumber  int    `json:"document_number"`
	DocumentExpDate string `json:"document_exp_date"`
	UserId          int    `json:"user_id"`
}
