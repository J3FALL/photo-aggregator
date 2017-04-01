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
	ID           int
	Name         string
	Surname      string
	VkURL        string
	InstagramURL string
}

type Tag struct {
	ID   int
	Name string
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
		ph := Photographer{ID: photographer.ID, Name: photographer.Name, Surname: photographer.Surname, VkURL: photographer.VkURL, InstagramURL: photographer.InstagramURL}
		photographers = append(photographers, ph)
	}
	return photographers, nil
}

func (interactor *PhotoInteractor) Photographer(id int) (Photographer, error) {
	photographerTmp := interactor.PhotographerRepository.FindById(id)
	fmt.Println(id)
	photographer := Photographer{ID: photographerTmp.ID, Name: photographerTmp.Name, Surname: photographerTmp.Surname, VkURL: photographerTmp.VkURL, InstagramURL: photographerTmp.InstagramURL}
	return photographer, nil
}

func (interactor *PhotoInteractor) NewPhotographer(photographer Photographer) {
	fmt.Println("from usecases")
	photographerToStore := domain.Photographer{ID: photographer.ID, Name: photographer.Name, Surname: photographer.Surname, VkURL: photographer.VkURL, InstagramURL: photographer.InstagramURL}
	interactor.PhotographerRepository.Store(photographerToStore)
	fmt.Println("usecases : good")
}

func (interactor *PhotoInteractor) UpdatePhotographer(photographer Photographer) bool {
	photographerToUpdate := domain.Photographer{ID: photographer.ID, Name: photographer.Name, Surname: photographer.Surname, VkURL: photographer.VkURL, InstagramURL: photographer.InstagramURL}
	interactor.PhotographerRepository.Update(photographerToUpdate)
	return true
}

func (interactor *PhotoInteractor) NewTag(tag Tag) {
	tagToStore := domain.Tag{ID: tag.ID, Name: tag.Name}
	fmt.Println(tagToStore)
	interactor.TagRepository.Store(tagToStore)
}

func (interactor *PhotoInteractor) Tag(id int) (Tag, error) {
	fmt.Println("from usecases")
	tagTmp := interactor.TagRepository.FindById(id)
	tag := Tag{ID: tagTmp.ID, Name: tagTmp.Name}
	return tag, nil
}

func (interactor *PhotoInteractor) UpdateTag(tag Tag) bool {
	tagToUpdate := domain.Tag{ID: tag.ID, Name: tag.Name}
	interactor.TagRepository.Update(tagToUpdate)
	return true
}

func (interactor *PhotoInteractor) Tags() ([]Tag, error) {
	tagsTmp := interactor.TagRepository.FindAll()
	tags := []Tag{}
	for _, tag := range tagsTmp {
		tg := Tag{ID: tag.ID, Name: tag.Name}
		tags = append(tags, tg)
	}
	return tags, nil
}
