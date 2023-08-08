FROM ubuntu:22.04

SHELL ["/bin/bash", "-c"]
WORKDIR /app

COPY package*.json ./
ADD ./node_modules ./node_modules
ADD ./.next ./.next
ADD ./.env.production ./.env.production

RUN DEBIAN_FRONTEND=noninteractive \
    echo 'APT::Install-Suggests "0";' >> /etc/apt/apt.conf.d/00-docker \
    && echo 'APT::Install-Recommends "0";' >> /etc/apt/apt.conf.d/00-docker \
    && apt-get update \
    && apt-get install -y ca-certificates wget \
    && rm -rf /var/lib/apt/lists/* \
    && update-ca-certificates \
    && useradd -ms /bin/bash user \
    && chown -R user:user /app

USER user
ENV NVM_DIR="/home/user/.nvm"
RUN wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.4/install.sh | bash \
    && . $NVM_DIR/nvm.sh && nvm install node

CMD . $NVM_DIR/nvm.sh && npm start
