From  golang:latest

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-ping

ENV PORT 9000

CMD ["/docker-gs-ping"]
