package model

// Product struct
type (
	Product struct {
		Name        string `json:"name" validate:"required min=2 max=32"`
		Description string `json:"description"`
		Category    string `json:"category"`
		Amount      int    `json:"amount"`
	}
)

// Products struct
type Products struct {
	Products []Product `json:"products"`
}
