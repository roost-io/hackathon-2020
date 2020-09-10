# pull official base image
FROM node:13.12.0-alpine
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio" \
    event="Roost Hackathon 2020"

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY package.json ./
COPY package-lock.json ./
RUN npm install --silent
RUN npm install react-scripts@3.4.3 -g --silent

# add app
COPY . ./

EXPOSE 3000
# start app
ENTRYPOINT ["npm", "start"]
