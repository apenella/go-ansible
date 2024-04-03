FROM python:bookworm


RUN adduser --disabled-password --gecos "" ansible

RUN apt-get update \
    && apt-get install -y \
        openssh-server \
    && rm -rf /var/lib/apt/lists/*

RUN mkdir /var/run/sshd \
    && ssh-keygen -A

COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

# Expose the SSH port
EXPOSE 22

CMD ["/usr/local/bin/entrypoint.sh"]
