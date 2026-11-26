package mixamen

import (
	"errors"
	"slices"
	"strings"
)

var (
	ErrCodigoVacio                = errors.New("el código no puede estar vacío")
	ErrDescripcionVacia           = errors.New("la descripción no puede estar vacía")
	ErrCriterioDuplicado          = errors.New("el criterio de evaluación ya está registrado")
	ErrCriterioNoEncontrado       = errors.New("el criterio de evaluación no existe en el catálogo")
	ErrSaberDuplicado             = errors.New("el saber básico ya está registrado")
	ErrSaberNoEncontrado          = errors.New("el saber básico no existe en el catálogo")
	ErrEnunciadoVacio             = errors.New("el enunciado del ejercicio no puede estar vacío")
	ErrEjercicioDuplicado         = errors.New("el ejercicio ya está catalogado")
	ErrSinJustificacionCurricular = errors.New("el ejercicio debe tener al menos un criterio y un saber básico")
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

func (c *Catalogo) RegistrarSaber(codigo, descripcion string) error {
	codigo = strings.TrimSpace(codigo)
	descripcion = strings.TrimSpace(descripcion)

	nuevo, err := nuevoSaberBasico(codigo, descripcion)
	if err != nil {
		return err
	}
	if _, existe := c.buscarSaber(codigo); existe {
		return ErrSaberDuplicado
	}

	c.saberes[nuevo] = struct{}{}
	return nil
}

func (c *Catalogo) CatalogarEjercicio(enunciado string, codigosCriterios, codigosSaberes []string) error {
	enunciado = strings.TrimSpace(enunciado)
	if enunciado == "" {
		return ErrEnunciadoVacio
	}
	if _, existe := c.ejerciciosPorEnunciado[enunciado]; existe {
		return ErrEjercicioDuplicado
	}

	criterios, err := c.resolverCriterios(codigosCriterios)
	if err != nil {
		return err
	}
	saberes, err := c.resolverSaberes(codigosSaberes)
	if err != nil {
		return err
	}

	ejercicio, err := nuevoEjercicioMatematico(criterios, saberes)
	if err != nil {
		return err
	}

	c.ejerciciosPorEnunciado[enunciado] = ejercicio
	return nil
}

func (c *Catalogo) BuscarEjerciciosPorCriterio(codigo string) ([]string, error) {
	codigo = strings.TrimSpace(codigo)
	if codigo == "" {
		return nil, ErrCodigoVacio
	}

	criterio, existe := c.buscarCriterio(codigo)
	if !existe {
		return nil, ErrCriterioNoEncontrado
	}

	var resultado []string
	for enunciado, ej := range c.ejerciciosPorEnunciado {
		if _, ok := ej.criterios[criterio]; ok {
			resultado = append(resultado, enunciado)
		}
	}
	slices.Sort(resultado)
	return resultado, nil
}

func (c *Catalogo) buscarCriterio(codigo string) (criterioEvaluacion, bool) {
	for crit := range c.criterios {
		if crit.codigo == codigo {
			return crit, true
		}
	}
	return criterioEvaluacion{}, false
}

func (c *Catalogo) buscarSaber(codigo string) (saberBasico, bool) {
	for sab := range c.saberes {
		if sab.codigo == codigo {
			return sab, true
		}
	}
	return saberBasico{}, false
}

func (c *Catalogo) resolverCriterios(codigos []string) ([]criterioEvaluacion, error) {
	res := make([]criterioEvaluacion, 0, len(codigos))
	for _, cod := range codigos {
		crit, ok := c.buscarCriterio(strings.TrimSpace(cod))
		if !ok {
			return nil, ErrCriterioNoEncontrado
		}
		res = append(res, crit)
	}
	return res, nil
}

func (c *Catalogo) resolverSaberes(codigos []string) ([]saberBasico, error) {
	res := make([]saberBasico, 0, len(codigos))
	for _, cod := range codigos {
		sab, ok := c.buscarSaber(strings.TrimSpace(cod))
		if !ok {
			return nil, ErrSaberNoEncontrado
		}
		res = append(res, sab)
	}
	return res, nil
}
