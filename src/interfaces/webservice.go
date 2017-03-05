package interfaces

import (
	"fmt"
	"io"
	"net/http"
	"photo-aggregator/src/usecases"
)

type PhotoInteractor interface {
	Photographers() ([]usecases.Photographer, error)
}
type WebServiceHandler struct {
	PhotoInteractor PhotoInteractor
}

func (handler WebServiceHandler) ShowAllPhotographers(res http.ResponseWriter, req *http.Request) {
	photographers, err := handler.PhotoInteractor.Photographers()
	if err != nil {
		fmt.Println(err)
	}
	for _, photographer := range photographers {
		io.WriteString(res, fmt.Sprintf("id: %d\n", photographer.ID))
		io.WriteString(res, fmt.Sprintf("name: %s\n", photographer.Name))
		io.WriteString(res, fmt.Sprintf("surname: %s\n", photographer.Surname))
		io.WriteString(res, fmt.Sprintf("phone: %s\n", photographer.Phone))
	}
}
