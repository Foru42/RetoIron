# Reto: RESTful API en Golang

## Descripción del Proyecto
Este proyecto es una RESTful API desarrollada en **Golang** utilizando el framework **Gin** para la gestión de usuarios. La API permite realizar las operaciones básicas de CRUD (Create, Read, Update, Delete) sobre un modelo de datos `Usuario`, con campos como `name`, `surname`, y `email`. Los datos se almacenan en una base de datos MongoDb.

### Características principales
- **Operaciones CRUD completas**:
  - GET: Obtener una lista de usuarios.
  - POST: Crear un nuevo usuario con validación de datos.
  - PUT: Actualizar un usuario existente.
  - DELETE: Eliminar un usuario mediante su ID.
- **Rate Limiting**: Limita las solicitudes a 10 por segundo para proteger contra ataques DoS.
- **Validaciones**: Validación estricta de los datos de entrada (formato de correo electrónico, campos obligatorios, etc...).
- **Persistencia**: Uso de MongoDb para almacenamiento ligero y seguro.
- **Contenerización**: Configuración de Docker y Docker Compose para despliegue sencillo.

---

## Requisitos Previos
1. **Software necesario**:
   - Docker y Docker Compose.
   - Visual Studio Code o cualquier editor compatible con Golang.
   - Golang 1.21 o superior.

2. **Instalación de dependencias**:
   - Todas las dependencias de Golang se manejan mediante `go mod`.

---

## Instrucciones para Ejecutar el Proyecto

### 1. Clonar el repositorio
```bash
git clone <URL_DEL_REPOSITORIO>
cd <NOMBRE_DEL_REPOSITORIO>
```

### 2. Construir y ejecutar con Docker Compose
```bash
docker-compose up --build
```
Esto levantará el servicio en el puerto **8080**.

### 3. Uso de la API
Puedes interactuar con la API utilizando herramientas como `curl`, Postman, o cualquier cliente HTTP. Aquí algunos ejemplos:

#### **GET /usuarios**
```bash
curl -X GET http://localhost:8080/usuarios
```
#### **GETbyID /usuarios**
```bash
curl -X GET http://localhost:8080/usuarios/id
```

#### **POST /usuarios**
```bash
curl -X POST http://localhost:8080/usuarios -H "Content-Type: application/json" -d '{"name":"Mokel", "surname":"Foruria", "email":"lol@gmail.com"}'
```

#### **PUT /usuarios**
```bash
curl -X PUT http://localhost:8080/usuarios/id -H "Content-Type: application/json" -d '{"name":"josul", "surname":"Smith", "email":"john.smith@example.com"}'
```

#### **DELETE /usuarios?id=1**
```bash
curl -X DELETE http://localhost:8080/usuarios/id

```

---

## Estructura del Proyecto
```
RetoIronChip/
├── src/
│   ├── main.go            # Configuración principal del servidor
│   ├── routes/            # Definición de rutas
│   ├── controllers/       # Controladores para cada operación CRUD
│   ├── models/            # Modelos de datos
│   ├── database/          # Conexión y configuración de la base de 
|   ├── utils/             # Estructura para los errores
|   ├── validators/        # Validaciones 
├── Dockerfile             # Configuración de Docker
├── docker-compose.yml     # Configuración de Docker Compose
├── .gitignore             # Configuración del ignore
├── .env                   # Configuración de las variables de entorno
```

---

## Medidas de Seguridad Implementadas
1. **Rate Limiting**:
   - Previene ataques de denegación de servicio (DoS) limitando las solicitudes a 10 por segundo.

2. **Validación de Datos**:
   - Verificación de formato de correo electrónico.
   - Validación de campos obligatorios y longitud máxima de texto.

3. **Errores Genéricos**:
   - Mensajes de error diseñados para evitar revelar detalles internos.

---

## Autor
**[Josu Foruria]**
