FROM node:15

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# Install app dependencies
# Bundle app source
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

COPY . /usr/src/app

EXPOSE 3000

CMD ["node", "index.js"]