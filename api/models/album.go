package models

import "time"

//Wager Model
type Wager struct {
	ID        int        `gorm:"id" json:"id"`
	UserId    int        `gorm:"user_id" json:"userId"`
	Title     string     `gorm:"title" json:"title"`
	CreatedAt time.Time  `gorm:"created_at" json:"-"`
	UpdatedAt time.Time  `gorm:"updated_at" json:"-"`
	DeletedAt *time.Time `gorm:"deleted_at" json:"-"`
}

//WagerRequest Type
type WagerRequest struct {
	ID int `gorm:"id" json:"id"`
}
