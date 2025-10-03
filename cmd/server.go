package cmd

import (
	repo2 "card-service/internal/repo"
	"card-service/internal/server"
	"card-service/internal/service"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start Server",
	RunE: func(cmd *cobra.Command, args []string) error {
		repo := repo2.NewMemRepo()
		cardSvc := service.NewCardService(repo)
		router := server.NewRouter(cardSvc)
		srv := &http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port),
			Handler: router,
		}

		errCh := make(chan error, 1)
		go func() {
			log.Printf("HTTP server listening on %s:%d", cfg.HTTP.Host, cfg.HTTP.Port)
			errCh <- srv.ListenAndServe()
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		select {
		case <-quit:
			log.Println("Shutting down HTTP server...")
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			return srv.Shutdown(ctx)
		case err := <-errCh:
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				return err
			}
			return nil
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
