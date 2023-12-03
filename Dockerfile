# build
ARG GO_VERSION=1.21.4

FROM golang:${GO_VERSION}-bullseye AS builder
LABEL version="1.0"

WORKDIR /server

RUN apt-get update && \
    apt-get install -y --no-install-recommends

COPY ./ ./
RUN go install github.com/cosmtrek/air@latest && \
    go mod download

ARG CGO_ENABLED=1
ARG GOOS=linux
RUN go build -ldflags="-s -w -linkmode external -extldflags -static" -o /server/bin ./cmd/server

CMD ["/server/docker/entrypoint.sh"]

# runtime
FROM scratch

ENV TZ Asia/Tokyo

COPY --from=builder /server/bin /server/bin
CMD ["/server/bin"]

EXPOSE 8080