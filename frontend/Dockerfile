FROM oven/bun:canary-alpine AS builder

WORKDIR /app
COPY package.json bun.lock ./
RUN bun install
COPY . .
RUN bun run build


FROM oven/bun:canary-alpine

WORKDIR /app

COPY --from=builder /app/build ./build
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json .

EXPOSE 3000
CMD ["bun", "build/index.js"]