# Mixamen

## Problema a tratar

Los profesores, especialmente en la ESO, se enfrentan a un problema enorme con la burocracia de las leyes educativas. Cada vez que usan un ejercicio en clase o en un examen, tienen que justificar exactamente qué "criterio de evaluación" oficial están aplicando.

Buscar y revisar a mano qué criterio encaja con cada problema de matemáticas es un proceso muy pesado. A un profesor le cuesta muchísimo tiempo y esfuerzo leer un ejercicio y deducir a qué apartado exacto de la ley corresponde. Necesitan una forma de introducir un ejercicio y que el sistema les diga qué criterios se están evaluando ahí.

## Conocimiento personal

Este problema viene de una charla que tuve con mi tía que es profesora. Ella lleva ya unos años dando clases de matemáticas en cursos de la ESO.  Me contó que se pierde muchísimo tiempo a lo largo del curso en buscar y revisar que criterio encaja con cada ejercicio. Me explicó que su mayor problema no es inventar o corregir los ejercicios, sino el tiempo que pierde peleándose con la ley educativa para "etiquetar" cada uno de ellos.

## ¿Cómo se obtienen los datos? 

Los datos provienen de dos fuentes reales y semiestructuradas:
1. **Los exámenes antiguos:** Archivos PDF o Latex que contienen el texto de los ejercicios de años pasados.
2. **Los criterios de evaluación:** Documentos PDF que contienen el temario estructurado oficial vigente. Principales enlaces: https://www.juntadeandalucia.es/boja/2023/104/39?utm_source=gemini, https://www.juntadeandalucia.es/boja/2023/90/?utm_source=gemini.

## ¿Por qué requiere una lógica de negocio y no solo almacenamiento?

El sistema tiene que coger el texto de un ejercicio de matemáticas, analizar las palabras y los conceptos que pide, y cruzar esa información con los textos legales de los criterios de evaluación para calcular y proponer qué criterio encaja mejor. Esto exige procesar y analizar texto, no es un simple guardar y buscar.

## Por qué requiere un despliegue en la nube

Un profesor prepara cosas en su portátil personal en casa, pero luego usa el ordenador del departamento en el instituto. Los ordenadores de los centros públicos suelen estar muy restringidos y no te dejan instalar bases de datos ni programas propios. Al estar en la nube, el profesor solo necesita abrir el navegador en cualquier ordenador para analizar sus ejercicios.


## Fotografía de las tarjetas de rol

![Fotografía de la tarjeta de rol](media/tarjetas_rol.jpg)

## Enlaces
* [Documentación adicional y de configuración](doc/objetivo0.md)