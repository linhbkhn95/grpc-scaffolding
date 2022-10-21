# compared to Ubuntu
FROM golang:1.18-alpine as builder

WORKDIR /rpc

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o /rpc/rpc-runtime cmd/server/main.go

FROM alpine:latest
EXPOSE 10080 10433
WORKDIR /rpc

COPY --from=builder /rpc/rpc-runtime /rpc/rpc-runtime
CMD ["/rpc/rpc-runtime", "server"]
