# Stage 1: Build
FROM golang:1.22-alpine AS builder
WORKDIR /app
RUN go mod init app && go get github.com/lib/pq
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Final
FROM alpine:3.19
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
