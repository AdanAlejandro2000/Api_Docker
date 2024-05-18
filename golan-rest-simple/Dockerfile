# Usar una imagen base de Golang
FROM golang:alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el código fuente al directorio de trabajo dentro del contenedor
COPY . /app

# Instalar las dependencias
RUN go mod tidy

# Compilar la aplicación
RUN go build -o main .

# Especificar el comando para ejecutar la aplicación
CMD ["./main"]
