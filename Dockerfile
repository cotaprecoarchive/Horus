FROM busybox:latest

MAINTAINER Andrey K. Vital <andreykvital@gmail.com>

COPY ./build/horus-linux-amd64 /app/horus

WORKDIR /app

CMD /app/horus
