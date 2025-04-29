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

func (h *homsaService) generateSetMinNightBody(roomID string, amount int, dates []string) any {
	switch h.service {
	case "homsa":
		if len(dates) > 1 {
			sort.Strings(dates)
		}
		return homsa_dto.HomsaSetMinNightBody{
			StartDate: dates[0],
			EndDate:   dates[len(dates)-1],
			Min:       amount,
			Max:       nil,
		}
	case "jajiga":
		return jajiga_dto.MinNightBody{
			RoomID:    roomID,
			Dates:     dates,
			MinNights: amount,
		}
	case "otaghak":
		return otaghak_dto.EditMinNightBody{
			MinNights:     amount,
			EffectiveDays: h.datesToIso(dates),
			RoomID:        roomID,
		}
	case "shab":
		return shab_dto.EditMinNightBody{
			Action:  "set_min_days",
			Dates:   h.datesToJalali(dates),
			MinDays: amount,
		}
	default:
		return nil
	}
}

func (h *homsaService) SetMinNight(fields dto.UpdateFields) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID)
	endpoint := h.getEndpoints().SetMinNight
	body := h.generateSetMinNightBody(fields.RoomID, fields.Amount, fields.Dates)
	err = h.handleUpdateResult(log, body, endpoint, fields)
	return log, err
}
