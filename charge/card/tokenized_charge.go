package card

import "time"

type TokenizedChargeRequest struct {
	Token     string `json:"token"`
	Currency  string `json:"currency"`
	Country   string `json:"country"`
	Amount    int    `json:"amount"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Ip        string `json:"ip"`
	Narration string `json:"narration"`
	TxRef     string `json:"tx_ref"`
}

type TokenizedChargeResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Id                int       `json:"id"`
		TxRef             string    `json:"tx_ref"`
		FlwRef            string    `json:"flw_ref"`
		RedirectUrl       string    `json:"redirect_url"`
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
		PaymentType       string    `json:"payment_type"`
		CreatedAt         time.Time `json:"created_at"`
		AccountId         int       `json:"account_id"`
		Customer          struct {
			Id          int         `json:"id"`
			PhoneNumber interface{} `json:"phone_number"`
			Name        string      `json:"name"`
			Email       string      `json:"email"`
			CreatedAt   time.Time   `json:"created_at"`
		} `json:"customer"`
		Card struct {
			First6Digits string `json:"first_6digits"`
			Last4Digits  string `json:"last_4digits"`
			Issuer       string `json:"issuer"`
			Country      string `json:"country"`
			Type         string `json:"type"`
			Expiry       string `json:"expiry"`
			Token        string `json:"token"`
		} `json:"card"`
	} `json:"data"`
}
