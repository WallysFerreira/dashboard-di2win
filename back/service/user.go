package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	Id      int
	Name    string
	Segment string
}

func (rp RepositorioPostgre) FindUser(id int) (User, error) {
	user := User{}

	db, err := sql.Open("postgres", rp.ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, segment FROM users WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Segment)
	}

	if user.Id == 0 {
		return User{}, errors.New("could not find user")
	}

	return user, nil
}

func (rp RepositorioPostgre) FindUsers(segment string) []User {
	filter := ""
	var result []User

	db, err := sql.Open("postgres", rp.ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if segment != "" {
		filter = fmt.Sprintf("WHERE segment LIKE '%s'", segment)
	}

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM users %s ORDER BY id", filter))
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		user := User{}

		rows.Scan(&user.Id, &user.Name, &user.Segment)

		result = append(result, user)
	}

	return result
}
