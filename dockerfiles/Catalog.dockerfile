FROM golang:1.24.6-bookworm

WORKDIR /app

COPY src/catalog/go.mod src/catalog/go.sum ./

RUN go mod download

COPY src/catalog/ .

WORKDIR /app/cmd/catalog
RUN CGO_ENABLED=0 GOOS=linux go build -o catalog

EXPOSE 3000
EXPOSE 8080

CMD ["./catalog"]
