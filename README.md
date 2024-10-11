# user-auth-api

API de autenticación y autorización de usuarios construida con Go y Fiber. Está diseñada siguiendo una arquitectura orientada a microservicios, con el objetivo de ser escalable, eficiente y fácil de integrar con otros servicios.

Los endpoints de autenticación y autorización implementan un sistema robusto utilizando **JWT** y **Refresh Tokens**. Cuando un usuario realiza una petición, el microservicio público (gateway) valida el token y redirige la solicitud al servicio correspondiente. En caso de que el **JWT** haya expirado, el servicio de autenticación permite renovar el token utilizando el **Refresh Token**, asegurando sesiones seguras y controladas. Además, el sistema garantiza que solo se permita una sesión activa por usuario, invalidando los tokens anteriores al iniciar sesión en un nuevo dispositivo.

- [Estructura del Proyecto](#estructura-del-proyecto)
- [Endpoints](#endpoints)
  - [Autenticación](#autenticación)
  - [Usuarios](#usuarios)
  - [Organizaciones](#organizaciones)
- [Estrategia de Autenticación y Autorización](#estrategia-de-autenticación-y-autorización)
  - [1. Creación de Tokens](#1-creación-de-tokens)
    - [Proceso de Login:](#proceso-de-login)
  - [2. Expiración y Renovación de Tokens](#2-expiración-y-renovación-de-tokens)
    - [Proceso de Renovación:](#proceso-de-renovación)
  - [3. Expiración de Sesión y Revocación de Tokens](#3-expiración-de-sesión-y-revocación-de-tokens)
  - [4. Control de Sesiones](#4-control-de-sesiones)
  - [5. Roles y Permisos](#5-roles-y-permisos)
- [Instalación y Configuración](#instalación-y-configuración)
- [Contribuciones](#contribuciones)


## Estructura del Proyecto

```bash
user-auth-api/
├── cmd/
│   ├── routes/           # Definición de rutas
│   ├── handlers/         # Controladores de la API
│   ├── middlewares/      # 
│   └── server.go         # Inicia el servidor y dependencias
├── config/
│   └── config.go         #  Cargar configuraciones (puerto, base de datos, etc.)
├── internal/
│   ├── models/           # Modelos de la base de datos (User, Organization)
│   ├── repositories/     # Conexión a la base de datos
│   └── services/         # Lógica de negocio (User, Organization)
├── utils/                # Funciones de utilidad (hashing, validación, etc.)
├── common/               # Definiciones comunes (errores, constantes, etc.)
├── main.go               # Punto de entrada de la aplicación
└── README.md             # Documentación del proyecto
```

## Endpoints

### Autenticación
1. **POST /auth/login**: Permite que los usuarios inicien sesión proporcionando credenciales válidas y reciban un par de tokens (JWT + Refresh Token).
2. **POST /auth/logout**: Invalida tanto el JWT como el Refresh Token del usuario, cerrando la sesión de forma segura.
3. **POST /auth/refresh**: Permite generar un nuevo JWT si el anterior ha expirado, usando un Refresh Token válido. Si el **Refresh Token** ha expirado, la sesión se cerrará. Si faltan menos de 24 horas para que el **Refresh Token** expire, se emitirá un nuevo Refresh Token junto con el JWT.

### Usuarios
1. **GET /users/:id**: Obtiene un usuario por ID.
2. **POST /users**: Crea un nuevo usuario.
2. **PUT /users/:id**: Actualiza los detalles de un usuario.
4. **DELETE /users/:id**: Elimina un usuario.

### Organizaciones
1. **GET /organizations**: Obtiene todas las organizaciones.
2. **GET /organizations/:orgId**: Obtiene una organización específica por su ID.
3. **GET /organizations/:orgId/users**: Obtiene todos los usuarios de una organización.
4. **POST /organizations**: Crea una nueva organización.
5. **PUT /organizations/:orgId**: Actualiza los detalles de una organización.

## Estrategia de Autenticación y Autorización

### 1. Creación de Tokens

Al iniciar sesión con el endpoint de **login**, se genera un **JWT** con un tiempo de vida limitado (15 minutos) y un **Refresh Token** que tiene una duración más larga (7 días). Esto permite mantener sesiones cortas para mayor seguridad, pero al mismo tiempo proporcionar una experiencia de usuario fluida con la renovación automática del **JWT** utilizando el **Refresh Token**.

#### Proceso de Login:

1. El usuario envía sus credenciales.
2. Si las credenciales son correctas, el servidor genera un **JWT** con información del usuario (como `user_id`, `roles` y `email`).
3. Se genera un **Refresh Token** que se almacena en la base de datos junto con el ID del usuario para gestionar su sesión.
4. Se devuelve el **JWT** y el **Refresh Token** al cliente.

### 2. Expiración y Renovación de Tokens

El **JWT** expira después de un periodo corto de tiempo (15 minutos). Cuando un **JWT** expira, el cliente puede utilizar el **Refresh Token** para obtener un nuevo **JWT** sin necesidad de volver a iniciar sesión.

#### Proceso de Renovación:

1. El cliente envía el **Refresh Token** al servidor.
2. El servidor verifica el **Refresh Token** y, si es válido y no ha sido revocado, genera un nuevo **JWT**.
3. El nuevo **JWT** se devuelve al cliente.

### 3. Expiración de Sesión y Revocación de Tokens

Para asegurar que un usuario solo tenga una sesión activa a la vez, cada vez que un usuario inicia sesión en un nuevo dispositivo:

- La sesión anterior se revoca y el **Refresh Token** asociado se elimina de la base de datos.
- Si el usuario intenta utilizar un **JWT** o **Refresh Token** de una sesión anterior, el servidor lo rechazará.

### 4. Control de Sesiones

El sistema de autenticación solo permite una sesión activa por usuario a la vez:

- Si un usuario intenta iniciar sesión en un nuevo dispositivo, su **Refresh Token** anterior se invalida y se le asigna un nuevo **Refresh Token** para esa nueva sesión.
- Si un usuario intenta utilizar un **JWT** o **Refresh Token** de una sesión anterior, ese token será rechazado porque la sesión anterior ya no es válida.

### 5. Roles y Permisos

Cada usuario tiene uno o más roles asignados. Los roles definen qué acciones un usuario puede realizar:

- **Roles**: Un usuario puede tener roles como `admin`, `editor`, `viewer`, etc.
- **Permisos**: Los permisos permiten limitar el acceso a ciertos recursos o endpoints de la API basándose en el rol del usuario.

La validación de roles y permisos ocurre en cada petición. Si el usuario no tiene el permiso adecuado para acceder a un recurso, se le devolverá un error de **403 Forbidden**.

## Instalación y Configuración

### Requisitos:

- **Go** 1.19+
- **Redis** (para gestionar sesiones y revocación de tokens)
- **PostgreSQL** 


### Clonar repositorio:

```bash
git clone https://github.com/juanMaAV92/user-auth-api.git
cd user-auth-api
```

### Configuración de Entorno

Configura las variables de entorno en un archivo .env:

```env
JWT_SECRET=your-secret-key
REDIS_ADDR=localhost:6379
POSTGRES_URL=postgres://user:password@localhost:5432/dbname
```


### Iniciar la Aplicación
Instalar dependencias y compilar el proyecto:

```bash
go mod tidy
go run main.go

# Ejecución de Pruebas
go test ./...
```

## Contribuciones
Las contribuciones son bienvenidas. Si deseas contribuir a este proyecto, por favor abre un issue o envía un pull request.

## Licencia
Este proyecto está licenciado bajo la MIT License.