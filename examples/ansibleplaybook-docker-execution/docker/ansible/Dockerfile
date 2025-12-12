<<<<<<< HEAD
<<<<<<< HEAD
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

# Install ansible and required python packages
=======
ARG golang_version=1.23
=======
FROM python:3.14-alpine3.23
>>>>>>> a2a1d16 (Add example for executing Ansible playbooks in Docker)

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

<<<<<<< HEAD
>>>>>>> f904139 (create an example to execute Ansible within a container)
=======
# Switch to the non-root user
USER ansible:ansible

# Install ansible and required python packages
>>>>>>> a2a1d16 (Add example for executing Ansible playbooks in Docker)
RUN pip3 install -U pip setuptools \
    && pip3 install --no-cache-dir \
        setuptools-rust \
        cryptography \
        # Required library to execute ansible community.general.dig plugin
        dnspython \
<<<<<<< HEAD
<<<<<<< HEAD
        ansible

ENV PATH /home/ansible/.local/bin:$PATH 
=======
        ansible \
    && ln /usr/local/bin/ansible-playbook /usr/bin/ansible-playbook

COPY --from=golang /usr/local/go /usr/local/go

# Configure Go
ENV GOROOT /usr/local/go
ENV PATH /usr/local/go/bin:/go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin
>>>>>>> f904139 (create an example to execute Ansible within a container)
=======
        ansible

ENV PATH /home/ansible/.local/bin:$PATH 
>>>>>>> a2a1d16 (Add example for executing Ansible playbooks in Docker)
