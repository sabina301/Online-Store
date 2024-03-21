package entity

type Cart struct {
	Id     int `json:"-" db:"id"`
	UserId int `json:"userId" binding:"required"`
}
