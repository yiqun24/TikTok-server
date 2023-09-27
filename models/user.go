package models

// gorm model
type User struct {
	ID             uint64 `gorm:"primary_key"`
	Username       string `gorm:"type:varchar(30);not null;unique"`
	FollowCount    uint64 `gorm:"type:int;not null;default:0"`
	FollowerCount  uint64 `gorm:"type:int;not null;default:0"`
	Avatar         string `gorm:"type:varchar(255);not null;default:''"`
	BackgroudImage string `gorm:"type:varchar(255);not null;default:''"`
	Signture       string `gorm:"type:varchar(255);not null;default:''"`
	TotalFavorited uint64 `gorm:"type:int;not null;default:0"`
	WorkCount      uint64 `gorm:"type:int;not null;default:0"`
	FavoriteCount  uint64 `gorm:"type:int;not null;default:0"`
}
