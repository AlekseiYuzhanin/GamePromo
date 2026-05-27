FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o items_svc cmd/app/main.go

FROM scratch

COPY --from=builder /app/items_svc /items_svc

ENTRYPOINT ["/items_svc"]