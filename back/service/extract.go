package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type Extract struct {
	Id             int
	CreatedAt      time.Time
	PagesProcessed int
	DocType        string
	UserId         int
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
		rows.Scan(&found.Id, &unparsed_date, &found.PagesProcessed, &found.DocType, &found.UserId)
	}

	if found.Id == 0 {
		return found, errors.New("could not find extract")
	} else {
		found.CreatedAt, err = time.Parse(time.RFC3339, unparsed_date)
		if err != nil {
			log.Fatal(err)
		}
	}

	return found, nil
}

func (rp RepositorioPostgre) FindExtracts(filter Filtro) []Extract {
	found_extracts := []Extract{}
	filter_string := filter.gerarFiltro()

	db, err := sql.Open("postgres", rp.ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM extracts %s ORDER BY id", filter_string))
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		found_extract := Extract{}

		rows.Scan(&found_extract.Id, &found_extract.CreatedAt, &found_extract.PagesProcessed, &found_extract.DocType, &found_extract.UserId)

		found_extracts = append(found_extracts, found_extract)
	}

	return found_extracts
}
