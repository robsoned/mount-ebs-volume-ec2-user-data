FROM ghcr.io/robsoned/docker-devbox:latest

WORKDIR /app

COPY devbox.* ./

RUN  devbox install && \
     devbox shellenv >> ~/.profile && \
     echo 'eval "$(devbox shellenv)"' >> ~/.bashrc