# Use la imagen de Go como base
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el código fuente al directorio de trabajo dentro del contenedor
COPY . .

# Descarga e instala las dependencias necesarias (en este caso, solo Gorilla Mux)
RUN go install -u github.com/gorilla/mux

# Compila el código de Go dentro del contenedor
RUN go build -o main .

# Exponer el puerto en el que se ejecutará la aplicación
EXPOSE 8080


# Ejecutar la aplicación cuando el contenedor se inicie
CMD ["./main"]
