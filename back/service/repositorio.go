package service

import "api/graph/model"

type Repositorio interface {
	FindUser(id int) (User, error)
	FindUsers(segment string) []User
	FindExtract(id int) (Extract, error)
	FindExtracts(filter Filtro) ([]Extract, int)
	CountExtracts(group_by string) []model.Count
}

type RepositorioPostgre struct {
	ConnStr string
}
