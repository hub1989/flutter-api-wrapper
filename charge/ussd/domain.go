package ussd

import "time"

type Request struct {
	AccountBank string `json:"account_bank"`
	Amount      int    `json:"amount"`
	Currency    string `json:"currency"`
	Email       string `json:"email"`
	Fullname    string `json:"fullname"`
	TxRef       string `json:"tx_ref"`
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
		PaymentType       string    `json:"payment_type"`
		FraudStatus       string    `json:"fraud_status"`
		ChargeType        string    `json:"charge_type"`
		CreatedAt         time.Time `json:"created_at"`
		AccountId         int       `json:"account_id"`
		Customer          struct {
			Id          int         `json:"id"`
			PhoneNumber interface{} `json:"phone_number"`
			Name        string      `json:"name"`
			Email       string      `json:"email"`
			CreatedAt   time.Time   `json:"created_at"`
		} `json:"customer"`
		PaymentCode string `json:"payment_code"`
	} `json:"data"`
	Meta struct {
		Authorization struct {
			Mode string `json:"mode"`
			Note string `json:"note"`
		} `json:"authorization"`
	} `json:"meta"`
}
