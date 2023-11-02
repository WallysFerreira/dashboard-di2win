package service

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	id      int
	name    string
	segment string
}

func (rp *RepositorioPostgre) findUser(id int) (User, error) {
	user := User{}

	db, err := sql.Open("postgres", rp.connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, segment FROM users WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		rows.Scan(&user.id, &user.name, &user.segment)
	}

	if user.id == 0 {
		return User{}, errors.New("could not find user")
	}

	return user, nil
}