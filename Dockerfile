# Use Alpine Linux as base image
FROM alpine:latest

# Install Go and required dependencies
RUN apk add --no-cache go gcc musl-dev

# Set Go environment variables
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH

# Install PostgreSQL and required dependencies
RUN apk add --no-cache postgresql

# Set up a PostgreSQL database and table
RUN mkdir -p /docker-entrypoint-initdb.d
COPY db/init-db.sql /docker-entrypoint-initdb.d/init-db.sql

# Copy the code of your Go application
COPY . /ms-user-auth
WORKDIR /ms-user-auth

# Build the Go application binary
RUN go build -o ms-user-auth .

# Expose the port the Go app listens on
EXPOSE 4000

# Start the Go application
CMD ["./ms-user-auth"]
