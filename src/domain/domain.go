package domain

type Photographer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Phone   string `json:"phone"`
}

type Tag struct {
	ID   int
	Name string
}

type Attachment struct {
	ID          int
	Description string
}

type PhotographerRepository interface {
	Store(photographer Photographer)
	FindById(id int) Photographer
	FindAll() []Photographer
}

type TagRepository interface {
	Store(tag Tag)
	FindById(id int)
}

type AttachmentRepository interface {
	Store(attachment Attachment)
	FindById(id int)
}
