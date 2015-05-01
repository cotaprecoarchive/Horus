FROM busybox:latest

MAINTAINER Andrey K. Vital <andreykvital@gmail.com>

COPY ./build/horus-linux-amd64 /app/horus

ENV UDP_RECEIVER_HOST 0.0.0.0
ENV UDP_RECEIVER_PORT 7600
ENV WS_HOST 0.0.0.0
ENV WS_PORT 8000

EXPOSE 7600
EXPOSE 8000

WORKDIR /app

CMD /app/horus
