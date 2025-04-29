package cloner

import (
	"sort"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	homsa_dto "github.com/amirhosseinf79/renthub_service/internal/dto/homsa"
	jajiga_dto "github.com/amirhosseinf79/renthub_service/internal/dto/jajiga"
	otaghak_dto "github.com/amirhosseinf79/renthub_service/internal/dto/otaghak"
	shab_dto "github.com/amirhosseinf79/renthub_service/internal/dto/shab"
)

func (h *homsaService) generateAddDiscountBody(roomID string, amount int, dates []string) any {
	switch h.service {
	case "homsa":
		if len(dates) > 1 {
			sort.Strings(dates)
		}
		return homsa_dto.HomsaAddDiscountBody{
			StartDate:    dates[0],
			EndDate:      dates[len(dates)-1],
			Discount:     amount,
			KeepDiscount: 0,
		}
	case "jajiga":
		return jajiga_dto.DiscountBody{
			RoomID:  roomID,
			Dates:   dates,
			Percent: amount,
		}
	case "otaghak":
		return otaghak_dto.EditDiscountBody{
			DiscountPercent: amount,
			EffectiveDays:   h.datesToIso(dates),
			RoomID:          roomID,
		}
	case "shab":
		return shab_dto.EditDiscountBody{
			Action:        "set_daily_discount",
			Dates:         h.datesToJalali(dates),
			DailyDiscount: amount,
		}
	}
	return nil
}

func (h *homsaService) AddDiscount(fields dto.UpdateFields) (log *models.Log, err error) {
	endpoint := h.getEndpoints().AddDiscount
	log = h.initLog(fields.UserID, fields.ClientID)
	body := h.generateAddDiscountBody(fields.RoomID, fields.Amount, fields.Dates)
	err = h.handleUpdateResult(log, body, endpoint, fields)
	return log, err
}
