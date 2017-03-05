package usecases

type User struct {
	Id    int
	Email string
}

type UserRepository interface {
	Store(user User)
	FindById(id int) User
}
