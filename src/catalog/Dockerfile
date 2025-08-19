FROM golang:1.24.6-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -C /app/cmd/server -o catalog


FROM alpine
WORKDIR /root/
COPY --from=builder /app ./app
EXPOSE 3000
WORKDIR /root/app
CMD ["./cmd/server/catalog"]
