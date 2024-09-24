#!/bin/sh

home=$(getent passwd "$(whoami)" |  cut -d: -f6)

mkdir -p "${home}/.ssh"
cat /ssh/id_rsa.pub > "${home}/.ssh/authorized_keys"

{
  echo "PasswordAuthentication no"
  echo "PubkeyAuthentication yes"
  echo "AllowUsers aleix"
} >> /etc/ssh/sshd_config

adduser aleix -D
cp -r ~/.ssh /home/aleix/
chown -R aleix:aleix /home/aleix/.ssh
echo 'aleix:12345' | chpasswd; 

apk update
apk add sudo

echo "aleix ALL=(ALL:ALL) ALL" >> /etc/sudoers

## -e option will print the logs to the console
/usr/sbin/sshd -D -e
