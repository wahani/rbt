FROM ubuntu:bionic

ADD . /root/go/src/rbt

RUN apt-get update \
  && apt-get install -y --no-install-recommends \
    golang-go \
    locales \
    ca-certificates \
#   && dpkg-reconfigure ca-certificates \
  && echo "en_US.UTF-8 UTF-8" >> /etc/locale.gen \
  && locale-gen en_US.utf8 \
  && /usr/sbin/update-locale LANG=en_US.UTF-8

RUN go install rbt \
  && ln -s /root/go/bin/rbt /usr/local/bin/rbt

CMD [ "rbt -install" ]

