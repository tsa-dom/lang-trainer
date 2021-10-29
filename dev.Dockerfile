FROM golang:1.17.2

WORKDIR /usr/app

RUN apt update

COPY . .

EXPOSE 8080
