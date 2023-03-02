package main

import (
	"github.com/Krokin/gocloudcamp/Part_Two/config"
	"net"

	p "github.com/Krokin/gocloudcamp/Part_Two/proto"
	s "github.com/Krokin/gocloudcamp/Part_Two/server"

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
		} else {
			srv.Logger.Log.Printf("Playlist saved")
		}
		srv.Logger.Log.Printf("Exit server")
	})
	//загрузка конфигурации
	conf, err := config.LoadConfig()
	if err != nil {
		srv.Logger.Log.Fatal(err)
	}
	serv := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpc_logrus.UnaryServerInterceptor(srv.Logger.Log, grpc_logrus.WithLevels(srv.Logger.LogConfig)),
	))
	p.RegisterPlaylistServer(serv, srv)
	srv.Logger.Log.Printf("Created server on port %s", conf.Port)
	l, err := net.Listen("tcp", conf.Port)
	if err != nil {
		srv.Logger.Log.Fatal(err)
	}
	err = serv.Serve(l)
	if err != nil {
		srv.Logger.Log.Error(err)
		closer.Fatalln()
	}
}
