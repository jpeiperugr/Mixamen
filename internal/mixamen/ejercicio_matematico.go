package mixamen

type ejercicioMatematico struct {
	criterios map[criterioEvaluacion]struct{}
	saberes   map[saberBasico]struct{}
}

func nuevoEjercicioMatematico(criterios []criterioEvaluacion, saberes []saberBasico) (ejercicioMatematico, error) {
	if len(criterios) == 0 || len(saberes) == 0 {
		return ejercicioMatematico{}, ErrSinJustificacionCurricular
	}

	conjCriterios := make(map[criterioEvaluacion]struct{}, len(criterios))
	for _, c := range criterios {
		conjCriterios[c] = struct{}{}
	}

	conjSaberes := make(map[saberBasico]struct{}, len(saberes))
	for _, s := range saberes {
		conjSaberes[s] = struct{}{}
	}

	return ejercicioMatematico{
		criterios: conjCriterios,
		saberes:   conjSaberes,
	}, nil
}
