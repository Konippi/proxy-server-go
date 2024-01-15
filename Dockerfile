# build
ARG GO_VERSION=1.21.4

FROM golang:${GO_VERSION}-bullseye AS build
LABEL version="1.0"

WORKDIR /proxy-server

RUN --mount=type=cache,target=/var/cache/apt \
    --mount=type=cache,target=/var/lib/apt \
    apt-get update && \
    apt-get install -y --no-install-recommends gcc-x86-64-linux-gnu libc6-dev-amd64-cross

COPY go.mod go.sum ./
RUN go install github.com/cosmtrek/air@latest && \
    go mod download

COPY ./ ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
ARG CC=x86_64-linux-gnu-gcc
RUN go build -ldflags="-s -w -linkmode external -extldflags -static" -o /proxy-server/bin ./cmd/proxy-server-go

CMD ["air", "-c", ".air.toml"]

# runtime
FROM scratch AS runtime

ENV TZ Asia/Tokyo

COPY --from=build /proxy-server/bin /proxy-server/bin

EXPOSE 8888

CMD ["/server/bin"]