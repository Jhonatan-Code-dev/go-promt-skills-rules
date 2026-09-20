---
name: golang-multitenant-check
description: Audita la seguridad y aislamiento de datos Multi-Tenant, verificando el campo empresa_id, índices compuestos y filtros automáticos en context.
---

# Skill: Auditoría de Seguridad Multi-Tenant (Golang Multi-Tenant Auditor)

Esta habilidad se activa cuando el usuario quiere auditar o asegurar que un módulo, repositorio o consulta cumpla con el aislamiento estricto por inquilino (`empresa_id`).

## Protocolo de Auditoría Multi-Tenant

Al revisar un esquema, handler o repositorio, sigue estas comprobaciones:

### 1. Verificación de Esquema ENT
- Confirma que todas las entidades pertenecientes a un inquilino posean el campo `empresa_id`.
- Revisa que los índices de ENT comiencen con `empresa_id` (`index.Fields("empresa_id", ...)`).

### 2. Extracción de `empresa_id` desde Token JWT
- Verifica que el `empresa_id` no provenga de parámetros de URL ni del body del JSON enviado por el usuario, sino exclusivamente de las `Locals` / `Cookie` autenticada del contexto.

### 3. Filtro Automático por Contexto
- Comprueba que los repositorios de datos consuman `empresa_id` desde el `context.Context` mediante un middleware o interceptor de ENT.

### 4. Reporte de Seguridad Multi-Tenant
- **Entidades Seguras Identificadas**.
- **Fugas de Datos Potenciales** (Consultas sin filtro por `empresa_id`).
- **Plan de Corrección**.
