## Guia Petstore Loneliness 2000 API
Las siguientes son instrucciones para desarrollar el taller Petstore Loneliness 2000 API.

### ¿Que vamos a hacer?
En este taller vamos a aplicar los conocimientos adquiridos del paquete estandar de Go `net/http`. La solución del mismo solo debe estar limitada al uso de este paquete, paquetes de terceros como `gorilla/mux` o `gobuffalo` no serán necesarios. La lista de requerimientos sobre la cual trabajaremos en este taller están definidas en la sección [**Petstore Loneliness 2000 API**](https://github.com/wawandco/Web-Dev-Community/blob/main/petstorehttp/README.md#petstore-loneliness-2000-api) del archivo [**README.md**](https://github.com/wawandco/Web-Dev-Community/blob/main/petstorehttp/README.md#readme).

### Clonar repositorio
* Si es la primera vez que descargas o clonas el repositorio donde se encuentra el presente taller, ejecuta el siguiente comando en la ubicación deseada:
```console
git clone git@github.com:wawandco/Web-Dev-Community.git
```

* Si ya tienes el repositorio descargado en tu computador personal, ubícate dentro del directorio y  ejecuta el siguiente comando para actualizar el contenido del mismo:
```console
git pull origin main
```

El sub-directorio donde se encuentra ubicado el presente taller será `Web-Dev-Community/petstorehttp`.


### Preparación de la base de datos.
Antes de iniciar la ejecución del taller debemos asegurarnos que la base de datos esté correctamente configurada e inicializada. Para instrucciones de como crear la base de datos y correr las migraciones recomendamos tener instalado **soda**. El archivo [**CHEATSHEET.md**](https://github.com/wawandco/Web-Dev-Community/blob/main/petstore/CHEATSHEET.md) del taller anterior provee instrucciones adicionales para la instalación de soda y su uso para crear la base de datos.

Una vez creada la base de datos, se deberán correr las migraciones ubicadas en el directorio `migrations` de este repositorio. Las migraciones agregarán las tablas necesarias y junto a ellas algunos registros que serán usados a lo largo del taller.  
```console
soda migrate -e development
```

### Acceder a la base de datos.
El contenido del taller trae consigo el paquete `petstore/manage` que contiene funciones y estructuras que permiten leer, modificar y crear registros en las tablas de la base de datos. No hay necesidad de escribir código para leer, modificar o crear registros en las tablas, asi como no hay necesidad de escribir código para establecer la conexión a la base de datos. 

A continuación se describen los 3 paquetes de soporte que facilitan operaciones y acceso a datos en la base de datos.

#### manage/pet
Este paquete permite la manipulación, lectura y creación de registros en la tabla `pets`.

```go
import petstore/manage/pet
```

```go
// list all pets
p, err := pet.DB(nil).List()
```

```go
// find a pet with id 1
p, err := pet.DB(nil).Find(pet.WithID(1))
```

```go
// destroy a pet
var p *pet.Pet
...
err := pet.DB(nil).Destroy(p)
```

#### manage/client
Este paquete permite la manipulación, lectura y creación de registros en la tabla `clients`.

```go
import petstore/manage/client
```

```go
// list all clients
c, err := client.DB(nil).List()
```

```go
// find a client with id 1
c, err := client.DB(nil).Find(client.WithID(1))
```

```go
// destroy a client
var c *client.Client
...
err := client.DB(nil).Destroy(c)
```

#### manage/sale
Este paquete permite la manipulación, lectura y creación de registros en la tabla `sales`.

```go
import petstore/manage/sale
```

```go
// create a sale
newSale, err := sale.DB(nil).Create(clientID, petID)
```

```go
// find all pets bought by a client with id 1
sale, err := sale.DB(nil).ForClient(1)
```

### [CHEATSHEET.md](https://github.com/wawandco/Web-Dev-Community/blob/main/petstorehttp/CHEATSHEETHTTP.md)
En este archivo se encuentran definiciones y usos comunes del paquete `net/http` que servirán de guía en medio de la ejecución del taller.

### Run
El siguiente comando inicia la ejecución de la aplicación. Hace uso del paquete estandar `net/http` para iniciar el servidor que estará escuchando por peticiones HTTP en el puerto 8080.
```console
go run cmd/petstore/main.go
```

