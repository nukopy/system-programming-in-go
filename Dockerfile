FROM ubuntu:20.04

ENV DEBIAN_FRONTEND noninteractive

ENV LANG en_US.UTF-8
ENV PATH $PATH:/usr/local/go/bin

# Setting envs related to GOENV
ENV GOOS linux
ENV GOVERSION 1.18beta1
ENV GOARCH amd64
ENV CGO_ENABLED 0
ENV GO111MODULE on

# Install Requisites
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        build-essential \
        git \
        vim \
        wget \
        # ref: https://qiita.com/shimacpyon/items/1af6d1ed69f6ad54c73c
        ca-certificates \
        strace \
	python3.9 && \
    apt-get clean

# Install Go
WORKDIR /opt
RUN wget https://go.dev/dl/go${GOVERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GOVERSION}.linux-amd64.tar.gz && \
    rm go${GOVERSION}.linux-amd64.tar.gz

# Copy src
WORKDIR /src
COPY . .

CMD ["/bin/bash"]
