FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go build -o loadtester .

ENTRYPOINT ["./loadtester", "loadtest"]
