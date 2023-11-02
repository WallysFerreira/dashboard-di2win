package service

type Repositorio interface {
	findUser(id int) (User, error)
	findUsers(segment string) []User
}

type RepositorioPostgre struct {
	connStr string
}
