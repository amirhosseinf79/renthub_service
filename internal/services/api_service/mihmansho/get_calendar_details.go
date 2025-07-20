package mihmansho

import (
	"errors"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/pkg"
	"gorm.io/gorm"
)

func (h *service) GetCalendarDetails(fields dto.UpdateFields, response any) (log *models.Log, err error) {
	log = h.initLog(fields.UserID, fields.ClientID, dto.SetPrice)
	model, err := h.apiAuthService.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = dto.ErrorApiTokenExpired
		}
		log.FinalResult = err.Error()
		return
	}
	if len(fields.Dates) == 0 {
		err = dto.ErrInvalidRequest
		return
	}

	endpoint := h.getEndpoints().GetCalendarDetails
	dates := pkg.DatesToJalali(fields.Dates, false)

	url, err := h.getFullURL(endpoint, fields.RoomID, dates[0])
	if err != nil {
		return
	}

	header := h.getHeader()
	extraHeader := h.getExtraHeader(model)
	request := h.request.New(endpoint.Method, url, header, extraHeader, log)
	err = request.Start(nil, endpoint.ContentType)

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
