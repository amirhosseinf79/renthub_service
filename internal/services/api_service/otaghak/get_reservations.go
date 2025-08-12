package otaghak

import (
	"github.com/amirhosseinf79/renthub_service/internal/domain/interfaces"
	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	otaghak_dto "github.com/amirhosseinf79/renthub_service/internal/dto/otaghak"
)

// bookStatuses=[NewBooking] , top=100

func (h *service) GetReservations(fields dto.RecieveFields) (*models.Log, interfaces.ReservationResponseInterface, error) {
	log := h.initLog(fields.UserID, fields.ClientID, dto.GetReservations)
	endpoint := h.getEndpoints().GetReservations

	fields.Filters["top"] = 100

	var response otaghak_dto.ReservationResponse
	err := h.handleGetResult(log, endpoint, fields, &response)
	return log, &response, err
}
