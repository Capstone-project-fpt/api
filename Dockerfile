FROM golang:alpine AS builder

WORKDIR /app 

COPY . .

RUN go mod download

RUN go build -o build ./cmd/server

FROM scratch

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/build /build
COPY ./config.yaml /config.yaml
COPY ./internal/locales /internal/locales
COPY ./templates /templates

# Set the entry point to the binary
ENTRYPOINT [ "/build" ]
