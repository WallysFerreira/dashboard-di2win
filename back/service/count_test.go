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
		expected := []*model.Count{
			{
				Name:  "1",
				Value: 3649,
			},
			{
				Name:  "5",
				Value: 3381,
			},
			{
				Name:  "4",
				Value: 3144,
			},
			{
				Name:  "3",
				Value: 3041,
			},
			{
				Name:  "2",
				Value: 2894,
			},
		}

		got := rep.CountExtracts(group_by, FiltroExtract{})

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})

	t.Run("count where user_id is 5 grouping by document", func(t *testing.T) {
		group_by := "doc_type"
		filter := FiltroExtract{
			UserId: 5,
		}

		expected := []*model.Count{
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

		got := rep.CountExtracts(group_by, filter)

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})

	t.Run("count grouping by user segment", func(t *testing.T) {
		expected := []*model.Count{
			{
				Name:  "banco",
				Value: 6038,
			},
			{
				Name:  "construtora",
				Value: 3649,
			},
			{
				Name:  "financeira",
				Value: 3381,
			},
			{
				Name:  "imobiliaria",
				Value: 3041,
			},
		}

		got := rep.CountExtracts("users.segment", FiltroExtract{})

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})

	t.Run("filter by segment grouping by month", func(t *testing.T) {
		filter := FiltroExtract{
			Segment: "financeira",
		}

		expected := []*model.Count{
			{Name: "7", Value: 376},
			{Name: "8", Value: 1023},
			{Name: "9", Value: 1020},
			{Name: "10", Value: 962},
		}

		got := rep.CountExtracts("EXTRACT(month FROM created_at::date)", filter)

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})
}
