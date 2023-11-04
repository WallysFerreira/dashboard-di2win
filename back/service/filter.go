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
	DataEnd        time.Time
	PagesProcessed int
	DocType        string
	UserId         int
}

func (fe FiltroExtract) gerarFiltro() string {
	filter := ""
	filtered := false
	filter_count := 0
	fe_reflected := reflect.ValueOf(fe)
	fe_length := fe_reflected.NumField()

	for i := 0; i < fe_length; i++ {
		field := fe_reflected.Field(i)
		field_name := fe_reflected.Type().Field(i).Name
		field_type := field.Type()

		if field_type.Name() == "Time" && !field.IsZero() {
			filter = checarContagem(filter_count, filter)

			switch field_name {
			case "DataStart":
				year, month, day := fe.DataStart.Date()
				filter = fmt.Sprintf("%s created_at > '%d-%d-%d'", filter, year, month, day)
			case "DataEnd":
				year, month, day := fe.DataEnd.Date()
				filter = fmt.Sprintf("%s created_at < '%d-%d-%d'", filter, year, month, day)
			}

			filter_count++
			filtered = true
		}

		if field_type.Name() == "int" && field.Int() > 0 {
			filter = checarContagem(filter_count, filter)

			switch field_name {
			case "PagesProcessed":
				filter = fmt.Sprintf("%s pages_process = %d", filter, field.Int())
			case "UserId":
				filter = fmt.Sprintf("%s user_id = %d", filter, field.Int())
			}

			filter_count++
			filtered = true
		}

		if field_type.Name() == "string" && field.String() != "" {
			filter = checarContagem(filter_count, filter)

			filter_count++
			filtered = true
		}
	}

	if filtered {
		filter = fmt.Sprintf("WHERE%s", filter)
	}

	return filter
}

func checarContagem(count int, filter string) string {
	if count > 0 {
		return fmt.Sprintf("%s AND", filter)
	}

	return ""
}
