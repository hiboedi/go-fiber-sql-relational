package models

type Locker struct {
	ID     int          `json:"id" form:"id" grom:"primaryKey"`
	Code   string       `json:"code" from:"code" gorm:"unique,not null"`
	UserID int          `json:"user_id" form:"user_id" gorm:"not null"`
	User   UserResponse `json:"user"`
}

type LockerResponse struct {
	ID     int    `json:"id" form:"id"`
	Code   string `json:"code" from:"code"`
	UserID int    `json:"-" form:"user_id"`
}

// convention GORM

func (LockerResponse) TableName() string {
	return "Lockers"
}
