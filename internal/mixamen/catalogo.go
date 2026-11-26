package mixamen

import (
	"errors"
	"strings"
)

var (
	ErrCodigoVacio       = errors.New("el código no puede estar vacío")
	ErrDescripcionVacia  = errors.New("la descripción no puede estar vacía")
	ErrCriterioDuplicado = errors.New("el criterio de evaluación ya está registrado")
)

type Catalogo struct {
	criterios              map[criterioEvaluacion]struct{}
	saberes                map[saberBasico]struct{}
	ejerciciosPorEnunciado map[string]ejercicioMatematico
}

func NuevoCatalogo() *Catalogo {
	return &Catalogo{
		criterios:              make(map[criterioEvaluacion]struct{}),
		saberes:                make(map[saberBasico]struct{}),
		ejerciciosPorEnunciado: make(map[string]ejercicioMatematico),
	}
}

func (c *Catalogo) RegistrarCriterio(codigo, descripcion string) error {
	codigo = strings.TrimSpace(codigo)
	descripcion = strings.TrimSpace(descripcion)

	nuevo, err := nuevoCriterioEvaluacion(codigo, descripcion)
	if err != nil {
		return err
	}
	if _, existe := c.buscarCriterio(codigo); existe {
		return ErrCriterioDuplicado
	}

	c.criterios[nuevo] = struct{}{}
	return nil
}

func (c *Catalogo) buscarCriterio(codigo string) (criterioEvaluacion, bool) {
	for crit := range c.criterios {
		if crit.codigo == codigo {
			return crit, true
		}
	}
	return criterioEvaluacion{}, false
}
