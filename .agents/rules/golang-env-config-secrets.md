# Regla Estricta: Gestión de Variables de Entorno y Cero Secretos Hardcodeados (Zero Hardcoded Secrets & Env Config)

> 🔐 **PROHIBICIÓN ABSOLUTA DE SECRETOS Y CONFIGURACIONES HARDCODEADAS**: Queda estrictamente prohibido escribir cadenas de conexión a bases de datos, claves secretas de JWT, contraseñas, tokens de API, puertos o URLs de servicios directamente en el código fuente Go. **Todo debe cargarse obligatoriamente mediante Variables de Entorno (`.env`)**.

---

## 1. Reglas de Configuración por Variables de Entorno

1. **Cero Hardcoding**:
   - ❌ **Prohibido (Hardcoded)**:
     ```go
     // Inseguro y rígido
     cadenaConexion := "postgres://usuario:secreto123@localhost:5432/mi_bd"
     claveSecretaJWT := "mi_clave_secreta_super_secreta"
     puertoServidor := ":8080"
     ```
   - ✅ **Obligatorio (Variables de Entorno con `.env`)**:
     ```go
     cadenaConexion := os.Getenv("DATABASE_URL")
     claveSecretaJWT := os.Getenv("JWT_SECRET")
     puertoServidor := os.Getenv("PUERTO_SERVIDOR")
     ```

2. **Struct de Configuración Tipada con Validación Fuerte**:
   - Toda aplicación debe tener un paquete de configuración (ejemplo `internal/plataforma/configuracion`) que cargue el archivo `.env` mediante `godotenv` y valide que todas las variables requeridas existan:
   ```go
   package configuracion

   import (
       "errors"
       "fmt"
       "os"
       "github.com/joho/godotenv"
   )

   type Configuracion struct {
       CadenaConexionBD  string
       ClaveSecretaJWT   string
       PuertoServidor    string
       EntornoAplicacion string
   }

   func CargarConfiguracion() (*Configuracion, error) {
       // Cargar archivo .env si existe en desarrollo
       _ = godotenv.Load(".env")

       config := &Configuracion{
           CadenaConexionBD:  os.Getenv("DATABASE_URL"),
           ClaveSecretaJWT:   os.Getenv("JWT_SECRET"),
           PuertoServidor:    os.Getenv("PUERTO_SERVIDOR"),
           EntornoAplicacion: os.Getenv("ENTORNO_APLICACION"),
       }

       // Validar variables críticas obligatorias
       if config.CadenaConexionBD == "" {
           return nil, errors.New("la variable de entorno DATABASE_URL es obligatoria")
       }
       if config.ClaveSecretaJWT == "" {
           return nil, errors.New("la variable de entorno JWT_SECRET es obligatoria")
       }
       if config.PuertoServidor == "" {
           config.PuertoServidor = "8080" // Valor por defecto seguro
       }

       return config, nil
   }
   ```

---

## 2. Archivos Obligatorios `.env.example` y `.gitignore`

1. **`.env.example` Obligatorio**:
   - Todo repositorio debe mantener un archivo `.env.example` actualizado que sirva de plantilla con nombres de variables descriptivas pero **SIN valores reales ni contraseñas**.
2. **Inclusión Obligatoria en `.gitignore`**:
   - El archivo `.env` con credenciales reales DEBE estar incluido en `.gitignore` para evitar filtraciones accidentales de secretos en repositorios Git.
