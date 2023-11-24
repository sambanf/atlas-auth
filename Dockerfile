FROM golang:alpine AS build-stage

RUN  apk add && apk update --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o ./bin/atlas-auth ./cmd

FROM alpine:latest AS release-stage

WORKDIR /

COPY --from=build-stage /app/bin/atlas-auth /atlas-auth

EXPOSE 8080

ENTRYPOINT ["/atlas-auth"]
