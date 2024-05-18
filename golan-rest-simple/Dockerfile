# Use la imagen de Go como base
FROM golang:latest AS builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el código fuente al directorio de trabajo dentro del contenedor
COPY golan-rest-simple /app

# Instala las dependencias
RUN go get -u github.com/gorilla/mux

# Compila el código de Go dentro del contenedor
RUN go build -o main .

# Exponer el puerto en el que se ejecutará la aplicación
EXPOSE 8080

# Ejecutar la aplicación cuando el contenedor se inicie
CMD ["./main"]
