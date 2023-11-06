package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type Extract struct {
	id              int
	created_at      time.Time
	pages_processed int
	doc_type        string
	user_id         int
}

func (rp RepositorioPostgre) FindExtract(id int) (Extract, error) {
	found := Extract{}
	var unparsed_date string

	db, err := sql.Open("postgres", rp.ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM extracts WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		rows.Scan(&found.id, &unparsed_date, &found.pages_processed, &found.doc_type, &found.user_id)
	}

	if found.id == 0 {
		return found, errors.New("could not find extract")
	} else {
		found.created_at, err = time.Parse(time.RFC3339, unparsed_date)
		if err != nil {
			log.Fatal(err)
		}
	}

	return found, nil
}

func (rp RepositorioPostgre) FindExtracts(filter Filtro) ([]Extract, int) {
	found_extracts := []Extract{}
	filter_string := filter.gerarFiltro()

	db, err := sql.Open("postgres", rp.ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM extracts %s", filter_string))
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		found_extract := Extract{}

		rows.Scan(&found_extract.id, &found_extract.created_at, &found_extract.pages_processed, &found_extract.doc_type, &found_extract.user_id)

		found_extracts = append(found_extracts, found_extract)
	}

	rows, err = db.Query(fmt.Sprintf("SELECT count(*) FROM extracts %s", filter_string))
	if err != nil {
		log.Fatal(err)
	}

	var count int
	for rows.Next() {
		rows.Scan(&count)
	}

	return found_extracts, count
}
