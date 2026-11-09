package mixamen

type idEjercicio string

type EjercicioMatematico struct {
	id idEjercicio
	enunciado string
	conceptos map[conceptoMatematico]struct{}
	criterios map[criterioEvaluacion]struct{}
	saberes map[saberBasico]struct{}
}
