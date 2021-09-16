# syntax=docker/dockerfile:1
FROM golang:1.16

WORKDIR /code

# COPY go.mod ./
# COPY go.sum ./

# RUN go mod tidy
# RUN go mod download
# RUN go mod vendor

COPY ./main .

CMD [ "./main", "--host=0.0.0.0", "--port", "38465" ]
