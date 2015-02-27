.PHONY: webserver
PKG=github.com/olivere/simple-webserver

default: webserver

webserver:
	go build -o webserver ./webserver.go

webserve: webserver
	./webserver -addr=":8000" -cert=./etc/star.go.crt -key=./etc/star.go.key
