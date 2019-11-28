package utils

import (
	"encoding/json"
	"net/http"
	"server/models"
)

func NewResult(code int,msg string,data interface{}) *models.Result {
	return &models.Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func ReturnData(w http.ResponseWriter, code int,msg string,data interface{})  {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(NewResult(code,msg,data))
	return
}

