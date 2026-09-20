---
name: golang-ent-schema-gen
description: Genera esquemas de ENT ORM (entgo.io/ent) con tablas y campos explícitamente en español, anotaciones de tabla, validaciones y relaciones.
---

# Skill: Generador de Esquemas ENT ORM en Español (Golang ENT Schema Generator)

Esta habilidad se activa cuando el usuario solicita crear un esquema de base de datos o entidad de ENT en Go (ejemplo: usuario, factura, producto, cliente).

## Protocolo de Generación de Esquema ENT

Al definir un esquema de ENT para una entidad (ejemplo: `<nombre_entidad>`):

### 1. Ubicación del Archivo
- Crear el archivo en `ent/schema/<nombre_entidad>.go`.

### 2. Estructura y Nombre de Tabla en Español
- Incluir la anotación `entsql.Annotation{Table: "<nombre_tabla_plural_en_español>"}` para asegurar que el nombre de la tabla en la base de datos esté en español.

### 3. Campos en Español, Validaciones y Manejo de Nulos (NULL)
- Todos los campos (`field.String`, `field.Int`, `field.Time`, etc.) deben nombrarse en **español** (`correo_electronico`, `nombre_completo`, `monto_total`, `estado`).
- Añadir `.Comment("...")` explicativo en español para cada campo.
- Aplicar validaciones de ENT (`.NotEmpty()`, `.Positive()`, `.Unique()`, `.Sensitive()`).
- **Campos Opcionales**: Todo campo opcional DEBE declararse obligatoriamente con `.Optional().Nillable()` para garantizar que los valores vacíos o ausentes se almacenen como `NULL` en la base de datos (prohibidas las cadenas vacías `""`).

### 4. Relaciones (`Edges`)
- Definir relaciones en español (`.To("pedidos", Pedido.Type)` o `.From("usuario", Usuario.Type)`).

