#!/bin/sh

home=$(getent passwd "$(whoami)" |  cut -d: -f6)

mkdir -p "${home}/.ssh"
cat /ssh/id_rsa.pub > "${home}/.ssh/authorized_keys"

/usr/sbin/sshd -D