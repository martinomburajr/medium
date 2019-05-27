FROM golang@sha256:d481168873b7516b9f34d322615d589fafb166ff5fd57d93e96f64787a58887c as build-env
MAINTAINER "Martin Ombura Jr."

RUN apk update && apk add --no-cache git ca-certificates tzdata bash && update-ca-certificates
RUN apk update && apk upgrade && apk add --no-cache git

# Set Working Directory
WORKDIR $GOPATH/src/github.com/martinomburajr/tutorials/cloudscheduler-cloudfunctions-cloudrun/cloudrun
COPY . .

RUN go build -o /go/bin/cloudrun

FROM alpine:3.9

COPY --from=build-env /go/bin/cloudrun /go/bin/cloudrun

EXPOSE 8080

ENTRYPOINT ["./go/bin/cloudrun"]
