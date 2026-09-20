---
name: golang-test-runner
description: Ejecuta y audita la suite centralizada de pruebas organizadas dentro de la carpeta principal pruebas/ (unitarias, integración, rendimiento).
---

# Skill: Ejecutor de Pruebas Centralizadas (Golang Test Runner)

Esta habilidad se activa cuando el usuario solicita ejecutar, crear o verificar pruebas en la carpeta principal `pruebas/`.

## Protocolo de Ejecución de Pruebas Centralizadas

Al crear o ejecutar pruebas:

### 1. Ubicación en `pruebas/`
- Asegura que la prueba esté dentro de `pruebas/unitarias/`, `pruebas/integracion/` o `pruebas/rendimiento/`.

### 2. Encabezado de Documentación
- Incluye el encabezado obligatorio especificando:
  - **¿Qué prueba?**
  - **Escenario**
  - **Comando exacto de consola**

### 3. Comando de Ejecución
Ejecuta la prueba correspondiente mediante `run_command`:
```bash
go test -v ./pruebas/unitarias/...
```
Verifica que el código de salida sea `0` y que todos los casos pasen en verde antes de reportar el resultado al usuario.
