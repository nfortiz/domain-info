# Challenge truora
Esta carpeta contiene la solución al reto de truora

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
Dentro de la carpeta truora-server ejecutar el siguiente comando
```bash
go run main.go
```

POor defecto el servidor corre en el puerto `5500`


## DataBase
Asegurese de tener cockroacDB ejecutando un nodo lo puede hacer con el comando 
```bash
cockroach start-single-node --insecure
```


