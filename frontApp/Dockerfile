FROM node:21-alpine

# Create app directory
ADD ./vul-app-front /usr/src/app

WORKDIR /usr/src/app

RUN npm install && npm run build

CMD npm run preview