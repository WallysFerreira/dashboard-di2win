package service

import (
	"database/sql"
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

func (rp *RepositorioPostgre) findExtract(id int) (Extract, error) {
	found := Extract{}
	var unparsed_date string

	db, err := sql.Open("postgres", rp.connStr)
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

	found.created_at, err = time.Parse(time.RFC3339, unparsed_date)
	if err != nil {
		log.Fatal(err)
	}

	return found, nil
}
