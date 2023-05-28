# Simple usage with a mounted data directory:
# > docker build -t cf-sre .

FROM golang:1.20-alpine AS build-env

# Install minimum necessary dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev
RUN apk add --no-cache $PACKAGES

# Set working directory for the build
WORKDIR /go/src/github.com/vedhavyas/cf-sre

# optimization: if go.sum didn't change, docker will use cached image
COPY go.mod go.sum ./

RUN go mod download

# Add source files
COPY . .

# Dockerfile Cross-Compilation Guide
# https://www.docker.com/blog/faster-multi-platform-builds-dockerfile-cross-compilation-guide
ARG TARGETOS TARGETARCH

RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o cf-sred ./cmd/cf-sred/main.go

# Use alpine:3 as a base image
FROM alpine:3

EXPOSE 26656 26657 1317 9090 2345
CMD ["cf-sred"]
STOPSIGNAL SIGTERM
WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/src/github.com/vedhavyas/cf-sre/cf-sred /usr/bin/cf-sred