package service

type Repositorio interface {
	findUser(id int) (User, error)
	findUsers() []User
}

type RepositorioPostgre struct {
	connStr string
}
