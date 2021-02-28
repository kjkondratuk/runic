FROM ubuntu:latest

COPY . ~/
WORKDIR ~/

ENV PATH="$PATH:/usr/local/go/bin"

RUN apt-get update && \
    apt-get -y install wget build-essential && \
    wget https://golang.org/dl/go1.16.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.16.linux-amd64.tar.gz && \
    make build

ENTRYPOINT ["bin/runic"]