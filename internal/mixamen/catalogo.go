package mixamen

type Catalogo struct {
	criterios map[criterioEvaluacion]struct{}
	saberes map[saberBasico]struct{}
	ejerciciosPorEnunciado map[string]ejercicioMatematico
}
