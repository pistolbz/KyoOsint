FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./goosint


FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/goosint .
EXPOSE 8088
ENTRYPOINT ["./goosint"]S
