---
name: golang-patterns-expert
description: Experto en patrones de diseño idiomáticos en Go (Fábrica, Singleton a prueba de hilos con sync.Once, Decorador) y corrección de anti-patrones.
---

# Skill: Experto en Patrones y Anti-Patrones de Diseño en Go (Golang Patterns Expert)

Esta habilidad se activa cuando el usuario solicita refactorizar o diseñar componentes aplicando patrones Go (Fábrica, Singleton, Decorador) o corregir anti-patrones de diseño.

## Protocolo de Aplicación de Patrones

Al auditar o diseñar componentes en Go:

### 1. Aplicación de Patrón Fábrica (Factory)
- Genera funciones constructoras `Nuevo<Struct>()` o `Fabrica<Interfaz>()` que validen parámetros y retornen interfaces o structs inicializados.

### 2. Aplicación de Singleton a Prueba de Hilos
- Implementa `sync.Once` para garantizar inicializaciones únicas concurrentemente seguras (Base de datos, Configuración, Loggers).

### 3. Aplicación de Patrón Decorador
- Implementa composición de interfaces para agregar responsabilidades cruzadas (Logging, Métricas, Caching, Tracing) sin alterar el struct base.

### 4. Corrección de Anti-Patrones
- Elimina funciones que retornen `nil` sin devolver `(T, error)`.
- Reemplaza algoritmos personalizados redundantes por funciones optimizadas de la Standard Library de Go (`slices`, `maps`, `strings`, `sync`, `context`).
