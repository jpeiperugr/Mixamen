package mixamen

type Catalogo struct {
	criterios  map[criterioEvaluacion]struct{}
	saberes    map[saberBasico]struct{}
	ejercicios []ejercicioMatematico
}
