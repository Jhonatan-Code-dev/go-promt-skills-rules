# Regla Estricta: Documentación Swagger / OpenAPI de Nivel Enterprise (Swagger Pro Standard)

> **ESTÁNDAR DE DOCUMENTACIÓN ENTERPRISE**: Todo endpoint HTTP expuesto en la aplicación DEBE estar minuciosamente documentado mediante anotaciones **Swagger / OpenAPI** de nivel profesional (estándar de empresas Big Tech). Toda la documentación de la API DEBE escribirse en **ESPAÑOL**, ser 100% precisa y carecer de emojis.

---

## 1. Reglas Mandatorias para Anotaciones Swagger

1. **Documentación Obligatoria en TODOS los Handlers**:
   - Todo método de presentación o controlador HTTP (Fiber v3) DEBE llevar su bloque de anotaciones Swagger completo encima de la declaración de la función.
2. **Idioma Mandatorio: ESPAÑOL**:
   - `@Summary`, `@Description`, `@Tags` y las descripciones de los parámetros DEBEN escribirse en español profesional.
3. **Mapeo Completo de Respuestas HTTP (200, 400, 401, 403, 404, 429, 500)**:
   - No limitar la documentación al caso exitoso (200/201). Se deben documentar explícitamente las respuestas de error esperadas con sus DTOs correspondientes.
4. **Seguridad y Autenticación Explicita**:
   - Indicar el esquema de seguridad utilizado (`@Security CookieAuth` para cookies HTTP-Only).

---

## 2. Plantilla Estándar Enterprise de Anotación Swagger

```go
// RegistrarNuevoUsuario godoc
// @Summary Registrar un nuevo usuario en la plataforma
// @Description Crea la cuenta de un usuario validando el correo electrónico y estableciendo la empresa asociada según la sesión.
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param datosUsuario body DTORegistroUsuario true "Datos requeridos para el registro del usuario"
// @Success 201 {object} RespuestaExitosaUsuario "Usuario creado exitosamente"
// @Failure 400 {object} RespuestaError "Datos de entrada inválidos o formato de correo incorrecto"
// @Failure 401 {object} RespuestaError "Sesión no autenticada o token inválido"
// @Failure 409 {object} RespuestaError "El correo electrónico ya se encuentra registrado"
// @Failure 429 {object} RespuestaError "Demasiadas peticiones enviadas, intente más tarde"
// @Failure 500 {object} RespuestaError "Error interno del servidor al procesar la solicitud"
// @Security CookieAuth
// @Router /api/v1/usuarios [post]
func (h *ManejadorUsuario) RegistrarNuevoUsuario(c fiber.Ctx) error {
    // ...
}
```

---

## 3. DTOs de Respuesta Estandarizados para Swagger

Todos los DTOs utilizados en `@Success` y `@Failure` deben incluir etiquetas `json:"..."`, `example:"..."` y comentarios explicativos en español:

```go
// RespuestaError representa la estructura estándar de error devuelta por la API.
type RespuestaError struct {
    Codigo   int    `json:"codigo" example:"400"`
    Mensaje  string `json:"mensaje" example:"El correo electrónico ingresado no es válido"`
    Detalle  string `json:"detalle,omitempty" example:"El campo correo_electronico debe contener un formato de email valido"`
}
```
