package controllers

import (
	"anekakios/helpers"
	"anekakios/models"
	"anekakios/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var dbConfig *gorm.DB

func InitiateDB(db *gorm.DB) {
	dbConfig = db
}

func GetDB() *gorm.DB {
	return dbConfig
}

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	newUser := &models.User{}
	err := json.NewDecoder(r.Body).Decode(newUser)
	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "Invalid Request")
	}
	_, resp := services.CreateAccount(GetDB(), newUser)
	if resp != nil {
		helpers.ResponseWithJson(w, 200, resp)
	}
}

var Login = func(w http.ResponseWriter, r *http.Request) {
	userLogin := &models.Login{}
	err := json.NewDecoder(r.Body).Decode(userLogin)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "Invalid Request")
	}
	resp, err := services.Login(GetDB(), userLogin)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "Invalid Request")
	}
	helpers.ResponseWithJson(w, 200, resp)
}

var GetUserByID = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
	}
	resp, err := services.GetUserByID(GetDB(), id)
	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "Invalid Request")
	}
	helpers.ResponseWithJson(w, 200, resp)
}

var GetUsers = func(w http.ResponseWriter, r *http.Request) {
	resp, err := services.GetAllUser(GetDB())
	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "Invalid Request")
	}
	helpers.ResponseWithJson(w, 200, resp)
}

var UpdateUser = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
	}
	user := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "Invalid Request")
	}
	resp, err := services.UpdateUser(GetDB(), id, user)
	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "Invalid Request")
	}
	helpers.ResponseWithJson(w, 200, resp)
}

var DeleteUser = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
	}
	resp, err := services.DeleteUser(GetDB(), id)
	if err != nil {
		helpers.ResponseWithJson(w, http.StatusBadGateway, resp)
	}
	helpers.ResponseWithJson(w, 200, resp)
}
