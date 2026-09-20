# Regla PRIMORDIAL: Aislamiento Absoluto Multi-Tenant en TODOS los Endpoints

> 🚨 **REGLA PRIMORDIAL Y MANDATORIA #1**: TODO endpoint HTTP o servicio que consulte, busque, cree, actualice o elimine datos DEBE VERIFICAR Y RESPETAR EL MULTI-TENANT (`empresa_id`). Está estrictamente prohibido retornar datos de un cliente a otro bajo cualquier circunstancia.

---

## 1. Verificación Primordial en Endpoints HTTP

Cada vez que se cree o modifique un Handler o Controller HTTP en Fiber v3:

1. **Middleware Multi-Tenant Obligatorio**:
   - Todo endpoint de datos debe estar protegido por el middleware de autenticación que extrae e inyecta `empresa_id` en el contexto.
2. **Prevención de Vulnerabilidades IDOR / Fuga entre Inquilinos**:
   - Cuando un endpoint pida un registro por ID (`/api/v1/pedidos/:id`), **NUNCA** se debe buscar solo por `id`.
   - ❌ **INSEGURO (Vulnerable a fuga de datos)**:
     ```go
     // Inseguro: Un usuario del Tenant A podría ver el pedido del Tenant B si adivina el ID
     pedido, err := client.Pedido.Get(ctx, id)
     ```
   - ✅ **SEGURO Y MANDATORIO (Aislado por Tenant)**:
     ```go
     // Seguro: Filtra estrictamente por ID Y empresa_id del usuario autenticado
     empresaID, _ := ObtenerEmpresaID(ctx)
     pedido, err := client.Pedido.Query().
         Where(
             pedido.IDEQ(id),
             pedido.EmpresaIDEQ(empresaID), // Verificación de pertenencia al Tenant
         ).
         Only(ctx)
     ```

---

## 2. Esquema de Base de Datos Multi-Tenant (`empresa_id`)

1. **Campo `empresa_id` Mandatorio**:
   - Toda tabla o entidad de ENT que pertenezca a un inquilino DEBE incluir el campo `empresa_id` como `field.Int("empresa_id")` o `field.UUID("empresa_id", uuid.UUID{})`.
2. **Índices Compuestos por Tenant**:
   - Las claves primarias y consultas de tablas multi-tenant deben incluir `empresa_id` en el primer término de sus índices compuestos para máxima velocidad y aislamiento a nivel de índice B-Tree:
   ```go
   func (Venta) Indexes() []ent.Index {
       return []ent.Index{
           index.Fields("empresa_id", "id"),            // Búsqueda por ID aislada por empresa
           index.Fields("empresa_id", "fecha_creacion"), // Reportes por empresa
           index.Fields("empresa_id", "estado"),         // Filtros de estado por empresa
       }
   }
   ```

---

## 3. Inyección y Propagación del Tenant mediante `context.Context`

- El middleware de autenticación de Fiber v3 DEBE extraer el `empresa_id` del token JWT verificado e inyectarlo en el `context.Context` de la solicitud:
  ```go
  type tenantKey struct{}

  func ConEmpresaID(ctx context.Context, empresaID int) context.Context {
      return context.Context(context.WithValue(ctx, tenantKey{}, empresaID))
  }

  func ObtenerEmpresaID(ctx context.Context) (int, error) {
      id, ok := ctx.Value(tenantKey{}).(int)
      if !ok || id == 0 {
          return 0, errors.New("identificador de empresa no presente en el contexto")
      }
      return id, nil
  }
  ```

---

## 4. Interceptores de ENT para Filtro Automático por Tenant

- Para prevenir errores humanos (olvidar añadir `.Where(factura.EmpresaIDEQ(empresaID))`), se deben configurar **Interceptors / Privacy Rules en ENT** que apliquen automáticamente el filtro de `empresa_id` a TODAS las consultas de lectura, actualización y eliminación:
  ```go
  client.Intercept(
      ent.InterceptFunc(func(next ent.Querier) ent.Querier {
          return ent.QuerierFunc(func(ctx context.Context, q ent.Query) (ent.Value, error) {
              empresaID, err := ObtenerEmpresaID(ctx)
              if err != nil {
                  return nil, fmt.Errorf("seguridad multi-tenant: %w", err)
              }
              // Inyectar automáticamente el filtro por empresa_id
              return next.Query(ctx, q)
          })
      }),
  )
  ```

---

## 5. Reglas de Validación y Seguridad Multi-Tenant

- 🚫 **Prohibido realizar consultas globales sin `empresa_id`** excepto para operaciones de super-administrador global explícitamente autorizadas.
- 🚫 **Prohibido recibir `empresa_id` desde el body de la petición HTTP**: El `empresa_id` SIEMPRE se obtiene del token verificado en la cookie HTTP-Only, NUNCA del payload enviado por el cliente frontend.
