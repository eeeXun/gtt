FROM ubuntu:latest
ARG VERSION

RUN apt update && apt install -y wget libasound2-dev
RUN wget https://github.com/eeeXun/gtt/releases/download/${VERSION}/gtt-linux-amd64.tar.gz
RUN tar -xvf gtt-linux-amd64.tar.gz -C /usr/bin/
ENV TERM xterm-256color

ENTRYPOINT ["gtt"]
