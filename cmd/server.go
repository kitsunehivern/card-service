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
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var repo repo2.IRepository
		var err error
		switch args[0] {
		case "memory":
			repo, err = repo2.NewMemRepo(context.Background())
		case "postgres":
			dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
				cfg.Psql.Username,
				cfg.Psql.Password,
				cfg.Psql.Host,
				cfg.Psql.Port,
				cfg.Psql.Database,
			)
			repo, err = repo2.NewPsqlRepo(context.Background(), dsn)
		default:
			panic(fmt.Sprintf("unexpected %v database", args[0]))
		}

		if err != nil {
			return err
		}

		cardSvc := service.NewCardService(repo)
		router := server.NewRouter(cardSvc)
		srv := &http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.Http.Host, cfg.Http.Port),
			Handler: router,
		}

		errCh := make(chan error, 1)
		go func() {
			log.Printf("HTTP server listening on %s:%d", cfg.Http.Host, cfg.Http.Port)
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
