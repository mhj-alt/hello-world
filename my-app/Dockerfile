FROM node:20-alpine

# layer union fs
WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .


# localhost : 3000
CMD ["npm", "start"]