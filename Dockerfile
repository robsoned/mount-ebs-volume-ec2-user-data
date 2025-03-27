FROM ghcr.io/robsoned/docker-devbox:latest

WORKDIR /app

COPY devbox.* ./

RUN  echo 'eval "$(devbox shellenv)"' >> ~/.profile