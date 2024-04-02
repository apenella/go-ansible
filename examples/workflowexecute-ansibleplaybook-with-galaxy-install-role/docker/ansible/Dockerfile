FROM python:3.12-alpine3.19

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
        ansible \
    && ln /usr/local/bin/ansible-playbook /usr/bin/ansible-playbook

RUN apk add --update --no-cache \
        go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin
