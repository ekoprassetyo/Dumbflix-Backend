package models

import "time"

type User struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"		gorm:"type: varchar(255)"`
	Email     string          `json:"email"		gorm:"type: varchar(255)"`
	Password  string          `json:"password"	gorm:"type: varchar(255)"`
	Status    string          `json:"status" gorm:type:"varchar(255)"`
	Profile   ProfileResponse `json:"profile"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type UserProfileResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (UserProfileResponse) TableName() string {
	return "users"
}
