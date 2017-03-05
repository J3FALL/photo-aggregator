package interfaces

import (
	"fmt"
	"photo-aggregator/src/usecases"
)

type DbHandler interface {
	Execute(statement string)
	Query(statement string) Row
}

type Row interface {
	Scan(dest ...interface{})
	Next() bool
}

type DbRepo struct {
	dbHandlers map[string]DbHandler
	dbHandler  DbHandler
}

type DbPhotographerRepo DbRepo
type DbTagRepo DbRepo
type DbAttachmentRepo DbRepo
type DbUserRepo DbRepo

func NewDbUserRepo(dbHandlers map[string]DbHandler) *DbUserRepo {
	dbUserRepo := new(DbUserRepo)
	dbUserRepo.dbHandlers = dbHandlers
	dbUserRepo.dbHandler = dbHandlers["DbUserRepo"]
	return dbUserRepo
}

func (repo *DbUserRepo) Store(user usecases.User) {
	repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO users (id, email)
                                      VALUES ('%d', '%s')`,
		user.Id, user.Email))
}

func (repo *DbUserRepo) FindById(id int) usecases.User {
	row := repo.dbHandler.Query(fmt.Sprintf(`SELECT email
                                           FROM users WHERE id = '%d' LIMIT 1`,
		id))
	var email string
	row.Next()
	row.Scan(&email)
	user := usecases.User{Id: id, Email: email}
	return user
}
