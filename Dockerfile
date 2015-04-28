FROM busybox:latest

MAINTAINER Andrey K. Vital <andreykvital@gmail.com>

COPY ./build/horus /app/horus

EXPOSE 7600
EXPOSE 8000

WORKDIR /app

CMD /app/horus
