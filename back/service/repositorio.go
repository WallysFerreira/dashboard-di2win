package service

import "time"

type Repositorio interface {
	findUser(id int) (User, error)
	findUsers(segment string) []User
	findExtract(id int) (Extract, error)
	findExtracts(date_start time.Time, date_end time.Time, created_at string, pages_processed int, doc_type string, user_id int) []Extract
}

type RepositorioPostgre struct {
	connStr string
}
