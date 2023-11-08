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
				Value: 1058,
			},
			{
				Name:  "4",
				Value: 1003,
			},
			{
				Name:  "5",
				Value: 987,
			},
			{
				Name:  "2",
				Value: 978,
			},
			{
				Name:  "3",
				Value: 974,
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

		got := rep.CountExtracts(group_by, filter)

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})
}
