package server

import (
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func RunGRPC(port int, registerServer func(server *grpc.Server)) error {
	s := NewGRPC(port)
	registerServer(s.Server)
	return s.Run()
}

func RunGRPCWithSecure(port int, secure SecureOptions, registerServer func(server *grpc.Server)) error {
	s, err := NewGRPCWithSecure(port, secure)
	if err != nil {
		return err
	}
	registerServer(s.Server)
	return s.Run()
}

type GRPC struct {
	*grpc.Server

	Addr string
}

type SecureOptions struct {
	CertFile string
	KeyFile  string
}

func NewGRPC(port int) *GRPC {
	addr := ":" + strconv.Itoa(port)

	return &GRPC{
		Server: grpc.NewServer(),
		Addr:   addr,
	}
}

func NewGRPCWithSecure(port int, secure SecureOptions) (*GRPC, error) {
	addr := ":" + strconv.Itoa(port)

	creds, err := credentials.NewServerTLSFromFile(secure.CertFile, secure.KeyFile)
	if err != nil {
		return nil, err
	}

	return &GRPC{
		Server: grpc.NewServer(grpc.Creds(creds)),
		Addr:   addr,
	}, nil
}

func (s *GRPC) Run() error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sig)

	quit := make(chan error)
	go s.start(quit)

	var err error
	select {
	case <-sig:
	case err = <-quit:
	}

	s.exit()
	return err
}

func (s *GRPC) start(quit chan error) {
	listen, err := net.Listen("tcp", s.Addr)
	if err != nil {
		quit <- err
		return
	}

	err = s.Serve(listen)
	if err != nil {
		quit <- err
		return
	}
	close(quit)
}

func (s *GRPC) exit() {
	s.GracefulStop()
}
