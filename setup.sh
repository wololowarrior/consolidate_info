#! /bin/bash


docker build --target pg_db -t pg_db .

echo "Removing old stuff"
docker rm assignment_db consolidation_service -f
docker volume rm pg_data -f

echo "Creating experiment"
docker volume create pg_data
docker network create app-network

docker run \
        --env-file .env \
        -d \
        --name assignment_db \
        --hostname assignment_db  \
        -v pg_data:/var/lib/postgresql/data \
        --network app-network \
        -p 5432:5432 \
        pg_db 
docker exec assignment_db /bin/bash db_init.sh > /dev/null
docker run --env-file .env --network app-network -p 10000:10000 -d --name consolidation_service  consolidation_service:latest
