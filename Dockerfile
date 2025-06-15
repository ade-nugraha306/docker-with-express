FROM node:23-alpine

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

# Default: production run
CMD ["node", "server.js"]
