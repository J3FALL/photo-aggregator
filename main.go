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
	photoInteractor.TagRepository = interfaces.NewDbTagRepo(handlers)
	webServiceHandler := interfaces.WebServiceHandler{}
	webServiceHandler.PhotoInteractor = photoInteractor

	router := mux.NewRouter()

	//Photographers API
	router.HandleFunc("/api/photographers", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		webServiceHandler.ShowAllPhotographers(res, req)
	})
	router.HandleFunc("/api/photographer/{id}", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			res.Header().Set("Content-Type", "application/json")
			webServiceHandler.GetPhotographerById(res, req)
		} else if req.Method == "PUT" {
			res.Header().Set("Content-Type", "application/json")
			webServiceHandler.UpdatePhotographer(res, req)
		}
	})
	router.HandleFunc("/api/photographer", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			res.Header().Set("Content-Type", "application/json")
			webServiceHandler.CreateNewPhotographer(res, req)
		}
	})

	//Tags API
	router.HandleFunc("/api/tag/{id}", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			res.Header().Set("Content-Type", "application/json")
			webServiceHandler.GetTagById(res, req)
		} else if req.Method == "PUT" {
			res.Header().Set("Content-Type", "application/json")
			webServiceHandler.UpdateTag(res, req)
		}
	})
	router.HandleFunc("/api/tag", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			res.Header().Set("Content-Type", "application/json")
			webServiceHandler.CreateNewTag(res, req)
		}
	})
	router.HandleFunc("/api/tags", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		webServiceHandler.ShowAllTags(res, req)
	})

	http.Handle("/api", router)

	//Front-end router setup
	router.PathPrefix("/assets/css/").Handler(http.StripPrefix("/assets/css/", http.FileServer(http.Dir("./assets/css"))))
	router.PathPrefix("/assets/templates/").Handler(http.StripPrefix("/assets/templates/", http.FileServer(http.Dir("./assets/templates"))))
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		webServiceHandler.ShowTemplates(res, req)
	})
	http.Handle("/", router)

	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
