FROM node:20-slim AS builder

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

WORKDIR /app

COPY package.json pnpm-* .
COPY packages ./packages
RUN --mount=type=cache,id=pnpm,target=/pnpm/store \
    pnpm install --frozen-lockfile

RUN mkdir dist && pnpm run build

FROM nginx:1.25-alpine3.18-slim

COPY deployment/static/nginx.conf /etc/nginx/nginx.conf
COPY deployment/static/templates /etc/nginx/templates
COPY --from=builder /app/dist /app
