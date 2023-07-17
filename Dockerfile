FROM golang as base

ENV GO111MODULE=on
ENV PROJECT accuknox

WORKDIR /$PROJECT
COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 cd main && go build
EXPOSE 10000
CMD ["/bin/bash","-c","cd main && ./main"]
