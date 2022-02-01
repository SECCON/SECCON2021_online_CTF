FROM ubuntu:20.04

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get -y update --fix-missing && apt-get -y upgrade
RUN apt-get -y install xinetd
RUN apt-get -y install gcc
RUN apt-get -y install python3.9
RUN groupadd -r pwn && useradd -r -g pwn pwn

ADD files/pyast64++.pwn/xinetd/pwn.xinetd /etc/xinetd.d/pwn
ADD files/pyast64++.pwn/xinetd/init.sh    /etc/init.sh
RUN chmod 700 /etc/init.sh
RUN chmod 1733 /tmp /var/tmp /dev/shm

WORKDIR /home/pwn
ADD app/flag.txt .
ADD files/pyast64++.pwn/pyast64.py .
RUN chmod 440 flag.txt
RUN chmod 550 pyast64.py
RUN chown -R root:pwn /home/pwn
RUN mv flag.txt flag-$(md5sum flag.txt | awk '{print $1}').txt

RUN service xinetd restart
