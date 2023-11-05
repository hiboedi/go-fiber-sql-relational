package models

type Tag struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}
