package entity

type Product struct {
	Id          int    `json:"-" db:"id"`
	Category    string `json:"category" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Price       string `json:"price" binding:"required"`
	Description string `json:"description" binding:"required"`
	Color       string `json:"color" binding:"required"`
}
