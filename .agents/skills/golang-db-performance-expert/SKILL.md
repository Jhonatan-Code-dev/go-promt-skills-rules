---
name: golang-db-performance-expert
description: Experto en análisis y optimización extrema de rendimiento en consultas SQL, esquema de ENT ORM, índices B-Tree, resolución N+1 y paginación por cursor.
---

# Skill: Experto en Rendimiento Extremo de Base de Datos y Consultas (Golang DB Performance Expert)

Esta habilidad se activa cuando el usuario solicita optimizar la velocidad de una consulta, analizar índices de base de datos, corregir lentitud bajo alta demanda o diseñar esquemas de ENT para alta escala.

## Protocolo de Análisis y Optimización

Al auditar o construir consultas/esquemas, sigue estos pasos rigurosos:

### 1. Auditoría de Índices y Claves Foráneas
- Verifica que todas las Foreign Keys (`usuario_id`, `empresa_id`, etc.) tengan un índice explícito en `Indexes()`.
- Revisa las consultas más frecuentes y recomienda **Índices Compuestos** si filtran o devuelven por múltiples columnas.

### 2. Detección y Eliminación del Problema N+1
- Escanea el código en busca de llamadas a la BD dentro de ciclos `for` o `range`.
- Sustituye iteraciones múltiples por **Eager Loading** (`With<Relación>()` en ENT) o consultas por lotes (`Where(in(IDs))`).

### 3. Optimización de Paginación en Tablas Grandes
- Reemplaza paginación basada en `OFFSET` / `LIMIT` por **Paginación basada en Cursor / Keyset** (`Where(id > ultimoID).Limit(...)`).

### 4. Proyección de Campos (Select Específico)
- Verifica que no se lean campos grandes o columnas irrelevantes si solo se necesitan 2 o 3 atributos (`Select(...)`).

### 5. Reporte de Rendimiento Extremo
Genera un diagnóstico claro con:
- **Problema de Rendimiento Detectado** (Ej: Falta de índice, N+1, OFFSET lento).
- **Impacto a Gran Escala** (Ej: Cuello de botella cuando la tabla alcance 100k+ filas).
- **Código Optimizado Propuesto**.
