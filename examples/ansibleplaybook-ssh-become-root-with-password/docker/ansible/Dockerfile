ARG golang_version=1.22

FROM golang:${golang_version}-bookworm as golang

FROM python:3.12-alpine3.19

ARG ansible_core_version=2.17.5

RUN apk add --update --no-cache \
        openssh-client \
        git \
    && rm -rf /var/cache/apk/*

RUN pip3 install -U pip setuptools \
    && pip3 install --no-cache-dir \
        setuptools-rust \
        cryptography \
        # Required library to execute ansible community.general.dig plugin
        dnspython \
        ansible-core==${ansible_core_version} \
    && ln /usr/local/bin/ansible-playbook /usr/bin/ansible-playbook \
    && ansible-galaxy collection install community.general

COPY --from=golang /usr/local/go /usr/local/go

# Configure Go
ENV GOROOT /usr/local/go
ENV PATH /usr/local/go/bin:/go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin
