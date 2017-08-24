package compass

import (
	"context"

	"github.com/golang/glog"
	pb "github.com/weiwei04/compass/pkg/api/services/compass"
)

type Server struct {
	config Config
}

func NewServer(config Config) *Server {
	return &Server{config}
}

func (s *Server) Serve() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	<-stopCh
	//	cancel()
	//}()

	//var err error
	var (
		srv pb.CompassServiceServer
		err error
	)
	if s.config.Mock {
		glog.Infof("Created mock server")
		srv = newFakeCompassServer()
	} else {
		srv, err = newCompassServer(ctx, s.config)
	}
	if err != nil {
		return err
	}
	grpcSrv := newRPCServer(ctx, s.config, srv)
	go func() {
		glog.Infof("Start grpc server")
		grpcSrv.Serve(s.config.RPCAddr)
	}()

	restSrv := newRESTServer(ctx, s.config)

	glog.Infof("Start rest server")
	return restSrv.Serve(s.config.RESTAddr)
}
