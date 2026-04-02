FROM node:22-alpine AS builder
WORKDIR /src

COPY . .
RUN npm ci

RUN npm run build

FROM node:22-alpine
WORKDIR /src

COPY --from=builder /src/.output ./.output

ENV HOST=0.0.0.0
ENV PORT=3000

EXPOSE 3000
CMD ["node", ".output/server/index.mjs"]