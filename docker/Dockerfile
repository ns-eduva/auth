FROM golang:1.24.4-bookworm AS base
RUN apt-get update && apt-get install -y --no-install-recommends \
    git \
    ca-certificates \
 && rm -rf /var/lib/apt/lists/*
RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/nsevenpack/mignosql/cmd/migrationcreate@latest
RUN mkdir -p /app/tmp/air
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

FROM base AS dev
WORKDIR /app
COPY . .
CMD ["air", "-c", ".air.toml"]