FROM golang:1.19

WORKDIR /app

COPY go.* ./

RUN go get -d -v ./..

RUN go install -v ./..

EXPOSE 8080
