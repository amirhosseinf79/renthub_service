package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	// اجرای یک مرورگر headless (می‌تونی headless رو برداری برای مشاهده)
	url := launcher.New().Headless(true).MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage()

	// آدرس checklogin با توکن
	token := "671dd13a-993f-4a49-b68f-c3041586e479"
	checkLoginUrl := fmt.Sprintf("https://www.mihmansho.com/myapi/v1/checklogin?token=%s&returnUrl=/account/home/manage", token)

	// باز کردن صفحه
	page.MustNavigate(checkLoginUrl).MustWaitLoad()

	// fmt.Println(page.HTML())

	cookies := page.MustCookies()

	for _, c := range cookies {
		fmt.Println("Cookie Name:", c.Name)
		if c.Name == "ASP.NET_SessionId" {
			fmt.Println("ASP.NET_SessionId:", c.Value)
		}
	}
}
