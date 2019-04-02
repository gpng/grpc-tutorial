FROM golang:1.12.1 as builder

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o vessel-service -a -installsuffix cgo main.go repository.go handler.go datastore.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/vessel-service .

CMD ["./vessel-service"]