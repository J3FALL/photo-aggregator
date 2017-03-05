package usecases

import (
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
	var photographers []Photographer
	photographers = make([]Photographer, 1)
	photographers[0] = Photographer{1, "ivan", "petrov", "+79992134567"}

	return photographers, nil
}
