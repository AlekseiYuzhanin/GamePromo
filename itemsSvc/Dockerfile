FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY pkg /pkg

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd/app

FROM scratch

COPY --from=builder /app/myapp /myapp

ENTRYPOINT ["/myapp"]