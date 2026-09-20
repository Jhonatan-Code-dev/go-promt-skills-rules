# Regla Estricta: Prohibición de Malas Prácticas y Anti-Patrones en Go (Anti-Patterns Prevention)

> 🚫 **PROHIBICIÓN DE CÓDIGO SUCIO Y ANTI-PATRONES**: Queda estrictamente prohibido incurrir en cualquiera de las siguientes malas prácticas comunes de desarrollo en Go. Todo código debe ser idiomático, seguro, eficiente e impecable.

---

## 1. Prohibición en Manejo de Errores

1. **Jamás Ignorar Errores (`_ = err`)**:
   - ❌ **Mal**: `_, err := funcion() // Omitir verificación`
   - ✅ **Bien**:
     ```go
     resultado, err := funcion()
     if err != nil {
         return nil, fmt.Errorf("fallo al ejecutar función: %w", err)
     }
     ```
2. **Prohibido el Abuso de `panic()`**:
   - Reservar `panic()` únicamente para fallos catastróficos e irrecuperables durante la inicialización (`main` / `init`). Para lógica de negocio siempre retornar `(T, error)`.

---

## 2. Prohibición en Concurrencia y Goroutines

1. **Cero Goroutines Huérfanas (Goroutine Leaks)**:
   - Toda goroutine DEBE tener un mecanismo de cancelación con `context.Context` o estar controlada por un `sync.WaitGroup`.
2. **Prevención Absoluta de Data Races**:
   - Toda variable mutable compartida entre goroutines DEBE protegerse con `sync.Mutex` / `sync.RWMutex` o comunicarse vía canales.
3. **Captura Correcta de Variables en Bucles**:
   - ❌ **Mal**: `for _, item := range lista { go func() { procesar(item) }() }`
   - ✅ **Bien**: `for _, item := range lista { go func(val Elemento) { procesar(val) }(item) }`

---

## 3. Abuso de Punteros e Interfaces

1. **Uso Eficiente de Punteros**:
   - No usar punteros en structs pequeñas o tipos primitivos a menos que sean mutables o representen campos opcionales/nulos. Evitar presión innecesaria sobre el Garbage Collector (GC).
2. **Interfaces Pequeñas y Justificadas**:
   - Prohibido crear interfaces gigantescas (> 3-4 métodos) o interfaces prematuras sin implementación real. *Acepta interfaces, retorna structs concretos*.

---

## 4. Diseño Limpio de Paquetes (Anti `utils`/`common`)

1. **Prohibido Paquetes Genéricos y Ambiguos**:
   - ❌ **Prohibido**: `package utils`, `package common`, `package helper`, `package base`.
   - ✅ **Bien**: Paquetes expresivos por dominio (`package autenticacion`, `package formateador_fecha`, `package validador_correo`).
2. **Cero Dependencias Circulares**:
   - Diseñar las capas de importación en una sola dirección (Dominio ➔ Caso de Uso ➔ Infraestructura).
