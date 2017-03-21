package main

import (
	"net/http"
	"os"
	"photo-aggregator/src/infrastructure"
	"photo-aggregator/src/interfaces"
	"photo-aggregator/src/usecases"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func main() {
	dbHandler := infrastructure.NewPgHandler(os.Getenv("DATABASE_URL"))
	handlers := make(map[string]interfaces.DbHandler)
	handlers["DbUserRepo"] = dbHandler
	handlers["DbTagRepo"] = dbHandler
	handlers["DbPhotographerRepo"] = dbHandler
	handlers["DbAttachmentRepo"] = dbHandler

	photoInteractor := new(usecases.PhotoInteractor)
	photoInteractor.UserRepository = interfaces.NewDbUserRepo(handlers)
	photoInteractor.PhotographerRepository = interfaces.NewDbPhotographerRepo(handlers)
	webServiceHandler := interfaces.WebServiceHandler{}
	webServiceHandler.PhotoInteractor = photoInteractor

	router := mux.NewRouter()
	router.HandleFunc("/photographers", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		webServiceHandler.ShowAllPhotographers(res, req)
	})
	router.HandleFunc("/photographer/{id}", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			res.Header().Set("Content-Type", "application/json")
			webServiceHandler.GetPhotographerById(res, req)
		}
	})

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}
