package jabama

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	jabama_dto "github.com/amirhosseinf79/renthub_service/internal/dto/jabama"
)

func (h *service) GetCalendarDetails(fields dto.UpdateFields) (*models.Log, *jabama_dto.CalendarRoomResponse, error) {
	// log := h.initLog(fields.UserID, fields.ClientID, dto.GetCalendar)
	log, err := h.updateRoomID(&fields)
	if err != nil {
		return log, nil, err
	}
	endpoint := h.getEndpoints().GetCalendarDetails
	log.Action = dto.GetCalendar

	var response jabama_dto.CalendarRoomResponse
	err = h.handleUpdateResult(log, nil, endpoint, fields, &response)
	return log, &response, err
}
