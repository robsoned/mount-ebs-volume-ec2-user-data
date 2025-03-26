FROM ghcr.io/robsoned/docker-devbox:latest

WORKDIR /app

COPY devbox.* ./

RUN devbox shellenv --init-hook >> ~/.profile