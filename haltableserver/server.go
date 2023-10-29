package haltableserver

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Server struct {
	*http.Server
}

func (srv *Server) ListenAndServe() error {
	sigtermMonitored := sync.WaitGroup{}
	sigtermMonitored.Add(1)
	go func() {
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGTERM, os.Interrupt)
		sigtermMonitored.Done()
		<-sigterm
		srv.Server.Shutdown(context.Background())
	}()
	sigtermMonitored.Wait()
	return srv.Server.ListenAndServe()
}
