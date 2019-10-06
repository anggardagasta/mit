package delivery

import (
	"encoding/json"
	"github.com/anggardagasta/mit/models"
	"github.com/anggardagasta/mit/service/repository/response"
	"github.com/asaskevich/govalidator"
	"net/http"
)

func (hd handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var request models.FormUserRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		response.ResultError(w, http.StatusInternalServerError, response.MessageInternalError, err)
		return
	}

	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		response.ResultError(w, http.StatusInternalServerError, response.MessageInternalError, err)
		return
	}

	err = hd.usersUsecase.RegisterUser(request)
	if err != nil {
		response.ResultError(w, http.StatusInternalServerError, response.MessageInternalError, err)
		return
	}
	response.Result(w, response.MessageSucceed, http.StatusOK)
}

func (hd handler) CheckUserPhone(w http.ResponseWriter, r *http.Request) {
	var request models.FormPhone

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		response.ResultError(w, http.StatusInternalServerError, response.MessageInternalError, err)
		return
	}

	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		response.ResultError(w, http.StatusInternalServerError, response.MessageInternalError, err)
		return
	}
	result, err := hd.usersUsecase.GetUserByPhoneNumber(request)
	if err != nil {
		response.ResultError(w, http.StatusInternalServerError, response.MessageInternalError, err)
		return
	}

	if result.Id <= 0 {
		response.ResultWithData(w, nil, response.MessageUserNotFount, http.StatusNotFound)
	} else {
		response.ResultWithData(w, result, response.MessageSucceed, http.StatusOK)
	}
}

func (hd handler) CheckUserEmail(w http.ResponseWriter, r *http.Request) {
	var request models.FormEmail

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		response.ResultError(w, http.StatusInternalServerError, response.MessageInternalError, err)
		return
	}

	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		response.ResultError(w, http.StatusInternalServerError, response.MessageInternalError, err)
		return
	}

	result, err := hd.usersUsecase.GetUserByEmail(request)
	if err != nil {
		response.ResultError(w, http.StatusInternalServerError, response.MessageInternalError, err)
		return
	}

	if result.Id <= 0 {
		response.ResultWithData(w, nil, response.MessageUserNotFount, http.StatusNotFound)
	} else {
		response.ResultWithData(w, result, response.MessageSucceed, http.StatusOK)
	}
}
