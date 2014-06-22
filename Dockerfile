#DOCKER-VERSION 1.0.0

FROM node:0.10.28
MAINTAINER berwyn.codeweaver@gmail.com

ADD . /src
WORKDIR /src

RUN npm install
RUN npm install -g grunt-cli
RUN grunt build
EXPOSE 8080

CMD ["node", "app.js"]