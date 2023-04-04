package payout

type SubAccountBalance struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Currency         string `json:"currency"`
		AvailableBalance int    `json:"available_balance"`
		LedgerBalance    int    `json:"ledger_balance"`
	} `json:"data"`
}
