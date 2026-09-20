# Regla Estricta: Código, Métodos, Variables y Documentación 100% en Español (Spanish Code Standard)

> 🇪🇸 **MANDATO DE CÓDIGO EN ESPAÑOL**: Para garantizar la comprensión total y facilitar el mantenimiento, **TODO EL CÓDIGO (nombres de métodos, funciones, structs, interfaces, variables, campos, constantes, errores y comentarios) DEBE ESCRIBIRSE EN ESPAÑOL**.

---

## 1. Identificadores de Código en Español (Spanish Identifiers)

1. **Nombres de Métodos y Funciones en Español**:
   - ❌ **Mal (en inglés)**: `GetUserByID()`, `CalculateTotal()`, `ValidateToken()`, `ProcessPayment()`.
   - ✅ **Bien (en español)**: `ObtenerUsuarioPorID()`, `CalcularMontoTotal()`, `ValidarToken()`, `ProcesarPago()`.

2. **Nombres de Structs, Interfaces y DTOs en Español**:
   - ❌ **Mal (en inglés)**: `UserRepository`, `AuthService`, `UserHandler`, `CreateUserDTO`.
   - ✅ **Bien (en español)**: `RepositorioUsuario`, `ServicioAutenticacion`, `ManejadorUsuario`, `DTOCrearUsuario`.

3. **Variables, Parámetros y Campos de Structs en Español**:
   - ❌ **Mal (en inglés)**: `email`, `passwordHash`, `createdAt`, `totalAmount`, `status`.
   - ✅ **Bien (en español)**: `correoElectronico`, `claveHash`, `fechaCreacion`, `montoTotal`, `estado`.

4. **Constantes y Errores Centinela en Español**:
   - ❌ **Mal (en inglés)**: `ErrUserNotFound = errors.New("user not found")`
   - ✅ **Bien (en español)**: `ErrUsuarioNoEncontrado = errors.New("el usuario solicitado no existe en el sistema")`

---

## 2. Nombres Claros, Expresivos y Sin Abreviaturas

- **Zero Abreviaturas Crípticas**:
  - ❌ **Mal**: `usr`, `cnt`, `calc()`, `dto`, `fn()`, `val`.
  - ✅ **Bien**: `usuario`, `cantidad`, `calcularMontoTotal()`, `objetoTransferenciaDatos`, `validarParametro()`.

---

## 3. Ejemplo Completo de Código Idiomático en Español

```go
package uso_caso

import (
    "context"
    "fmt"
    "github.com/Jhonatan-Code-dev/mi-proyecto/internal/domain"
)

// InterfazRepositorioUsuario define las operaciones de almacenamiento para la entidad Usuario.
type InterfazRepositorioUsuario interface {
    ObtenerPorID(ctx context.Context, empresaID int, usuarioID int) (*domain.Usuario, error)
    Guardar(ctx context.Context, usuario *domain.Usuario) error
}

// CasoUsoObtenerUsuario gestiona la lógica de negocio para consultar usuarios.
type CasoUsoObtenerUsuario struct {
    repositorioUsuario InterfazRepositorioUsuario
}

// NuevoCasoUsoObtenerUsuario crea una nueva instancia del caso de uso.
func NuevoCasoUsoObtenerUsuario(repo InterfazRepositorioUsuario) *CasoUsoObtenerUsuario {
    return &CasoUsoObtenerUsuario{
        repositorioUsuario: repo,
    }
}

// Ejecutar busca un usuario por su identificador asegurando el aislamiento por empresa.
func (c *CasoUsoObtenerUsuario) Ejecutar(ctx context.Context, empresaID int, usuarioID int) (*domain.Usuario, error) {
    if usuarioID <= 0 {
        return nil, fmt.Errorf("el identificador de usuario es inválido: %d", usuarioID)
    }

    usuario, err := c.repositorioUsuario.ObtenerPorID(ctx, empresaID, usuarioID)
    if err != nil {
        return nil, fmt.Errorf("fallo al consultar usuario en la base de datos: %w", err)
    }

    return usuario, nil
}
```
