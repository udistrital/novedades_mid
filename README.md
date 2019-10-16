# novedades_mid
novedades_mid, MID para el negocio de novedades, el proyecto está programado en el lenguaje Go y creado con el [framework beego](https://beego.me/).

***Instlaciones Previas:***
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)

## Configuración del Proyecto.

### Opción 1
Ejecutar desde la terminal 'go get repositorio':
```shell 
go get github.com/udistrital/novedades_mid
```
### Opción 2
Para instalar el proyecto realizar los siguientes pasos:
- Para clonar el proyecto en la carpeta local go/src/github.com/udistrital ir a la consola y ejecutar:
```shell 
    cd go/src/github.com/udistrital
```
- Ejecutar:
```shell 
    git clone https://github.com/udistrital/novedades_mid.git
```

- Ir a la carpeta del proyecto:
```shell 
    cd novedades_mid
```

- Instalar dependencias del proyecto:
```shell 
    go get
```

## Variables de entorno 

* El puerto por el que se expone la api **httpport = 8502**; si se cambia de puerto se debe editar la configuración en el [cliente](https://github.com/udistrital/novedades_cliente), especificamente la varible de entorno ARGO_NOSQL_SERVICE.
* La variable de entorno corresponde al puerto en donde se desplegará el API y corresponde a la siguiente :
```shell 
    NOVEDADES_API_HTTP_PORT=8502
```

## Ejecución del proyecto

* Ubicado en la raíz del proyecto, ejecutar:
```shell 
    NOVEDADES_API_HTTP_PORT=8502 bee run
```
* O si se quiere ejecutar el swager:
```shell 
    NOVEDADES_API_HTTP_PORT=8502 bee run -downdoc=true -gendoc=true
```

### Puertos

* El servidor se expone en el puerto: localhost:8502

* Para ver la documentación de swagger: [localhost:8502/swagger/](http://localhost:8502/swagger/)
    *Nota*: En el swagger sale un error, hacer caso omiso.

### EndPoints

Cada controlador tiene los metodos :
* Post
 los endpoint a los cuales apuntar son los siguientes:


||End Point|
|----------------|------------------------|
| **registroNovedad** | `[host de la maquina]:[puerto]/v1/registro_novedad` |



