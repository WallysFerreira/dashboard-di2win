package service

import (
	"reflect"
	"time"
)

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

func (fe FiltroExtract) gerarFiltro() string {
	fe_reflected := reflect.ValueOf(fe)
	fe_length := fe_reflected.NumField()

	for i := 0; i < fe_length; i++ {
		field_type := fe_reflected.Field(i).Type()

		if field_type.Name() == "Time" {
		}

		if field_type.Name() == "int" {
		}

		if field_type.Name() == "string" {
		}
	}

	return ""
}
