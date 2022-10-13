package cmd

import (
	//"google.golang.org/grpc/grpclog"
	"log"
	"net"
	"spider/services"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start  service",
	Long:  `start  service`,
	Run: func(cmd *cobra.Command, args []string) {
		serveMain()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func startGRPCServer() {
	address := "127.0.0.1:12345"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	logger, _ := zap.NewProduction()
	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger),
			grpc_recovery.StreamServerInterceptor(),
		),
	)
	services.RegisterUserServiceServer(s)
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func serveMain() {
	log.Println("spider starting ...")
	log.Println("spider start init db ...")
	log.Println("spider init db finish.")
	log.Println("spider start grpc server ...")
	startGRPCServer()
}
