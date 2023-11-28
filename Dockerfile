FROM golang as service

ENV GO111MODULE=on
ENV PROJECT emotorad

WORKDIR /$PROJECT
COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 cd cmd && go build main.go
EXPOSE 10000
CMD ["/bin/bash","-c","cd cmd && ./main"]

FROM postgres as pg_db
COPY bin/db_init.sh /db_init.sh
EXPOSE 5432
# CMD [ "bash", "db_init.sh" ]
