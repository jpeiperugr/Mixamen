package mixamen

type EjercicioMatematico struct {
	enunciado string
	conceptos map[conceptoMatematicas]struct{}
	criterios map[criterioEvaluacion]struct{}
	saberes map[saberBasico]struct{}
}
