# set up the build environment
FROM golang:1.21-alpine AS builder

# set the working directory in the container
WORKDIR /app

# copy files that manage dependencies in the first  step
COPY go.mod go.sum ./

# resolve dependencies in the container
RUN go mod download

# copy the source code
COPY ./ ./

# build the application in the same container and app directory 
# under banking-service executable
RUN go build -o api

# set up the runtime environment
FROM alpine:latest

# set the working directory in the container
WORKDIR /app

# copy the executable from the builder container
COPY --from=builder /app/api .

# run the application
CMD ["./api"]