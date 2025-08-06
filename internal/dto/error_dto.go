package dto

import "errors"

var ErrEmailExists = errors.New("ایمیل وارد شده وجود دارد")
var ErrInvalidCredentials = errors.New("نام کاربری یا رمز عبور نامعتبر است")
var ErrInvalidPhoneNumber = errors.New("شماره وارد شده صحیح نمیباشد")
var ErrInvalidCode = errors.New("کد تایید نامعتبر است")
var ErrUserNotFound = errors.New("کاربر مورد نظر یافت نشد")
var ErrorApiTokenExpired = errors.New("token expired")
var ErrorUnauthorized = errors.New("توکن یافت نشد")
var ErrorSessionNotFound = errors.New("کد دسترسی میهمان شو یافت نشد")
var ErrorPermission = errors.New("دسترسی مجاز نیست")
var ErrInvalidRequest = errors.New("ورودی نامعتبر است")
var ErrInvalidDate = errors.New("تاریخ ورودی نامعتبر است")
var ErrRoomNotFound = errors.New("اتاق مورد نظر یافت نشد")
var ErrServiceUnavailable = errors.New("سرویس مورد نظر فعال نمی باشد")
var ErrInvalidPrice = errors.New("هزینه وارد شده باید بزرگتر از 0 و به تومان باشد")
var ErrTimeOut = errors.New("در این لحظه امکان بروزرسانی وجود ندارد")
var ErrInvalidDay = errors.New("روز انتخاب شده معتبر نیست")
var ErrUnknownMsg = errors.New("خطایی در بروزرسانی رخ داد")
var ErrJabamaViaTicket = errors.New("محدویت رزرو خود را از طریق تیکت در بخش پشتیبانی جاباما اطلاع دهید")
