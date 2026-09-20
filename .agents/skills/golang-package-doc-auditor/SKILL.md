---
name: golang-package-doc-auditor
description: Audita y garantiza el cumplimiento de la documentación de paquetes en Go (Staticcheck ST1000 y Godoc en español), verificando que cada paquete posea comentarios explicativos.
---

# Skill: Auditor de Documentación de Paquetes en Go (Golang Package Doc Auditor)

Esta habilidad se activa cuando el usuario solicita auditar, verificar o corregir la documentación de paquetes en Go, resolver la advertencia de linter `ST1000` (at least one file in a package should have a package comment) o garantizar el estándar de comentarios Godoc en español.

## Protocolo de Auditoría y Corrección ST1000

Al ejecutar una auditoría de documentación de paquetes Go:

### 1. Escaneo de Paquetes Sin Documentar
- Examina todos los directorios que contengan archivos `.go` en el proyecto.
- Verifica si al menos un archivo `.go` dentro de cada directorio incluye un comentario que inicie con `// Package <nombre>` previo a la sentencia `package`.
- Ejecuta `staticcheck ./...` o inspecciona los avisos del compilador/linter para identificar infracciones ST1000.

### 2. Validación de Formato y Estándares
- **Idioma**: Asegura que la descripción esté redactada 100% en español.
- **Formato Godoc**: Confirma que el comentario empiece con la palabra `Package` seguida del nombre exacto del paquete.
- **Sin Emojis**: Garantiza la ausencia total de emojis en los comentarios.

### 3. Estrategia de Corrección Automática
- Si un paquete carece de comentario de paquete:
  1. Identifica el archivo principal del paquete (por ejemplo, `dominio.go`, `handler.go` o el archivo homónimo a la carpeta).
  2. En caso de paquetes de gran volumen, crea un archivo `doc.go` dentro de dicho paquete.
  3. Agrega la cabecera explicativa respetando el estándar:
     ```go
     // Package <nombre_paquete> proporciona la funcionalidad necesaria para ...
     package <nombre_paquete>
     ```

### 4. Verificación de Cierre
- Ejecuta nuevamente las pruebas de compilación y linters para confirmar que la advertencia `ST1000` ha quedado resuelta sin afectar el código ejecutable.
