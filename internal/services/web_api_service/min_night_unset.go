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

func (h *homsaService) generateUnsetMinNightBody(roomID string, dates []string) any {
	switch h.service {
	case "homsa":
		if len(dates) > 1 {
			sort.Strings(dates)
		}
		return homsa_dto.HomsaUnsetMinNightBody{
			StartDate: dates[0],
			EndDate:   dates[len(dates)-1],
		}
	case "jajiga":
		return jajiga_dto.MinNightBody{
			RoomID:    roomID,
			Dates:     dates,
			MinNights: 1,
		}
	case "otaghak":
		return otaghak_dto.EditMinNightBody{
			MinNights:     1,
			EffectiveDays: h.datesToIso(dates),
			RoomID:        roomID,
		}
	case "shab":
		return shab_dto.EditMinNightBody{
			Action:  "set_min_days",
			Dates:   h.datesToJalali(dates, true),
			MinDays: 1,
		}
	}
	return nil
}

func (h *homsaService) UnsetMiniNight(fields dto.UpdateFields) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID)
	endpoint := h.getEndpoints().UnsetMinNight
	body := h.generateUnsetMinNightBody(fields.RoomID, fields.Dates)
	err = h.handleUpdateResult(log, body, endpoint, fields)
	return log, err
}
