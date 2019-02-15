FROM alpine:3.8

RUN apk update

RUN apk add --update nodejs nodejs-npm

RUN adduser -D reactwebserver

#USER reactwebserver

WORKDIR /usr/app/src

RUN npm install
RUN npm cache clean
RUN npm install -g npm
RUN npm install -g create-react-app

ENTRYPOINT npm start

