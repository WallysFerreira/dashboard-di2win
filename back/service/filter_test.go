package service

import (
	"testing"
	"time"
)

func TestFilter(t *testing.T) {
	t.Run("generate empty filter", func(t *testing.T) {
		filtro := FiltroExtract{}

		expected := ""

		got := filtro.gerarFiltro()

		if expected != got {
			t.Errorf("Expected '%s', got '%s'", expected, got)
		}
	})

	t.Run("generate filter with int", func(t *testing.T) {
		filtro := FiltroExtract{}
		filtro.UserId = 8

		expected := "WHERE user_id = 8"

		got := filtro.gerarFiltro()

		if expected != got {
			t.Errorf("Expected '%s', got '%s'", expected, got)
		}
	})

	t.Run("generate filter with two dates", func(t *testing.T) {
		filtro := FiltroExtract{}
		filtro.DataStart = time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
		filtro.DataEnd = time.Date(2023, 10, 23, 0, 0, 0, 0, time.UTC)

		expected := "WHERE created_at > '2023-10-1' AND created_at < '2023-10-23'"

		got := filtro.gerarFiltro()

		if expected != got {
			t.Errorf("Expected '%s', got '%s'", expected, got)
		}
	})

	t.Run("generate filter with one date", func(t *testing.T) {
		filtro := FiltroExtract{}
		filtro.DataEnd = time.Date(2022, 9, 22, 0, 0, 0, 0, time.UTC)

		expected := "WHERE created_at < '2022-9-22'"

		got := filtro.gerarFiltro()

		if expected != got {
			t.Errorf("Expected '%s', got '%s'", expected, got)
		}
	})

	t.Run("generate filter with string", func(t *testing.T) {
		filtro := FiltroExtract{}
		filtro.DocType = "CNH"

		expected := "WHERE doc_type = 'CNH'"

		got := filtro.gerarFiltro()

		if expected != got {
			t.Errorf("Expected '%s', got '%s'", expected, got)
		}
	})

	t.Run("generate filter with multiple fields", func(t *testing.T) {
		filter := FiltroExtract{
			DataStart: time.Date(2023, 7, 12, 0, 0, 0, 0, time.UTC),
			UserId:    2,
			DocType:   "FATURA_ENERGIA",
		}

		expected := "WHERE created_at > '2023-7-12' AND doc_type = 'FATURA_ENERGIA' AND user_id = 2"

		got := filter.gerarFiltro()

		if expected != got {
			t.Errorf("Expected '%s', got '%s'", expected, got)
		}
	})
}
