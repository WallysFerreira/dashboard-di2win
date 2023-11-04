package service

type Repositorio interface {
	findUser(id int) (User, error)
	findUsers(segment string) []User
	findExtract(id int) (Extract, error)
	findExtracts(filter Filtro) []Extract
}

type RepositorioPostgre struct {
	ConnStr string
}
