package service

import "testing"

func TestUser(t *testing.T) {
	rep := RepositorioPostgre{}

	t.Run("find one user", func(t *testing.T) {
		expected := User{id: 3, name: "Caio", segment: "imobiliaria"}
		got := rep.findUser(3)

		if expected != got {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})
}
