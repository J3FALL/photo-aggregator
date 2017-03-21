package infrastructure

import (
	"database/sql"
	"fmt"
	"photo-aggregator/src/interfaces"

	_ "github.com/lib/pq"
)

type PgHandler struct {
	Conn *sql.DB
}

func (handler *PgHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *PgHandler) Query(statement string) interfaces.Row {
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		fmt.Println(err)
		return new(PgRow)
	}
	row := new(PgRow)
	row.Rows = rows
	return row
}

type PgRow struct {
	Rows *sql.Rows
}

func (r PgRow) Scan(dest ...interface{}) {
	r.Rows.Scan(dest...)
}

func (r PgRow) Next() bool {
	return r.Rows.Next()
}

func NewPgHandler(dbUrl string) *PgHandler {
	conn, _ := sql.Open("postgres", dbUrl)
	pgHandler := new(PgHandler)
	pgHandler.Conn = conn
	return pgHandler
}
