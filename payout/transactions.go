package payout

type SubAccountTransactionsResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Transactions []interface{} `json:"transactions"`
		PageInfo     struct {
			Total       int `json:"total"`
			CurrentPage int `json:"current_page"`
			TotalPages  int `json:"total_pages"`
		} `json:"page_info"`
		Message string `json:"message"`
	} `json:"data"`
}
