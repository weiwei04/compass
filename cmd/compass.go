package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"github.com/weiwei04/compass/pkg/compass"
	"github.com/weiwei04/compass/pkg/proto/compass/services"
)

func main() {
	fmt.Println("vim-go")
}

func start() {
	var opts = []grpc.ServerOption{}

	var grpcAddr *string

	grpcSrv := grpc.NewServer(opts...)
	lstn, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		return
	}

	srvErrCh := make(chan error)

	go func() {
		compassSrv := compass.NewCompassServer()
		services.RegisterCompassServiceServer(grpcSrv, compassSrv)
		if err := grpcSrv.Serve(lstn); err != nil {
			srvErrCh <- err
		}
 	} ()

	<-srvErrCh
}
