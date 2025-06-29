package chromium

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

type ChromiumService struct {
	browser *rod.Browser
}

func NewChromiumService() interfaces.ChromeService {
	confUrl := launcher.New().Headless(true).NoSandbox(true).MustLaunch()
	browser := rod.New().ControlURL(confUrl)
	err := browser.Connect()
	if err != nil {
		log.Fatal("Failed to connect to the browser:", err)
	}
	return &ChromiumService{
		browser: browser,
	}
}

func (s *ChromiumService) Close() {
	s.browser.MustClose()
}

func (s *ChromiumService) GetMihmanshoSessionID(token string, log *models.Log) (string, error) {
	page := s.browser.MustPage()
	// token := "671dd13a-993f-4a49-b68f-c3041586e479"
	log.StatusCode = 200
	log.RequestBody = "Normal GET"
	log.RequestURL = fmt.Sprintf("https://www.mihmansho.com/myapi/v1/checklogin?token=%s&returnUrl=/account/home/manage", token)

	page.MustSetExtraHeaders("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	page.MustSetExtraHeaders("ASP.NET_SessionId", token)

	page.MustNavigate(log.RequestURL).MustWaitLoad()
	page.MustNavigate(log.RequestURL).MustWaitLoad()

	log.ResponseBody = page.MustInfo().URL
	loggedIn := strings.Contains(log.ResponseBody, "account/home/manage")
	if !loggedIn {
		return "", dto.ErrorSessionNotFound
	}

	cookies := page.MustCookies()
	for _, c := range cookies {
		fmt.Println("Cookie Name:", c.Name)
		if c.Name == "ASP.NET_SessionId" {
			fmt.Println("ASP.NET_SessionId:", c.Value)
			return c.Value, nil
		}
	}
	return "", dto.ErrorSessionNotFound
}

func (s *ChromiumService) GetJajigaHeaders(log *models.Log) (map[string]string, error) {
	page := s.browser.MustPage("https://www.jajiga.com")

	headers := make(map[string]string)
	targetRequestSubstring := "api.jajiga.com"
	found := false

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	page.EachEvent(func(e *proto.NetworkRequestWillBeSent) {
		select {
		case <-ctx.Done():
			return
		default:
		}
		if strings.Contains(e.Request.URL, targetRequestSubstring) {
			fmt.Println("Request Found:", e.Request.URL)
			for k, v := range e.Request.Headers {
				headers[k] = fmt.Sprintf("%v", v)
			}
			found = true
			cancel()
		}
	})()

	page.MustNavigate("https://www.jajiga.com").MustWaitLoad()
	<-ctx.Done()

	if !found {
		return nil, dto.ErrInvalidRequest
	}
	return headers, nil
}
