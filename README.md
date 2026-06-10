# 📊 API HPC

Este proyecto ofrece una API desarrollada en **Go** utilizando **gorilla mux**.  

---

## Requisitos previos

1. Descargar la última versión del lenguage [Go](https://go.dev/dl/) según el sistema operativo que uses.
2. Instalar y dar siguiente hasta el final.
3. Añadir **%USERPROFILE%\go\bin** en la variable de entorno Path (en caso uses Windows).
4. Ejecutar en la terminal el comando `go version`, si todo está correctamente instalado debería salir lo siguiente:
```bash

go version go1.26.4 windows/amd64
```

## ⚙️ Instalación de librerías

1. Clona este repositorio.  
2. Abre el proyecto en tu editor de código.
3. Abre una nueva terminal en la ubicación de tu proyecto
4. Ejecuta el siguiente comando:

```bash
go get -u github.com/gorilla/mux
```

5. para ejecutar la API, en la terminal ejecuta el siguiente comando:

```bash
go run cmd/api/main.go
```
Saldrá una ventana emergente en Windows, darle en permitir acceso.
Ahora la terminal debería estar mostrando lo siguiente:

```bash

server executing on port :8080...
```

Probar la api en la ruta:

```bash
localhost:8080/api/results
```