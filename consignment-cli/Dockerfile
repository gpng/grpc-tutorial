FROM golang:1.12.1 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

ADD consignment.json /app/consignment.json
ADD consignment-cli /app/consignment-cli

COPY --from=builder /go/src/github.com/gpng/grpc-tutorial/consignment-cli/consignment-cli .

CMD ["./consignment-cli"]
