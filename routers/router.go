package routers

import (
	"net/http"

	"github.com/MauricioUhlig/go-rest-api/controllers"
	"github.com/MauricioUhlig/go-rest-api/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/personalidades", controllers.Home).Methods("Get")
	r.HandleFunc("/personalidades/{id}", controllers.GetById).Methods("Get")
	r.HandleFunc("/personalidades", controllers.CreateNewPersinalidade).Methods("Post")
	r.HandleFunc("/personalidades/{id}", controllers.DeletePersonalidade).Methods("Delete")
	r.HandleFunc("/personalidades/{id}", controllers.EditPersonalidade).Methods("Put")

	http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r))
}
