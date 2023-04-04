package paypal

import "time"

type Request struct {
	TxRef             string `json:"tx_ref"`
	Amount            string `json:"amount"`
	Currency          string `json:"currency"`
	Country           string `json:"country"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phone_number"`
	Fullname          string `json:"fullname"`
	ClientIp          string `json:"client_ip"`
	RedirectUrl       string `json:"redirect_url"`
	DeviceFingerprint string `json:"device_fingerprint"`
	BillingAddress    string `json:"billing_address"`
	BillingCity       string `json:"billing_city"`
	BillingZip        string `json:"billing_zip"`
	BillingState      string `json:"billing_state"`
	BillingCountry    string `json:"billing_country"`
	ShippingName      string `json:"shipping_name"`
	ShippingAddress   string `json:"shipping_address"`
	ShippingCity      string `json:"shipping_city"`
	ShippingZip       string `json:"shipping_zip"`
	ShippingState     string `json:"shipping_state"`
	ShippingCountry   string `json:"shipping_country"`
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
		ChargedAmount     int       `json:"charged_amount"`
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
