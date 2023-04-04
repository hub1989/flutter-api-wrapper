package payout

type StaticAccountResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		StaticAccount string `json:"static_account"`
		BankName      string `json:"bank_name"`
		BankCode      string `json:"bank_code"`
		Currency      string `json:"currency"`
	} `json:"data"`
}
