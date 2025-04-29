package cloner

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	homsa_dto "github.com/amirhosseinf79/renthub_service/internal/dto/homsa"
	jabama_dto "github.com/amirhosseinf79/renthub_service/internal/dto/jabama"
	jajiga_dto "github.com/amirhosseinf79/renthub_service/internal/dto/jajiga"
	mihmansho_dto "github.com/amirhosseinf79/renthub_service/internal/dto/mihmansho"
	otaghak_dto "github.com/amirhosseinf79/renthub_service/internal/dto/otaghak"
	shab_dto "github.com/amirhosseinf79/renthub_service/internal/dto/shab"
	"github.com/amirhosseinf79/renthub_service/internal/services/requests"
	"gorm.io/gorm"
)

func (h *homsaService) generateCalendarBody(roomID string, setOpen bool, dates []string) any {
	switch h.service {
	case "homsa":
		if len(dates) > 1 {
			sort.Strings(dates)
		}
		return homsa_dto.HomsaCalendarBody{
			StartDate: dates[0],
			EndDate:   dates[len(dates)-1],
		}
	case "jabama":
		return jabama_dto.OpenClosCalendar{
			Dates: dates,
		}
	case "jajiga":
		var num int
		if !setOpen {
			num = 1
		}
		return jajiga_dto.CalendarBody{
			RoomID:       roomID,
			Dates:        dates,
			DisableCount: num,
		}
	case "otaghak":
		if setOpen {
			return otaghak_dto.CalendarBody{
				RoomID:        roomID,
				UnblockedDays: h.datesToIso(dates),
			}
		}
		return otaghak_dto.CalendarBody{
			RoomID:      roomID,
			BlockedDays: h.datesToIso(dates),
		}
	case "shab":
		status := "set_disabled"
		if setOpen {
			status = "unset_disabled"
		}
		return shab_dto.CalendarBody{
			Action: status,
			Dates:  h.datesToJalali(dates, true),
		}
	case "mihmansho":
		fDate := mihmansho_dto.Calendar{
			ProductId: roomID,
		}
		jdates := h.datesToJalali(dates, false)
		for _, jdate := range jdates {
			fDate.Dates = append(fDate.Dates, mihmansho_dto.CalendarDates{Date: jdate, IsReserve: !setOpen, RequestId: 0})
		}
		bdata, err := json.Marshal(fDate)
		if err != nil {
			return err
		}
		mainBody := mihmansho_dto.FormBody{
			"Dates": string(bdata),
		}
		mbody, err := json.Marshal(mainBody)
		if err != nil {
			return err
		}
		return mbody
	}
	return nil
}

func (h *homsaService) handleUpdateResult(log *models.Log, body any, endpoint dto.EndP, fields dto.UpdateFields) (err error) {
	model, err := h.apiAuthRepo.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = dto.ErrorUnauthorized
		}
		log.FinalResult = err.Error()
		return err
	}
	url, err := h.getFullURL(endpoint, fields.RoomID, fields.Amount)
	if err != nil {
		log.FinalResult = err.Error()
		return err
	}
	request := requests.New(endpoint.Method, url, h.getHeader(), h.getExtraHeader(model), log)
	err = request.Start(body, endpoint.ContentType)
	if err != nil {
		log.FinalResult = err.Error()
		return err
	}
	ok, err := request.Ok()
	fmt.Println("Request Result:", ok, err)
	if ok && h.service != "mihmansho" {
		log.FinalResult = "success"
		log.IsSucceed = true
		return nil
	} else if h.service != "mihmansho" {
		log.FinalResult = err.Error()
	}
	response := h.generateUpdateErrResponse()
	if response != nil {
		err2 := request.ParseInterface(response)
		if err2 == nil {
			ok, result := response.GetResult()
			if !ok && result != "" {
				err = errors.New(result)
				log.FinalResult = result
			} else if result != "" {
				log.FinalResult = result
			}
		}
	}
	return err
}
