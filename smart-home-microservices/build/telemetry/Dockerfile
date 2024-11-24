FROM golang:1.22-alpine as builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o main cmd/telemetry/main.go

FROM alpine:3.18
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/main /app/main
WORKDIR /app
EXPOSE 8080
CMD ["./main"]
