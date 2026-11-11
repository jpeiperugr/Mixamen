package mixamen

type EjercicioMatematico struct {
	enunciado string
	conceptos map[conceptoMatematico]struct{}
	criterios map[criterioEvaluacion]struct{}
	saberes map[saberBasico]struct{}
}
