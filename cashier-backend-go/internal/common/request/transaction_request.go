package request

type CreateTransactionRequest struct {
	CustomerID         uint                       `json:"customer_id" binding:"required"`
	TransactionDetails []TransactionDetailRequest `json:"transaction_details" binding:"required"`
}

type TransactionDetailRequest struct {
	MenuID   uint    `json:"menu_id" binding:"required"`
	Quantity int     `json:"quantity" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Total    float64 `json:"total" binding:"required"`
}
