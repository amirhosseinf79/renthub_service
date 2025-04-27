package dto

import "errors"

var ErrEmailExists = errors.New("email already exists")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrUserNotFound = errors.New("user not found")
var ErrEmptyAuth = errors.New("username or password cannot be empty")
var ErrEmptyPhone = errors.New("phone number cannot be empty")
var ErrEmptyCode = errors.New("code field cannot be empty")
