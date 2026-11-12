# Gestores de dependencias y tareas (Issue #21)

En Go, el gestor de dependencias estándar y nativo es `go mod`. Al ser el estándar oficial del  desde la versión 1.16, además no requiere instalación externa.

Para el gestor de tareas he desarrollado los siguientes requisitos:

- Que no añada herramientas de terceros
- Que se ejecute de forma rápida 
- Que a la hora de integrar nuevos desarrolladores suponga el menor problema posible
- Que sea lo más ligera posible en cuanto a peso y en cuanto a líneas necesarias para una misma tarea

Me voy a centrar en comparar *Make* y *Task* que son los principales de la industria.

## Pruebas objetivas 

He lanzado `time <herramienta> check` para medir el tiempo de compilación, para ello he creado un ejecutable con un simple `echo`.

![Medición](media/medicion.png)

*Tabla para comparar herramientas según los requisitos:*

| Requisito Objetivo | Make (`Makefile`) | Task (`Taskfile.yml`) |
| :--- | :--- | :--- |
| **Tiempo de ejecución**| 0.004s | ~0.121s |
| **Dependencia Terceros** | Viene preinstalado en Linux y Mac | Descarga externa en cualquier sistema |
| **Integración desarrolladores** | Configuración cero | Requiere paso de instalación |
| **Ligereza de herramienta** | 250 KB | 15MB |
| **Ligereza de configuración** | 2 líneas | 5 líneas |


*Decisión:* He elegido make porque como se aprecia en la captura, make ya viene preinstalado en el sistema mientras que task no, haciendo que para cualquier persona nueva que tuviera que trabajar en este proyecto ya tuviese instalado el gestor de tareas.