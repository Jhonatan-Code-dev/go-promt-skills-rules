---
name: golang-anti-pattern-auditor
description: Audita y detecta malas prácticas y anti-patrones en Go (errores ignorados, goroutines huérfanas, data races, paquetes utils/common, captura en bucles y abuso de punteros).
---

# Skill: Auditor de Anti-Patrones y Malas Prácticas en Go (Golang Anti-Pattern Auditor)

Esta habilidad se activa cuando el usuario solicita escanear, auditar o corregir malas prácticas de programación en un repositorio Go.

## Protocolo de Escaneo de Anti-Patrones

Al auditar código Go, escanea minuciosamente los siguientes 5 bloques:

### 1. Detección de Errores Ignorados y Panic Excesivo
- Busca asignaciones con guion bajo `_, err :=` o invocaciones donde se ignore el valor devuelto de tipo `error`.
- Revisa el uso de `panic()` fuera de `main` / `init`.

### 2. Detección de Fugas en Concurrencia (Goroutine Leaks & Data Races)
- Verifica goroutines sin `context.Context` ni `sync.WaitGroup`.
- Detecta lectura/escritura concurrente de variables compartidas sin `sync.Mutex`.
- Revisa bucles `for` que lancen goroutines pasando variables de iteración por referencia.

### 3. Auditoría de Punteros e Interfaces
- Detecta punteros innecesarios en structs pequeñas y primitivos.
- Identifica interfaces con más de 4 métodos o interfaces no consumidas.

### 4. Auditoría de Nombres de Paquetes
- Detecta carpetas o paquetes llamados `utils`, `common`, `helpers` y sugiere nombres de dominio específicos.
- Escanea en busca de dependencias circulares.

### 5. Reporte de Corrección de Anti-Patrones
Genera una lista con:
- **Línea de Código / Archivo afectado**.
- **Anti-Patrón Detectado**.
- **Riesgo** (Memory Leak, Data Race, Crash en producción, Dificultad de Mantenimiento).
- **Código Refactorizado Sugerido**.
