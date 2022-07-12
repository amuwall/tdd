package web

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
	"tdd/config"
	"tdd/web/handler"
	"time"
)

type Server struct {
	Config config.WebTemplate

	instance *http.Server
	exit     *sync.WaitGroup
}

func NewServer(config config.WebTemplate) *Server {
	gin.SetMode(config.Mode)

	engine := gin.New()
	handler.Register(engine)

	return &Server{
		Config: config,
		instance: &http.Server{
			Addr:              fmt.Sprintf("%s:%d", config.Host, config.Port),
			Handler:           engine,
			ReadTimeout:       2 * time.Minute,
			ReadHeaderTimeout: 30 * time.Second,
			WriteTimeout:      2 * time.Minute,
			IdleTimeout:       1 * time.Minute,
		},
		exit: &sync.WaitGroup{},
	}
}

func (server *Server) Start() {
	go func() {
		server.exit.Add(1)
		defer server.exit.Done()

		err := server.instance.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("servr: listen and serve error %s\n", err)
			return
		}
	}()
}

func (server *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.instance.Shutdown(ctx)
	if err != nil {
		log.Printf("server: shutdown error %s\n", err)
		return
	}

	server.exit.Wait()
}
