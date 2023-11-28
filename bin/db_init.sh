#! /bin/bash

set -e
readonly TIMEOUT_S=10
readonly INTERVAL=0.5

printf 'waiting for postgres...\n'
start_ts=$(date +%s)
set +e
while :
do
  pg_isready
  result=$?
  end_ts=$(date +%s)
  duration=$((end_ts - start_ts))
  if [[ $result -eq 0 ]]; then
    echo "  ...${duration}s"
    set -e
    break
  fi
  if [[ ${duration} -gt ${TIMEOUT_S} ]]; then
    echo "  ...timeout (${TIMEOUT_S}s)"
    set -e
    exit ${duration}
  fi
  sleep ${INTERVAL}
done


# Create database
PGPASSWORD=$POSTGRES_PASSWORD psql  --username=$POSTGRES_USER << EOF
CREATE DATABASE $POSTGRES_DATABASE;
EOF

# Connect to the new database and execute SQL commands
PGPASSWORD=$POSTGRES_PASSWORD psql --username=$POSTGRES_USER --dbname=$POSTGRES_DATABASE <<EOF
CREATE TYPE linkPrecedenceType AS ENUM ('primary', 'secondary');

CREATE TABLE Contact (
    id BIGSERIAL PRIMARY KEY,
    phonenumber INT NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    linkedID INTEGER[],
    linkPrecedence linkPrecedenceType DEFAULT 'primary',
    ipv4 VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX email_phone_number_idx ON Contact (email, phonenumber)
EOF
