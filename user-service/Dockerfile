# We use the official golang image, which contains all the
# correct build tools and libraries, Notice `as builder`,
# this gives the container a name that we can reference later on.
FROM golang:1.12.1 as builder

# Set our workdir to our current service in the gopath
RUN mkdir /app
WORKDIR /app

# Copy .mod and .sum for dependancies list
COPY go.mod .
COPY go.sum .

# Download dependancies
RUN go mod download

# Copy the current code into the workdir
COPY . .

# Build the binary, with a few flags which will allow
# us to run this binary in Alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o user-service -a -installsuffix cgo main.go repository.go handler.go database.go

# Here we're using a second FROM statement, which is strange,
# but this tells Docker to start a new build process with this
# image.
FROM alpine:latest

# Security related package, good to have.
RUN apk --no-cache add ca-certificates

# Same as before, create a directory for our app.
RUN mkdir /app
WORKDIR /app

# Here, instead of copying the binary from our host machine,
# we pull the binary from the container named `builder`, within
# this build context. This reaches into our previous image, finds
# the binary we build, and pulls it into this container. Amazing!
COPY --from=builder /app/user-service .

# Run the binary as per usual! This time with a binary build in a
# separate container, with all of the correct dependancies and
# run time libraries.
CMD [ "./user-service" ]