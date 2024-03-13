# ARG image_from_registry_host=dev-hub.odigeoconnect.com
# ARG image_from_registry_namespace=alpha
# ARG image_from_name=ubuntu
# ARG image_from_tag=20.04

# FROM ${image_from_registry_host}/${image_from_registry_namespace}/${image_from_name}:${image_from_tag}

# USER root

# RUN apt-get update \
#   && apt-get install --no-install-recommends --yes \
#     build-essential \
#     libssl-dev \
#     libffi-dev \
#     python3 \
#     python3-dev \
#     python3-pip \
#     openssh-client \
#     git \
#   && apt-get clean \
#   && apt-get autoremove --yes \
#   && rm -rf /var/lib/apt/lists/* \
#   && ln -s /usr/bin/python3 /usr/bin/python

# ARG ansible_version
# ARG ansible_strategy=mitogen
# ARG ansible_stdout_callback=yaml

# RUN pip3 install -U pip setuptools \
#   && pip3 install --no-cache-dir \
#       setuptools-rust \
#       cryptography \
#       boto \
#       boto3 \
#       botocore \
#       mitogen \
#       # Required library to execute ansible community.general.dig plugin
#       dnspython \
#       ansible==${ansible_version} \
#   && ln /usr/local/bin/ansible-playbook /usr/bin/ansible-playbook

# RUN apt-get update \
#   && apt-get install --no-install-recommends --yes \
#     curl \
#     unzip \
#     dnsutils \
#   && curl -sL https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip -o /tmp/awscliv2.zip \
#   && cd /tmp \
#   && unzip awscliv2.zip \
#   && aws/install \
#   && rm -rf \
#     /tmp/aws \
#     /tmp/awscliv2.zip \
#   && apt-get purge --yes \
#     unzip \
#   && apt-get clean \
#   && apt-get autoremove --yes \
#   && rm -rf /var/lib/apt/lists/*

# ENV ANSIBLE_FORCE_COLOR=true
# ENV ANSIBLE_HOST_KEY_CHECKING=False
# ENV ANSIBLE_STDOUT_CALLBACK=${ansible_stdout_callback}
# ENV ANSIBLE_STRATEGY_PLUGINS=/usr/local/lib/python3.8/dist-packages/ansible_mitogen/plugins/strategy
# ENV ANSIBLE_STRATEGY=${ansible_strategy}
# ENV PYTHONUNBUFFERED=1

# ENTRYPOINT ["/usr/local/bin/ansible-playbook"]
# CMD ["--version"]

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