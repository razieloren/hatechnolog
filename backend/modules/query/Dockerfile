FROM ubuntu:22.04

ADD ./out/linux/query /query
ADD ./config.yaml /config.yaml

RUN DEBIAN_FRONTEND=noninteractive \
    echo 'APT::Install-Suggests "0";' >> /etc/apt/apt.conf.d/00-docker \
    && echo 'APT::Install-Recommends "0";' >> /etc/apt/apt.conf.d/00-docker \
    && apt-get update \
    && apt-get install -y ca-certificates \
    && rm -rf /var/lib/apt/lists/* \
    && update-ca-certificates \
    && useradd -ms /bin/bash user \
    && chown user:user /query /config.yaml \
    && chmod +x /query

USER user
ENV PATH="${PATH}:/"

CMD ["query"]