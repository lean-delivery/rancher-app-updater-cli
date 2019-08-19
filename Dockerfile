FROM golang:latest as build-stage

# METADATA
LABEL maintainer="team@lean-delivery.com" 

WORKDIR /go/src/
COPY src .

RUN go get github.com/go-resty/resty
RUN go get github.com/sirupsen/logrus

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o charts-updater-cli main/main.go

FROM alpine

RUN apk --no-cache --update add ca-certificates
RUN update-ca-certificates

WORKDIR /
COPY --from=build-stage /go/src/charts-updater-cli /bin/charts-updater-cli

CMD ["charts-updater-cli"]
