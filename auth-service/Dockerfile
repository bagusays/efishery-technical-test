FROM node:14.18-alpine3.15 AS builder

WORKDIR /app

COPY . .

RUN npm install && npm run build


FROM node:12.17.0-alpine

WORKDIR /app

COPY package.json ./
RUN npm install --only=production

COPY --from=builder /app/dist .
COPY --from=builder /app/config /app/config

RUN npm install pm2 -g

EXPOSE 8080

CMD ["pm2-runtime","index.js"]