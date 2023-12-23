# build
ARG GO_VERSION=1.21.4

FROM golang:${GO_VERSION}-bullseye AS builder
LABEL version="1.0"

WORKDIR /server

RUN apt-get update && \
    apt-get install -y --no-install-recommends gcc-x86-64-linux-gnu libc6-dev-amd64-cross

COPY ./ ./
RUN go install github.com/cosmtrek/air@latest && \
    go mod download

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
ARG CC=x86_64-linux-gnu-gcc
RUN go build -ldflags="-s -w -linkmode external -extldflags -static" -o /server/bin ./cmd/proxy-server-go

CMD ["air", "-c", "docker/proxy-server-go/.air.toml"]

# runtime
FROM scratch AS runtime

ENV TZ Asia/Tokyo

COPY --from=builder /server/bin /server/bin

EXPOSE 8888

CMD ["/server/bin"]