FROM alpine:latest

RUN apk update
RUN apk add bash 
RUN apk add git 

COPY *.go/ /opt/nuage-metro/

