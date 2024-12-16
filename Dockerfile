FROM --platform=linux/amd64 debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

COPY .env /usr/bin/

WORKDIR /usr/bin

ADD notely /usr/bin/notely

CMD ["notely"]
