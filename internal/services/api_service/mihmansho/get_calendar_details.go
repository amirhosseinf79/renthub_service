package mihmansho

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	"github.com/amirhosseinf79/renthub_service/pkg"
)

func (h *service) GetCalendarDetails(fields dto.UpdateFields, response any) (log *models.Log, err error) {
	if len(fields.Dates) == 0 {
		err = dto.ErrInvalidRequest
		return
	}

	log = h.initLog(fields.UserID, fields.ClientID, dto.SetPrice)
	endpoint := h.getEndpoints().GetCalendarDetails
	dates := pkg.DatesToJalali(fields.Dates, false)

	url, err := h.getFullURL(endpoint, fields.RoomID, dates[0])
	if err != nil {
		return
	}

	header := h.getHeader()
	extraHeader := map[string]string{}
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
