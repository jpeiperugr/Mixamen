package mixamen

type ejercicioMatematico struct {
	enunciado string
	conceptos map[string]struct{}
	criterios map[criterioEvaluacion]struct{}
	saberes   map[saberBasico]struct{}
}
