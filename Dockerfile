FROM golang:latest

WORKDIR /app

RUN apt-get update

COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /opsrelay

EXPOSE 8080

ENTRYPOINT ["/opsrelay"]