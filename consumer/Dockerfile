FROM golang:1.22.0-alpine AS build

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./

RUN go build -o /consumer

FROM scratch

WORKDIR /
COPY --from=build /consumer /consumer
EXPOSE 8080

ENTRYPOINT ["/consumer"]
