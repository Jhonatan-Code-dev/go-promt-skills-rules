# Regla Estricta: Tech Stack (ENT ORM, Fiber v3, Wire, Swagger, JWT & Cookies)

> **ESTÁNDAR TECNOLÓGICO DE LA APLICACIÓN**: Todo desarrollo backend se construye usando **ENT ORM**, **Fiber v3**, **Google Wire**, **Swagger**, **JWT en Cookies HTTP-Only** y **CORS Seguro**.

---

## 1. ENT ORM (`entgo.io/ent`) - Nombres y Esquemas en Español

1. **Tablas y Campos en Español**:
   - Todo esquema de ENT debe definir el nombre exacto de la tabla en español usando `ent.Table(...)`.
   - Los campos (`ent.Field`) deben nombrarse en español claro con notación `snake_case`.
   ```go
   // ent/schema/usuario.go
   package schema

   import (
       "time"
       "entgo.io/ent"
       "entgo.io/ent/dialect/entsql"
       "entgo.io/ent/schema"
       "entgo.io/ent/schema/field"
       "entgo.io/ent/schema/edge"
   )

   type Usuario struct {
       ent.Schema
   }

   func (Usuario) Annotations() []schema.Annotation {
       return []schema.Annotation{
           entsql.Annotation{Table: "usuarios"}, // Nombre explícito de la tabla en Español
       }
   }

   func (Usuario) Fields() []ent.Field {
       return []ent.Field{
           field.String("nombre_completo").NotEmpty().Comment("Nombre completo del usuario"),
           field.String("correo_electronico").Unique().NotEmpty().Comment("Correo electrónico de acceso"),
           field.String("clave_hash").Sensitive().Comment("Hash de la contraseña"),
           field.Enum("rol").Values("ADMINISTRADOR", "CLIENTE", "OPERADOR").Default("CLIENTE"),
           field.Time("fecha_creacion").Default(time.Now).Immutable(),
       }
   }
   ```

2. **Desacoplamiento de ENT**:
   - Prohibido exponer las entidades generadas por ENT (`*ent.Usuario`) fuera de la capa de `infrastructure`. Se deben mapear a Entidades puras de Dominio en `domain/`.

---

## 2. Fiber v3 (`github.com/gofiber/fiber/v3`) & CORS

1. **Servidor HTTP con Fiber v3**:
   - Configurar Fiber v3 con límites de cuerpo (body limit), manejo global de errores y timeouts de lectura/escritura.
2. **Configuración de CORS Seguro**:
   ```go
   app := fiber.New(fiber.Config{
       ErrorHandler: ManejadorErroresGlobal,
   })

   app.Use(cors.New(cors.Config{
       AllowOrigins:     []string{"https://mi-dominio.com"},
       AllowCredentials: true, // Requerido para envío de Cookies HTTP-Only
       AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
       AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
   }))
   ```

---

## 3. Autenticación con JWT & Cookies HTTP-Only

1. **Almacenamiento de JWT en Cookie HTTP-Only**:
   - **Prohibido enviar JWTs en el cuerpo JSON para guardarse en localStorage**. Los tokens de acceso y refresh DEBEN viajar en Cookies Seguras `HttpOnly`.
   ```go
   c.Cookie(&fiber.Cookie{
       Name:     "token_acceso",
       Value:    tokenJWT,
       Expires:  time.Now().Add(24 * time.Hour),
       HTTPOnly: true, // Previene ataques XSS
       Secure:   true, // Requiere HTTPS en producción
       SameSite: "Lax",
   })
   ```
2. **Middleware de Autenticación**:
   - Extrae el token JWT desde la cookie `token_acceso`, lo valida e inyecta el `IDUsuario` en el contexto local de la petición `c.Locals("usuario_id", claims.IDUsuario)`.

---

## 4. Documentación Automática con Swagger (`swaggo/swag`)

- Todo handler de Fiber v3 DEBE tener anotaciones Swagger explicativas escritas en **Español**:
```go
// IniciarSesion godoc
// @Summary Iniciar sesión de usuario
// @Description Autentica al usuario con correo y clave, estableciendo una cookie HTTP-Only con el token JWT.
// @Tags Autenticación
// @Accept json
// @Produce json
// @Param credenciales body DTOInicioSesion true "Credenciales de acceso"
// @Success 200 {object} RespuestaExitosa "Sesión iniciada correctamente"
// @Failure 400 {object} RespuestaError "Datos de entrada inválidos"
// @Failure 401 {object} RespuestaError "Credenciales incorrectas"
// @Router /api/v1/auth/login [post]
func (h *ManejadorAutenticacion) IniciarSesion(c fiber.Ctx) error {
    // ...
}
```

---

## 5. Integración con Google Wire

- Inyectar el cliente `*ent.Client` y los servicios de Fiber v3 a través de los ProviderSets de Wire en `provider.go`.
