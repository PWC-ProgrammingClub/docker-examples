FROM golang:1.20-alpine3.18 AS getdeps
# link to image: https://hub.docker.com/_/golang/

WORKDIR /usr/src/server/


COPY ./.docker/go/go.mod ./
COPY ./.docker/go/go.sum ./

# Get the dependencies in the first image layer to cache them for faster builds
RUN go mod download && go mod verify

FROM getdeps

COPY ./.docker/go ./

RUN go build .

ARG username="root"
ARG password="password"
ARG db="db"
ARG host="localhost"
ARG httpport="8080"

ENV DB_USER=${username}
ENV DB_PASSWORD=${password}
ENV DB_HOST=${host}
ENV DB_DB=${db}
ENV HTTP_PORT=${httpport}

CMD ["./gopo"]