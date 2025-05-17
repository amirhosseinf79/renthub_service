package mihmansho

import (
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func (s *service) getSession(token string) (string, error) {
	url := launcher.New().Headless(true).NoSandbox(true).MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage()
	// token := "671dd13a-993f-4a49-b68f-c3041586e479"
	checkLoginUrl := fmt.Sprintf("https://www.mihmansho.com/myapi/v1/checklogin?token=%s&returnUrl=/account/home/manage", token)
	page.MustNavigate(checkLoginUrl).MustWaitLoad()
	cookies := page.MustCookies()

	for _, c := range cookies {
		fmt.Println("Cookie Name:", c.Name)
		if c.Name == "ASP.NET_SessionId" {
			fmt.Println("ASP.NET_SessionId:", c.Value)
			return c.Value, nil
		}
	}
	return "", dto.ErrorUnauthorized
}
