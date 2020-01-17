FROM ubuntu:xenial
RUN apt-get update
RUN apt-get install tzdata
RUN echo "Asia/Shanghai" > /etc/timezone
RUN rm -f /etc/localtime
RUN dpkg-reconfigure -f noninteractive tzdata
RUN apt-get install -y ca-certificates
#RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai/etc/localtime
#RUN ln -fs /usr/share/zoneinfo/US/Pacific-New /etc/localtime && dpkg-reconfigure -f noninteractive tzdata
WORKDIR /root/deploy/az-util

RUN mkdir -p /root/deploy/az-util
COPY config.toml /root/deploy/az-util/config.toml
COPY golang-az-util /root/deploy/az-util/
RUN chmod +x /root/deploy/golang-az-util/az-util

ENTRYPOINT ["/root/deploy/az-util/az-util"]
