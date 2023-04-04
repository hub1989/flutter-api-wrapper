package card

import "time"

type ValidateChargeRequest struct {
	Otp    string `json:"otp"`
	FlwRef string `json:"flw_ref"`
}

type ValidateChargeResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Id                int         `json:"id"`
		TxRef             string      `json:"tx_ref"`
		FlwRef            string      `json:"flw_ref"`
		DeviceFingerprint string      `json:"device_fingerprint"`
		Amount            int         `json:"amount"`
		ChargedAmount     int         `json:"charged_amount"`
		AppFee            float64     `json:"app_fee"`
		MerchantFee       int         `json:"merchant_fee"`
		ProcessorResponse string      `json:"processor_response"`
		AuthModel         string      `json:"auth_model"`
		Currency          string      `json:"currency"`
		Ip                string      `json:"ip"`
		Narration         string      `json:"narration"`
		Status            string      `json:"status"`
		AuthUrl           string      `json:"auth_url"`
		PaymentType       string      `json:"payment_type"`
		Plan              interface{} `json:"plan"`
		FraudStatus       string      `json:"fraud_status"`
		ChargeType        string      `json:"charge_type"`
		CreatedAt         time.Time   `json:"created_at"`
		AccountId         int         `json:"account_id"`
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
		} `json:"card"`
	} `json:"data"`
}
