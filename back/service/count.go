package service

import (
	"api/graph/model"
	"database/sql"
	"fmt"
	"log"
)

func (rp RepositorioPostgre) CountExtracts(group_by string) []model.Count {
	result := []model.Count{}

	db, err := sql.Open("postgres", rp.ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("SELECT %s, count(*) FROM extracts GROUP BY %s", group_by, group_by))
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		count := model.Count{}

		rows.Scan(&count.Name, &count.Value)

		result = append(result, count)
	}

	return result
}
