# Regla Estricta: Inspección Completa e Incondicional del Código y Veracidad Absoluta (Complete Code Inspection Standard)

> 🔍 **PROHIBICIÓN ABSOLUTA DE REVISIÓN PARCIAL (NO SNIPPET TUNNEL VISION)**: Queda estrictamente prohibido que la IA audite, revise o diagnostique código leyendo únicamente fragmentos parciales, las primeras líneas de un archivo o haciendo suposiciones sobre funciones no leídas. **SI EL USUARIO PIDE REVISAR SU CÓDIGO, SE DEBE INSPECCIONAR EL CONTENIDO COMPLETO Y TOTAL DE LOS ARCHIVOS**.

---

## 1. Inspección Completa y Total del Código

1. **Prohibido la Vista de Túnel de Fragmentos (No Snippet Tunnel Vision)**:
   - Al auditar o analizar un archivo de código, la IA DEBE leerlo **DE PRINCIPIO A FIN** (usando desplazamientos/offsets de lectura si el archivo supera los límites).
   - Queda prohibido sacar conclusiones o diagnosticar un error habiendo visto solo las primeras 15 o 30 líneas de un archivo.
2. **Rastreo Completo de Dependencias e Invocaciones**:
   - Si una función o tipo se utiliza en múltiples capas (Domain, UseCase, Infrastructure, Presentation, ENT Schema), la IA debe verificar todas las llamadas y usos reales en la base de código.
3. **Auditoría Holística (Sin Omisiones)**:
   - En auditorías de arquitectura, seguridad o rendimiento, revisar cada archivo del paquete o módulo sin omitir constructores, métodos auxiliares privados ni archivos de configuración.

---

## 2. Prohibido Confiar en Cambios Sin Probar (Mandato de Verificación)

1. **Editar un Archivo NO Equivale a Completar la Tarea**:
   - Tras realizar cualquier modificación o refactorización en el código, la IA DEBE ejecutar comandos de verificación reales (`go test ./...`, `go build ./...`, `go run ...`) para validar que el código compila y ejecuta sin errores.
2. **Cero Declaraciones Prematuras de Éxito**:
   - Queda estrictamente prohibido afirmar *"El problema está resuelto"*, *"El código funciona perfectamente"* o *"Se aplicó la corrección"* si no se han inspeccionado los logs reales de ejecución y verificado un código de salida `0` (Exit Code 0).
3. **Pruebas Reales y Cobertura**:
   - Toda validación debe basarse en ejecuciones reales de la suite de pruebas o ejecutables Go. No se permiten pruebas simuladas o ficticias.

---

## 3. Transparencia Absoluta ante Fallos y Errores

1. **Reporte Honesto e Inmediato de Errores**:
   - Si un comando de prueba o compilación devuelve un error o fallas en tests, la IA DEBE mostrar de inmediato el log exacto y la falla real. Queda prohibido ocultar, ignorar o disfrazar errores.
2. **Cero Suposiciones o Alucinaciones**:
   - No adivinar firmas de funciones, respuestas de BD o rutas de archivos. Si hay incertidumbre, investigar primero inspeccionando los archivos y ejecutando comandos antes de concluir.
3. **Cero Parches Superficiales**:
   - Queda prohibido silenciar errores con `try/catch` vacíos, suprimir aserciones o ignorar variables devueltas (`_ = err`) para simular que algo funciona.

---

## 4. Protocolo Obligatorio de 4 Pasos para Toda Revisión y Modificación

```text
1. LEER COMPLETO ➔ Leer de principio a fin el archivo o módulo completo sin omisiones.
2. MODIFICAR     ➔ Aplicar el cambio de código o refactorización.
3. EJECUTAR      ➔ Correr comandos reales de construcción y pruebas (go build / go test).
4. VERIFICAR     ➔ Leer el log de salida real y confirmar 100% éxito antes de responder.
```
