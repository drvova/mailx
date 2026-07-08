FROM oven/bun:1 AS builder

WORKDIR /build

ARG VITE_API_URL
ARG VITE_DOMAINS
ARG VITE_APP_NAME
ARG VITE_RESYNC_URL
ARG VITE_PRICING_URL

ENV VITE_API_URL=${VITE_API_URL}
ENV VITE_DOMAINS=${VITE_DOMAINS}
ENV VITE_APP_NAME=${VITE_APP_NAME}
ENV VITE_RESYNC_URL=${VITE_RESYNC_URL}
ENV VITE_PRICING_URL=${VITE_PRICING_URL}

COPY app/package.json app/bun.lock* ./
RUN bun install --frozen-lockfile

COPY app/ .
RUN bun run build

FROM nginx:alpine

RUN rm /etc/nginx/conf.d/default.conf

COPY --from=builder /build/dist /usr/share/nginx/html
COPY app/nginx.conf.template /etc/nginx/conf.d/default.conf.template
COPY app/docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh

ENV API_UPSTREAM=http://api.zeabur.internal:3000/

EXPOSE 80

CMD ["/docker-entrypoint.sh"]
