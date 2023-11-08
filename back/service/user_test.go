package service

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/joho/godotenv"
)

func TestUser(t *testing.T) {
	conn_str := os.Getenv("TEST_POSTGRE_URL")

	if conn_str == "" {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatal(err)
		}

		conn_str = os.Getenv("TEST_POSTGRE_URL")
	}

	rep := RepositorioPostgre{
		ConnStr: conn_str,
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
				Id:      2,
				Name:    "Madeira",
				Segment: "banco",
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
		}
		got := rep.FindUsers("")

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})

	t.Run("find all users using filter", func(t *testing.T) {
		expected := []User{
			{
				Id:      2,
				Name:    "Madeira",
				Segment: "banco",
			},
			{
				Id:      4,
				Name:    "Augusto",
				Segment: "banco",
			},
		}

		got := rep.FindUsers("banco")

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})
}
