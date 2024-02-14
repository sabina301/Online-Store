package entity

type Cart struct {
	Id       int     `json:"-" db:"id"`
	User     User    `json:"user" binding:"required"`
	Products Product `json:"product" binding:"required"`
}
