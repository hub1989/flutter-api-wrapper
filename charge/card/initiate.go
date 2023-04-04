package card

type InitiateRequest struct {
	CardNumber  string `json:"card_number"`
	Cvv         string `json:"cvv"`
	ExpiryMonth string `json:"expiry_month"`
	ExpiryYear  string `json:"expiry_year"`
	Currency    string `json:"currency"`
	Amount      string `json:"amount"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	TxRef       string `json:"tx_ref"`
	RedirectUrl string `json:"redirect_url"`
}

type InitiateResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Meta    struct {
		Authorization struct {
			Mode   string   `json:"mode"`
			Fields []string `json:"fields"`
		} `json:"authorization"`
	} `json:"meta"`
}
