# Desafio de Cache en Go

En este apartado, se muestra un servicio de cache hecho en Go.
Es un servicio HTTP Rest el cual permite guardar, obtener y eliminar
registros por un mecanismo de llave-valor o mapa, los cuales persisten al reinicio de la aplicación, como una base de datos. 

El mapa se manipula directamente en la memoria usando mecanismos que lo hacen seguro para la concurrencia.
Como consecuencia, se obtienen tiempos de respuesta espectaculares.

Además, se agregó un mecanismo de persistencia que replica la cache en el sistema de archivos. 
Se trata de un worker que espera ser notificado y actualiza acordemente los archivos.
Estos archivos se pueden cargar en el momento de iniciar la aplicación.

## Requerimientos

* Guardar un registro por llave
* Obtener datos por llave
* Eliminar una una registro particular del mapa
* Guardar los valores de la base de datos en el sistema de archivos
* Cargar la base de datos en memoria desde el sistema de archivos

## Uso

Ejemplos dados en formato .http

### Obtener registro

Devuelve 400 si no se especifica la key, 404 si no existe el registro. Si todo esta correcto devuelve el registro en el cuerpo de la respuesta.

```http request
GET http://localhost:8080/{key}
```

### Eliminar registro

Devuelve 400 si no se especifica la key, 404 si no existe el registro. Si todo esta correcto devuelve código 202 (aceptado) en la respuesta.

```http request
GET http://localhost:8080/{key}
```

### Guardar registro

Devuelve 400 si no se especifica la key, 500 hubo un error leyendo el cuerpo de la petición.
Si todo esta correcto devuelve código 202 (aceptado) en la respuesta, significado que se efectuará con prontitud.

```http request
GET http://localhost:8080/{key}
```