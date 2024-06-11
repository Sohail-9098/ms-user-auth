# Use Alpine Linux as the base image
FROM golang:1.20-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Set Go environment variables
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# Copy the Go application source code
WORKDIR /src
COPY . .

# Build the Go application binary
RUN go build -o ms-user-auth .

# Use a minimal base image for the final image
FROM alpine:latest

# Install necessary runtime dependencies
RUN apk add --no-cache postgresql postgresql-contrib

# Copy the Go application binary from the builder image
COPY --from=builder /src/ms-user-auth /usr/local/bin/ms-user-auth

# Copy the config.yaml file
COPY config.yaml /usr/local/bin/config.yaml

# Set up PostgreSQL environment variables
ENV POSTGRES_USER user
ENV POSTGRES_PASSWORD password
ENV POSTGRES_DB mydb

# Initialize the PostgreSQL database
RUN mkdir -p /docker-entrypoint-initdb.d
COPY db/init-db.sql /docker-entrypoint-initdb.d/init-db.sql

# Copy the PostgreSQL configuration file
COPY postgresql.conf /etc/postgresql/postgresql.conf

# Copy the entrypoint script
COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

# Expose the port the Go app listens on
EXPOSE 4000

# Set the entrypoint
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
