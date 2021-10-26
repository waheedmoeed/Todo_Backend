# Build stage
FROM golang:1.16-alpine3.13 AS builder
WORKDIR /app
COPY . .
RUN go build
RUN apk --no-cache add curl

# Run stage
FROM golang:1.16-alpine3.13
WORKDIR /app
COPY --from=builder /app/todobackend .



EXPOSE 8001
CMD [ "/app/todobackend" ]