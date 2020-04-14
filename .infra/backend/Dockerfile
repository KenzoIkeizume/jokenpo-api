FROM golang

ENV TZ=UTC

RUN apt-get update && apt-get -y install netcat && apt-get clean

COPY . /app

WORKDIR /app

RUN go build -o dist

ENTRYPOINT ["/bin/bash", "./.infra/wait-for-mysql.sh", "./dist"]
