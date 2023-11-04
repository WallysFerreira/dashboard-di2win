package service

import (
	"testing"
	"time"
)

func TestFilter(t *testing.T) {
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
}