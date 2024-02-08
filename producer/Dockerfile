FROM golang:1.22.0-alpine AS build

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./

RUN go build -o /producer

FROM scratch

WORKDIR /
COPY --from=build /producer /producer
EXPOSE 8080

ENTRYPOINT ["/producer"]
