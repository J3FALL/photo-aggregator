package interfaces

import (
	"fmt"
	"photo-aggregator/src/domain"
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

func NewDbPhotographerRepo(dbHandlers map[string]DbHandler) *DbPhotographerRepo {
	dbPhotographerRepo := new(DbPhotographerRepo)
	dbPhotographerRepo.dbHandlers = dbHandlers
	dbPhotographerRepo.dbHandler = dbHandlers["DbPhotographerRepo"]
	return dbPhotographerRepo
}

func (repo *DbUserRepo) Store(user usecases.User) {
	repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO users (id, email)
                                      VALUES ('%d', '%s')`,
		user.Id, user.Email))
}

func (repo *DbUserRepo) FindById(id int) usecases.User {
	row := repo.dbHandler.Query(fmt.Sprintf(`SELECT email
                                           FROM users WHERE id = '%d' LIMIT 1`, id))
	var email string
	row.Next()
	row.Scan(&email)
	user := usecases.User{Id: id, Email: email}
	return user
}

func (repo *DbPhotographerRepo) Store(photographer domain.Photographer) {
	fmt.Println("from repositories")
	repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO photographers (id, name, surname, phone)
                                      VALUES ('%d', '%s', '%s', '%s')`,
		photographer.ID, photographer.Name, photographer.Surname, photographer.Phone))

	fmt.Println("repositories : good")
}

//to-do: 0 rows, what to return?
func (repo *DbPhotographerRepo) FindById(id int) domain.Photographer {
	row := repo.dbHandler.Query(fmt.Sprintf(`SELECT name, surname, phone
                                           FROM photographers WHERE id = '%d' LIMIT 1`, id))
	var (
		name    string
		surname string
		phone   string
	)
	row.Next()
	row.Scan(&name, &surname, &phone)

	//return id = -1 if 0 rows
	if name == "" {
		return domain.Photographer{ID: -1, Name: name, Surname: surname, Phone: phone}
	} else {
		return domain.Photographer{ID: id, Name: name, Surname: surname, Phone: phone}
	}
}

func (repo *DbPhotographerRepo) Update(photographer domain.Photographer) bool {
	repo.dbHandler.Execute(fmt.Sprintf(`UPDATE photographers SET name = '%s', surname = '%s', phone = '%s'
																			WHERE id = '%d'`,
		photographer.Name, photographer.Surname, photographer.Phone, photographer.ID))
	return true
}

func (repo *DbPhotographerRepo) FindAll() []domain.Photographer {
	row := repo.dbHandler.Query(fmt.Sprintf(`SELECT id, name, surname, phone
                                           FROM photographers`))
	var (
		id      int
		name    string
		surname string
		phone   string
	)

	photographers := []domain.Photographer{}
	for row.Next() {
		row.Scan(&id, &name, &surname, &phone)
		photographer := domain.Photographer{ID: id, Name: name, Surname: surname, Phone: phone}
		photographers = append(photographers, photographer)
	}
	return photographers
}
