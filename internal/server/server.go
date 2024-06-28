package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	getuserinfo "dl-user-profile/internal/server/handlers/get_user_info"
	selectitems "dl-user-profile/internal/server/handlers/select_items"
)

type Server struct {
	httpServer *http.Server
}

func New() *Server {
	mux := http.NewServeMux()

	httpServer := http.Server{
		Handler:      mux,
		Addr:         fmt.Sprintf("%s:%d", "0.0.0.0", 8080),
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	mux.HandleFunc("/api/1/user-profile/user-info", getuserinfo.New().GetUserInfo)
	mux.HandleFunc("/api/1/user-profile/user-info/select-items", selectitems.New().SelectItems)

	server := Server{
		httpServer: &httpServer,
	}

	return &server
}

func (s *Server) Start() error {
	idleConnClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := s.httpServer.Shutdown(ctx); err != nil {
			return
		}

		close(idleConnClosed)
	}()

	if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("failed to listen and serve: %v", err)
	}

	<-idleConnClosed

	return nil
}
