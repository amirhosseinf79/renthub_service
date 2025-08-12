package jajiga

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	jajiga_dto "github.com/amirhosseinf79/renthub_service/internal/dto/jajiga"
)

// section=inprogress , page=1

func (h *service) GetReservations(fields dto.RecieveFields) (*models.Log, interfaces.ReservationResponseInterface, error) {
	log := h.initLog(fields.UserID, fields.ClientID, dto.GetReservations)
	endpoint := h.getEndpoints().GetReservations

	fields.Filters["per_page"] = 30

	var response jajiga_dto.ReservationResponse
	err := h.handleGetResult(log, endpoint, fields, &response)
	return log, &response, err
}
