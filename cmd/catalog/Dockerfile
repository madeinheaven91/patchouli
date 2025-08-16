FROM golang:1.24.6-bookworm

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/catalog
RUN CGO_ENABLED=0 GOOS=linux go build -o catalog

EXPOSE 3000

WORKDIR /app/cmd/catalog
CMD ["./catalog"]
