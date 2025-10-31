# Frontend Dockerfile
FROM node:22-bullseye-slim

WORKDIR /app

COPY . .

RUN npm install --include=dev

EXPOSE 5173

CMD ["npm", "run", "dev"]
