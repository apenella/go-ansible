FROM python:alpine3.18

ARG sshd_port=2222

RUN apk add --update --no-cache \
        openssh-server \
    # create a user
    && adduser -D ansible \
    && rm -rf /var/cache/apk/* \
    && ssh-keygen -A

RUN echo "Port ${sshd_port}" >> /etc/ssh/sshd_config

COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

# Expose the SSH port
EXPOSE ${sshd_port}

CMD ["/usr/local/bin/entrypoint.sh"]
