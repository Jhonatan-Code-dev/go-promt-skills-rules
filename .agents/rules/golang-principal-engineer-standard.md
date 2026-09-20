# Regla Suprema: Estándar de Ingeniería de Elite (Go Core Team & Big Tech Principal Engineers)

> **PERSPECTIVA DE INGENIERÍA DE CLASE MUNDIAL**: Todo código, arquitectura, consulta y diseño generado debe responder con la excelencia, rigor y maestría combinada del **Go Core Team de Google** (creadores del lenguaje Go) y de los **Principal Software Engineers & Architects de las principales empresas tecnológicas del mundo** (Google, Microsoft, Meta, Amazon, Apple).

---

## 1. Principios de Diseño del Go Core Team (Simplicidad y Robustez)

1. **Simplicidad sobre Complejidad Innecesaria**:
   - Código claro y directo sobre abstracciones excesivas. Las soluciones deben ser elegantes, legibles y fáciles de mantener.
2. **Concurrencia Segura y Cero Goroutine Leaks**:
   - Diseñar pipelines de concurrencia usando canales y `sync` con garantías matemáticas de no estancamiento (Deadlock-Free).
3. **Manejo de Memoria Eficiente**:
   - Reducir asignaciones innecesarias en el Heap (`zero-allocation mindset`). Aprovechar la optimización del compilador de Go (Escape Analysis).

---

## 2. Rigor de Arquitectura de las Big Tech (Escalabilidad & Resiliencia)

1. **Clean Architecture & Monolito Modular**:
   - Desacoplamiento total de componentes de acuerdo a las mejores prácticas de microservicios y monolitos modulares de Google y Meta.
2. **Defensa en Profundidad y Seguridad Zero-Trust**:
   - Aislamiento absoluto por inquilino (`empresa_id`) en cada consulta y endpoint HTTP.
   - Manejo seguro de credenciales con Cookies `HttpOnly`, `Secure` y `SameSite`.
3. **Alto Rendimiento en Tiempo de Compilación (Wire & ENT)**:
   - Verificación de dependencias en tiempo de compilación con Google Wire para evitar fallas en tiempo de ejecución.
   - Uso avanzado de ENT ORM con generación de código tipado y prevención de problemas N+1.

---

## 3. Compromiso de Calidad Impecable

- **Cero Deuda Técnica**: Todo archivo generado es de calidad de producción final para sistemas críticos de alta demanda.
- **Formateo e Idioma Impecable**: Estricto cumplimiento del orden de archivo, nombres expresivos y documentación profesional en **Español**.
