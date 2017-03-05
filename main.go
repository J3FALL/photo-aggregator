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

	webServiceHandler := interfaces.WebServiceHandler{}
	webServiceHandler.PhotoInteractor = photoInteractor

	http.HandleFunc("/photographers", func(res http.ResponseWriter, req *http.Request) {
		webServiceHandler.ShowAllPhotographers(res, req)
	})
	http.ListenAndServe(":8080", nil)
	/*db, err := sql.Open("postgres", "postgres://postgres:vqislemaro1@localhost/photo?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}*/
	/*rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var id int
		var email string
		errr := rows.Scan(&id, &email)
		if errr != nil {
			fmt.Println(errr)
		}
		fmt.Println(id, " ", email)
	}*/

}
