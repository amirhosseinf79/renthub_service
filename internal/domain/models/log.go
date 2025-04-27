package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	Service      string
	UserID       uint
	ClientID     string
	IsSucceed    bool
	RequestBody  string
	StatusCode   int
	ResponseBody string
	FinalResult  string
	User         User
}
