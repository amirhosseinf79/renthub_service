package jabama

import (
	"errors"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"gorm.io/gorm"
)

func (h *service) GetCalendarDetails(fields dto.UpdateFields, response any) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID, dto.GetCalendar)
	model, err := h.apiAuthService.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = dto.ErrorApiTokenExpired
		}
		log.StatusCode = 401
		log.FinalResult = err.Error()
		return
	}

	endpoint := h.getEndpoints().GetCalendarDetails
	url, err := h.getFullURL(endpoint, fields.RoomID)
	if err != nil {
		return
	}

	header := h.getHeader()
	extraHeader := h.getExtraHeader(model)
	request := h.request.New(endpoint.Method, url, header, extraHeader, log)
	body := h.generateCalendarDetailsBody(fields.Dates)
	err = request.Start(body, endpoint.ContentType)

	if err != nil {
		return
	}

	_, err = request.Ok()
	if err != nil {
		return
	}
	err = request.ParseInterface(response)
	return
}
