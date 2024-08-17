FROM golang:alpine AS builder

WORKDIR /app 

COPY . .

RUN go mod download

RUN go build -o build ./cmd/server

FROM scratch

COPY --from=builder /app/build /build
COPY ./config.yaml /config.yaml

# Set the entry point to the binary
ENTRYPOINT [ "/build" ]
