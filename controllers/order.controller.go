package controllers

import (
	"anekakios/helpers"
	"anekakios/models"
	"anekakios/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var Kulak = func(w http.ResponseWriter, r *http.Request) {
	var productKulak models.OrderKulak
	idUser := r.Header.Get("user_id")
	i, err := strconv.Atoi(idUser)
	if err != nil {
		fmt.Println(err)
	}
	if err := json.NewDecoder(r.Body).Decode(&productKulak); err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "Invalid Request")
	}
	productKulak.TanggalBeli = time.Now()
	err = services.Kulak(GetDB(), productKulak, i)
	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "Cannot Kulak")
	}
	helpers.ResponseWithJson(w, 200, map[string]string{
		"message": "Kulak telah berhasil dibeli",
	})
}
