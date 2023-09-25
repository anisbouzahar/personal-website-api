FROM arm32v7/golang:1.21-alpine AS builder
ARG VERSION=latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN  CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o portfolio-api ./cmd/app

FROM arm32v7/alpine:latest

COPY --from=builder /app/portfolio-api .
EXPOSE 8080

ENTRYPOINT ["./portfolio-api"]