FROM golang:alpine AS build-env

WORKDIR /go/src/github.com/martinomburajr/medium/go/tutorials/building-a-simple-server

COPY . .

RUN go build -o main

FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*

WORKDIR /go/src/github.com/martinomburajr/medium/go/tutorials/building-a-simple-server

COPY --from=build-env /go/src/github.com/martinomburajr/medium/go/tutorials/building-a-simple-server* /go/src/github.com/martinomburajr/medium/go/tutorials/building-a-simple-server

EXPOSE 8080

ENTRYPOINT [ "./main" ]