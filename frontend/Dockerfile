FROM node:22-alpine AS builder
WORKDIR /src

ARG NUXT_PUBLIC_API_HOST=''
ARG NUXT_PUBLIC_BOT_NAME=''

ENV NUXT_PUBLIC_API_HOST=${NUXT_PUBLIC_API_HOST}
ENV NUXT_PUBLIC_BOT_NAME=${NUXT_PUBLIC_BOT_NAME}

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