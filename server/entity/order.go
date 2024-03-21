package entity

type Order struct {
	Id     int `json:"-" db:"id"`
	CartId int `json:"cartId" binding:"required"`
}
