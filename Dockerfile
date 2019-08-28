FROM alpine:3.10

COPY bin/m3uproxy /usr/local/bin

CMD m3uproxy