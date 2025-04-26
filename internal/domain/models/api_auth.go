package models

import "gorm.io/gorm"

type ApiAuth struct {
	gorm.Model
	UserID       uint
	User         User `gorm:"constraint:OnDelete:SET NULL;"`
	ClientID     string
	Service      string
	Username     string
	Password     string
	AccessToken  string
	RefreshToken string
	Ucode        string
}
