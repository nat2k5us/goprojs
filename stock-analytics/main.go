package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/marvin5064/stock-analytics/lib/logger"
	"github.com/marvin5064/stock-analytics/lib/stockfetch"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type Server struct {
	grpcSrv           *grpc.Server
	stockFetchManager stockfetch.Manager
	quit              chan os.Signal
}

func main() {
	srv := &Server{
		grpcSrv: grpc.NewServer(),
		quit:    make(chan os.Signal),
	}
	srv.sigtermHandler()

	defer logger.Sync()
	viper.SetConfigType("json")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("fail to load config")
	}

	apikey := viper.GetString("apikey")
	if apikey != "" {
		logger.Info("successfully loaded api key from config")
	}

	url := viper.GetString("url")
	if url != "" {
		logger.Info("successfully loaded url from config")
	}

	srv.stockFetchManager = stockfetch.New(url, apikey)

	go RunGrpcServer(srv, viper.GetString("grpc.hostname"), viper.GetInt("grpc.port"))
	<-srv.quit
	srv.grpcSrv.GracefulStop()
	logger.Info("Server Closed Gracefully!")
}

func (s *Server) sigtermHandler() {
	signal.Notify(s.quit, syscall.SIGTERM)
	signal.Notify(s.quit, syscall.SIGINT)
}
