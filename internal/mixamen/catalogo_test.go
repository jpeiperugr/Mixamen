package mixamen

import (
	"errors"
	"testing"
)

func TestRegistrarCriterio(t *testing.T) {
	t.Run("registra criterio valido", func(t *testing.T) {
		cat := NuevoCatalogo()
		err := cat.RegistrarCriterio("1.1", "Interpretar problemas matemáticos de la vida cotidiana")
		if err != nil {
			t.Fatalf("no se esperaba error, se obtuvo: %v", err)
		}
	})

	t.Run("falla con codigo vacio", func(t *testing.T) {
		cat := NuevoCatalogo()
		err := cat.RegistrarCriterio("", "Descripción válida")
		if !errors.Is(err, ErrCodigoVacio) {
			t.Fatalf("se esperaba %v, se obtuvo: %v", ErrCodigoVacio, err)
		}
	})

	t.Run("falla con descripcion vacia", func(t *testing.T) {
		cat := NuevoCatalogo()
		err := cat.RegistrarCriterio("1.1", "   ")
		if !errors.Is(err, ErrDescripcionVacia) {
			t.Fatalf("se esperaba %v, se obtuvo: %v", ErrDescripcionVacia, err)
		}
	})

	t.Run("falla con criterio duplicado", func(t *testing.T) {
		cat := NuevoCatalogo()
		_ = cat.RegistrarCriterio("1.1", "Descripción inicial")
		err := cat.RegistrarCriterio("1.1", "Otra descripción")
		if !errors.Is(err, ErrCriterioDuplicado) {
			t.Fatalf("se esperaba %v, se obtuvo: %v", ErrCriterioDuplicado, err)
		}
	})
}
