# Challenge Leal

Este proyecto es una aplicación para gestionar campañas de puntos o cashback para comercios y sucursales.

## Requisitos

- [Go](https://golang.org/dl/) (para ejecutar y desarrollar el código)
- [MySQL](https://dev.mysql.com/downloads/) (como base de datos)

## Configuración del Proyecto

### 1. Clonar el Repositorio

```bash
git clone https://github.com/jair921/challenge-leal.git
cd challenge-leal
```

### 2. Configurar la Base de Datos

Crear una base de datos MySQL llamada lealcoins (o el nombre que prefieras).

Configurar el archivo de configuración con los detalles de tu base de datos. 
Modifica el archivo config.deveploment.yml o el de producci{on
para incluir los detalles de la conexión a la base de datos, como el nombre de la base de datos, usuario y contraseña.

### 2. Configurar la Base de Datos


Instalar golang-migrate
```bash
go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest 
```

Ejecutar (ajsutando los datos de tu conexión)
```bash
migrate -source file://migrations -database 'root:root@tcp(localhost:3306)/lealcoins?charset=utf8&parseTime=True&loc=Local' up 
```

4. Ejecutar la Aplicación

Instalar dependencias:
```bash
go mod download 
```

Ejecutar la aplicación:

```bash
go run cmd/main.go
```