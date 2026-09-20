# Regla Estricta: Documentación de Paquetes en Go y Cumplimiento Linter ST1000 (Golang Package Documentation)

> **PRINCIPIO FUNDAMENTAL**: Todo paquete Go dentro de la solución DEBE incluir obligatoriamente comentarios de documentación claros, expresivos y 100% en español. Queda prohibida la presencia de paquetes sin comentar que disparen advertencias de linters como Staticcheck (ST1000), Revive o Golangci-lint.

---

## 1. Regla Mandatoria ST1000 (Package Comment)

Cualquier directorio o módulo que contenga un paquete Go DEBE tener al menos un archivo `.go` con un comentario de paquete formal inmediatamente antes de la declaración `package`.

### Formato Estándar Obligatorio:

1. El comentario debe comenzar exactamente con la frase `// Package <nombre>` (o en su defecto utilizar el bloque `/* ... */`).
2. El comentario debe redactarse **100% en español**, describiendo con precisión la responsabilidad técnica o de dominio del paquete.
3. Se prohíbe el uso de emojis en los comentarios de documentación (cumpliendo con `no-emojis-standard.md`).

#### Ejemplo Correcto en Archivo Principal del Paquete (`dominio.go`):

```go
// Package dominio contiene las entidades de negocio, interfaces de repositorio
// y reglas de dominio fundamentales para la gestión de clientes.
package dominio
```

#### Ejemplo Correcto con Archivo Dedicado (`doc.go`):

Para paquetes extensos o de infraestructura compleja, se permite centralizar la documentación del paquete en un archivo dedicado `doc.go`:

```go
// Package repositorio implementa la persistencia de datos y acceso a base de datos
// para el módulo de clientes utilizando ENT ORM.
package repositorio
```

---

## 2. Anti-Patrones de Documentación Prohibidos

- 🚫 **Ausencia de Comentario de Paquete (ST1000)**: Dejar un paquete sin ningún comentario previa al `package <nombre>`.
- 🚫 **Comentarios Genéricos o Vacíos**: Escribir comentarios redundantes como `// Package cliente es el paquete cliente.`
- 🚫 **Documentación en Inglés**: Escribir comentarios en inglés u otros idiomas. Todo comentario debe ser en español.
- 🚫 **Uso de Emojis**: Incluir emoticonos o emojis dentro de la documentación Godoc.
- 🚫 **Desconexión del Nombre**: Iniciar el comentario sin hacer referencia explícita al nombre del paquete (ejemplo incorrecto: `// Este paquete maneja usuarios`).

---

## 3. Verificación Automática en Build y Linters

Antes de dar por finalizada cualquier tarea o pull request, se debe verificar que `staticcheck ./...` o `golangci-lint run` finalice con cero advertencias de tipo `ST1000`.
