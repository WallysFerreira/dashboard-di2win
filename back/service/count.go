package service

import (
	"api/graph/model"
	"database/sql"
	"fmt"
	"log"
)

func (rp RepositorioPostgre) CountExtracts(group_by string, filter Filtro) []*model.Count {
	result := []*model.Count{}

	db, err := sql.Open("postgres", rp.ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	filter_string := filter.gerarFiltro()

	rows, err := db.Query(fmt.Sprintf("SELECT %s, sum(extracts.pages_process) FROM extracts JOIN users on users.id = extracts.user_id %s GROUP BY %s ORDER BY sum(pages_process) DESC", group_by, filter_string, group_by))

	if group_by == "EXTRACT(month FROM created_at::date)" {
		rows, err = db.Query(fmt.Sprintf("SELECT %s, sum(extracts.pages_process) FROM extracts JOIN users on users.id = extracts.user_id %s GROUP BY %s ORDER BY %s", group_by, filter_string, group_by, group_by))
	}

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		count := model.Count{}

		rows.Scan(&count.Name, &count.Value)

		result = append(result, &count)
	}

	return result
}
