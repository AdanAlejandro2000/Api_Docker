# Usa una imagen de Go como base para compilar la aplicación
FROM golang:latest AS builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el código fuente al directorio de trabajo dentro del contenedor
COPY . .

# Descarga e instala las dependencias necesarias (en este caso, solo Gorilla Mux)
RUN go mod init app && go get -u github.com/gorilla/mux

# Compila el código de Go dentro del contenedor
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Usa una imagen ligera como base para la ejecución de la aplicación
FROM alpine:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /root/

# Copia el binario compilado desde el primer contenedor
COPY --from=builder /app/main .

# Exponer el puerto en el que se ejecutará la aplicación
EXPOSE 8080

# Ejecutar la aplicación cuando el contenedor se inicie
CMD ["./main"]
