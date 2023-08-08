FROM ubuntu:22.04

ADD ./out/linux/api /api
ADD ./config.yaml /config.yaml

RUN DEBIAN_FRONTEND=noninteractive \
    echo 'APT::Install-Suggests "0";' >> /etc/apt/apt.conf.d/00-docker \
    && echo 'APT::Install-Recommends "0";' >> /etc/apt/apt.conf.d/00-docker \
    && apt-get update \
    && apt-get install -y ca-certificates \
    && rm -rf /var/lib/apt/lists/* \
    && update-ca-certificates \
    && mkdir -p /etc/letsencrypt/archive/hatechnolog.com \
    && mkdir -p /etc/letsencrypt/live/hatechnolog.com \
    && chmod +x /api

ENV PATH="${PATH}:/"

CMD ["api"]