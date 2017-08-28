package compass

import (
	"net"

	"github.com/golang/glog"
	//"time"

	//"github.com/grpc-ecosystem/go-grpc-middleware"
	//"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	//"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	//"go.uber.org/zap"
	//"go.uber.org/zap/zapcore"
	context "golang.org/x/net/context"

	//"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	compassapi "github.com/weiwei04/compass/pkg/api/services/compass"
	"google.golang.org/grpc"
)

type rpcServer struct {
	config  Config
	grpcSrv *grpc.Server
}

func newRPCServer(ctx context.Context, config Config, srv compassapi.CompassServiceServer) rpcServer {
	//grpc_zap.ReplaceGrpcLogger(config.Logger)

	//zapOpts := []grpc_zap.Option{
	//	grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
	//		return zap.Int64("grpc.time_ns", duration.Nanoseconds())
	//	}),
	//}

	grpcSrv := grpc.NewServer(
	//grpc_middleware.WithUnaryServerChain(
	//	grpc_ctxtags.UnaryServerInterceptor(),
	//	grpc_zap.UnaryServerInterceptor(config.Logger, zapOpts...),
	//),
	//grpc_middleware.WithStreamServerChain(
	//	grpc_ctxtags.StreamServerInterceptor(),
	//	grpc_zap.StreamServerInterceptor(config.Logger, zapOpts...),
	//),
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
