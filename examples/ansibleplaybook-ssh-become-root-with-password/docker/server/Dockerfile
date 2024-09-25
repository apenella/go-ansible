FROM python:alpine3.18

RUN apk add --update --no-cache \
        openssh-server \
    # create a user
    && adduser -D ansible \
    && rm -rf /var/cache/apk/* \
    && ssh-keygen -A

COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

# Expose the SSH port
EXPOSE 22

CMD ["/usr/local/bin/entrypoint.sh"]
