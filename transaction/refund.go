package transaction

import "time"

type RefundResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Id             int    `json:"id"`
		AccountId      int    `json:"account_id"`
		TxId           int    `json:"tx_id"`
		FlwRef         string `json:"flw_ref"`
		WalletId       int    `json:"wallet_id"`
		AmountRefunded int    `json:"amount_refunded"`
		Status         string `json:"status"`
		Destination    string `json:"destination"`
		Meta           struct {
			Source string `json:"source"`
		} `json:"meta"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"data"`
}
