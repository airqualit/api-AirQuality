# Air Quality API

La **Air Quality API** es un servicio RESTful desarrollado en Go que permite gestionar y consultar datos sobre la calidad del aire. La API ofrece dos endpoints principales para insertar nuevos datos y obtener datos específicos mediante un ID. Los datos se almacenan en una base de datos MongoDB y se pueden visualizar mediante Grafana.

## URL Base
`https://api-airquality.onrender.com`

## Endpoints

### 1. Obtener datos por ID (GET)
- **URL:** `/iotdevice/{id}`
- **Método:** GET
- **Descripción:** Obtiene los datos de calidad del aire asociados a un ID específico.
- **Parámetros de URL:**
  - `id` (string): El ID del registro de calidad del aire.
- **Respuesta Exitosa:**
  - **Código:** 200 OK
  - **Contenido:** JSON con los datos de calidad del aire.
    ```json
    {
      "id": "a97faf7b-fdc9-4da5-a6ac-0375ae4e4b04",
      "data": {
        "temperature": 56,
        "relativehumidity": 150,
        "barometricpressure": 134,
        "rainflow": 155,
        "PM2.5": 5.5,
        "PM10": 47,
        "CO": 66,
        "C2O": 90
      }
    }
    ```
- **Errores Comunes:**
  - **Código:** 404 Not Found
  - **Descripción:** El ID proporcionado no existe.

### 2. Insertar datos de calidad del aire (POST)
- **URL:** `/iotdevice`
- **Método:** POST
- **Descripción:** Inserta nuevos datos de calidad del aire en la base de datos.
- **Cuerpo de la Solicitud (JSON):**
  ```json
  {
    "data": {
        "Temperature": 56,
        "RelativeHumidity": 150,
        "BarometricPressure": 134,
        "RainFlow": 155,
        "PM2.5": 5.5,
        "PM10": 47,
        "CO": 66,
        "C2O": 90
    }
  }
