
# Series Tracker Backend

Objetivo del Proyecto
En este proyecto construimos una aplicación **full stack real**, manteniendo una estricta separación de responsabilidades:
* Backend (Este repositorio)Expone una API REST que devuelve datos en formato JSON. No sabe nada de HTML ni de cómo se presentan los datos.

---

## Estructura del Proyecto

El código está organizado de manera modular para mantener todo ordenado y fácil de encontrar:

*   `Backend/`
    *   `db/`: Configuración y conexión a la base de datos PostgreSQL.
    *   `docs/`: Archivos generados automáticamente por Swagger para la documentación de la API.
    *   `handlers/`: Controladores que reciben las peticiones HTTP, validan los datos y envían respuestas JSON.
    *   `model/`: Estructuras de datos (Structs) como `Serie`, `Response` y `ErrorResponse`.
    *   `repository/`: Lógica de interacción directa con la base de datos (Queries CRUD).
    *   `services/`: Lógica de negocio adicional y middleware (como la configuración de CORS y guardado de imágenes).
    *   `uploads/`: Carpeta donde se guardan las imágenes (portadas) de las series de forma local.
    *   `main.go`: Punto de entrada de la aplicación, donde se registran las rutas y arranca el servidor.


## Endpoints de la API REST

El API responde en la ruta principal y maneja los siguientes métodos HTTP con respuestas apropiadas y validaciones:

| Método | Endpoint | Descripción |
| :--- | :--- | :--- |
| `GET` | `/series` | Obtiene todas las series. Soporta paginación (`?page=`, `?limit=`), búsqueda (`?q=`), y ordenamiento (`?sort=`, `?order=asc` o `desc`) |
| `GET` | `/series/:id` | Obtiene los detalles de una serie específica por su ID. |
| `POST` | `/series` | Crea una nueva serie. Soporta la subida de una imagen mediante `multipart/form-data`. |
| `PUT` | `/series/:id` | Edita una serie existente (incluyendo actualización de su portada). |
| `DELETE` | `/series/:id` | Elimina una serie por su ID (Retorna código HTTP de éxito). |
| `POST` | `/series/:id/ratings` | Agrega un rating a una serie |
| `GET`  | `/series/:id/ratings` | Obtiene el promedio de rating |

---

## Manejo de Imágenes

El backend permite subir imágenes mediante `multipart/form-data`.

- Las imágenes se almacenan localmente en la carpeta `/uploads`
- Si no se envía imagen, se usa una imagen por defecto
- Las imágenes se sirven como archivos estáticos en `/uploads/:filename`

---

## Documentación con Swagger

El API está documentada usando **Swagger**! 
Esto permite tener un contrato claro para los endpoints.

**¿Cómo usarlo?**
1. Levanta el backend.
2. Ingresa a la ruta `http://localhost:3005/swagger/index.html` o la ruta que exponga el servidor web.


---

## Variables de Entorno (`.env`)

Para ejecutar este proyecto, necesitas configurar las variables de entorno. El archivo `.env.example` contiene un ejemplo.

```env
DB_USER=root
DB_PASSWORD=root_2021131
DB_NAME=SeriesTracker
DB_PORT=5432

DB_HOST=db
BACKEND_PORT=3005
```

---

## Cómo ejecutar el proyecto

1. Clonar el repositorio
2. Crear archivo `.env` basado en `.env.example`
3. Levantar el proyecto con docker-compose:

```bash
docker-compose up --build
```

---

## Reflexión
Go resultó ser un lenguaje eficiente para construir API, pero requiere más trabajo manual en el manejo de rutas comparado con frameworks de otros lenguajes como typescript con express que al menos para mi se hace más sencillo.

El uso de Docker facilitó la portabilidad del proyecto y la configuración del entorno. Lo que reduce en gran medida el tiempo y esfuerzo en procesos de despliegue.

Se exploraron conceptos como CORS, separación cliente-servidor, manejo de imágenes y documentación con Swagger.

Por mi parte, nunca me había involucrado en el despliegue de aplicaciones, pero este proyecto me permitió aprender mucho sobre el tema y obtener la capacidad de hacerlo por mi cuenta.

La experiencia fue positiva y estos conceptos serían utilizadas nuevamente en proyectos futuros.