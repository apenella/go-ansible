ARG golang_version=1.22

FROM golang:${golang_version}-bookworm as golang

FROM python:3.12-bookworm as builder

RUN apt-get update \
    && apt-get install -y \
        openssh-client \
        git \
    && rm -rf /var/lib/apt/lists/*

COPY --from=golang /usr/local/go /usr/local/go

# Configure Go
ENV GOROOT /usr/local/go
ENV PATH /usr/local/go/bin:/go/bin:$PATH

WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

## go generate is used to install go-embed-python
RUN go generate ./...

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /usr/local/bin/app ansibleplaybook-embed-python.go

FROM debian:bookworm-slim

COPY --from=builder /usr/local/bin/app /usr/local/bin/app

CMD ["/usr/local/bin/app"]

