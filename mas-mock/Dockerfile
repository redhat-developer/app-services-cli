# Stage 1 - the build process
FROM node:10 as compile-server
WORKDIR /usr/src/app
COPY . .
RUN npm install

CMD [ "npm", "start" ]