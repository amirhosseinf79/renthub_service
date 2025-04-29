package cloner

import (
	"encoding/json"
	"sort"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	homsa_dto "github.com/amirhosseinf79/renthub_service/internal/dto/homsa"
	jabama_dto "github.com/amirhosseinf79/renthub_service/internal/dto/jabama"
	jajiga_dto "github.com/amirhosseinf79/renthub_service/internal/dto/jajiga"
	mihmansho_dto "github.com/amirhosseinf79/renthub_service/internal/dto/mihmansho"
	otaghak_dto "github.com/amirhosseinf79/renthub_service/internal/dto/otaghak"
	shab_dto "github.com/amirhosseinf79/renthub_service/internal/dto/shab"
)

func (h *homsaService) generatePriceBody(roomID string, amount int, dates []string) any {
	switch h.service {
	case "homsa":
		if len(dates) > 1 {
			sort.Strings(dates)
		}
		return homsa_dto.HomsaPriceBody{
			StartDate:    dates[0],
			EndDate:      dates[len(dates)-1],
			Price:        amount,
			KeepDiscount: 0,
		}
	case "jabama":
		return jabama_dto.EditPricePerDay{
			Type:  nil,
			Days:  dates,
			Value: amount * 10,
		}
	case "jajiga":
		return jajiga_dto.PriceBody{
			RoomID: roomID,
			Dates:  dates,
			Price:  amount,
		}
	case "otaghak":
		formattedDates := h.datesToIso(dates)
		var formattedDays []otaghak_dto.DayPricePair
		for _, item := range formattedDates {
			formattedDays = append(formattedDays, otaghak_dto.DayPricePair{Day: item, Price: amount})
		}
		return otaghak_dto.EditPriceBody{
			RoomID:       roomID,
			PerDayPrices: formattedDays,
		}
	case "shab":
		return shab_dto.EditPriceBody{
			KeepDiscount: false,
			Price:        amount / 1000,
			Dates:        h.datesToJalali(dates, true),
		}
	case "mihmansho":
		pbody := mihmansho_dto.FormBody{}
		jdates := h.datesToJalali(dates, false)
		for _, date := range jdates {
			pbody["Dates"] = date
		}
		mbody, err := json.Marshal(pbody)
		if err != nil {
			return err
		}
		return mbody
	default:
		return nil
	}
}

func (h *homsaService) EditPricePerDays(fields dto.UpdateFields) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID)
	endpoint := h.getEndpoints().EditPricePerDay
	body := h.generatePriceBody(fields.RoomID, fields.Amount, fields.Dates)
	err = h.handleUpdateResult(log, body, endpoint, fields)
	return log, err
}
