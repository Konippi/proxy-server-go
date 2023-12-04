FROM redis:latest

COPY ./redis.conf /usr/local/etc/redis/redis.conf
CMD [ "redis-server", "/usr/local/etc/redis/redis.conf"]

RUN mkdir -p /var/log/redis && \
    chmod 777 /var/log/redis

EXPOSE 6379