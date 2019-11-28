package handlers

import (
	"net/http"
	"server/models"
	"server/utils"
)

func CountHandler(w http.ResponseWriter,r *http.Request)  {
	res := utils.GetCount()
	utils.ReturnData(w,models.Success,"",res)
	return
}
