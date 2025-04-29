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

func (h *homsaService) generateRemoveDiscountBody(roomID string, dates []string) any {
	switch h.service {
	case "homsa":
		if len(dates) > 1 {
			sort.Strings(dates)
		}
		return homsa_dto.HomsaRemoveDiscountBody{
			StartDate:    dates[0],
			EndDate:      dates[len(dates)-1],
			KeepDiscount: 0,
		}
	case "jajiga":
		return jajiga_dto.DiscountBody{
			RoomID:  roomID,
			Dates:   dates,
			Percent: 0,
		}
	case "otaghak":
		return otaghak_dto.EditDiscountBody{
			DiscountPercent: 0,
			EffectiveDays:   h.datesToIso(dates),
			RoomID:          roomID,
		}
	case "shab":
		return shab_dto.EditDiscountBody{
			Action: "unset_daily_discount",
			Dates:  h.datesToJalali(dates),
		}
	default:
		return nil
	}
}

func (h *homsaService) RemoveDiscount(fields dto.UpdateFields) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID)
	endpoint := h.getEndpoints().RemoveDiscount
	body := h.generateRemoveDiscountBody(fields.RoomID, fields.Dates)
	err = h.handleUpdateResult(log, body, endpoint, fields)
	return log, err
}
