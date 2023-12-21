# Challenge 11

Esta aplicación web está construida utilizando Go, Gorm y Gin para gestionar datos relacionados con Usuarios, Proyectos y Tareas. El objetivo principal es demostrar competencia en la configuración de rutas, la definición de múltiples modelos de datos y el establecimiento de relaciones entre ellos.

## Requisitos Previos

- Go instalado en tu máquina
- PostgreSQL u otro sistema de base de datos configurado

## Instalación

1. Clona el repositorio: `git clone https://github.com/tu_usuario/tu_proyecto.git`
2. Navega al directorio del proyecto: `cd tu_proyecto`
3. Instala las dependencias: `go mod tidy`

## Configuración

1. Configura la conexión a la base de datos en el archivo `.env`.
2. Asegúrate de tener una base de datos PostgreSQL disponible.

## Ejecución

1. Inicia la aplicación: `go run main.go`
2. La aplicación se ejecutará en `http://localhost:5000` (o el puerto configurado).

## Rutas y Ejemplos de Uso

### Crear un usuario
curl -X POST -H "Content-Type: application/json" -d '{"name":"NombreUsuario", "Password":" Password "}' http://localhost:5000/users

### Obtener todos los usuarios:
curl http://localhost:5000/users

### Obtener un usuario específico
curl http://localhost:5000/users/1

### Actualizar un usuario:
curl -X PUT -H "Content-Type: application/json" -d '{"name":"NombreUsuario", "Password":" Password "}' http://localhost:5000/users/1

### Eliminar un usuario:
curl -X DELETE http://localhost:5000/users/1

### Crear un proyecto:
curl -X POST -H "Content-Type: application/json" -d '{"Title":"NombreProyecto"}' http://localhost:5000/users/1/project

### Obtener todos los proyectos de un usuario:
curl http://localhost:5000/users/1/project

### Crear una tarea:
curl -X POST -H "Content-Type: application/json" -d '{"Title":"NombreTarea", "Description":, "DescripciónTarea"}' http://localhost:5000/users/1/project/2/tasks

### Obtener todas las tareas de un proyecto:
curl http://localhost:5000/users/1/project/2/tasks

# API-RESTFULL-Prueba-Vue-Go
