package main

import (
	"context"
	"fmt"
	"net"

	"github.com/marvin5064/stock-analytics/lib/logger"
	stock "github.com/marvin5064/stock-analytics/protobuf/stock"
	"google.golang.org/grpc/reflection"
)

func RunGrpcServer(srv *Server, host string, port int) {
	address := fmt.Sprintf("%v:%v", host, port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		logger.Fatal("failed to listen", err)
	}
	stock.RegisterStockManagerServer(srv.grpcSrv, srv)
	// Register reflection service on gRPC server.
	reflection.Register(srv.grpcSrv)
	logger.Info("GRPC serving at", address)
	if err := srv.grpcSrv.Serve(lis); err != nil {
		logger.Fatal("failed to serve", err)
	}
}

func (s *Server) GetStockPrices(
	ctx context.Context,
	request *stock.StockPriceRequest) (*stock.StockPriceResponse, error) {
	logger.Info("processing GetStockPrices", request)
	return s.stockFetchManager.GetData(request)
}
