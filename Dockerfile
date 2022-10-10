#build stage
FROM golang:alpine 
COPY ./ /go/src/shorturllinkapi
#RUN apk add --no-cache git
WORKDIR /go/src/shorturllinkapi
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go  ./
RUN go build -o /shorturllinkapi
EXPOSE 3000
CMD ["/shorturllinkapi"]
