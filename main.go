package main

import (
	"net/http"
	"photo-aggregator/src/infrastructure"
	"photo-aggregator/src/interfaces"
	"photo-aggregator/src/usecases"

	_ "github.com/lib/pq"
)

func main() {
	dbHandler := infrastructure.NewPgHandler("postgres://postgres:vqislemaro1@localhost/photo?sslmode=disable")
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

	http.HandleFunc("/photographers", func(res http.ResponseWriter, req *http.Request) {
		webServiceHandler.ShowAllPhotographers(res, req)
	})
	http.ListenAndServe(":8080", nil)
}
