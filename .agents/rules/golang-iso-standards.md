# Regla Estricta: Estándares Internacionales ISO/IEC de Calidad y Seguridad de Software

> **CUMPLIMIENTO ISO OBLIGATORIO**: Todo el software desarrollado en Go debe ser diseñado, implementado y auditado bajo el marco de los estándares internacionales de calidad y seguridad de la información **ISO/IEC 25010**, **ISO/IEC 27001** e **ISO/IEC 5055**.

---

## 1. ISO/IEC 25010 - Modelo de Calidad del Producto de Software

1. **Eficiencia del Rendimiento (Performance Efficiency)**:
   - Tiempos de respuesta ultra bajos (latencia sub-100ms).
   - Uso eficiente de recursos (CPU, memoria, descriptores de archivo, pool de conexiones a BD).
2. **Mantenibilidad (Maintainability)**:
   - **Modularidad**: Separación limpia en Monolito Modular.
   - **Reutilización y Modificabilidad**: Inyección de dependencias con Google Wire e inversión de control.
   - **Capacidad de Prueba (Testability)**: Table-driven tests con `testify` y mocks para lograr >80% de cobertura.
3. **Fiabilidad (Reliability)**:
   - **Tolerancia a Fallos**: Captura limpia de errores con `%w` e interceptores de `recover()` para prevenir caídas del servidor.
   - **Recuperabilidad**: Manejo explícito de contextos con timeouts (`context.WithTimeout`).
4. **Seguridad (Security)**:
   - **Confidencialidad y Autenticidad**: Aislamiento estricto por `empresa_id` (Multi-Tenancy) y autenticación con JWT en Cookies HTTP-Only (`Secure`, `HttpOnly`, `SameSite`).

---

## 2. ISO/IEC 27001 - Seguridad de la Información

1. **Protección de Datos en Tránsito y Reposo**:
   - Todo campo sensible (`clave_hash`, `tokens`) en ENT ORM debe tener la anotación `.Sensitive()` para no filtrarse en logs ni respuestas JSON.
   - Comunicaciones encriptadas vía TLS/HTTPS y cifrado de datos sensibles.
2. **Control de Acceso y Aislamiento por Inquilino**:
   - Aislamiento absoluto Multi-Tenant en cada consulta de base de datos e interceptores de seguridad.

---

## 3. ISO/IEC 5055 - Medición Automatizada de Calidad del Código

- **Prevención de Vulnerabilidades de Código**: Cero punteros nulos sin validar (`nil check`), cero accesos a memoria fuera de límites, cero condiciones de carrera en goroutines (`sync.Mutex`).
