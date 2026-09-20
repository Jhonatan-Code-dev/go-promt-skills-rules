---
name: golang-iso-compliance-audit
description: Audita el cumplimiento de los estándares internacionales ISO/IEC 25010 (Calidad), ISO/IEC 27001 (Seguridad) e ISO/IEC 5055 en proyectos Go.
---

# Skill: Auditoría de Cumplimiento ISO/IEC (Golang ISO Compliance Auditor)

Esta habilidad se activa cuando el usuario solicita verificar que el proyecto cumpla con los estándares internacionales ISO de calidad de software y seguridad de la información.

## Protocolo de Auditoría ISO

Al ejecutar la auditoría ISO en una base de código Go:

### 1. Evaluación ISO/IEC 25010 (Calidad y Mantenibilidad)
- Evalúa la modularidad del código, arquitectura en capas, inyección con Wire y cobertura de pruebas unitarias.

### 2. Evaluación ISO/IEC 27001 (Seguridad)
- Audita el aislamiento Multi-Tenant (`empresa_id`), el manejo de JWT en Cookies HTTP-Only y el etiquetado `.Sensitive()` de campos en ENT ORM.

### 3. Evaluación ISO/IEC 5055 (Concurrencia y Fiabilidad)
- Escanea el código en busca de posibles `nil pointer dereference`, Data Races, Goroutine Leaks o falta de timeouts en `context.Context`.

### 4. Informe de Certificación ISO Interno
Genera un informe estructurado indicando:
- **Nivel de Cumplimiento ISO (%)**.
- **Puntos Críticos de No Conformidad**.
- **Plan de Acción de Mitigación**.
