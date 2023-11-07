package service

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestExtracts(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	conn_str := os.Getenv("TEST_POSTGRE_URL")

	rep := RepositorioPostgre{
		ConnStr: conn_str,
	}

	t.Run("find one extract", func(t *testing.T) {
		expected := Extract{
			Id:             4012,
			CreatedAt:      time.Date(2023, 7, 17, 17, 50, 52, 363000000, time.UTC),
			PagesProcessed: 1,
			DocType:        "CNH",
			UserId:         2,
		}

		got, err := rep.FindExtract(4012)
		if err != nil {
			t.Error("Got an error when it wasn't expected")
		}

		if expected != got {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})

	t.Run("could not find one extract", func(t *testing.T) {
		_, err := rep.FindExtract(2)

		if err == nil {
			t.Error("Expected an error and didn't get one")
		}
	})

	t.Run("find extracts using doc_type as filter", func(t *testing.T) {
		doc_type := "COMPROVANTE_RESIDENCIA"
		filter := FiltroExtract{
			DocType: doc_type,
		}

		got := rep.FindExtracts(filter)

		if len(got) == 0 {
			t.Errorf("Did not expect an empty slice")
		}

		for _, value := range got {
			if value.DocType != doc_type {
				t.Errorf("Expected all extracts to have %s as doc_type", doc_type)
			}
		}
	})

	t.Run("find extracts using user_id as filter", func(t *testing.T) {
		user_id := 1
		filter := FiltroExtract{
			UserId: user_id,
		}

		got := rep.FindExtracts(filter)

		if len(got) == 0 {
			t.Errorf("Did not expect an empty slice")
		}

		for _, value := range got {
			if value.UserId != user_id {
				t.Errorf("Expected all extracts to have %d as user_id", user_id)
			}
		}
	})

	t.Run("find extracts using dates as filter", func(t *testing.T) {
		date_start := time.Date(2023, 3, 1, 10, 0, 0, 0, time.UTC)
		date_end := time.Date(2023, 7, 17, 10, 0, 0, 0, time.UTC)
		filter := FiltroExtract{
			DataStart: date_start,
			DataEnd:   date_end,
		}

		got := rep.FindExtracts(filter)

		for _, value := range got {
			if value.CreatedAt.Compare(date_start) == -1 {
				t.Errorf("Expected all extracts to be created after %v", date_start)
				break
			}

			if value.CreatedAt.Compare(date_end) == 1 {
				t.Errorf("Expected all extracts to be created before %v", date_end)
				break
			}
		}
	})

	t.Run("find extracts using only start date", func(t *testing.T) {
		date_start := time.Date(2023, 10, 1, 10, 0, 0, 0, time.UTC)
		filter := FiltroExtract{
			DataStart: date_start,
		}

		got := rep.FindExtracts(filter)

		for _, value := range got {
			if value.CreatedAt.Compare(date_start) == -1 {
				t.Errorf("Expected all extracts to be created after %v", date_start)
				break
			}
		}
	})

	t.Run("find extracts using only end date", func(t *testing.T) {
		date_end := time.Date(2023, 7, 17, 10, 0, 0, 0, time.UTC)
		filter := FiltroExtract{
			DataEnd: date_end,
		}

		got := rep.FindExtracts(filter)

		for _, value := range got {
			if value.CreatedAt.Compare(date_end) == 1 {
				t.Errorf("Expected all extracts to be created before %v", date_end)
				break
			}
		}
	})

	t.Run("find extracts with multiple filters", func(t *testing.T) {
		data_start := time.Date(2023, 7, 17, 0, 0, 0, 0, time.UTC)
		doc_type := "CAPA_SERASA"
		pages_processed := 3
		filter := FiltroExtract{
			DataStart:      data_start,
			DocType:        doc_type,
			PagesProcessed: pages_processed,
		}

		got := rep.FindExtracts(filter)

		for _, value := range got {
			if value.CreatedAt.Compare(data_start) == -1 {
				t.Errorf("Expected all extracts to be created after %v", data_start)
			}

			if value.DocType != doc_type {
				t.Errorf("Expected all extracts to have %s as doc_type", doc_type)
			}

			if value.PagesProcessed != pages_processed {
				t.Errorf("Expected all extracts to have %d as pages_process", pages_processed)
			}
		}
	})
}
