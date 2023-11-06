package service

type Repositorio interface {
	FindUser(id int) (User, error)
	FindUsers(segment string) []User
	FindExtract(id int) (Extract, error)
	FindExtracts(filter Filtro) []Extract
}

type RepositorioPostgre struct {
	ConnStr string
}
