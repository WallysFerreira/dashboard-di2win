package service

import "api/graph/model"

type Repositorio interface {
	FindUser(id int) (User, error)
	FindUsers(segment string) []User
	FindExtract(id int) (Extract, error)
	FindExtracts(filter Filtro) []Extract
	CountExtracts(what_to_count int, group_by string, filter Filtro) []model.Count
}

type RepositorioPostgre struct {
	ConnStr string
}
