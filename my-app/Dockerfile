FROM node:20-alpine

# layer union fs
WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

# 브라우저 자동 열기 비활성화 (Docker/WSL2 환경에서 필요)
ENV BROWSER=none

# localhost : 3000
CMD ["npm", "start"]