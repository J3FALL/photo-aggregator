package interfaces

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"photo-aggregator/src/usecases"
	"strconv"

	"github.com/gorilla/mux"
)

type PhotoInteractor interface {
	Photographers() ([]usecases.Photographer, error)
	Photographer(id int) (usecases.Photographer, error)
}
type WebServiceHandler struct {
	PhotoInteractor PhotoInteractor
}

func (handler WebServiceHandler) ShowAllPhotographers(res http.ResponseWriter, req *http.Request) {
	photographers, err := handler.PhotoInteractor.Photographers()
	if err != nil {
		fmt.Println(err)
	}

	body, err := json.Marshal(photographers)
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(res, string(body))
	/*for _, photographer := range photographers {
		/*io.WriteString(res, fmt.Sprintf("id: %d\n", photographer.ID))
		io.WriteString(res, fmt.Sprintf("name: %s\n", photographer.Name))
		io.WriteString(res, fmt.Sprintf("surname: %s\n", photographer.Surname))
		io.WriteString(res, fmt.Sprintf("phone: %s\n", photographer.Phone))
	}*/
}

func (handler WebServiceHandler) GetPhotographerById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}
	photographer, err := handler.PhotoInteractor.Photographer(id)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(photographer.ID)
	//dirty code of 0 rows issue
	if photographer.ID == -1 {
		body, err := json.Marshal(nil)
		if err != nil {
			fmt.Println(err)
		}
		io.WriteString(res, string(body))
	} else {
		fmt.Println(photographer.ID)
		body, err := json.Marshal(photographer)

		if err != nil {
			fmt.Println(err)
		}

		io.WriteString(res, string(body))
	}

}
