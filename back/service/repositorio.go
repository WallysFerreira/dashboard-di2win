package service

type Repositorio interface {
	findUser(id int) (User, error)
	findUsers() ([]User, error)
}

type RepositorioPostgre struct {
	connStr string
}
