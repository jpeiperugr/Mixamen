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

func TestRegistrarSaber(t *testing.T) {
	t.Run("registra saber basico valido", func(t *testing.T) {
		cat := NuevoCatalogo()
		err := cat.RegistrarSaber("2.A.2.3.", "Números enteros, fraccionarios, decimales y raíces")
		if err != nil {
			t.Fatalf("no se esperaba error, se obtuvo: %v", err)
		}
	})

	t.Run("falla con codigo vacio", func(t *testing.T) {
		cat := NuevoCatalogo()
		err := cat.RegistrarSaber("", "Descripción válida")
		if !errors.Is(err, ErrCodigoVacio) {
			t.Fatalf("se esperaba %v, se obtuvo: %v", ErrCodigoVacio, err)
		}
	})

	t.Run("falla con descripcion vacia", func(t *testing.T) {
		cat := NuevoCatalogo()
		err := cat.RegistrarSaber("2.A.2.3.", "   ")
		if !errors.Is(err, ErrDescripcionVacia) {
			t.Fatalf("se esperaba %v, se obtuvo: %v", ErrDescripcionVacia, err)
		}
	})

	t.Run("falla con saber duplicado", func(t *testing.T) {
		cat := NuevoCatalogo()
		_ = cat.RegistrarSaber("2.A.2.3.", "Números enteros")
		err := cat.RegistrarSaber("2.A.2.3.", "Duplicado")
		if !errors.Is(err, ErrSaberDuplicado) {
			t.Fatalf("se esperaba %v, se obtuvo: %v", ErrSaberDuplicado, err)
		}
	})
}

func TestCatalogarEjercicio(t *testing.T) {
	nuevoCatalogoBase := func(t *testing.T) *Catalogo {
		t.Helper()
		cat := NuevoCatalogo()
		if err := cat.RegistrarCriterio("1.1", "Interpretar problemas matemáticos"); err != nil {
			t.Fatal(err)
		}
		if err := cat.RegistrarSaber("2.A.2.3.", "Números enteros y fraccionarios"); err != nil {
			t.Fatal(err)
		}
		return cat
	}

	t.Run("cataloga ejercicio correctamente", func(t *testing.T) {
		cat := nuevoCatalogoBase(t)
		err := cat.CatalogarEjercicio("Calcula 3/4 de 120 euros.", []string{"1.1"}, []string{"2.A.2.3."})
		if err != nil {
			t.Fatalf("no se esperaba error, se obtuvo: %v", err)
		}
	})

	t.Run("falla con enunciado vacio", func(t *testing.T) {
		cat := nuevoCatalogoBase(t)
		err := cat.CatalogarEjercicio("   ", []string{"1.1"}, []string{"2.A.2.3."})
		if !errors.Is(err, ErrEnunciadoVacio) {
			t.Fatalf("se esperaba %v, se obtuvo: %v", ErrEnunciadoVacio, err)
		}
	})

	t.Run("falla con ejercicio duplicado", func(t *testing.T) {
		cat := nuevoCatalogoBase(t)
		_ = cat.CatalogarEjercicio("Calcula 3/4 de 120 euros.", []string{"1.1"}, []string{"2.A.2.3."})
		err := cat.CatalogarEjercicio("Calcula 3/4 de 120 euros.", []string{"1.1"}, []string{"2.A.2.3."})
		if !errors.Is(err, ErrEjercicioDuplicado) {
			t.Fatalf("se esperaba %v, se obtuvo: %v", ErrEjercicioDuplicado, err)
		}
	})

	t.Run("falla si falta criterio o saber basico", func(t *testing.T) {
		cat := nuevoCatalogoBase(t)
		err := cat.CatalogarEjercicio("Ejercicio sin criterios", nil, []string{"2.A.2.3."})
		if !errors.Is(err, ErrSinJustificacionCurricular) {
			t.Fatalf("se esperaba %v, se obtuvo: %v", ErrSinJustificacionCurricular, err)
		}
	})

	t.Run("falla si el criterio no existe en el catalogo", func(t *testing.T) {
		cat := nuevoCatalogoBase(t)
		err := cat.CatalogarEjercicio("Otro ejercicio", []string{"9.9"}, []string{"2.A.2.3."})
		if !errors.Is(err, ErrCriterioNoEncontrado) {
			t.Fatalf("se esperaba %v, se obtuvo: %v", ErrCriterioNoEncontrado, err)
		}
	})

	t.Run("falla si el saber no existe en el catalogo", func(t *testing.T) {
		cat := nuevoCatalogoBase(t)
		err := cat.CatalogarEjercicio("Otro ejercicio", []string{"1.1"}, []string{"9.Z.9."})
		if !errors.Is(err, ErrSaberNoEncontrado) {
			t.Fatalf("se esperaba %v, se obtuvo: %v", ErrSaberNoEncontrado, err)
		}
	})
}
