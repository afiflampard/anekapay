package main

import (
	"anekakios/config"
	"anekakios/controllers"
	"anekakios/middleware"
	"anekakios/migrations"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	fmt.Println(port)
	db := config.Init()
	migrations.Migrate(db)
	controllers.InitiateDB(db)

	router := mux.NewRouter()
	router.HandleFunc("/", HelloWorld).Methods("GET")
	subRouterAuth := router.PathPrefix("/auth").Subrouter()
	subRouterUser := router.PathPrefix("/user").Subrouter()
	subRouterProduct := router.PathPrefix("/product").Subrouter()
	subRouterKulak := router.PathPrefix("/kulak").Subrouter()

	subRouterAuth.HandleFunc("/v1/login", controllers.Login).Methods("POST")
	subRouterAuth.HandleFunc("/v1/signup", controllers.CreateAccount).Methods("POST")

	subRouterUser.Use(middleware.JwtVerifyToken)
	subRouterUser.HandleFunc("/v1/{id}", controllers.GetUserByID).Methods("GET")
	subRouterUser.HandleFunc("/v1/", controllers.GetUsers).Methods("GET")
	subRouterUser.HandleFunc("/v1/{id}", controllers.UpdateUser).Methods("PUT")
	subRouterUser.HandleFunc("/v1/{id}", controllers.DeleteUser).Methods("DELETE")

	subRouterProduct.Use(middleware.JwtVerifyToken)
	subRouterProduct.HandleFunc("/v1/create", controllers.CreateProduct).Methods("POST")
	subRouterProduct.HandleFunc("/v1/{id}", controllers.GetProductById).Methods("GET")
	subRouterProduct.HandleFunc("/v1/", controllers.GetProductAll).Methods("GET")
	subRouterProduct.HandleFunc("/v1/{id}", controllers.DeleteProduct).Methods("DELETE")
	subRouterProduct.HandleFunc("/v1/{id}", controllers.UpdateProduct).Methods("PUT")

	subRouterKulak.Use(middleware.JwtVerifyToken)
	subRouterKulak.HandleFunc("/v1/kulak", controllers.Kulak).Methods("POST")

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err)
	}

}

func HelloWorld(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running"))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running"))
}
