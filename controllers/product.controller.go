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
)

var CreateProduct = func(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	idUser := r.Header.Get("user_id")
	i, err := strconv.Atoi(idUser)
	if err != nil {
		fmt.Println(err)
	}
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "Invalid Request")
	}
	resp, err := services.CreateProduct(GetDB(), product, i)
	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "Invalid Request")
	}
	helpers.ResponseWithJson(w, 200, resp)

}

var GetProductById = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idProduct, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
	}
	idUser, err := strconv.Atoi(r.Header.Get("user_id"))
	if err != nil {
		fmt.Println(err)
	}
	resp, err := services.GetProductID(GetDB(), idUser, idProduct)
	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "Invalid Request")
	}
	helpers.ResponseWithJson(w, 200, resp)

}

var GetProductAll = func(w http.ResponseWriter, r *http.Request) {
	idUser, err := strconv.Atoi(r.Header.Get("user_id"))
	if err != nil {
		fmt.Println(err)
	}
	resp, err := services.GetAllProduct(GetDB(), idUser)
	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "Invalid Request")
	}
	helpers.ResponseWithJson(w, 200, resp)
}

var DeleteProduct = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idProduct, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
	}
	idUser, err := strconv.Atoi(r.Header.Get("user_id"))
	if err != nil {
		fmt.Println(err)
	}
	resp, err := services.DeleteProduct(GetDB(), idUser, idProduct)
	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "Invalid Request")
	}
	helpers.ResponseWithJson(w, 200, resp)
}

var UpdateProduct = func(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	params := mux.Vars(r)

	idProduct, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
	}
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "Invalid Request")
	}
	idUser, err := strconv.Atoi(r.Header.Get("user_id"))
	if err != nil {
		fmt.Println(err)
	}
	product.IDUser = uint(idUser)
	resp, err := services.UpdateProduct(GetDB(), product, idProduct)
	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "Invalid Request")
	}
	helpers.ResponseWithJson(w, 200, resp)
}
