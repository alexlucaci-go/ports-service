FROM alpine

COPY ports-service /app/ports-service
COPY ports.json /app/ports.json
COPY start.sh /app/start.sh
RUN chmod +x /app/ports-service

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

EXPOSE 8000

WORKDIR /app

CMD sh start.sh