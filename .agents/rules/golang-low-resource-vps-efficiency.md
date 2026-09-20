# Regla Estricta: Eficiencia en VPS de Bajos Recursos & Principios Clean Code (Low-Resource VPS Standard)

> 💡 **DISEÑO PARA VPS DE BAJOS RECURSOS**: El código debe funcionar de manera ultrarrápida, fluida y estable incluso en VPS pequeños con escasa memoria RAM (ejemplo: 512MB - 1GB RAM, 1 vCPU). **Está estrictamente prohibido cargar grandes volúmenes de datos en memoria RAM**.

---

## 1. Principios de Código Limpio (Clean Code & SOLID)

1. **S - Single Responsibility**: Cada paquete, struct y función tiene una sola razón para cambiar.
2. **O - Open/Closed**: Abierto a extensión mediante interfaces, cerrado a modificación.
3. **L - Liskov Substitution**: Sustitución transparente de implementaciones concretas.
4. **I - Interface Segregation**: Interfaces pequeñas y especializadas.
5. **D - Dependency Inversion**: Inversión de dependencias con Google Wire.
6. **KISS (Keep It Simple, Stupid) & YAGNI (You Aren't Gonna Need It)**: Evitar sobre-ingeniería innecesaria.

---

## 2. Optimización Estricta de Memoria en VPS Pequeños

1. **Cero Cargas Masivas en Memoria (No In-Memory Buffering)**:
   - Prohibido hacer `All(ctx)` en tablas con miles de registros. Se exige usar **Paginación basada en Cursor** (`Where(id > ultimoID).Limit(50)`) o procesar mediante lecturas por lotes (chunking/streaming).
2. **Control Estricto de Goroutines (Prevención de Memory Leaks)**:
   - Cada goroutine iniciada DEBE tener un ciclo de vida definido y cancelable vía `context.Context` con timeout.
   - Usar `sync.WaitGroup` o worker pools con tamaño limitado para evitar instanciar miles de goroutines simultáneas que consuman la RAM de la VPS.
3. **Optimización del Pool de Conexiones a la Base de Datos**:
   - Ajustar el pool de ENT/SQL según la capacidad de la VPS:
     ```go
     db.SetMaxOpenConns(15)                 // Máximo 15 conexiones abiertas
     db.SetMaxIdleConns(5)                  // Máximo 5 inactivas
     db.SetConnMaxLifetime(15 * time.Minute) // Reciclaje de conexiones
     ```
