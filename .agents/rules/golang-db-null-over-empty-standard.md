# Regla Estricta: Almacenamiento de Campos Opcionales como NULL (Golang DB NULL Standard)

> **PRINCIPIO FUNDAMENTAL**: Queda estrictamente PROHIBIDO almacenar cadenas vacías `""` o espacios en blanco en la base de datos para campos opcionales. Todo valor opcional que no sea proporcionado o esté vacío DEBE persistirse obligatoriamente como `NULL` en la base de datos SQL.

---

## 1. Definición en Esquemas ENT ORM (`.Optional().Nillable()`)

Todo campo opcional en los esquemas de ENT (`ent/schema/*.go`) DEBE usar obligatoriamente la combinación `.Optional().Nillable()`.

### Razón Técnica:
- `.Optional()` indica a ENT que el campo no es requerido en las mutaciones de creación.
- `.Nillable()` convierte el tipo de dato en la entidad generada en un **puntero Go** (`*string`, `*int`, `*time.Time`), permitiendo diferenciar entre un valor ausente (`nil`), que se traduce a `NULL` en la base de datos, y un valor presente.

#### Ejemplo Correcto en ENT Schema:

```go
// ent/schema/cliente.go
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)

type Cliente struct {
    ent.Schema
}

func (Cliente) Fields() []ent.Field {
    return []ent.Field{
        // Campo Obligatorio (NOT NULL)
        field.String("nombre_razon_social").
            NotEmpty().
            Comment("Nombre o razón social requerida"),

        // Campo Opcional (NULL en BD si no se envía o viene vacío)
        field.String("telefono_secundario").
            Optional().
            Nillable().
            Comment("Teléfono de contacto secundario opcional. Almacena NULL si se omite"),

        // Campo Opcional Único (NULL permite múltiples nulos sin violar UNIQUE)
        field.String("numero_identificacion_tributaria").
            Optional().
            Nillable().
            Unique().
            Comment("Número de identificación tributaria opcional"),
    }
}
```

---

## 2. Sanitización y Normalización de Entrada (DTOs & Handlers)

Antes de enviar cualquier dato a la capa de dominio o infraestructura:

1. **Campos String en DTOs**: Usar punteros `*string` para campos opcionales.
2. **Normalización de Cadenas**: Aplicar limpieza de espacios en blanco (`strings.TrimSpace`).
3. **Conversión a NULL**: Si el valor resultante de una cadena es `""` (cadena vacía), se debe transformar inmediatamente a `nil`.

#### Ejemplo de Función Auxiliar de Sanitización:

```go
// NormalizarTextoOpcional convierte cadenas vacías o compuestas por solo espacios a nil (NULL)
func NormalizarTextoOpcional(valor *string) *string {
    if valor == nil {
        return nil
    }
    textoLimpio := strings.TrimSpace(*valor)
    if textoLimpio == "" {
        return nil
    }
    return &textoLimpio
}
```

---

## 3. Prevención de Fallas en Restricciones Únicas (`UNIQUE`)

- **Problema de la Cadena Vacía `""`**: Múltiples registros con `""` en una columna con restricción `UNIQUE` causarán una colisión de clave duplicada (`duplicate key value violates unique constraint`).
- **Solución con `NULL`**: Según la especificación SQL estándar, los valores `NULL` no son iguales entre sí, por lo que la base de datos permite múltiples registros con `NULL` en columnas `UNIQUE`.

---

## 4. Prohibiciones Explícitas

- 🚫 **Cero Cadenas Vacías**: Prohibido usar `field.String("campo").Default("")` para campos opcionales.
- 🚫 **Cero Omisión de `.Nillable()`**: Prohibido usar `.Optional()` sin `.Nillable()` en campos de tipo texto opcionales, ya que ENT asignaría por defecto la cadena vacía `""`.
- 🚫 **Cero Espacios en Blanco**: Prohibido guardar cadenas compuestas únicamente por espacios `"   "`. Debe aplicarse siempre `TrimSpace`.
