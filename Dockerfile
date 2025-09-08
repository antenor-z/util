FROM golang:latest

RUN apt-get update && apt-get install -y \
    whois \
    dnsutils \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o util

EXPOSE 5200

CMD ["./util"]