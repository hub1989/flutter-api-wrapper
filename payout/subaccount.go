package payout

import "time"

type SubAccountsResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    []SubAccount `json:"data"`
}

type SubAccountResponse struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Data    SubAccount `json:"data"`
}

type SubAccount struct {
	Id               int       `json:"id"`
	AccountReference string    `json:"account_reference"`
	AccountName      string    `json:"account_name"`
	BarterId         string    `json:"barter_id"`
	Email            string    `json:"email"`
	Mobilenumber     string    `json:"mobilenumber"`
	Country          string    `json:"country"`
	Nuban            string    `json:"nuban"`
	BankName         string    `json:"bank_name"`
	BankCode         string    `json:"bank_code"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
}

type CreateSubAccountRequest struct {
	AccountName  string `json:"account_name"`
	Email        string `json:"email"`
	Mobilenumber string `json:"mobilenumber"`
	Country      string `json:"country"`
}
