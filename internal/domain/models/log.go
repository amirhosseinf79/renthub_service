package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	Service      string
	Action       string
	UserID       uint
	ClientID     string
	IsSucceed    bool
	RequestURL   string
	RequestBody  string
	StatusCode   int
	ResponseBody string
	FinalResult  string
	User         User
}
