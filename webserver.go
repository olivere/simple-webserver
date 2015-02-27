package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	addr = flag.String("addr", ":8000", "HTTP address, e.g. :8000")
	cert = flag.String("cert", "", "Certificate file")
	key  = flag.String("key", "", "Key file")
)

func main() {
	flag.Parse()

	s, err := NewServer(
		SetAddr(*addr),
		SetTLS(*cert, *key),
		SetLogger(log.New(os.Stdout, "", log.LstdFlags)))
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

type Server struct {
	addr string

	tls  bool
	cert string
	key  string

	logger *log.Logger
}

type ServerOption func(*Server) error

func NewServer(opts ...ServerOption) (*Server, error) {
	s := &Server{
		addr: ":8000",
	}
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}
	return s, nil
}

func SetAddr(addr string) ServerOption {
	return func(s *Server) error {
		s.addr = addr
		return nil
	}
}

func SetTLS(cert, key string) ServerOption {
	return func(s *Server) error {
		s.cert = cert
		s.key = key
		s.tls = s.cert != "" && s.key != ""
		return nil
	}
}

func SetLogger(logger *log.Logger) ServerOption {
	return func(s *Server) error {
		s.logger = logger
		return nil
	}
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	if s.tls {
		cert, err := tls.LoadX509KeyPair(s.cert, s.key)
		if err != nil {
			return err
		}
		cfg := &tls.Config{
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: true, // WARNING: only use for self-signed certs
		}
		lis = tls.NewListener(lis, cfg)
		srv := http.Server{
			Addr:    s.addr,
			Handler: s,
		}
		s.printf("Listen on %s via TLS", s.addr)
		return srv.Serve(lis)
	} else {
		s.printf("Listen on %s", s.addr)
		srv := http.Server{
			Addr:    s.addr,
			Handler: s,
		}
		return srv.Serve(lis)
	}
}

func (s *Server) printf(format string, args ...interface{}) {
	if s.logger != nil {
		s.logger.Printf(format, args...)
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello\n")
}
