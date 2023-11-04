package service

import "time"

type Filtro interface {
	gerarFiltro() string
}

type FiltroExtract struct {
	DataStart      time.Time
	DateEnd        time.Time
	PagesProcessed int
	DocType        string
	UserId         int
}
