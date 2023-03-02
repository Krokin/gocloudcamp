package main

import (
	"net"

	p "main/Part_Two/proto"
	s "main/Part_Two/server"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/xlab/closer"
	"google.golang.org/grpc"
)

func main() {
	srv := s.NewServer()
	//сохранение данных при получении сигнала завершения от os
	closer.Bind(func() {
		err := srv.P.SavePlaylist()
		if err != nil {
			srv.Logger.Log.Error(err)
		}
		srv.Logger.Log.Printf("Close server")
	})
	serv := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpc_logrus.UnaryServerInterceptor(srv.Logger.Log, grpc_logrus.WithLevels(srv.Logger.LogConfig)),
	))
	p.RegisterPlaylistServer(serv, srv)
	srv.Logger.Log.Printf("Created server on port :%d", 8080)
	l, err := net.Listen("tcp", "8080")
	if err != nil {
		srv.Logger.Log.Error(err)
		closer.Fatalln()
	}
	err = serv.Serve(l)
	if err != nil {
		srv.Logger.Log.Error(err)
		closer.Fatalln()
	}
}
