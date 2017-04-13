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

func NewDbTagRepo(dbHandlers map[string]DbHandler) *DbTagRepo {
	dbTagRepo := new(DbTagRepo)
	dbTagRepo.dbHandlers = dbHandlers
	dbTagRepo.dbHandler = dbHandlers["DbTagRepo"]
	return dbTagRepo
}

func NewDbAttachmentRepo(dbHandlers map[string]DbHandler) *DbAttachmentRepo {
	dbAttachmentRepo := new(DbAttachmentRepo)
	dbAttachmentRepo.dbHandlers = dbHandlers
	dbAttachmentRepo.dbHandler = dbHandlers["DbAttachmentRepo"]
	return dbAttachmentRepo
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
	repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO photographers (id, name, surname, description, sub_description, vk_url, instagram_url)
                                      VALUES (nextval('photographers_seq'), '%s', '%s', '%s', '%s', '%s', '%s')`,
		photographer.Name, photographer.Surname, photographer.Description, photographer.SubDescription, photographer.VkURL, photographer.InstagramURL))

	fmt.Println("repositories : good")
}

//to-do: 0 rows, what to return?
func (repo *DbPhotographerRepo) FindById(id int) domain.Photographer {
	row := repo.dbHandler.Query(fmt.Sprintf(`SELECT name, surname, description, sub_description, vk_url, instagram_url
                                           FROM photographers WHERE id = '%d' LIMIT 1`, id))
	var (
		name           string
		surname        string
		description    string
		subDescription string
		vkUrl          string
		instagramUrl   string
	)
	row.Next()
	row.Scan(&name, &surname, &description, &subDescription, &vkUrl, &instagramUrl)

	//return id = -1 if 0 rows
	if name == "" {
		return domain.Photographer{ID: -1, Name: name, Surname: surname, Description: description, SubDescription: subDescription, VkURL: vkUrl, InstagramURL: instagramUrl}
	} else {
		return domain.Photographer{ID: id, Name: name, Surname: surname, Description: description, SubDescription: subDescription, VkURL: vkUrl, InstagramURL: instagramUrl}
	}
}

func (repo *DbPhotographerRepo) Update(photographer domain.Photographer) bool {
	repo.dbHandler.Execute(fmt.Sprintf(`UPDATE photographers SET name = '%s', surname = '%s', description = '%s', sub_description = '%s', vk_url = '%s', instagram_url = '%s'
																			WHERE id = '%d'`,
		photographer.Name, photographer.Surname, photographer.Description, photographer.SubDescription, photographer.VkURL, photographer.InstagramURL, photographer.ID))
	return true
}

func (repo *DbPhotographerRepo) FindAll() []domain.Photographer {
	row := repo.dbHandler.Query(fmt.Sprintf(`SELECT id, name, surname, description, sub_description, vk_url, instagram_url
                                           FROM photographers`))
	var (
		id             int
		name           string
		surname        string
		description    string
		subDescription string
		vkUrl          string
		instagramUrl   string
	)

	photographers := []domain.Photographer{}
	for row.Next() {
		row.Scan(&id, &name, &surname, &description, &subDescription, &vkUrl, &instagramUrl)
		fmt.Println(id)
		photographer := domain.Photographer{ID: id, Name: name, Surname: surname, Description: description, SubDescription: subDescription, VkURL: vkUrl, InstagramURL: instagramUrl}
		photographers = append(photographers, photographer)
	}
	fmt.Println(photographers)
	return photographers
}

func (repo *DbTagRepo) Store(tag domain.Tag) {
	repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO tags (id, name)
                                      VALUES (nextval('tags_seq'), '%s')`,
		tag.Name))
}

func (repo *DbTagRepo) FindById(id int) domain.Tag {
	fmt.Println("findById")
	row := repo.dbHandler.Query(fmt.Sprintf(`SELECT name
                                           FROM tags WHERE id = '%d' LIMIT 1`, id))
	var name string

	fmt.Println("1")
	row.Next()
	fmt.Println("2")
	row.Scan(&name)

	fmt.Println("from repos ", name)
	//return id = -1 if 0 rows
	if name == "" {
		return domain.Tag{ID: -1, Name: name}
	} else {
		return domain.Tag{ID: id, Name: name}
	}
}

func (repo *DbTagRepo) Update(tag domain.Tag) bool {
	repo.dbHandler.Execute(fmt.Sprintf(`UPDATE tags SET name = '%s'
																			WHERE id = '%d'`,
		tag.Name, tag.ID))
	return true
}

func (repo *DbTagRepo) FindAll() []domain.Tag {
	row := repo.dbHandler.Query(fmt.Sprintf(`SELECT id, name
                                           FROM tags`))
	var (
		id   int
		name string
	)

	tags := []domain.Tag{}
	for row.Next() {
		row.Scan(&id, &name)
		tag := domain.Tag{ID: id, Name: name}
		tags = append(tags, tag)
	}
	return tags
}

func (repo *DbAttachmentRepo) Store(attach domain.Attachment) {
	repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO attachments (id, description, url)
                                      VALUES (nextval('attachments_seq'), '%s', '%s')`,
		attach.Description, attach.Url))
}

func (repo *DbAttachmentRepo) FindById(id int) domain.Attachment {
	row := repo.dbHandler.Query(fmt.Sprintf(`SELECT description, url
                                           FROM attachments WHERE id = '%d' LIMIT 1`, id))
	var (
		description string
		url         string
	)
	row.Next()
	row.Scan(&description, &url)
	fmt.Println("!", description)
	//return id = -1 if 0 rows
	if url == "" {
		return domain.Attachment{ID: -1, Description: description, Url: url}
	} else {
		return domain.Attachment{ID: id, Description: description, Url: url}
	}
}

func (repo *DbAttachmentRepo) Update(attachment domain.Attachment) bool {
	repo.dbHandler.Execute(fmt.Sprintf(`UPDATE attachments SET description = '%s', url = '%s'
																			WHERE id = '%d'`,
		attachment.Description, attachment.Url, attachment.ID))
	return true
}

func (repo *DbAttachmentRepo) FindAll() []domain.Attachment {
	row := repo.dbHandler.Query(fmt.Sprintf(`SELECT id, description, url
                                           FROM attachments`))
	var (
		id          int
		description string
		url         string
	)

	attachments := []domain.Attachment{}
	for row.Next() {
		row.Scan(&id, &description, &url)
		attach := domain.Attachment{ID: id, Description: description, Url: url}
		attachments = append(attachments, attach)
	}
	return attachments
}
