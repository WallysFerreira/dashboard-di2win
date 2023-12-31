package service

import (
	"api/graph/model"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/joho/godotenv"
)

func TestCount(t *testing.T) {
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

	t.Run("count grouping by user_id", func(t *testing.T) {
		group_by := "user_id"
		expected_pages_processed := []*model.Count{
			{Name: "1", Value: 3649},
			{Name: "5", Value: 3381},
			{Name: "4", Value: 3144},
			{Name: "3", Value: 3041},
			{Name: "2", Value: 2894},
		}
		expected_rows_count := []*model.Count{
			{Name: "1", Value: 1058},
			{Name: "4", Value: 1003},
			{Name: "5", Value: 987},
			{Name: "2", Value: 978},
			{Name: "3", Value: 974},
		}

		got_pages_processed := rep.CountExtracts(false, group_by, FiltroExtract{})
		got_rows_count := rep.CountExtracts(true, group_by, FiltroExtract{})

		if !reflect.DeepEqual(expected_pages_processed, got_pages_processed) {
			t.Errorf("Pages processed: Expected %v, got %v", expected_pages_processed, got_pages_processed)
		}

		if !reflect.DeepEqual(expected_rows_count, got_rows_count) {
			t.Errorf("Rows counted: Expected %v, got %v", expected_rows_count, got_rows_count)
		}
	})

	t.Run("count where user_id is 5 grouping by document", func(t *testing.T) {
		group_by := "doc_type"
		filter := FiltroExtract{
			UserId: 5,
		}

		expected_pages_processed := []*model.Count{
			{Name: "POSICAO_CONSOLIDADA", Value: 651},
			{Name: "BALANCO_PATRIMONIAL", Value: 603},
			{Name: "DECLARACAO_IR", Value: 432},
			{Name: "ENDIVIDAMENTO", Value: 415},
			{Name: "CAPA_SERASA", Value: 288},
			{Name: "COMPROVANTE_RESIDENCIA", Value: 243},
			{Name: "CONTRATO_SOCIAL", Value: 179},
			{Name: "FATURAMENTO", Value: 165},
			{Name: "CNH", Value: 155},
			{Name: "FATURA_ENERGIA", Value: 147},
			{Name: "RECIBOS", Value: 103},
		}
		expected_rows_count := []*model.Count{
			{Name: "BALANCO_PATRIMONIAL", Value: 174},
			{Name: "DECLARACAO_IR", Value: 158},
			{Name: "ENDIVIDAMENTO", Value: 139},
			{Name: "POSICAO_CONSOLIDADA", Value: 138},
			{Name: "CAPA_SERASA", Value: 90},
			{Name: "CNH", Value: 62},
			{Name: "FATURAMENTO", Value: 52},
			{Name: "COMPROVANTE_RESIDENCIA", Value: 49},
			{Name: "CONTRATO_SOCIAL", Value: 47},
			{Name: "FATURA_ENERGIA", Value: 43},
			{Name: "RECIBOS", Value: 35},
		}

		got_pages_processed := rep.CountExtracts(false, group_by, filter)
		got_rows_count := rep.CountExtracts(true, group_by, filter)

		if !reflect.DeepEqual(expected_pages_processed, got_pages_processed) {
			t.Errorf("Pages processed: Expected %v, got %v", expected_pages_processed, got_pages_processed)
		}

		if !reflect.DeepEqual(expected_rows_count, got_rows_count) {
			t.Errorf("Rows counted: Expected %v, got %v", expected_rows_count, got_rows_count)
		}
	})

	t.Run("count grouping by user segment", func(t *testing.T) {
		expected_pages_processed := []*model.Count{
			{Name: "banco", Value: 6038},
			{Name: "construtora", Value: 3649},
			{Name: "financeira", Value: 3381},
			{Name: "imobiliaria", Value: 3041},
		}
		expected_rows_count := []*model.Count{
			{Name: "banco", Value: 1981},
			{Name: "construtora", Value: 1058},
			{Name: "financeira", Value: 987},
			{Name: "imobiliaria", Value: 974},
		}

		got_pages_processed := rep.CountExtracts(false, "users.segment", FiltroExtract{})
		got_rows_count := rep.CountExtracts(true, "users.segment", FiltroExtract{})

		if !reflect.DeepEqual(expected_pages_processed, got_pages_processed) {
			t.Errorf("Pages processed: Expected %v, got %v", expected_pages_processed, got_pages_processed)
		}

		if !reflect.DeepEqual(expected_rows_count, got_rows_count) {
			t.Errorf("Rows counted: Expected %v, got %v", expected_rows_count, got_rows_count)
		}
	})

	t.Run("filter by segment grouping by month", func(t *testing.T) {
		filter := FiltroExtract{
			Segment: "financeira",
		}

		expected_pages_processed := []*model.Count{
			{Name: "7", Value: 376},
			{Name: "8", Value: 1023},
			{Name: "9", Value: 1020},
			{Name: "10", Value: 962},
		}
		expected_rows_count := []*model.Count{
			{Name: "7", Value: 172},
			{Name: "8", Value: 301},
			{Name: "9", Value: 280},
			{Name: "10", Value: 234},
		}

		got_pages_processed := rep.CountExtracts(false, "EXTRACT(month FROM created_at::date)", filter)
		got_rows_count := rep.CountExtracts(true, "EXTRACT(month FROM created_at::date)", filter)

		if !reflect.DeepEqual(expected_pages_processed, got_pages_processed) {
			t.Errorf("Pages processed: Expected %v, got %v", expected_pages_processed, got_pages_processed)
		}

		if !reflect.DeepEqual(expected_rows_count, got_rows_count) {
			t.Errorf("Rows counted: Expected %v, got %v", expected_rows_count, got_rows_count)
		}
	})
}
