FROM golang

ADD . /go/src/github.com/olivere/simple-webserver
WORKDIR /go/src/github.com/olivere/simple-webserver
RUN apt-get install openssl
RUN ./make-certs.sh
RUN go get github.com/bradfitz/http2
RUN go build -o webserver ./webserver.go
ENTRYPOINT ./webserver -addr=":8000" -cert=./etc/star.go.crt -key=./etc/star.go.key

EXPOSE 8000

