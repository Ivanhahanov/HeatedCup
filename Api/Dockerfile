FROM golang:1.14.3-alpine AS build
WORKDIR $GOPATH/src/HeatedCup/Api
RUN apk --no-cache add build-base git bzr mercurial gcc
RUN go get github.com/eclipse/paho.mqtt.golang
COPY . .
RUN go build -o /bin/goCup .
FROM alpine
WORKDIR /app
COPY --from=build /bin/goCup /app/
CMD ["./goCup"]