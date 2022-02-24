package main

import (
	"context"

	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/zknow/my-arch/config"
	"github.com/zknow/my-arch/internal/route"
	_ "github.com/zknow/my-arch/pkg/logger"
	"github.com/zknow/my-arch/pkg/shutdown"
	"github.com/zknow/my-arch/pkg/utility"
)

func main() {
	cfg := config.GetConfig()
	port := cfg.Server.Port
	srv := &http.Server{
		Addr:    port,
		Handler: route.NewRouter(),
	}
	go func() {
		log.Println("Start listen", utility.ResolveHostIpV4(), port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	shutdown.NewHook().Close(
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			srv.Shutdown(ctx)
		},
	)
}
