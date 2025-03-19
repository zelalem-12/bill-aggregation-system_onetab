FROM kong:2.8.5-alpine

ARG KONG_PROXY_PORT
ARG KONG_ADMIN_PORT

WORKDIR /app

COPY config/kong.conf /etc/kong/kong.conf

EXPOSE ${KONG_PROXY_PORT} ${KONG_ADMIN_PORT}

CMD ["kong", "start", "-c", "/etc/kong/kong.conf"]
