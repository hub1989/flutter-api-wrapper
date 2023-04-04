package apple_pay

import "time"

type Request struct {
	TxRef             string `json:"tx_ref"`
	Amount            string `json:"amount"`
	Currency          string `json:"currency"`
	Email             string `json:"email"`
	Fullname          string `json:"fullname"`
	Narration         string `json:"narration"`
	RedirectUrl       string `json:"redirect_url"`
	ClientIp          string `json:"client_ip"`
	DeviceFingerprint string `json:"device_fingerprint"`
	BillingZip        string `json:"billing_zip"`
	BillingCity       string `json:"billing_city"`
	BillingAddress    string `json:"billing_address"`
	BillingState      string `json:"billing_state"`
	BillingCountry    string `json:"billing_country"`
	PhoneNumber       string `json:"phone_number"`
	Meta              struct {
		Metaname  string `json:"metaname"`
		Metavalue string `json:"metavalue"`
	} `json:"meta"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Id                int       `json:"id"`
		TxRef             string    `json:"tx_ref"`
		FlwRef            string    `json:"flw_ref"`
		DeviceFingerprint string    `json:"device_fingerprint"`
		Amount            int       `json:"amount"`
		ChargedAmount     float64   `json:"charged_amount"`
		AppFee            float64   `json:"app_fee"`
		MerchantFee       int       `json:"merchant_fee"`
		ProcessorResponse string    `json:"processor_response"`
		AuthModel         string    `json:"auth_model"`
		Currency          string    `json:"currency"`
		Ip                string    `json:"ip"`
		Narration         string    `json:"narration"`
		Status            string    `json:"status"`
		AuthUrl           string    `json:"auth_url"`
		PaymentType       string    `json:"payment_type"`
		FraudStatus       string    `json:"fraud_status"`
		ChargeType        string    `json:"charge_type"`
		CreatedAt         time.Time `json:"created_at"`
		AccountId         int       `json:"account_id"`
		Customer          struct {
			Id          int       `json:"id"`
			PhoneNumber string    `json:"phone_number"`
			Name        string    `json:"name"`
			Email       string    `json:"email"`
			CreatedAt   time.Time `json:"created_at"`
		} `json:"customer"`
		Meta struct {
			Authorization struct {
				Mode     string `json:"mode"`
				Redirect string `json:"redirect"`
			} `json:"authorization"`
		} `json:"meta"`
	} `json:"data"`
}
