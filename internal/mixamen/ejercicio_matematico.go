package mixamen

type ejercicioMatematico struct {
	criterios map[criterioEvaluacion]struct{}
	saberes   map[saberBasico]struct{}
}
