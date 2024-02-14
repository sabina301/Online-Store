package entity

type Product struct {
	Id      int     `json:"-" db:"id"`
	Name    string  `json:"name" binding:"required"`
	Company Company `json:"company" binding:"required"`
	Price   uint32  `json:"price" binding:"required"`
}
