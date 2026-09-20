# Regla Estricta: Patrones de Diseño Idiomáticos y Prevención de Anti-Patrones en Go (Go Design Patterns Standard)

> 🎨 **PATRONES Y ANTI-PATRONES EN GO**: El diseño de la aplicación debe aplicar estrictamente los **Patrones de Diseño Idiomáticos de Go** (Fábrica, Singleton a prueba de hilos con `sync.Once`, Decorador por Composición) y evitar cualquier anti-patrón de arquitectura o concurrencia.

---

## 1. Patrones Go Esenciales (Best Practices)

### 1. Patrón Fábrica (Factory Pattern)
- Usar funciones constructoras que reciban parámetros de configuración y retornen la **interfaz del dominio** (o puntero al struct concreto en capas de infraestructura).
```go
type Forma interface {
    Dibujar() string
}

type Circulo struct{}
func (c *Circulo) Dibujar() string { return "Círculo" }

type Cuadrado struct{}
func (c *Cuadrado) Dibujar() string { return "Cuadrado" }

// FabricaForma es la función constructora que retorna la interfaz Forma
func FabricaForma(tipoForma string) (Forma, error) {
    switch tipoForma {
    case "circulo":
        return &Circulo{}, nil
    case "cuadrado":
        return &Cuadrado{}, nil
    default:
        return nil, fmt.Errorf("tipo de forma no soportado: %s", tipoForma)
    }
}
```

### 2. Patrón Singleton a Prueba de Hilos (`sync.Once`)
- Garantizar instancias únicas globales (pools de conexión, configuraciones de la app, loggers) utilizando `sync.Once`:
```go
type InstanciaUnicaConfiguracion struct {
    CadenaConexionBD string
}

var (
    instanciaGlobal *InstanciaUnicaConfiguracion
    ejecutarUnaVez   sync.Once
)

func ObtenerInstanciaConfiguracion() *InstanciaUnicaConfiguracion {
    ejecutarUnaVez.Do(func() {
        instanciaGlobal = &InstanciaUnicaConfiguracion{
            CadenaConexionBD: "postgres://...",
        }
    })
    return instanciaGlobal
}
```

### 3. Patrón Decorador por Composición de Interfaces
- Extender funcionalidad de componentes (middleware HTTP, caching, logging, auditoría) mediante composición de interfaces sin modificar el componente original:
```go
type ComponenteServicio interface {
    Ejecutar(ctx context.Context) string
}

type ServicioBase struct{}
func (s *ServicioBase) Ejecutar(ctx context.Context) string {
    return "Ejecución Base"
}

// DecoradorLogging envuelve cualquier ComponenteServicio para agregar logs
type DecoradorLogging struct {
    servicioBase ComponenteServicio
}

func NuevoDecoradorLogging(servicio ComponenteServicio) *DecoradorLogging {
    return &DecoradorLogging{servicioBase: servicio}
}

func (d *DecoradorLogging) Ejecutar(ctx context.Context) string {
    fmt.Println("LOG: Iniciando ejecución de servicio...")
    resultado := d.servicioBase.Ejecutar(ctx)
    fmt.Println("LOG: Ejecución finalizada.")
    return resultado
}
```

---

## 2. Anti-Patrones Adicionales a Evitar

1. **Prohibido Devolver `nil` en lugar de un `error`**:
   - ❌ **Anti-Patrón**:
     ```go
     func ObtenerRecurso() *Recurso {
         if noExiste { return nil } // Quién llama no sabe si falló o si no existe
         return &Recurso{}
     }
     ```
   - ✅ **Forma Correcta**:
     ```go
     func ObtenerRecurso() (*Recurso, error) {
         if noExiste {
             return nil, errors.New("el recurso solicitado no existe")
         }
         return &Recurso{}, nil
     }
     ```

2. **Prohibido Reinventar la Rueda**:
   - Siempre aprovechar la rica librería estándar de Go (`net/http`, `sync`, `context`, `time`, `crypto`, `os`, `encoding/json`) y paquetes consolidados antes de escribir utilidades caseras propensas a errores.

3. **Uso Obligatorio de Primitivas `sync` en Concurrencia**:
   - Toda sección crítica de memoria compartida DEBE usar `sync.Mutex`, `sync.RWMutex`, `sync.Once` o canales para evitar condiciones de carrera (Data Races).
