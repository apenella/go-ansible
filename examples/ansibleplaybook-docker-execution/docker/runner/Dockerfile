# Based on https://github.com/timberio/docker-client/blob/master/Dockerfile
ARG golang_version="1.23"
ARG docker_version="27.1"
ARG program_name="example"

FROM golang:${golang_version}-alpine AS builder

WORKDIR /code

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN cd /code/examples/ansibleplaybook-docker-execution && \
    go build -o app

FROM docker:${docker_version}-dind

COPY --from=builder /code/examples/ansibleplaybook-docker-execution/docker/runner/entrypoint.sh /usr/local/bin/entrypoint.sh
COPY --from=builder /code/examples/ansibleplaybook-docker-execution/docker/runner/wait-for-dockerd.sh /usr/local/bin/wait-for-dockerd.sh
COPY --from=builder /code/examples/ansibleplaybook-docker-execution /code

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["/code/app"]
