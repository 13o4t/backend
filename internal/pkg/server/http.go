package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func RunHTTP(port int, createRouter func(*gin.Engine)) error {
	s := NewHTTP(port)
	createRouter(s.Engine)
	return s.Run()
}

func RunHTTPS(port int, opt HTTPSOptions, createRouter func(*gin.Engine)) error {
	s := NewHTTPS(port, opt)
	createRouter(s.Engine)
	return s.Run()
}

type HTTP struct {
	http.Server
	Engine *gin.Engine

	ExitTimeout int
	UseHTTPS    bool
	HTTPS       HTTPSOptions
}

type HTTPSOptions struct {
	CertFile string
	KeyFile  string
	Secure   secure.Options
}

func NewHTTP(port int) *HTTP {
	addr := ":" + strconv.Itoa(port)

	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	return &HTTP{
		Server: http.Server{
			Addr:    addr,
			Handler: engine,
		},
		Engine:      engine,
		ExitTimeout: 5,
		UseHTTPS:    false,
	}
}

func NewHTTPS(port int, opt HTTPSOptions) *HTTP {
	addr := ":" + strconv.Itoa(port)

	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.Use(TLSMiddleware(opt.Secure))

	return &HTTP{
		Server: http.Server{
			Addr:    addr,
			Handler: engine,
		},
		Engine:      engine,
		ExitTimeout: 5,
		UseHTTPS:    true,
		HTTPS:       opt,
	}
}

func (s *HTTP) Run() error {
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

func (s *HTTP) start(quit chan error) {
	var err error

	if s.UseHTTPS {
		err = s.ListenAndServeTLS(s.HTTPS.CertFile, s.HTTPS.KeyFile)
	} else {
		err = s.ListenAndServe()
	}

	if err != nil && err != http.ErrServerClosed {
		quit <- err
		return
	}
	close(quit)
}

func (s *HTTP) exit() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.ExitTimeout)*time.Second)
	defer cancel()
	_ = s.Shutdown(ctx)
}

func TLSMiddleware(opt secure.Options) gin.HandlerFunc {
	return func(c *gin.Context) {
		m := secure.New(opt)
		if err := m.Process(c.Writer, c.Request); err != nil {
			return
		}
		c.Next()
	}
}
