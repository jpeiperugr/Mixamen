package mixamen

type ejercicioMatematico struct {
	enunciado string
	conceptos map[conceptoMatematico]struct{}
	criterios map[criterioEvaluacion]struct{}
	saberes map[saberBasico]struct{}
}
