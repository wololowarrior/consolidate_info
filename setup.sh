#! /bin/bash


docker build --target pg_db -t pg_db .

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
