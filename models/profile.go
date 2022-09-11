package models

import (
	"time"
)

type Profile struct {
	ID        int                 `json:"id" 		gorm:"primary_key:auto_increment"`
	Phone     string              `json:"phone"		gorm:"type: varchar(255)"`
	Gender    string              `json:"gender"	gorm:"type: varchar(255)"`
	Address   string              `json:"address"	gorm:"type: varchar(255)"`
	Subscribe bool                `json:"subscribe" gorm:"type: boolean"`
	UserID    int                 `json:"user_id"`
	User      UserProfileResponse `json:"user"		 gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time           `json:"-"`
	UpdatedAt time.Time           `json:"-"`
}

// hadiah buat user atau apa yang dikirim profile untuk user
type ProfileResponse struct {
	Phone   string `json:"phone"`
	Gender  string `json:"gender"`
	Address string `json:"address"`
	UserID  int    `json:"user_id"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
