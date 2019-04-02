FROM golang:1.12.1 as builder

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o consignment-cli -a -installsuffix cgo cli.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

ADD consignment.json /app/consignment.json

COPY --from=builder /app/consignment-cli .

CMD ["./consignment-cli"]
