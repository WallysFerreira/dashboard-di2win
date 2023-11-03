package service

type Repositorio interface {
	findUser(id int) (User, error)
	findUsers(segment string) []User
	findExtract(id int) (Extract, error)
}

type RepositorioPostgre struct {
	connStr string
}
