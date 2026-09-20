# REGLAS MANDATORIAS Y OBLIGATORIAS DEL PROYECTO (AGENTS.MD)

> **ATENCIÓN A TODA INTELIGENCIA ARTIFICIAL (Antigravity, Cursor, Windsurf, Claude, Copilot, Gemini)**: Este proyecto está regido por un conjunto de **Reglas Mandatorias e Inviolables** ubicadas en `.agents/rules/` y **Habilidades Específicas** en `.agents/skills/`.
> DEBES leer y respetar estrictamente cada una de estas reglas ANTES de escribir, revisar, modificar o auditar código en este repositorio.

---

## Lista de Reglas Mandatorias de Cumplimiento Obligatorio

1. **PROHIBICIÓN ABSOLUTA DE EMOJIS**: [`no-emojis-standard.md`](.agents/rules/no-emojis-standard.md)
   - **MANDATO OBLIGATORIO**: Cero emojis en documentación, comentarios, logs o código. Estilo sobrio y profesional.
2. **DOCUMENTACIÓN SWAGGER / OPENAPI ENTERPRISE**: [`golang-swagger-documentation.md`](.agents/rules/golang-swagger-documentation.md)
   - **MANDATO OBLIGATORIO**: Documentación Swagger Pro completa en español para cada handler HTTP (Fiber v3) mapeando códigos de respuesta (200, 400, 401, 404, 500) y esquemas DTO.
3. **INSPECCIÓN COMPLETA Y VERDAD ABSOLUTA**: [`ai-truth-transparency.md`](.agents/rules/ai-truth-transparency.md)
   - **MANDATO OBLIGATORIO**: Leer los archivos DE PRINCIPIO A FIN (sin revisiones parciales ni vista de túnel). Probar empíricamente todo cambio ejecutando `go build` / `go test` antes de responder al usuario.
4. **CÓDIGO Y NOMBRES 100% EN ESPAÑOL**: [`golang-naming-spanish.md`](.agents/rules/golang-naming-spanish.md)
   - **MANDATO OBLIGATORIO**: Todos los métodos, funciones, structs, interfaces, variables, campos, errores y comentarios DEBEN escribirse en ESPAÑOL.
5. **VERIFICACIÓN MULTI-TENANT EN CADA ENDPOINT**: [`golang-multitenancy.md`](.agents/rules/golang-multitenancy.md)
   - **REGLA PRIMORDIAL #1**: Todo endpoint y consulta DEBE verificar `empresa_id` extraído del token JWT en la cookie HTTP-Only. Prohibido filtrar datos entre empresas.
6. **CERO SECRETOS HARDCODEADOS**: [`golang-env-config-secrets.md`](.agents/rules/golang-env-config-secrets.md)
   - Cadenas de conexión, secretos y puertos DEBEN cargarse obligatoriamente desde `.env`.
7. **CIBERSEGURIDAD EXTREMA & DEFENSA HACKER**: [`golang-security-hacker-defense.md`](.agents/rules/golang-security-hacker-defense.md)
   - OWASP Top 10, Rate Limiting en Fiber v3, prevención IDOR, XSS, CSRF y sanitización de archivos.
8. **EFICIENCIA EN VPS DE BAJOS RECURSOS (512MB RAM)**: [`golang-low-resource-vps-efficiency.md`](.agents/rules/golang-low-resource-vps-efficiency.md)
   - Cero cargas masivas en RAM. Paginación basada en cursor (`Where(id > ultimo_id)`).
9. **MONOLITO MODULAR, CLEAN ARCHITECTURE & GOOGLE WIRE**: [`golang-modular-monolith.md`](.agents/rules/golang-modular-monolith.md)
   - Aislamiento en `internal/modules/`, interfaces puras en dominio y `provider.go` para Google Wire.
10. **STACK TECNOLÓGICO ENT ORM & FIBER V3**: [`golang-tech-stack-ent-fiber.md`](.agents/rules/golang-tech-stack-ent-fiber.md)
   - Tablas y campos de ENT en español, Fiber v3 y documentación Swagger en español.
11. **ORDEN Y LIMPIEZA OBSESIVA DE CÓDIGO**: [`golang-code-order.md`](.agents/rules/golang-code-order.md)
   - Layout interno estricto de archivos `.go` (Package ➔ Imports en 3 bloques ➔ Const ➔ Var ➔ Interface ➔ Struct ➔ New() ➔ Métodos).
12. **PATRONES Y PREVENCIÓN DE ANTI-PATRONES**: [`golang-design-patterns.md`](.agents/rules/golang-design-patterns.md) y [`golang-anti-patterns-prevention.md`](.agents/rules/golang-anti-patterns-prevention.md)
   - Patrón Fábrica, Singleton (`sync.Once`), Decorador. Prohibido `_ = err`, `panic()` en negocio o data races.
13. **PRUEBAS CENTRALIZADAS EN CARPETA `pruebas/`**: [`golang-test-organization.md`](.agents/rules/golang-test-organization.md)
   - Pruebas ordenadas en `pruebas/unitarias/`, `integracion/` y `rendimiento/` documentando qué prueban y cómo se ejecutan.
14. **ESTÁNDARES ISO/IEC (25010, 27001, 5055)**: [`golang-iso-standards.md`](.agents/rules/golang-iso-standards.md)
   - Garantía de calidad, mantenibilidad y seguridad certificable.
15. **SCRUM MASTER & PLANIFICACIÓN ÁGIL**: [`agile-scrum-master-planner.md`](.agents/rules/agile-scrum-master-planner.md)
   - Descomposición en Historias de Usuario y Criterios de Aceptación (DoD).
16. **ESTÁNDAR DE INGENIERÍA GO CORE TEAM (GOOGLE)**: [`golang-principal-engineer-standard.md`](.agents/rules/golang-principal-engineer-standard.md)
   - Rigor técnico de clase mundial.
17. **DOCUMENTACIÓN DE PAQUETES GO Y CUMPLIMIENTO ST1000**: [`golang-package-documentation.md`](.agents/rules/golang-package-documentation.md)
   - **MANDATO OBLIGATORIO**: Todo paquete DEBE incluir al menos un archivo con comentario de paquete en español (`// Package <nombre> ...`) eliminando advertencias ST1000 de linters.
18. **ALMACENAMIENTO DE CAMPOS OPCIONALES COMO NULL**: [`golang-db-null-over-empty-standard.md`](.agents/rules/golang-db-null-over-empty-standard.md)
   - **MANDATO OBLIGATORIO**: Prohibido almacenar cadenas vacías `""` o en blanco. Todo campo opcional DEBE declararse con `.Optional().Nillable()` en ENT ORM y guardarse como `NULL` en SQL.


