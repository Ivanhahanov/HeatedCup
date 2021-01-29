FROM golang:1.14.3-alpine AS build
WORKDIR /src
COPY Api .
RUN go build -o /bin/goCup .
FROM alpine
WORKDIR /app
COPY --from=build /bin/goCup /app/
CMD ["./goCup"]