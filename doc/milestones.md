# Milestones

## Milestone 0: Estructuras base del problema

1. Producto interno (PMV 0). Se entregará el código inicial con las clases que representan los conceptos (como el Ejercicio, el Criterio o el Saber Básico). Este código se desarrolla creando la base lógica sin depender todavía de ninguna base de datos ni interfaz.

2. Se considerará válido si se ha seguido la metodología de diseño dirigido por el dominio (Domain Driven Design) a partir de la HU001. Se comprobará en el pull request de la siguiente forma:

- Los issues se asignan a la HU001 y son problemas y no tareas.
- Nada de lógica de negocio, no es el objetivo de este PMV.
- Documentación breve sobre las decisiones tomadas.

## Milestone 1: Extractor y procesador de ejercicios

1. Producto interno (PMV 1). Se entregará una librería de funciones que se apoya en las clases del hito anterior. Este paquete de código podrá recibir el texto en bruto sacado de los apuntes de los profesores, procesarlo y aislar la información útil.

2. Se considerará válido cuando este código supere una serie de tests automáticos, demostrando que es capaz de procesar el texto y aplicar la lógica correctamente sin romperse.