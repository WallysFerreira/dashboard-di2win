package service

import (
	"reflect"
	"testing"
)

func TestUser(t *testing.T) {
	rep := RepositorioPostgre{
		ConnStr: "user=postgres dbname=database sslmode=disable",
	}

	t.Run("find one user", func(t *testing.T) {
		expected := User{Id: 3, Name: "Caio", Segment: "imobiliaria"}
		got, err := rep.FindUser(3)

		if err != nil {
			t.Errorf(err.Error())
		}

		if expected != got {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})

	t.Run("could not find one user", func(t *testing.T) {
		_, err := rep.FindUser(10)

		if err == nil {
			t.Error("Should've failed but didn't")
		}
	})

	t.Run("find all users without filter", func(t *testing.T) {
		expected := []User{
			{
				Id:      1,
				Name:    "Rômulo",
				Segment: "construtora",
			},
			{
				Id:      3,
				Name:    "Caio",
				Segment: "imobiliaria",
			},
			{
				Id:      4,
				Name:    "Augusto",
				Segment: "banco",
			},
			{
				Id:      5,
				Name:    "Inoa",
				Segment: "financeira",
			},
			{
				Id:      2,
				Name:    "Madeira",
				Segment: "banco",
			},
		}
		got := rep.FindUsers("")

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})

	t.Run("find all users using filter", func(t *testing.T) {
		expected := []User{
			{
				Id:      4,
				Name:    "Augusto",
				Segment: "banco",
			},
			{
				Id:      2,
				Name:    "Madeira",
				Segment: "banco",
			},
		}

		got := rep.FindUsers("banco")

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})
}
