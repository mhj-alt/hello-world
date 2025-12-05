FROM node:20-alpine

# layer union fs
WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

CMD ["npm", "start"]