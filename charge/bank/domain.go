package bank

type TransferRequest struct {
	TxRef             string `json:"tx_ref"`
	Amount            string `json:"amount"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phone_number"`
	Currency          string `json:"currency"`
	ClientIp          string `json:"client_ip"`
	DeviceFingerprint string `json:"device_fingerprint"`
	Narration         string `json:"narration"`
	IsPermanent       bool   `json:"is_permanent"`
}

type TransferResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Meta    struct {
		Authorization struct {
			TransferReference string `json:"transfer_reference"`
			TransferAccount   string `json:"transfer_account"`
			TransferBank      string `json:"transfer_bank"`
			AccountExpiration int    `json:"account_expiration"`
			TransferNote      string `json:"transfer_note"`
			TransferAmount    string `json:"transfer_amount"`
			Mode              string `json:"mode"`
		} `json:"authorization"`
	} `json:"meta"`
}
