# Build stage
FROM golang:1.20.8-alpine3.18 AS BuildStage

RUN apk add --no-cache make

WORKDIR /app

COPY . .

RUN go mod download

RUN make build


# Deploy stage
FROM alpine:latest

WORKDIR /app

COPY --from=BuildStage /app/bin/ /app/bin

COPY --from=BuildStage /app/cmd/ /app/cmd

CMD ["/app/bin/auth"]
