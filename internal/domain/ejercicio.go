package domain

type Ejercicio struct { //Representa el problema matemático y actúa como la entidad principal del problema
	enunciado string
	conceptos map[conceptoMatematico]struct{}
	criterios map[criterioEvaluacion]struct{}
	saberes map[saberBasico]struct{}
}
