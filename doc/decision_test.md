# Elección de herramientas de test (Issue #28)

En este archivo se argumenta la decisión del test runner y la biblioteca de aserciones que se van a utilizar en el proyecto.

Para que se entienda la decisión primero voy a explicar que es y para que sirve cada herramienta:

- **Test runner:** busca los ficheros _test.go, ejecuta las pruebas y subtests, y muestra los resultados. En Go se hace con el comando **go test** junto con el paquete estándar *testing*, o con runners externos como gotestsum.
- **Biblioteca de aserciones:** son las funciones que comparan el valor obtenido con el esperado y marcan el fallo si no coinciden. Go no incluye una biblioteca de aserciones propiamente dicha, sino que emplea comparaciones nativas del lenguaje (if) y se apoya en funciones de la biblioteca estándar como **errors.Is** o **slices.Equal**, o con bibliotecas externas como testify/assert.

## Criterios

Para buscar opciones he filtrado por herramientas mantenidas y compatibles con la versión de Go elegida para este proyecto (1.27). Para aceptar o descartar cada opción he establecido tres criterios:

- **Sin librerías externas:** la herramienta no puede añadir paquetes de terceros en go.mod.
- **Sin escritura en el directorio de código:** al lanzar los tests no se puede crear ningún fichero temporal ni *lockfile* dentro del directorio donde están los fuentes.
- **Agrupar pruebas, comprobar errores y ver cobertura:** debe permitir organizar varios casos dentro de una misma función de test, comprobar fácilmente que devuelve error cuando los datos están mal y calcular qué parte del código se ha probado.

## Decisión

He elegido **la cadena estándar de Go (go test + paquete testing) como test runner** y **las comprobaciones nativas del lenguaje (if, errors.Is, slices.Equal) como mecanismo de aserción**.

Descarto gotestsum y testify/assert porque obligan a meter dependencias externas en go.mod. Con lo que trae Go de serie no se añade nada a go.mod, la caché de ejecución se guarda fuera de la carpeta del proyecto (GOCACHE) y podemos agrupar pruebas (t.Run), verificar errores (errors.Is) y medir cobertura (go test -cover) sin instalar nada más.

## Principios F.I.R.S.T. aplicados en los tests

- **Fast (Rápidos):** los tests se ejecutan en memoria sobre Catalogo, sin leer ficheros ni usar red, así que tardan milisegundos.
- **Independent (Independientes):** cada test crea su propio Catalogo nuevo, por lo que una prueba nunca afecta a las demás.
- **Repeatable (Repetibles):** usan siempre los mismos datos de entrada y ordenan los resultados (ya que en Go los map no tienen orden fijo) para dar la misma salida en cualquier entorno.
- **Self-validating (Autovalidables):** cada test comprueba por sí solo si pasa o falla con t.Fatalf y errors.Is, sin tener que revisar nada a mano.
- **Timely (Oportunos):** se escriben junto a cada *issue* antes de dar por terminada cada función de la lógica de negocio.