# Regla Estricta: Ciberseguridad Extrema, Defensa Hacker y Prevención de Saturación (Zero-Trust & Hacker Defense)

> 🛡️ **DEFENSA EN EL PEOR ESCENARIO**: Todo desarrollo en Go debe programarse bajo la mentalidad de un **Hacker Ético y Arquitecto de Ciberseguridad**. Asumir que el sistema estará expuesto a intentos constantes de intrusión, ataques de denegación de servicio (DDoS), saturación de memoria y explotación de vulnerabilidades.

---

## 1. Protección contra los Top 10 Riesgos OWASP y Ataques Comunes

1. **Prevención de Inyección (SQL/NoSQL/Command Injection)**:
   - Uso exclusivo de consultas parametrizadas con ENT ORM. Queda prohibida la concatenación manual de strings SQL.
2. **Aislamiento Absoluto de Datos (Prevención IDOR & Cross-Tenant)**:
   - Verificación obligatoria de `empresa_id` en cada consulta y endpoint (Regla Primordial Multi-Tenant).
3. **Protección contra CSRF, XSS y Hijacking**:
   - Tokens JWT almacenados únicamente en Cookies HTTP-Only (`HttpOnly: true`, `Secure: true`, `SameSite: "Lax"`).
4. **Prevención de Path Traversal & File Exposure**:
   - Sanitización rigurosa de nombres de archivo cargados usando `filepath.Base` y validación de extensiones permitidas.

---

## 2. Rate Limiting y Protección contra Saturación / DDoS en Fiber v3

- **Middleware Limitador de Tasa (Rate Limiter)**:
  - Todo endpoint público o sensible (Login, Registro, API endpoints) DEBE incluir un Rate Limiter para prevenir ataques de fuerza bruta y saturación del servidor:
  ```go
  app.Use("/api/v1/auth/iniciar-sesion", limiter.New(limiter.Config{
      Max:        5,               // Máximo 5 peticiones
      Expiration: 1 * time.Minute, // por minuto por IP
      LimitReached: func(c fiber.Ctx) error {
          return c.Status(429).JSON(fiber.Map{
              "error": "demasiados intentos de inicio de sesión, intente más tarde",
          })
      },
  }))
  ```

---

## 3. Prevención de Saturación y Congelamiento de VPS

- **Límites de Carga en Body (`BodyLimit`)**:
  - Limitar el tamaño máximo del payload HTTP a 4MB-10MB para evitar ataques de desbordamiento de memoria por peticiones gigantescas.
- **Timeouts de Lectura y Escritura**:
  - Configurar timeouts de conexión en Fiber v3 (`ReadTimeout: 10 * time.Second`, `WriteTimeout: 10 * time.Second`) para evitar ataques de conexiones lentas (Slowloris).
