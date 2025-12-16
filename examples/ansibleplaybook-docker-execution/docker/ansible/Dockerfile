FROM python:3.14-alpine3.23

# Create a non-root user and group (ansible)
RUN addgroup -S ansible \
    && adduser -S ansible -G ansible \
    && mkdir -p /home/ansible/.ansible \
    && chown -R ansible:ansible /home/ansible

# Install required packages
RUN apk add --update --no-cache \
        openssh-client \
        git \
        && rm -rf /var/cache/apk/*

# Switch to the non-root user
USER ansible:ansible

RUN pip3 install -U pip setuptools \
    && pip3 install --no-cache-dir \
        setuptools-rust \
        cryptography \
        # Required library to execute ansible community.general.dig plugin
        dnspython \
        ansible

ENV PATH /home/ansible/.local/bin:$PATH 

