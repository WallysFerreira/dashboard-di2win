package service

import (
	"reflect"
	"testing"
	"time"
)

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

	t.Run("find all users without filter", func(t *testing.T) {
		expected := []User{
			{
				id:      1,
				name:    "Rômulo",
				segment: "construtora",
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
			{
				id:      2,
				name:    "Madeira",
				segment: "banco",
			},
		}
		got := rep.findUsers("")

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})

	t.Run("find all users using filter", func(t *testing.T) {
		expected := []User{
			{
				id:      4,
				name:    "Augusto",
				segment: "banco",
			},
			{
				id:      2,
				name:    "Madeira",
				segment: "banco",
			},
		}

		got := rep.findUsers("banco")

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})
}

func TestExtracts(t *testing.T) {
	rep := RepositorioPostgre{
		connStr: "user=postgres dbname=database sslmode=disable",
	}

	t.Run("find one extract", func(t *testing.T) {
		expected := Extract{
			id:              4012,
			created_at:      time.Date(2023, 7, 17, 17, 50, 52, 0, time.Local),
			pages_processed: 1,
			doc_type:        "CNH",
			user_id:         2,
		}

		got, err := rep.findExtract(4012)
		if err != nil {
			t.Error("Got an error when it wasn't expected")
		}

		if expected != got {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})
}
