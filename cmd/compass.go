package main

import (
	"flag"
	"net"
	"os"
	"syscall"

	"github.com/weiwei04/compass/pkg/compass"
	"github.com/weiwei04/compass/pkg/proto/compass/services"
	"google.golang.org/grpc"

	"os/signal"
)

var (
	grpcAddr      = flag.String("listen", ":8910", "address:port to listen on")
	enableTracing = flag.Bool("trace", false, "enable rpc tracing")
	tillerAddr    = flag.String("tiller", "127.0.0.1:44134", "tiller")
)

func main() {
	flag.Parse()
	runGRPCServer()
}

func runGRPCServer() {
	var opts = []grpc.ServerOption{}

	grpcSrv := grpc.NewServer(opts...)
	lstn, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		return
	}

	compassSrv := compass.NewCompassServer()
	if err := compassSrv.Start(); err != nil {
		return
	}
	defer compassSrv.Shutdown()

	services.RegisterCompassServiceServer(grpcSrv, compassSrv)
	srvErrCh := make(chan error)
	go func() {
		srvErrCh <- grpcSrv.Serve(lstn)
	}()

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGTERM)
	signal.Notify(stopCh, syscall.SIGINT)
	go func() {
		<-stopCh
		grpcSrv.GracefulStop()
	}()

	if err := <-srvErrCh; err != nil {
	} else {
	}
}
