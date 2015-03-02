# Simple web server in Go

This is just a simple web server used as reference for teaching.
It illustrates the following things:

* Connect with either http or https.
* How to use TLS configuration.
* How to use http.Serve with net.Listener.
* Functional parameters in setting up the server.

## Usage

Run `./make-certs.sh` to generate a wildcard cert for `*.go`.

Run `make webserver` to build the webserver binary.

Run `./webserver -h` for help.

Run `./webserver -addr=":8000" -cert=./etc/star.go.crt -key=./etc/star.go.key` to run the webserver with TLS. Open e.g. `https://rpc.go:8000` and skip the browser warning.

Run `./webserver -addr=":8000"` to use standard HTTP. Open e.g. `http://127.0.0.1:8000` in your browser.

## Docker

You can use this repository with Docker.

1. `docker build -t simple-webserver .`
2. `docker run --publish 8000:8000 --name simple-webserver --rm
   simple-webserver`
3. `open https://<docker-ip>:8000`

Run `docker stop simple-webserver` to stop the docker container.

# LICENSE

MIT
