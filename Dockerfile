FROM ubuntu:latest

COPY . ~/
WORKDIR ~/

RUN apt-get update && \
    apt-get -y install wget build-essential && \
    wget https://golang.org/dl/go1.16.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.16.linux-amd64.tar.gz && \
    export PATH=$PATH:/usr/local/go/bin && \
    go version

RUN make build

RUN ls -al bin

ENTRYPOINT ["bash"]