package request

type CreateMenuRequest struct {
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Description string  `json:"description"`
	Category    string  `json:"category" binding:"required"`
	Active      bool    `json:"active"`
}

type UpdateMenuRequest struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Active      bool    `json:"active"`
}
