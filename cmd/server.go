package cmd

import (
	cardpb "card-service/gen/proto"
	repo2 "card-service/internal/repo"
	"card-service/internal/server"
	"card-service/internal/service"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
		httpAddr := fmt.Sprintf("%s:%d", cfg.Http.Host, cfg.Http.Port)
		grpcAddr := fmt.Sprintf("%s:%d", cfg.Grpc.Host, cfg.Grpc.Port)

		go func() {
			log.Printf("gRPC server listening on %v\n", grpcAddr)
			if err := server.NewRouter(cardSvc, grpcAddr); err != nil {
				log.Fatalf("failed to set up gRPC server: %v\n", err)
			}
		}()

		dialCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(
			dialCtx,
			grpcAddr, // <— use the gRPC host:port here
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
		)
		if err != nil {
			log.Fatalf("failed to dial gRPC at %s: %v", grpcAddr, err)
		}
		defer conn.Close()

		mux := runtime.NewServeMux()
		err = cardpb.RegisterCardServiceHandler(context.Background(), mux, conn)
		if err != nil {
			log.Fatalf("failed to register gateway: %v", err)
		}

		httpServer := &http.Server{
			Addr:    httpAddr,
			Handler: mux,
		}

		errCh := make(chan error, 1)
		go func() {
			log.Printf("HTTP server listening on %v\n", httpAddr)
			if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				errCh <- err
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-quit:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			_ = httpServer.Shutdown(ctx)
			return nil
		case err := <-errCh:
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
