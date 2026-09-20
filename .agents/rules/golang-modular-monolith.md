# Regla Estricta: Monolito Modular, Clean Architecture e Inyección con Google Wire

> **PRINCIPIO DE DESACOPLAMIENTO TOTAL**: El sistema debe estructurarse como un **Monolito Modular con Clean Architecture**. Cada módulo es autónomo e independiente. Si mañana se cambia una base de datos, una librería externa o un módulo completo, **ningún otro módulo o lógica de negocio debe sufrir cambios**.

---

## 1. Estructura de Monolito Modular (`internal/modules/`)

Cada funcionalidad de negocio habita en su propio módulo aislado bajo `internal/modules/[nombre_modulo]/`:

```text
internal/
├── modules/
│   ├── auth/
│   │   ├── domain/           # Entidades puras, Value Objects, Errores e Interfaces del repositorio
│   │   ├── usecase/          # Casos de uso de autenticación
│   │   ├── infrastructure/   # Repositorios (GORM/SQL/Redis), adaptadores externos
│   │   ├── presentation/     # Handlers HTTP/gRPC, DTOs
│   │   └── provider.go       # Wire ProviderSet del módulo Auth
│   ├── users/
│   │   ├── domain/
│   │   ├── usecase/
│   │   ├── infrastructure/
│   │   ├── presentation/
│   │   └── provider.go       # Wire ProviderSet del módulo Users
│   └── order/
│       └── ...
├── platform/                 # Servicios globales compartidos (Database, Logger, Config, EventBus)
└── cmd/
    └── api/
        ├── wire.go           # Inyección de dependencias global con Google Wire
        ├── wire_gen.go       # Código generado por Wire (NO EDITAR MANUALMENTE)
        └── main.go
```

---

## 2. Reglas Estrictas de Desacoplamiento

1. **Aislamiento de Dominio (`domain/`)**:
   - `domain/` no puede importar nada de `usecase`, `infrastructure`, `presentation` ni bibliotecas de terceros (salvo stdlib).
   - Las entidades de dominio se comunican con el exterior **únicamente a través de interfaces**.
2. **Comunicación Inter-Módulos**:
   - Prohibido que el módulo `order` acceda a la base de datos o structs internos del módulo `users`.
   - Si `order` necesita información de `users`, debe hacerlo a través de una **Interface de Servicio** expuesta por `users` o emitiendo **Eventos de Dominio** (Event-Driven / PubSub).
3. **Sustituibilidad de Infraestructura**:
   - Cambiar PostgreSQL por MongoDB o Firebase solo requiere crear una nueva implementación de la interfaz `domain.Repository` en `infrastructure/` y actualizar la inyección en Wire.

---

## 3. Inyección de Dependencias con Google Wire (`github.com/google/wire`)

- **Prohibido el Cableado Manual en `main.go`**: Toda la construcción del árbol de dependencias DEBE manejarse mediante Google Wire (`google/wire`).
- **ProviderSets por Módulo**: Cada módulo debe exponer un `ProviderSet` exportado en su archivo `provider.go`:
  ```go
  // internal/modules/users/provider.go
  package users

  import (
      "github.com/google/wire"
      "github.com/Jhonatan-Code-dev/mi-proyecto/internal/modules/users/infrastructure"
      "github.com/Jhonatan-Code-dev/mi-proyecto/internal/modules/users/presentation"
      "github.com/Jhonatan-Code-dev/mi-proyecto/internal/modules/users/usecase"
  )

  var ProviderSet = wire.NewSet(
      infrastructure.NewUserRepository,
      wire.Bind(new(domain.UserRepository), new(*infrastructure.UserRepositoryImpl)),
      usecase.NewUserUseCase,
      presentation.NewUserHandler,
  )
  ```
- **Inyector Principal (`cmd/api/wire.go`)**:
  ```go
  //go:build wireinject
  // +build wireinject

  package main

  import (
      "github.com/google/wire"
      "github.com/Jhonatan-Code-dev/mi-proyecto/internal/modules/auth"
      "github.com/Jhonatan-Code-dev/mi-proyecto/internal/modules/users"
  )

  func InitializeApp() (*App, func(), error) {
      wire.Build(
          users.ProviderSet,
          auth.ProviderSet,
          NewApp,
      )
      return nil, nil, nil
  }
  ```
- Tras modificar dependencias, siempre ejecutar `wire ./cmd/api` para compilar `wire_gen.go`.
