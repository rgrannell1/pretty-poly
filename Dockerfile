FROM ubuntu
MAINTAINER Ryan Grannell

RUN apt-get update
RUN apt-get install golang-go git --assume-yes
RUN mkdir /usr/local/app
RUN mkdir /usr/local/app/src

ADD src /usr/local/app/src
ADD Makefile /usr/local/app/Makefile

WORKDIR /usr/local/app

ENV GOPATH /usr/local/app
RUN go get github.com/gonum/matrix
RUN go get github.com/gonum/floats
RUN go get github.com/franela/goblin
RUN go get github.com/docopt/docopt-go

CMD [ "/usr/bin/make", "test" ]