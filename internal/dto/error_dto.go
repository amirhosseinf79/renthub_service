package dto

import "errors"

var ErrEmailExists = errors.New("ایمیل وارد شده وجود دارد")
var ErrInvalidCredentials = errors.New("نام کاربری یا رمز عبور نامعتبر است")
var ErrUserNotFound = errors.New("کاربر مورد نظر یافت نشد")
var ErrorUnauthorized = errors.New("توکن یافت نشد")
var ErrorPermission = errors.New("دسترسی غیر مجاز")
var ErrInvalidRequest = errors.New("ورودی نامعتبر است")
var ErrRoomNotFound = errors.New("اتاق مورد نظر یافت نشد")
var ErrServiceUnavailable = errors.New("سرویس مورد نظر فعال نمی باشد")
var ErrInvalidPrice = errors.New("هزینه وارد شده باید بزرگتر از 0 و به تومان باشد")
