package entity

type Order struct {
	Id     int `json:"id" binding:"required"`
	CartId int `json:"cartId" binding:"required"`
}
