package todoapi

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (srv *Server) RunServer(port string, handler http.Handler) error {
	srv.httpServer = &http.Server{
		Addr: ":" + port,
		MaxHeaderBytes: 1 << 20,
		Handler: handler,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return srv.httpServer.ListenAndServe()
}

func (srv *Server) ShutdownServer(ctx context.Context) error {
	return srv.httpServer.Shutdown(ctx)
}



