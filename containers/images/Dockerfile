# Build stage
FROM golang:1.18-alpine3.15 AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o /bin/app ./cmd

# Run stage
FROM alpine:3.15
WORKDIR /src
COPY --from=builder /bin/app /bin/app
COPY --from=builder /src/views /src/views
COPY --from=builder /src/assets /src/assets

ENTRYPOINT [ "/bin/app" ]