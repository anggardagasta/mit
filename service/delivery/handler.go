package delivery

import (
	"encoding/json"
	"github.com/anggardagasta/mit/models"
	"github.com/anggardagasta/mit/service/repository/response"
	"github.com/asaskevich/govalidator"
	"net/http"
)

func (hd handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var request models.FormUser

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {

	}

	_, err = govalidator.ValidateStruct(request)
	if err != nil {

	}

	err = hd.usersUsecase.RegisterUser(request)
	if err != nil {

	}
	response.ResultOk(w)
}

func (hd handler) CheckUserPhone(w http.ResponseWriter, r *http.Request) {
	var request models.FormPhone

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {

	}

	_, err = govalidator.ValidateStruct(request)
	if err != nil {

	}
	result, err := hd.usersUsecase.GetUserByPhoneNumber(request)
	if err != nil {

	}
	response.ResultOkWithData(w, result)
}

func (hd handler) CheckUserEmail(w http.ResponseWriter, r *http.Request) {
	var request models.FormEmail

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {

	}

	_, err = govalidator.ValidateStruct(request)
	if err != nil {

	}

	result, err := hd.usersUsecase.GetUserByEmail(request)
	if err != nil {

	}
	response.ResultOkWithData(w, result)
}
