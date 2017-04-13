package interfaces

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"photo-aggregator/src/usecases"
	"strconv"

	"github.com/gorilla/mux"
)

type PhotoInteractor interface {
	Photographers() ([]usecases.Photographer, error)
	Photographer(id int) (usecases.Photographer, error)
	NewPhotographer(photographer usecases.Photographer)
	UpdatePhotographer(photographer usecases.Photographer) bool

	NewTag(tag usecases.Tag)
	Tag(id int) (usecases.Tag, error)
	UpdateTag(tag usecases.Tag) bool
	Tags() ([]usecases.Tag, error)

	NewAttachment(attach usecases.Attachment)
	Attachment(id int) (usecases.Attachment, error)
	UpdateAttachment(attach usecases.Attachment) bool
	Attachments() ([]usecases.Attachment, error)
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

func (handler WebServiceHandler) CreateNewPhotographer(res http.ResponseWriter, req *http.Request) {
	var photographer usecases.Photographer
	_ = json.NewDecoder(req.Body).Decode(&photographer)

	fmt.Println(photographer)
	photographerTmp, err := handler.PhotoInteractor.Photographer(photographer.ID)
	fmt.Println("before ok")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(photographerTmp)
	if photographerTmp.ID == -1 {
		fmt.Println("ok")
		handler.PhotoInteractor.NewPhotographer(photographer)
	}

}

func (handler WebServiceHandler) UpdatePhotographer(res http.ResponseWriter, req *http.Request) {
	var photographer usecases.Photographer
	_ = json.NewDecoder(req.Body).Decode(&photographer)
	photographerTmp, err := handler.PhotoInteractor.Photographer(photographer.ID)
	if err != nil {
		fmt.Println(err)
	}
	if photographerTmp.ID != -1 {
		handler.PhotoInteractor.UpdatePhotographer(photographer)
	}
}

func (handler WebServiceHandler) CreateNewTag(res http.ResponseWriter, req *http.Request) {
	var tag usecases.Tag
	_ = json.NewDecoder(req.Body).Decode(&tag)

	fmt.Println("from webservice")
	tagTmp, err := handler.PhotoInteractor.Tag(tag.ID)
	if err != nil {
		fmt.Println(err)
	}

	if tagTmp.ID == -1 {
		handler.PhotoInteractor.NewTag(tag)
	}
}

func (handler WebServiceHandler) GetTagById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}
	tag, err := handler.PhotoInteractor.Tag(id)

	if err != nil {
		fmt.Println(err)
	}
	//dirty code of 0 rows issue
	if tag.ID == -1 {
		body, err := json.Marshal(nil)
		if err != nil {
			fmt.Println(err)
		}
		io.WriteString(res, string(body))
	} else {
		body, err := json.Marshal(tag)

		if err != nil {
			fmt.Println(err)
		}

		io.WriteString(res, string(body))
	}
}

func (handler WebServiceHandler) UpdateTag(res http.ResponseWriter, req *http.Request) {
	var tag usecases.Tag
	_ = json.NewDecoder(req.Body).Decode(&tag)
	tagTmp, err := handler.PhotoInteractor.Tag(tag.ID)
	if err != nil {
		fmt.Println(err)
	}
	if tagTmp.ID != -1 {
		handler.PhotoInteractor.UpdateTag(tag)
	}
}

func (handler WebServiceHandler) ShowAllTags(res http.ResponseWriter, req *http.Request) {
	tags, err := handler.PhotoInteractor.Tags()
	if err != nil {
		fmt.Println(err)
	}

	body, err := json.Marshal(tags)
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(res, string(body))
}

func (handler WebServiceHandler) CreateNewAttachment(res http.ResponseWriter, req *http.Request) {
	var attach usecases.Attachment
	_ = json.NewDecoder(req.Body).Decode(&attach)

	fmt.Println("from webservice")
	attachTmp, err := handler.PhotoInteractor.Attachment(attach.ID)
	if err != nil {
		fmt.Println(err)
	}

	if attachTmp.ID == -1 {
		handler.PhotoInteractor.NewAttachment(attach)
	}
}

func (handler WebServiceHandler) GetAttachmentById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}
	attach, err := handler.PhotoInteractor.Attachment(id)

	if err != nil {
		fmt.Println(err)
	}
	//dirty code of 0 rows issue
	if attach.ID == -1 {
		body, err := json.Marshal(nil)
		if err != nil {
			fmt.Println(err)
		}
		io.WriteString(res, string(body))
	} else {
		body, err := json.Marshal(attach)

		if err != nil {
			fmt.Println(err)
		}

		io.WriteString(res, string(body))
	}
}

func (handler WebServiceHandler) UpdateAttachment(res http.ResponseWriter, req *http.Request) {
	var attach usecases.Attachment
	_ = json.NewDecoder(req.Body).Decode(&attach)
	attachTmp, err := handler.PhotoInteractor.Attachment(attach.ID)
	if err != nil {
		fmt.Println(err)
	}
	if attachTmp.ID != -1 {
		handler.PhotoInteractor.UpdateAttachment(attach)
	}
}

func (handler WebServiceHandler) ShowAllAttachments(res http.ResponseWriter, req *http.Request) {
	attachments, err := handler.PhotoInteractor.Attachments()
	if err != nil {
		fmt.Println(err)
	}

	body, err := json.Marshal(attachments)
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(res, string(body))
}

//Front-end methods

func (handler WebServiceHandler) ShowTemplates(res http.ResponseWriter, req *http.Request) {
	t, err := template.New("").Delims("{{%", "%}}").ParseFiles("assets/templates/index.html", "assets/templates/header.html", "assets/templates/footer.html")
	if err != nil {
		fmt.Fprintf(res, err.Error())
		return
	}
	t.ExecuteTemplate(res, "index", nil)
}
