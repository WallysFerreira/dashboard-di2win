package service

import (
	"api/graph/model"
	"testing"
)

func TestCount(t *testing.T) {
	rep := RepositorioPostgre{
		ConnStr: "user=postgres dbname=database sslmode=disable",
	}

	t.Run("count grouping by user_id", func(t *testing.T) {
		group_by := "user_id"
		expected := []model.Count{
			{
				Name:  "1",
				Value: 1058,
			},
			{
				Name:  "3",
				Value: 974,
			},
			{
				Name:  "5",
				Value: 987,
			},
			{
				Name:  "4",
				Value: 1003,
			},
			{
				Name:  "2",
				Value: 978,
			},
		}

		got := rep.CountExtracts(group_by)

		if expected != got {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})
}
