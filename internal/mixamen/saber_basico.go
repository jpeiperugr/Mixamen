package mixamen

type saberBasico struct {
	codigo      string
	descripcion string
}

func nuevoSaberBasico(codigo, descripcion string) (saberBasico, error) {
	if codigo == "" {
		return saberBasico{}, ErrCodigoVacio
	}
	if descripcion == "" {
		return saberBasico{}, ErrDescripcionVacia
	}
	return saberBasico{
		codigo:      codigo,
		descripcion: descripcion,
	}, nil
}
