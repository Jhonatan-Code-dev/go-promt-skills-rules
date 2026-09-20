# Regla Estricta: Rendimiento Extremo en Consultas, Indexación y Escalabilidad BD (High Performance & DB Scalability)

> **PRINCIPIO DE ALTO RENDIMIENTO**: Toda consulta a base de datos y diseño de esquema debe realizarse desde el rol de **Arquitecto de Bases de Datos Experto en Escalabilidad Masiva**. Las consultas deben ejecutarse en milisegundos incluso bajo millones de registros y alta concurrencia.

---

## 1. Indexación Estratégica y Claves Foráneas (Foreign Keys)

1. **Búsqueda Ultrarrápida por ID / Clave Primaria**:
   - Toda búsqueda individual DEBE realizarse utilizando su ID o clave primaria indexada para garantizar velocidad $O(1)$ o búsqueda logarítmica instantánea en índices B-Tree.
2. **Índices Obligatorios en Claves Foráneas (`Foreign Keys`)**:
   - Todo campo que sirva como clave foránea o relación con otra tabla (`usuario_id`, `empresa_id`, `orden_id`) DEBE tener un **Índice B-Tree explícito** en el esquema de ENT.
   ```go
   func (Factura) Indexes() []ent.Index {
       return []ent.Index{
           index.Fields("usuario_id"),                            // Búsqueda rápida por usuario
           index.Fields("empresa_id", "estado"),                  // Índice compuesto para consultas combinadas
           index.Fields("fecha_creacion", "id"),                  // Para paginación por fecha/cursor
       }
   }
   ```
3. **Índices Compuestos para Filtros Frecuentes**:
   - Si una consulta filtra por `empresa_id` Y `estado` Y `fecha`, se debe definir un índice compuesto cubriente en ese orden exacto de selectividad.

---

## 2. Prevención Absoluta del Problema N+1 (Eager Loading con ENT)

- 🚫 **Prohibido realizar consultas dentro de bucles `for`**:
  - ❌ **Mal (N+1)**:
    ```go
    usuarios, _ := client.Usuario.Query().All(ctx)
    for _, u := range usuarios {
        pedidos, _ := u.QueryPedidos().All(ctx) // Ejecuta N consultas adicionales
    }
    ```
  - ✅ **Bien (Eager Loading)**:
    ```go
    usuarios, _ := client.Usuario.Query().
        WithPedidos(). // Trae la relación en 1 o 2 consultas optimizadas
        All(ctx)
    ```

---

## 3. Paginación por Cursor (Keyset Pagination) vs OFFSET

- 🚫 **Evitar `OFFSET` en tablas masivas**: `OFFSET 100000` fuerza a la base de datos a leer y descartar 100,000 filas antes de retornar el resultado.
- ✅ **Usar Paginación basada en Cursor / ID**:
  ```go
  // Paginación rápida por ID (Keyset)
  client.Factura.Query().
      Where(factura.IDGT(ultimoIDVisto)).
      Order(ent.Asc(factura.FieldID)).
      Limit(limitePagina).
      All(ctx)
  ```

---

## 4. Selección Exclusiva de Campos (Projection & Select)

- Evitar traer columnas pesadas (`TEXT`, `BLOB`, payloads JSON grandes) si solo se requieren campos específicos.
- Usar `Select(...)` en ENT para traer únicamente las columnas necesarias en memoria.

---

## 5. Estrategia para Alta Demanda (Concurrencia Masiva)

1. **Configuración de Connection Pool**:
   - Ajustar `SetMaxOpenConns`, `SetMaxIdleConns` y `SetConnMaxLifetime` según el número de núcleos y réplicas.
2. **Caché de Consultas de Lectura con Redis**:
   - Para datos de alta lectura y baja modificación (catalogos, configuraciones, usuarios activos), consultar primero la memoria en Redis antes de tocar la base de datos.
