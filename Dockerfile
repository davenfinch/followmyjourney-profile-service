# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY . .
RUN apk add --no-cache git
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/profile-service ./cmd/server

# Final image
FROM gcr.io/distroless/static-debian11
COPY --from=builder /out/profile-service /profile-service
EXPOSE 8081
ENTRYPOINT ["/profile-service"]
