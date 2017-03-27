package usecases

import (
	"fmt"
	"photo-aggregator/src/domain"
)

type User struct {
	Id    int
	Email string
}

type Photographer struct {
	ID      int
	Name    string
	Surname string
	Phone   string
}

type UserRepository interface {
	Store(user User)
	FindById(id int) User
}

type PhotoInteractor struct {
	UserRepository         UserRepository
	PhotographerRepository domain.PhotographerRepository
	TagRepository          domain.TagRepository
	AttachmentRepository   domain.AttachmentRepository
}

func (interactor *PhotoInteractor) Photographers() ([]Photographer, error) {
	/*var photographers []Photographer
	photographers = make([]Photographer, 1)
	photographers[0] = Photographer{1, "ivan", "petrov", "+79992134567"}*/
	photographersTmp := interactor.PhotographerRepository.FindAll()

	photographers := []Photographer{}
	for _, photographer := range photographersTmp {
		ph := Photographer{ID: photographer.ID, Name: photographer.Name, Surname: photographer.Surname, Phone: photographer.Phone}
		photographers = append(photographers, ph)
	}
	return photographers, nil
}

func (interactor *PhotoInteractor) Photographer(id int) (Photographer, error) {
	photographerTmp := interactor.PhotographerRepository.FindById(id)
	fmt.Println(id)
	photographer := Photographer{ID: photographerTmp.ID, Name: photographerTmp.Name, Surname: photographerTmp.Surname, Phone: photographerTmp.Phone}
	return photographer, nil
}

func (interactor *PhotoInteractor) NewPhotographer(photographer Photographer) {
	fmt.Println("from usecases")
	photographerToStore := domain.Photographer{ID: photographer.ID, Name: photographer.Name, Surname: photographer.Surname, Phone: photographer.Phone}
	interactor.PhotographerRepository.Store(photographerToStore)
	fmt.Println("usecases : good")
}

func (interactor *PhotoInteractor) UpdatePhotographer(photographer Photographer) bool {
	photographerToUpdate := domain.Photographer{ID: photographer.ID, Name: photographer.Name, Surname: photographer.Surname, Phone: photographer.Phone}
	return interactor.PhotographerRepository.Update(photographerToUpdate)
}
