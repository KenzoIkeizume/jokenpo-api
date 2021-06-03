FROM golang

ENV TZ=UTC

RUN apt-get update && \
    apt-get -y install netcat && \
    apt-get clean \
    apt install -y git \
    make openssh-client

COPY . /app

WORKDIR /app

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

ENTRYPOINT ["/bin/bash", "./.infra/wait-for-mysql.sh", "air"]
