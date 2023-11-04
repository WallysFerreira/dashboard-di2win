package service

import (
	"testing"
	"time"
)

func TestExtracts(t *testing.T) {
	rep := RepositorioPostgre{
		ConnStr: "user=postgres dbname=database sslmode=disable",
	}

	t.Run("find one extract", func(t *testing.T) {
		expected := Extract{
			id:              4012,
			created_at:      time.Date(2023, 7, 17, 17, 50, 52, 363000000, time.UTC),
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

	t.Run("could not find one extract", func(t *testing.T) {
		_, err := rep.findExtract(2)

		if err == nil {
			t.Error("Expected an error and didn't get one")
		}
	})

	t.Run("find extracts using doc_type as filter", func(t *testing.T) {
		filter := "COMPROVANTE_RESIDENCIA"

		got := rep.findExtracts(time.Time{}, time.Time{}, 0, filter, 0)

		if len(got) == 0 {
			t.Errorf("Did not expect an empty slice")
		}

		for _, value := range got {
			if value.doc_type != filter {
				t.Errorf("Expected all extracts to have %s as doc_type", filter)
			}
		}
	})

	t.Run("find extracts using user_id as filter", func(t *testing.T) {
		filter := 1

		got := rep.findExtracts(time.Time{}, time.Time{}, 0, "", filter)

		if len(got) == 0 {
			t.Errorf("Did not expect an empty slice")
		}

		for _, value := range got {
			if value.user_id != filter {
				t.Errorf("Expected all extracts to have %d as user_id", filter)
			}
		}
	})

	t.Run("find extracts using dates as filter", func(t *testing.T) {
		date_start := time.Date(2023, 3, 1, 10, 0, 0, 0, time.UTC)
		date_end := time.Date(2023, 7, 17, 10, 0, 0, 0, time.UTC)

		got := rep.findExtracts(date_start, date_end, 0, "", 0)

		for _, value := range got {
			if value.created_at.Compare(date_start) == -1 {
				t.Errorf("Expected all extracts to be created after %v", date_start)
				break
			}

			if value.created_at.Compare(date_end) == 1 {
				t.Errorf("Expected all extracts to be created before %v", date_end)
				break
			}
		}
	})

	t.Run("find extracts using only start date", func(t *testing.T) {
		date_start := time.Date(2023, 10, 1, 10, 0, 0, 0, time.UTC)

		got := rep.findExtracts(date_start, time.Time{}, 0, "", 0)

		for _, value := range got {
			if value.created_at.Compare(date_start) == -1 {
				t.Errorf("Expected all extracts to be created after %v", date_start)
				break
			}
		}
	})

	t.Run("find extracts using only end date", func(t *testing.T) {
		date_end := time.Date(2023, 7, 17, 10, 0, 0, 0, time.UTC)

		got := rep.findExtracts(time.Time{}, date_end, 0, "", 0)

		for _, value := range got {
			if value.created_at.Compare(date_end) == 1 {
				t.Errorf("Expected all extracts to be created before %v", date_end)
				break
			}
		}
	})
}
