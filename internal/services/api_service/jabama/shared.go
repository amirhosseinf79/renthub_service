package jabama

import (
	"errors"
	"fmt"

	"github.com/amirhosseinf79/renthub_service/internal/domain/models"
	"github.com/amirhosseinf79/renthub_service/internal/dto"
	jabama_dto "github.com/amirhosseinf79/renthub_service/internal/dto/jabama"
	"gorm.io/gorm"
)

func (h *service) handleUpdateResult(log *models.Log, body any, endpoint dto.EndP, fields dto.UpdateFields) (err error) {
	model, err := h.apiAuthService.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = dto.ErrorApiTokenExpired
		}
		log.FinalResult = err.Error()
		return err
	}
	url, err := h.getFullURL(endpoint, fields.RoomID, fields.Amount)
	if err != nil {
		log.FinalResult = err.Error()
		return err
	}
	request := h.request.New(endpoint.Method, url, h.getHeader(), h.getExtraHeader(model), log)
	err = request.Start(body, endpoint.ContentType)
	if err != nil {
		log.FinalResult = err.Error()
		return err
	}
	ok, err := request.Ok()
	if !ok {
		log.FinalResult = err.Error()
		response := h.generateUpdateErrResponse()
		err2 := request.ParseInterface(response)
		if err2 != nil {
			return err
		}
		_, result := response.GetResult()
		log.FinalResult = result
		err = errors.New(result)
		return err
	}
	log.FinalResult = "success"
	log.IsSucceed = true
	return nil
}

func (h *service) handleGetResult(log *models.Log, endpoint dto.EndP, fields dto.RecieveFields, response any) error {
	model, err := h.apiAuthService.GetByUnique(fields.UserID, fields.ClientID, h.service)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = dto.ErrorApiTokenExpired
		}
		log.FinalResult = err.Error()
		return err
	}

	url, err := h.getFullURL(endpoint)
	if err != nil {
		log.FinalResult = err.Error()
		return err
	}

	request := h.request.New(endpoint.Method, url, h.getHeader(), h.getExtraHeader(model), log)
	err = request.Start(fields.Filters, endpoint.ContentType)
	if err != nil {
		log.FinalResult = err.Error()
		return err
	}
	ok, err := request.Ok()
	if !ok {
		log.FinalResult = err.Error()
		response := h.generateUpdateErrResponse()
		err2 := request.ParseInterface(response)
		if err2 != nil {
			return err
		}
		_, result := response.GetResult()
		log.FinalResult = result
		err = errors.New(result)
		return err
	}
	err = request.ParseInterface(response)
	if err != nil {
		log.FinalResult = err.Error()
		return err
	}
	log.FinalResult = "success"
	log.IsSucceed = true
	return nil
}

func (h *service) updateRoomID(fields *dto.UpdateFields) (log *models.Log, err error) {
	counter := 0
	found := false
	for !found && counter < 2 {
		counter++
		result := jabama_dto.RoomListResponse{}
		getFields := dto.RecieveFields{
			RequiredFields: fields.RequiredFields,
		}
		log, err = h.GetRoomList(getFields, &result)
		for _, room := range result.Result.Items {
			if fields.RoomID == fmt.Sprintf("%v", room.Code) {
				fields.RoomID = room.ID
				found = true
				break
			}
		}
	}
	if err != nil {
		log.FinalResult = err.Error()
		return
	}
	if !found {
		err = dto.ErrRoomNotFound
		log.FinalResult = err.Error()
		return
	}
	return
}

func (h *service) getCurrentPrice(fields dto.UpdateFields) (price int, log *models.Log, err error) {
	response := jabama_dto.CalendarResponse{}
	log, err = h.GetCalendarDetails(fields, &response)
	if err != nil {
		log.FinalResult = err.Error()
		return
	}
	price = response.Result.Price.CurrentPrice
	if price == 0 {
		err = dto.ErrInvalidDay
		log.FinalResult = err.Error()
	}
	return
}
