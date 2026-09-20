# Go Agent Rules & Skills (`go-promt-skills-rules`)

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)](LICENSE)
[![Agents Ready](https://img.shields.io/badge/AI%20Agents-Antigravity%20%7C%20Cursor%20%7C%20Windsurf-blueviolet?style=for-the-badge)](https://github.com/Jhonatan-Code-dev/go-promt-skills-rules)

Suite estándar de Reglas y Habilidades para desarrollo backend de alto nivel en Go (Golang), optimizada para arquitectura empresarial, seguridad de vanguardia y máxima eficiencia.

---

## Instalación Rápida (1 Solo Comando)

Elige tu método preferido para instalar automáticamente las reglas y habilidades en la carpeta `.agents/` de tu proyecto Go actual:

### Opción 1: Vía Go CLI (Recomendado)
```bash
go run github.com/Jhonatan-Code-dev/go-promt-skills-rules@latest
```

### Opción 2: Vía PowerShell (Windows)
```powershell
iwr -useb https://raw.githubusercontent.com/Jhonatan-Code-dev/go-promt-skills-rules/main/install.ps1 | iex
```

### Opción 3: Vía Bash (Linux / macOS)
```bash
curl -fsSL https://raw.githubusercontent.com/Jhonatan-Code-dev/go-promt-skills-rules/main/install.sh | bash
```

---

## Contenido Incluido

### Reglas Estándar (`.agents/rules/`)

| Regla | Descripción |
| :--- | :--- |
| [`no-emojis-standard.md`](.agents/rules/no-emojis-standard.md) | **Prohibición Absoluta de Emojis**: Cero emojis en documentación, README, comentarios de código, logs o mensajes de error. Formato sobrio y profesional. |
| [`golang-swagger-documentation.md`](.agents/rules/golang-swagger-documentation.md) | **Documentación Swagger / OpenAPI Enterprise**: Documentación Swagger Pro completa en español para cada handler HTTP (Fiber v3) mapeando respuestas (200, 400, 401, 404, 500). |
| [`golang-env-config-secrets.md`](.agents/rules/golang-env-config-secrets.md) | **Cero Secretos Hardcodeados & `.env`**: Carga estricta de variables de entorno, validación al iniciar y cero cadenas de conexión o claves secretas en código. |
| [`golang-design-patterns.md`](.agents/rules/golang-design-patterns.md) | **Patrones y Anti-Patrones de Diseño**: Patrón Fábrica, Singleton a prueba de hilos con `sync.Once`, Decorador por composición y prohibición de retornar `nil` sin error. |
| [`golang-security-hacker-defense.md`](.agents/rules/golang-security-hacker-defense.md) | **Ciberseguridad Extrema & Defensa Hacker**: OWASP Top 10, Rate Limiting, prevención de saturación DDoS, IDOR, XSS, CSRF y Path Traversal. |
| [`golang-low-resource-vps-efficiency.md`](.agents/rules/golang-low-resource-vps-efficiency.md) | **Eficiencia en VPS de Bajos Recursos**: Optimización para 512MB RAM / 1 vCPU, paginación por cursor, cero cargas masivas en memoria y Clean Code (SOLID). |
| [`golang-anti-patterns-prevention.md`](.agents/rules/golang-anti-patterns-prevention.md) | **Prevención de Anti-Patrones**: Prohibición de errores ignorados (`_ = err`), abuso de `panic()`, goroutines huérfanas, data races y paquetes `utils/common`. |
| [`golang-test-organization.md`](.agents/rules/golang-test-organization.md) | **Organización Centralizada de Pruebas (`pruebas/`)**: Pruebas estructuradas en `pruebas/unitarias/`, `integracion/` y `rendimiento/` documentando escenario y comandos. |
| [`golang-iso-standards.md`](.agents/rules/golang-iso-standards.md) | **Cumplimiento ISO/IEC**: Estándares internacionales ISO/IEC 25010 (Calidad), 27001 (Seguridad) y 5055 (Medición de código). |
| [`ai-truth-transparency.md`](.agents/rules/ai-truth-transparency.md) | **Veracidad Absoluta y Transparencia**: Prohibido mentir, asumir o falsear resultados. Inspección de código completa de principio a fin. |
| [`agile-scrum-master-planner.md`](.agents/rules/agile-scrum-master-planner.md) | **Scrum Master & Planificación Ágil**: Rol de Scrum Master en la descomposición de Historias de Usuario, Criterios de Aceptación (DoD) y gestión de riesgos. |
| [`golang-principal-engineer-standard.md`](.agents/rules/golang-principal-engineer-standard.md) | **Estándar de Ingeniería de Elite**: Rigor técnico del Go Core Team (Google) y Principal Engineers de Big Tech. |
| [`golang-multitenancy.md`](.agents/rules/golang-multitenancy.md) | **Regla Primordial Multi-Tenant**: Aislamiento estricto de datos en TODO endpoint por `empresa_id`, índices compuestos y filtros automáticos. |
| [`golang-db-performance.md`](.agents/rules/golang-db-performance.md) | **Rendimiento Extremo & Escalabilidad BD**: Consultas ultra veloces por ID, índices B-Tree en Foreign Keys, solución N+1 con Eager Loading y paginación por cursor. |
| [`golang-tech-stack-ent-fiber.md`](.agents/rules/golang-tech-stack-ent-fiber.md) | **ENT, Fiber v3, Swagger, JWT & CORS**: Tablas y campos de ENT en español, Fiber v3, JWT en Cookies HTTP-Only, CORS seguro y documentación Swag. |
| [`golang-naming-spanish.md`](.agents/rules/golang-naming-spanish.md) | **Nombres Claros y Español Mandatorio**: Exige nombres claros y auto-explicativos. Código, métodos, variables, comentarios y errores SIEMPRE en Español. |
| [`golang-modular-monolith.md`](.agents/rules/golang-modular-monolith.md) | **Monolito Modular & Google Wire**: Arquitectura modular con Clean Architecture en `internal/modules/`, desacoplamiento total e inyección con `google/wire`. |
| [`golang-code-order.md`](.agents/rules/golang-code-order.md) | **Orden y Limpieza Obsesiva**: Regla estricta de estructura interna de archivo, grupos de imports, cero código muerto y cero números mágicos. |

---

## Habilidades Interactivas (`.agents/skills/`)

| Habilidad | Descripción |
| :--- | :--- |
| [`golang-env-auditor`](.agents/skills/golang-env-auditor/SKILL.md) | Audita el código Go en busca de credenciales, cadenas de conexión o secretos hardcodeados, verificando la carga correcta de `.env` y `.gitignore`. |
| [`golang-patterns-expert`](.agents/skills/golang-patterns-expert/SKILL.md) | Experto en patrones de diseño idiomáticos en Go (Fábrica, Singleton a prueba de hilos con `sync.Once`, Decorador) y corrección de anti-patrones. |
| [`golang-anti-pattern-auditor`](.agents/skills/golang-anti-pattern-auditor/SKILL.md) | Audita y detecta malas prácticas y anti-patrones en Go (errores ignorados, goroutines huérfanas, data races, paquetes `utils/common`, captura en bucles y abuso de punteros). |
| [`golang-test-runner`](.agents/skills/golang-test-runner/SKILL.md) | Ejecuta y audita la suite centralizada de pruebas organizadas dentro de la carpeta principal `pruebas/`. |
| [`golang-iso-compliance-audit`](.agents/skills/golang-iso-compliance-audit/SKILL.md) | Audita el cumplimiento de los estándares internacionales ISO/IEC 25010 (Calidad), ISO/IEC 27001 (Seguridad) e ISO/IEC 5055 en proyectos Go. |
| [`agile-sprint-planner`](.agents/skills/agile-sprint-planner/SKILL.md) | Actúa como Scrum Master y Technical Program Manager para estructurar planes de Sprint, Historias de Usuario y Criterios de Aceptación. |
| [`golang-multitenant-check`](.agents/skills/golang-multitenant-check/SKILL.md) | Audita la seguridad y aislamiento de datos Multi-Tenant, verificando el campo `empresa_id`, índices compuestos y filtros automáticos. |
| [`golang-db-performance-expert`](.agents/skills/golang-db-performance-expert/SKILL.md) | Experto en análisis y optimización extrema de rendimiento en consultas SQL, índices B-Tree, resolución N+1 y paginación por cursor. |
| [`golang-ent-schema-gen`](.agents/skills/golang-ent-schema-gen/SKILL.md) | Genera esquemas de ENT ORM con tablas y campos explícitamente en español, anotaciones de tabla y relaciones. |
| [`golang-module-gen`](.agents/skills/golang-module-gen/SKILL.md) | Genera la estructura completa de un nuevo módulo en Monolito Modular con Clean Architecture y `provider.go` para Google Wire. |
| [`golang-order-enforcer`](.agents/skills/golang-order-enforcer/SKILL.md) | Reorganiza y limpia de forma obsesiva cualquier archivo Go siguiendo la jerarquía estricta de código. |
| [`golang-audit`](.agents/skills/golang-audit/SKILL.md) | Realiza auditorías automatizadas de arquitectura, concurrencia, seguridad y calidad de código. |
| [`golang-test-gen`](.agents/skills/golang-test-gen/SKILL.md) | Genera pruebas unitarias e integración usando Table-Driven Tests y mocks con `testify`. |

---

## Compatibilidad con Agentes de IA

Esta configuración sigue el estándar universal de agentes agenticos y es compatible out-of-the-box con:
- Google Antigravity / Gemini CLI
- Cursor IDE (`.cursorrules` / `.agents`)
- Windsurf / Cascade
- Claude Code / GitHub Copilot Agent

---

## Cómo Actualizar

Para actualizar tus reglas a la versión más reciente del repositorio, simplemente vuelve a ejecutar:

```bash
go run github.com/Jhonatan-Code-dev/go-promt-skills-rules@latest
```

---

## Licencia

Este proyecto está bajo la Licencia [MIT](LICENSE). Libre uso comercial y privado.
