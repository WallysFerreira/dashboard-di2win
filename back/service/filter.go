package service

import (
	"fmt"
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
	filter := ""
	filter_count := 0
	fe_reflected := reflect.ValueOf(fe)
	fe_length := fe_reflected.NumField()

	for i := 0; i < fe_length; i++ {
		field := fe_reflected.Field(i)
		field_type := field.Type()

		if filter_count > 0 {
			filter = fmt.Sprintf("%s AND ", filter)
		}

		if field_type.Name() == "Time" && !field.IsZero() {
		}

		if field_type.Name() == "int" && field.Int() > 0 {
		}

		if field_type.Name() == "string" && field.String() != "" {
		}
	}

	return filter
}
