FROM golang

ENV TZ=UTC

COPY . /app

WORKDIR /app

RUN go get

CMD ["go", "run", "main.go"]

