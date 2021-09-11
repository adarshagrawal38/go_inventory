# syntax=docker/dockerfile:1
FROM golang
WORKDIR /code
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
RUN go mod download
RUN go mod vendor
EXPOSE 5000
COPY . .
CMD [ "make", "start-app" ]
