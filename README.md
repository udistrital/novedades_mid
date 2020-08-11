# novedades_mid

El Api tiene la funcionalidad de cominicar el cliente de novedades con el CRUD, el mid provee un manejo de datos para proveerle al cliente los datos en la estructura que este requiere ya que antes se usaba un Api para una base de datos en Mongo.

adicionalmente el api proveecomunicacion con un api de administrativa_amazon_api con el fin de que cierta novedades especificas aparezcan en ela aplicativo de condor.

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno
```shell
NOVEDADES_API_HTTP_PORT=[puerto en el que quiere ejcutar el api]
NOVEDADES_CRUD_SERVICE=[dereccion donde se encuentra el api crud de novedades inluyendo el puerto]
ADMINISTRATIVA_AMAZON_SERVICE=[dereccion donde se encuentra el api de administrativa_amazon_service inluyendo el puerto]
JBPM_SERVICE=[direccion del servicio de jbpm]
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/novedades_mid

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/novedades_mid

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
NOVEDADES_API_HTTP_PORT=8080 NOVEDADES_CRUD_SERVICE=127.0.0.1:27017 ADMINISTRATIVA_AMAZON_SERVICE=some_value bee run
```

### Ejecución Dockerfile
```shell
# docker build --tag=novedades_mid . --no-cache
# docker run -p 80:80 novedades_mid
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/novedades_mid

#2. Moverse a la carpeta del repositorio
cd novedades_mid

#3. Crear un fichero con el nombre **custom.env**
# En windows ejecutar:* ` ni custom.env`
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas

Pruebas unitarias
```shell
# Not Data
```
## Estado CI

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/novedades_mid/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/novedades_mid) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/novedades_mid/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/novedades_mid) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/novedades_mid/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/novedades_mid) |

## Licencia

This file is part of novedades_mid

novedades_mid is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

novedades_mid is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with novedades_mid. If not, see https://www.gnu.org/licenses/.
