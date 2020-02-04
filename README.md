# novedades_mid
novedades_mid, MID para el negocio de novedades, el proyecto está programado en el lenguaje Go y creado con el [framework beego](https://beego.me/).

El Api tiene la funcionalidad de cominicar el cliente de novedades con el CRUD, el mid provee un manejo de datos para proveerle al cliente los datos en la estructura que este requiere ya que antes se usaba un Api para una base de datos en Mongo.

adicionalmente el api proveecomunicacion con un api de administrativa_amazon_api con el fin de que cierta novedades especificas aparezcan en ela aplicativo de condor.

***Instlaciones Previas:***
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)


## Especificaciones Técnicas

### Instalación

#### Opción 1
```shell
go get github.com/udistrital/novedades_mid
```
#### Opción 2
```shell
# Para clonar el proyecto en la carpeta local go/src/github.com/udistrital
cd go/src/github.com/udistrital

# clonar repo
git clone https://github.com/udistrital/novedades_mid.git

# Ir a la carpeta del proyecto:
cd novedades_mid

# Instalar dependencias del proyecto
go get
```

### Variables de Entorno
```shell
NOVEDADES_API_HTTP_PORT=[puerto en el que quiere ejcutar el api]
NOVEDADES_CRUD_SERVICE=[dereccion donde se encuentra el api crud de novedades inluyendo el puerto]
ADMINISTRATIVA_AMAZON_SERVICE=[dereccion donde se encuentra el api de administrativa_amazon_service inluyendo el puerto]
JBPM_SERVICE=[direccion del servicio de jbpm]
```

### Ejecución del proyecto

Ubicado en la raíz del proyecto, ejecutar:
```shell
NOVEDADES_API_HTTP_PORT=XXX NOVEDADES_CRUD_SERVICE=XXX ADMINISTRATIVA_AMAZON_SERVICE=XXX JBPM_SERVICE=XXX bee run
```
O si se quiere ejecutar el swager:
```shell
NOVEDADES_API_HTTP_PORT=XXX NOVEDADES_CRUD_SERVICE=XXX ADMINISTRATIVA_AMAZON_SERVICE=XXX JBPM_SERVICE=XXX bee run -downdoc=true -gendoc=true
```

### Puertos

El servidor se expone por defecto en el puerto: localhost:8080

Para ver la documentación de swagger: [localhost:8080/swagger/](http://localhost:8080/swagger/)

**Nota**: *En el swagger sale un error, hacer caso omiso.*

### EndPoints

Al ejecutar el swagger se puede tener mayor apreciacion de los diferentes metodos de peticion por cada endpoint cuales son los distinpos endpoint disponibles y como usarlos.


## Licencia

This file is part of novedades_mid

novedades_mid is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

Foobar is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with Foobar. If not, see https://www.gnu.org/licenses/.
