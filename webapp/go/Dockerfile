FROM golang:1.23.3-bullseye

WORKDIR /tmp
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update
RUN apt-get -y upgrade
RUN apt-get install -y wget gcc g++ make sqlite3
RUN wget -q https://dev.mysql.com/get/mysql-apt-config_0.8.33-1_all.deb
RUN apt-get -y install ./mysql-apt-config_*_all.deb
RUN apt-get -y update
RUN apt-get -y install default-mysql-client

RUN useradd --uid=1001 --create-home isucon
USER isucon

RUN mkdir -p /home/isucon/webapp/go
WORKDIR /home/isucon/webapp/go
COPY --chown=isucon:isucon ./ /home/isucon/webapp/go/

ENV GOPATH=/home/isucon/tmp/go
ENV GOCACHE=/home/isucon/tmp/go/.cache

CMD ["go", "run", "./cmd/isuports/main.go"]
