FROM golang:1.21.3-alpine3.18 as build

COPY routes /build/routes
COPY todos /build/todos
COPY go.mod /build/go.mod
COPY go.sum /build/go.sum
COPY main.go /build/main.go

WORKDIR /build
RUN go build

FROM alpine:3.18

COPY --from=build /build/web-dev /app/web-dev
COPY templates /app/templates
WORKDIR /app
CMD ["./web-dev"]