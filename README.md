# Challenge Domain Info
Esta carpeta contiene la solución al reto de progrmación

## Requerimientos
- Go
- Node JS
- CockroanchDB

## Frontend
### Deploy
La aplicación fue desarrollada en Vue, para desplegar la aplicación ejecute los siguientes comandos:
```bash
npm ci 
npm run build
```
Esto generara una carpeta con el build de la aplicación

### Development
La aplicacion usa webpack dev server que es un servidor para desarrollo que permite ver en tiempo real los cambios en la aplicación.

```bash
npm ci
npm run serve
```
Por defecto el servidor corre en el puerto `8080`

## Backend

### Deploy
Se requiere agregar el archivo .env en la carpeta raiz `server/` con las siguientes configuraciones
- `PORT`: Puerto que escuchara la applicación
- `DB_URI`: URI de la base de datos cockroachDB

Dentro de la carpeta server ejecutar el siguiente comando
```bash
go run main.go
```

Por defecto el servidor corre en el puerto `5500`


## DataBase
Asegurese de tener cockroachDB ejecutando un nodo lo puede hacer con el comando 
```bash
cockroach start-single-node --insecure
```


