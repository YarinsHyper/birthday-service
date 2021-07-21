#build stage
FROM golang:alpine AS builder
ENV GO111MODULE=on
RUN apk add --no-cache git make
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build

#final stage
FROM scratch
COPY --from=builder /go/src/app/birthday-service /birthday-service
COPY --from=builder /bin/grpc_health_probe /bin/grpc_health_probe
LABEL Name=birthday-service Version=0.0.1
EXPOSE 8000
ENTRYPOINT ["/birthday-service"]