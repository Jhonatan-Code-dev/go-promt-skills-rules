# Regla Estricta: Organización Centralizada de Pruebas en Carpeta `pruebas/` (Centralized Testing Standard)

> 🧪 **ESTRUCTURA ORDENADA DE PRUEBAS**: Para mantener el código fuente limpio y ultra organizado, **todas las pruebas del proyecto DEBEN organizarse dentro de la carpeta principal `pruebas/`** (dividida en subcarpetas especializadas) y documentar claramente qué prueban y cómo se ejecutan.

---

## 1. Estructura Centralizada de Pruebas (`pruebas/`)

```text
mi-proyecto/
├── internal/               # Código fuente limpio de la aplicación
├── pruebas/                # CARPETA PRINCIPAL DE PRUEBAS
│   ├── unitarias/          # Pruebas unitarias de casos de uso y lógica pura
│   │   ├── usuario_test.go
│   │   └── orden_test.go
│   ├── integracion/        # Pruebas de integración con BD ENT y Fiber v3
│   │   ├── autenticacion_api_test.go
│   │   └── pago_api_test.go
│   └── rendimiento/        # Pruebas de carga y estrés para VPS
│       └── carga_endpoints_test.go
```

---

## 2. Documentación Explicativa Obligatoria en Cada Archivo de Prueba

Cada archivo de prueba dentro de `pruebas/` DEBE comenzar con un encabezado en **Español** que especifique:

1. **¿Qué prueba?**: Componente o caso de uso evaluado.
2. **Escenario de Prueba**: Entradas, condiciones iniciales y resultado esperado.
3. **Comando Exacto de Ejecución**: El comando de terminal para ejecutar la prueba individual o el paquete completo.

### Ejemplo de Archivo de Prueba Documentado (`pruebas/unitarias/usuario_test.go`):

```go
// ============================================================================
// ARCHIVO DE PRUEBAS UNITARIAS: LÓGICA DE USUARIOS
// ============================================================================
// ¿QUÉ PRUEBA?: Validación de registro y consulta de usuario aislado por empresa.
// ESCENARIO: Creación de usuario válido e intento de acceso entre tenants.
//
// COMANDO DE EJECUCIÓN (Ejecutar en la terminal desde la raíz del proyecto):
//   go test -v ./pruebas/unitarias/usuario_test.go
//   o para todas las unitarias:
//   go test -v ./pruebas/unitarias/...
// ============================================================================

package unitarias

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestDeberiaCrearUsuarioValido(t *testing.T) {
    t.Parallel()
    // Caso de prueba...
    assert.True(t, true)
}
```
