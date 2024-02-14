package entity

type Company struct {
	Id   int    `json:"-" db:"id"`
	Name string `json:"name" binding:"required"`
	City string `json:"city" binding:"required"`
}
