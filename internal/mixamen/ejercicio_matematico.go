package mixamen

type ejercicioMatematico struct {
	conceptos map[string]struct{}
	criterios map[criterioEvaluacion]struct{}
	saberes   map[saberBasico]struct{}
}
