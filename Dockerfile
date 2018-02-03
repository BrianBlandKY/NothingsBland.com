FROM centos:centos7

# Web Port
EXPOSE 8080/tcp 

VOLUME ["/apps/nothingsbland"]

ENV GOPATH /go
ENV GOBIN /go/bin
ENV PATH $PATH:/go/bin
ENV PATH=$PATH:/usr/local/go/bin
ENV container=docker

# updates and installs
RUN yum install -y epel-release && \
    yum update -y && \
    yum install -y wget git nano && \
    yum clean all

# Go 1.9.2
WORKDIR /tmp
RUN wget https://redirector.gvt1.com/edgedl/go/go1.9.2.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.9.2.linux-amd64.tar.gz

# Make volume log directory
RUN mkdir -P "/apps/nothingsbland/logs"

# Copy over source
# TODO: Update to pull from github!
RUN mkdir -p /go/src/NothingsBland.com
COPY . /go/src/NothingsBland.com
WORKDIR /go/src/NothingsBland.com

# Install Go Dep
RUN go get -u github.com/golang/dep/cmd/dep

# Install Application Dependencies
RUN dep ensure

# Build and Run minifier
WORKDIR /go/src/NothingsBland.com/minifier
RUN go build
RUN ./minifier

# Build Web
WORKDIR /go/src/NothingsBland.com/web
RUN go build

ENTRYPOINT [ "./web", "--config", "app.prod.yaml" ]

# Build
# docker build -t nothingsbland-img -f Dockerfile .

# Run
# docker run -tdi -p 8081:8081 --name nothingsbland nothingsbland-img

# Exec
# docker exec -ti nothingsbland sh -c "go run main.go"
# docker exec -ti nothingsbland sh -c "go-wrapper run"

# Terminal 
# docker exec -ti nothingsbland /bin/bash

# Stop Containers
# docker stop $(docker ps -aq)

# Remove Containers
# docker rm $(docker ps -aq)

# Remove Images
# docker rmi $(docker images -q)

# Helpful Commands
# docker stop $(docker ps -aq) && docker rm $(docker ps -aq) && docker rmi $(docker images -q)
# docker exec -t -i 50f331760ba7 /bin/bash
# docker start -a -i `docker ps -q -l`