## Build
FROM golang:1.19-alpine AS build

WORKDIR /distributed-cfg-service-mk

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

COPY ./server/main.go ./
COPY ./proto/schema_grpc.pb.go ./proto/schema_grpc.pb.go
COPY ./proto/schema.pb.go      ./proto/schema.pb.go

RUN go build -o /cfg-service-mk ./main.go

## Deploy
FROM alpine

COPY --from=build /cfg-service-mk /app/cfg-service-mk

EXPOSE 50051

ENTRYPOINT ["/app/cfg-service-mk"]

