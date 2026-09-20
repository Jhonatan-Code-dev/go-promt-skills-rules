# Regla Estricta: Orden, Limpieza y Estructura Obsesiva en el Código (Golang Code Order)

> **PRINCIPIO FUNDAMENTAL**: El código debe ser impecable, exageradamente ordenado, modular y libre de desorden ("cero código sucio"). Cada archivo debe parecer escrito por un arquitecto obsesionado con la perfección visual y estructural.

---

## 1. Orden Obsesivo de Archivos `.go` (Layout Interno Mandatorio)

Todo archivo de Go DEBE seguir estrictamente la siguiente jerarquía de arriba a abajo:

1. **Declaración del Package** (`package ...`)
2. **Importaciones Agrupadas y Separadas** en 3 bloques distintos:
   ```go
   import (
       // 1. Standard Library (Librería estándar de Go)
       "context"
       "fmt"
       "time"

       // 2. Dependencias Externas (Third-Party)
       "github.com/google/uuid"

       // 3. Modulos Internos del Proyecto
       "github.com/Jhonatan-Code-dev/mi-proyecto/internal/domain"
   )
   ```
3. **Constantes y Enums** (Agrupadas en `const (...)` con nombres auto-explicativos).
4. **Errores Centinela** (Agrupados en `var (...)` como `ErrUserNotFound = errors.New(...)`).
5. **Interfaces** (Definiciones claras y concisas).
6. **Structs y DTOs** (Campos ordenados por tamaño/tipo o agrupamiento lógico).
7. **Constructores** (`func New[NombreStruct](...) *[NombreStruct]`).
8. **Métodos Exportados / Públicos** (Ordenados por flujo de ejecución principal).
9. **Métodos Privados / Auxiliares** (Ubicados inmediatamente debajo de las funciones que los invocan).

---

## 2. Reglas de Limpieza e Higiene del Código

- 🚫 **Cero Código Muerto**: Prohibido dejar funciones no utilizadas, variables sin usar o imports sobrantes.
- 🚫 **Cero Números Mágicos**: Todo número o string literal que no sea evidente debe extraerse a una constante nominada.
- 🚫 **Cero Comentarios Sucios**: Prohibido dejar bloques de código comentados (`// func oldCode() {...}`). Si no se usa, SE ELIMINA.
- 📏 **Límite de Tamaño por Archivo**: Un archivo no debe exceder las 250-300 líneas. Si crece más, se debe modularizar en archivos independientes con responsabilidades únicas.
- 📏 **Límite de Tamaño por Función**: Ninguna función debe superar las 30 líneas. Dividir en sub-funciones auxiliares privadas.

---

## 3. Estructura de Directorios Limpia
- Directorios en minúsculas, separados por guiones si es necesario (`user-management`, `order-processing`).
- Cada carpeta debe tener una sola responsabilidad técnica o de dominio.
- Ningún archivo suelto en la raíz excepto `main.go`, `go.mod`, `go.sum`, `README.md` y configuraciones globales.
