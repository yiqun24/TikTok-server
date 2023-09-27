package models

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	ID            uint64    `gorm:"primary_key;auto_increment"`
	Author        User      `gorm:"foreignkey:AuthorId"`
	AuthorId      uint64    `gorm:"not null;index:idx_author_id"`
	PlayUrl       string    `gorm:"type:varchar(255);not null;default:''"`
	CoverUrl      string    `gorm:"type:varchar(255);not null;default:''"`
	Title         string    `gorm:"type:varchar(255);not null;default:''"`
	FavoriteCount uint64    `gorm:"type:int;not null;default:0"`
	CommentCount  uint64    `gorm:"type:int;not null;default:0"`
	CreatedAt     time.Time `gorm:"not null;index:idx_create_at;sort:desc;"`
}
