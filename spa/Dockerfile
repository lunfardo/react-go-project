FROM alpine:3.8

RUN apk update

RUN apk add --update nodejs nodejs-npm
RUN apk add bash
RUN adduser -D reactwebserver

#USER reactwebserver

WORKDIR /usr/app/src

SHELL ["/bin/bash", "-c"]
RUN npm install -g npm
COPY ./src/ /usr/app/src/
WORKDIR /usr/app/src
#just in case
RUN rm -rf node_modules
RUN npm install

ENTRYPOINT npm start

