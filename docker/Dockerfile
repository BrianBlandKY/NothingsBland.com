# GOAL
# To build a development environment managed entirely within docker.
# Golang + Docker as an alternative to python (virtualenv) and vagrant.

# Concerns: host IDE is going to eat shit when searching for dependencies locally.

FROM debian:latest

# Web Ports
EXPOSE 8080/tcp

ENV GOPATH /go:/
ENV GOBIN /go/bin
ENV PATH $PATH:/go/bin
ENV PATH=$PATH:/usr/local/go/bin

# updates
RUN apt-get update -qq && \
    apt-get install -y -qq wget git nginx

# install golang
WORKDIR /tmp

# Go 1.9.2
RUN wget https://redirector.gvt1.com/edgedl/go/go1.9.2.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.9.2.linux-amd64.tar.gz

# Configure nginx
# Remove default site
RUN rm -r /etc/nginx/sites-enabled/default
RUN rm -r /etc/nginx/sites-available/default
RUN rm -r /var/www/html
# Setup NothingsBland.com config
COPY nothingsbland.com.conf /etc/nginx/sites-available/
RUN ln -s /etc/nginx/sites-available/nothingsbland.com.conf /etc/nginx/sites-enabled/nothingsbland.com.conf
RUN service nginx restart

COPY nothingsbland/ /apps
WORKDIR /apps

COPY go-wrapper /usr/local/bin
RUN chmod 775 /usr/local/bin/go-wrapper

RUN go-wrapper download
RUN go-wrapper install

# simple entry point for continuous server
# ENTRYPOINT ["tail", "-f", "/dev/null"]
ENTRYPOINT ["go-wrapper", "run"]

# ENTRYPOINT [ "watcher" ]
# CMD bash -c "watcher > /logs/docker.log 2>&1"

# Build
# docker build -t nothingsbland-img -f Dockerfile .

# Run
# docker run -tdi -p 8080:8080 --name nothingsbland nothingsbland-img

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