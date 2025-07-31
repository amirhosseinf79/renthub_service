package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	Service      string `json:"service"`
	Action       string `json:"action"`
	UserID       uint   `json:"-"`
	ClientID     string `json:"clientID"`
	IsSucceed    bool   `json:"isSucceed"`
	RequestURL   string `json:"requestURL"`
	RequestBody  string `json:"requestBody"`
	StatusCode   int    `json:"statusCode"`
	ResponseBody string `json:"responseBody"`
	FinalResult  string `json:"finalResult"`
	User         User   `json:"-"`
}
