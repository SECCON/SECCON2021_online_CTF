# FROM ubuntu:20.10
FROM ubuntu@sha256:a7b08558af07bcccca994b01e1c84f1d14a2156e0099fcf7fcf73f52d082791e

RUN \
  apt-get update && \
  apt-get -y install xinetd && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/

RUN chmod a-w /tmp /var/tmp /dev/shm
RUN chmod ug-s /bin/* /sbin/* /usr/bin/* /usr/sbin/*

RUN useradd average

COPY average.xinetd /etc/xinetd.d/average

WORKDIR /home/average

COPY average.sh average flag.txt ./
RUN \
  chmod 550 average.sh average && \
  chmod 440 flag.txt

RUN chown -R root:average /home/average

CMD ["xinetd", "-dontfork"]
