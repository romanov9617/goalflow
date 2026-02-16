package di

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"pglease/internal/adapter/config"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewHTTPServer(lc fx.Lifecycle, cfg *config.Config, logger *zap.Logger) *http.Server {
	addr := fmt.Sprintf("%s:%v", cfg.API.Host, cfg.API.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: http.FileServer(http.Dir(".")),
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			logger.Info("Starting HTTP server at", zap.String("addr", srv.Addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}
