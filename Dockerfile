# Imagem base para compilação
FROM golang:1.20.5-alpine3.17 AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia o código fonte para o diretório de trabalho
COPY go.* ./

RUN go mod download

# Copy local code to the container image.
COPY s3-microservice ./

# Build the binary.
RUN go build -v -o server

# Imagem mínima para execução
FROM alpine:latest

# Copy the binary to the production image from the builder stage.
COPY --from=builder app/server /app/server

# Expõe a porta utilizada pelo microserviço
EXPOSE 8080

# Run the web service on container startup.
CMD ["/app/server"]