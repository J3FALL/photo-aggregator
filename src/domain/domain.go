package domain

type Photographer struct {
	ID      int
	Name    string
	Surname string
	Phone   string
}

type PhotographerRepository interface {
	Store(photographer Photographer)
	FindById(id int)
}
