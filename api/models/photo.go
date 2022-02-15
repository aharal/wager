package models

import "time"

//BuyWager Model
type BuyWager struct {
	ID           int        `gorm:"id" json:"id"`
	WagerId      int        `gorm:"album_id" json:"albumId"`
	BuyWagerId   string     `gorm:"photo_id" json:"photoId"`
	Title        string     `gorm:"title" json:"title"`
	Url          string     `gorm:"url" json:"url"`
	ThumbnailUrl string     `gorm:"thumbnail_url" json:"thumbnailUrl"`
	CreatedAt    time.Time  `gorm:"created_at" json:"-"`
	UpdatedAt    time.Time  `gorm:"updated_at" json:"-"`
	DeletedAt    *time.Time `gorm:"deleted_at" json:"-"`
}

//BuyWagerRequest Type
type BuyWagerRequest struct {
	WagerId int `gorm:"albumId" json:"albumId"`
}
