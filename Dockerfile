FROM golang:1.18.8-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./
RUN go mod tidy

## executable of our Go program
RUN go build -o main .
## Our start command which kicks off
## our newly created binary executable

EXPOSE 3000

CMD ["/app/main"]