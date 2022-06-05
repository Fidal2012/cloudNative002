FROM ubuntu:20.04
ADD ./bin/amd64/httpServer /httpServer
EXPOSE 80/tcp
ENTRYPOINT /httpServer

