# Mixamen

## Problema a tratar

El profesorado de distintas etapas académicas acumula a lo largo de los cursos cientos de exámenes, ejercicios y problemas prácticos en documentos de texto y PDF's sin estructurar. El problema surge cuando el temario o la ley educativa cambia (como ocurre frecuentemente en la ESO). 

Al cambiar el temario, el profesor no sabe qué preguntas de sus exámenes antiguos siguen siendo válidas o a qué nuevo bloque de criterios pertenecen. Revisar manualmente cada pregunta de cada uno de los PDF's antiguos y evaluarlo con el nuevo temario oficial para saber qué sirve y qué no, es un proceso molesto y propenso a errores. El problema, por tanto, no es generar un examen, sino extraer, analizar y catalogar automáticamente el conocimiento que ya está en ficheros semiestructurados de años anteriores.

## Conocimiento personal

Este problema viene de una charla que tuve con mi tía que es profesora. Ella lleva ya unos años dando clases de matemáticas en cursos de la ESO. Me contó que tenía un problema a la hora de aprovechar el material de años anteriores. Cuando cambian las guías docentes o los criterios de evaluación, tiene que abrir PDF's antiguos uno a uno, leer las preguntas y decidir a qué parte del nuevo temario corresponde cada ejercicio.

## ¿Cómo se obtienen los datos? 

Los datos provienen de dos fuentes reales y semiestructuradas:
1. **Los exámenes antiguos:** Archivos PDF y documentos de Word (.docx) que contienen el texto de los ejercicios de años pasados.
2. **Las Guías Docentes:** Documentos (PDFs o webs de la consejería de educación) que contienen el temario estructurado oficial vigente.

## ¿Por qué requiere una lógica de negocio y no solo almacenamiento?

Porque los datos actuales no tienen metadatos. El problema exige extraer el texto en bruto de los exámenes antiguos, aislar las preguntas individuales y aplicar una lógica de procesamiento para cruzar el contenido de esas preguntas con las palabras clave y conceptos de las guías docentes actuales para poder catalogarlas. 

## Por qué requiere un despliegue en la nube

Hay varias razones por la que este problema necesita ser desplegado en la nube. La primera es por las restricciones de instalación en centros educativos, en la mayoría de centros ya sean públicos, privados o concertados los profesores no pueden instalar las aplicaciones o programas que quieran y suelen ser bastante restrictivos. La segunda razón es porque muchos profesores tienen que trabajar con varios ordenadores, suelen tener su ordenador personal donde tienen la mayoría de sus ejercicios pero también necesitan acceder a esos ejercicios desde los ordenadores de los centros para bien exponerlos en clase, mandarlos a imprimir, etc, por lo que centralizar el procesamiento en la nube permite que el análisis del material esté accesible en cualquier momento y lugar.

## Fotografía de las tarjetas de rol

![Fotografía de la tarjeta de rol](media/tarjetas_rol.jpg)

## Enlaces
* [Documentación adicional y de configuración](doc/objetivo0.md)
