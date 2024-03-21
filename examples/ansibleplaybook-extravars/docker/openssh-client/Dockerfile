FROM alpine:3.19

RUN apk add --update --no-cache \
        openssh-client \
        git \
    && rm -rf /var/cache/apk/*

ENTRYPOINT [ "/usr/bin/ssh-keygen" ]
CMD [ "--help" ]