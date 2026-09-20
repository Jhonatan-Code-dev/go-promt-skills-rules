---
name: golang-null-schema-auditor
description: Audita y aplica la regla de campos opcionales NULL en esquemas ENT ORM y DTOs, eliminando cadenas vacías "" en favor de valores NULL en la base de datos.
---

# Skill: Auditor de Valores NULL y Esquemas en Go (Golang NULL Schema Auditor)

Esta habilidad se activa cuando el usuario solicita auditar, refactorizar o verificar el manejo de campos opcionales y valores nulos (`NULL`) en esquemas de ENT ORM, DTOs o la base de datos, garantizando que nunca se persistan cadenas vacías `""`.

## Protocolo de Auditoría y Corrección de Campos NULL

Al ejecutar una auditoría de valores nulos y esquemas:

### 1. Escaneo de Esquemas ENT ORM (`ent/schema/*.go`)
- Inspecciona cada definición de campo (`field.String`, `field.Int`, `field.Time`, etc.) en los archivos del directorio `ent/schema/`.
- Identifica campos que utilicen `.Optional()` pero **no** incluyan `.Nillable()`.
- Detecta campos con valores por defecto vacíos como `.Default("")`.
- Modifica la definición agregando `.Nillable()` para garantizar que ENT maneje punteros Go (`*string`) y los persista como `NULL` en SQL:
  ```go
  // Corrección en esquema ENT:
  field.String("observaciones").Optional().Nillable().Comment("Observaciones adicionales")
  ```

### 2. Auditoría de DTOs y Manejadores HTTP
- Revisa los structs de DTOs de petición en Fiber v3.
- Asegura que los campos opcionales usen punteros (ejemplo: `*string` en lugar de `string`).
- Implementa o aplica funciones de sanitización antes de invocar la capa de servicio/dominio:
  ```go
  if dto.TelefonoSecundario != nil {
      dto.TelefonoSecundario = NormalizarTextoOpcional(dto.TelefonoSecundario)
  }
  ```

### 3. Verificación de la Generación de ENT
- Ejecuta `go generate ./ent` o `go test ./...` para regenerar las entidades y verificar que no existan errores de compilación en los mapeadores ni en la capa de datos.
