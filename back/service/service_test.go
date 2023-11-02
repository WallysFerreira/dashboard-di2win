package service

import "testing"

func TestUser(t *testing.T) {
	rep := RepositorioPostgre{
		connStr: "user=postgres dbname=database sslmode=disable",
	}

	t.Run("find one user", func(t *testing.T) {
		expected := User{id: 3, name: "Caio", segment: "imobiliaria"}
		got, err := rep.findUser(3)

		if err != nil {
			t.Errorf(err.Error())
		}

		if expected != got {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})

	t.Run("could not find one user", func(t *testing.T) {
		_, err := rep.findUser(10)

		if err == nil {
			t.Error("Should've failed but didn't")
		}
	})

	t.Run("find all users", func(t *testing.T) {
		expected := []User{
			{
				id:      1,
				name:    "Rômulo",
				segment: "construtora",
			},
			{
				id:      2,
				name:    "Madeira",
				segment: "banco",
			},
			{
				id:      3,
				name:    "Caio",
				segment: "imobiliaria",
			},
			{
				id:      4,
				name:    "Augusto",
				segment: "banco",
			},
			{
				id:      5,
				name:    "Inoa",
				segment: "financeira",
			},
		}
		got := rep.findUsers()

		if expected != got {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})
}
