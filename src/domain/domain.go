package domain

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

type Attachment struct {
	ID          int
	Description string
}

type PhotographerRepository interface {
	Store(photographer Photographer)
	FindById(id int) Photographer
	Update(photographer Photographer) bool
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
