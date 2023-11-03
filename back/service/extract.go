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

func (rp *RepositorioPostgre) findExtracts(date_start time.Time, date_end time.Time, pages_processed int, doc_type string, user_id int) []Extract {
	found_extracts := []Extract{}
	filter := ""
	year_start, month_start, day_start := date_start.Date()
	year_end, month_end, day_end := date_end.Date()

	db, err := sql.Open("postgres", rp.connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if !date_start.IsZero() && !date_end.IsZero() {
		start_filter := fmt.Sprintf("%d-%d-%d", year_start, month_start, day_start)
		end_filter := fmt.Sprintf("%d-%d-%d", year_end, month_end, day_end)

		filter = fmt.Sprintf("WHERE created_at > '%s' AND created_at < '%s'", start_filter, end_filter)
	} else if !date_start.IsZero() {
		start_filter := fmt.Sprintf("%d-%d-%d", year_start, month_start, day_start)

		filter = fmt.Sprintf("WHERE created_at > '%s'", start_filter)
	} else if pages_processed > 0 {
		filter = fmt.Sprintf("WHERE pages_processed = %d", pages_processed)
	} else if doc_type != "" {
		filter = fmt.Sprintf("WHERE doc_type LIKE '%s'", doc_type)
	} else if user_id > 0 {
		filter = fmt.Sprintf("WHERE user_id =  %d", user_id)
	}

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM extracts %s", filter))
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		found_extract := Extract{}

		rows.Scan(&found_extract.id, &found_extract.created_at, &found_extract.pages_processed, &found_extract.doc_type, &found_extract.user_id)

		found_extracts = append(found_extracts, found_extract)
	}

	return found_extracts
}
