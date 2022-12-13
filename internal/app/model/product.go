package model

type ProductID struct {
	ID string `json:"id,omitempty" db:"id"`
}

type Product struct {
	ProductID *ProductID `json:"productId,omitempty" db:"productId"`
	Name      string     `json:"name,omitempty" db:"name"`
}
