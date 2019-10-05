package response

import (
	"encoding/json"
	"net/http"
)

type (
	responseBody struct {
		Status  int64             `json:"status"`
		Message map[string]string `json:"message"`
		Data    interface{}       `json:"data,omitempty"`
	}
)

func ResultOk(w http.ResponseWriter) {
	result := responseBody{}
	result.Status = http.StatusOK
	result.Message = map[string]string{
		"id": "Berhasil",
		"en": "Success",
	}

	_ = json.NewEncoder(w).Encode(result)
}

func ResultOkWithData(w http.ResponseWriter, data interface{}) {
	result := responseBody{}
	result.Status = http.StatusOK
	result.Message = map[string]string{
		"id": "Berhasil",
		"en": "Success",
	}
	result.Data = data

	_ = json.NewEncoder(w).Encode(result)
}

func ResultError(w http.ResponseWriter, statusCode int64, err error) {
	result := responseBody{}
	result.Status = statusCode
	result.Message = map[string]string{
		"id": "Terdapat kesalahan: " + err.Error(),
		"en": "Something when wrong: " + err.Error(),
	}
	_ = json.NewEncoder(w).Encode(result)
}
