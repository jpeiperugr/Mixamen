package mixamen

type criterioEvaluacion struct {
	codigo      string
	descripcion string
}

func nuevoCriterioEvaluacion(codigo, descripcion string) (criterioEvaluacion, error) {
	if codigo == "" {
		return criterioEvaluacion{}, ErrCodigoVacio
	}
	if descripcion == "" {
		return criterioEvaluacion{}, ErrDescripcionVacia
	}
	return criterioEvaluacion{
		codigo:      codigo,
		descripcion: descripcion,
	}, nil
}
