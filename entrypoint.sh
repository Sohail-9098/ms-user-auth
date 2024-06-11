#!/bin/sh

# Remove existing PostgreSQL data directory if it exists
if [ -d "/var/lib/postgresql/data" ]; then
    echo "Removing existing PostgreSQL data directory..."
    rm -rf /var/lib/postgresql/data
fi

# Initialize PostgreSQL data directory
echo "Initializing PostgreSQL data directory..."
su postgres -c "initdb -D /var/lib/postgresql/data"

# Set ownership and permissions for the PostgreSQL data directory
chown -R postgres:postgres /var/lib/postgresql/data
chmod 700 /var/lib/postgresql/data

# Start PostgreSQL and log output
echo "Starting PostgreSQL..."
su postgres -c "postgres -D /var/lib/postgresql/data -k /var/run/postgresql -h '' -c config_file=/etc/postgresql/postgresql.conf" &

# Wait for PostgreSQL to start
echo "Waiting for PostgreSQL to start..."
until su postgres -c "pg_isready"; do
    sleep 1
done
echo "PostgreSQL started successfully!"

# Run the initialization script (if needed)
echo "Running database initialization script..."
psql -U $POSTGRES_USER -d $POSTGRES_DB -f /docker-entrypoint-initdb.d/init-db.sql

# Start the Go application
echo "Starting Go application..."
/usr/local/bin/ms-user-auth
