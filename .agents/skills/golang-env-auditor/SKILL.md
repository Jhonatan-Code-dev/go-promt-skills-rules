---
name: golang-env-auditor
description: Audita el código Go en busca de credenciales, cadenas de conexión o secretos hardcodeados, verificando la carga correcta de .env y .gitignore.
---

# Skill: Auditor de Variables de Entorno y Secretos (Golang Env & Secrets Auditor)

Esta habilidad se activa cuando el usuario solicita verificar que no existan datos sensibles hardcodeados en el código Go y que la configuración por `.env` sea correcta.

## Protocolo de Auditoría de Secretos y Configuración

Al auditar un proyecto Go:

### 1. Detección de Hardcoding
- Escanea todo el código Go en busca de literales de strings que contengan:
  - Cadenas de conexión a BD (`postgres://`, `mysql://`, `redis://`).
  - Claves secretas de JWT o contraseñas.
  - Puertos o IP escritas directamente en el código.
  - API Keys o tokens.

### 2. Verificación de Carga de `.env`
- Comprueba la existencia del paquete de configuración que lea las variables de entorno.
- Verifica que la aplicación falle si faltan variables críticas (`DATABASE_URL`, `JWT_SECRET`).

### 3. Verificación de `.env.example` y `.gitignore`
- Asegura que exista un archivo `.env.example` con los nombres de las variables sin valores sensibles.
- Confirma que `.env` esté listado en `.gitignore`.

### 4. Reporte de Auditoría de Secretos
Genera un informe con:
- **Secretos Hardcodeados Detectados** (Línea de código y archivo).
- **Variables Faltantes en `.env.example`**.
- **Código Refactorizado Propuesto**.
