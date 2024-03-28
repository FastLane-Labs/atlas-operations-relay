FROM golang:latest

WORKDIR /app

RUN apt-get update

COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . ./

ARG VERSION=""
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.Version=${VERSION}" -o /opsrelay

EXPOSE 8080

ENTRYPOINT ["/opsrelay"]