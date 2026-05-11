# 1. Use the modern Go compiler (>= 1.25)
FROM golang:1.26-alpine

RUN apk add --no-cache git build-base

WORKDIR /app

# 2. Pinned to the specific latest version instead of floating
RUN go install github.com/air-verse/air@v1.65.1

# 3. Optional go.sum to prevent fresh-project build failures
COPY go.mod go.sum* ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]