package compass

import (
	"net"

	"time"

	"github.com/golang/glog"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	context "golang.org/x/net/context"

	compassapi "github.com/weiwei04/compass/pkg/api/services/compass"
	"google.golang.org/grpc"
)

type rpcServer struct {
	config  Config
	grpcSrv *grpc.Server
}

func newRPCServer(ctx context.Context, config Config, srv compassapi.CompassServiceServer) rpcServer {
	logger, _ := zap.NewProduction()
	grpc_zap.ReplaceGrpcLogger(logger)

	zapOpts := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
	}

	grpcSrv := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger, zapOpts...),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger, zapOpts...),
		),
	)
	glog.Infof("Registered CompassService")
	compassapi.RegisterCompassServiceServer(grpcSrv, srv)

	return rpcServer{config, grpcSrv}
}

func (s rpcServer) Serve(addr string) error {
	lst, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return s.grpcSrv.Serve(lst)
}
