FROM ubuntu:16.04

RUN apt-get update \
    && apt-get -y upgrade \
    && apt-get -y install make build-essential git curl wget vim

ENV GO_VERSION 1.11.2

RUN wget -q https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && rm go${GO_VERSION}.linux-amd64.tar.gz

ENV GOROOT /usr/local/go
ENV GOPATH /gopath
ENV PATH "${PATH}:${GOPATH}/bin:${GOROOT}/bin"

RUN mkdir -p ${GOPATH}/src/github.com/id9383 \
    && cd ${GOPATH}/src/github.com/id9383 \
    && git clone https://github.com/id9383/webserver

RUN cd ${GOPATH}/src/github.com/id9383/webserver && make

