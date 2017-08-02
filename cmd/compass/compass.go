package main

import (
	"flag"
	"net"
	"os"
	"syscall"

	"time"

	"os/signal"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/weiwei04/compass/pkg/services/compass"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	tiller "k8s.io/helm/pkg/proto/hapi/services"
)

var (
	grpcAddr      = flag.String("listen", ":8910", "address:port to listen on")
	enableTracing = flag.Bool("trace", false, "enable rpc tracing")
	tillerAddr    = flag.String("tiller", "127.0.0.1:44134", "tiller address, default: 127.0.0.1:44134")
	registryAddr  = flag.String("registry", "http://127.0.0.1:8900", "registry address, default: http://127.0.0.1:8900")
)

func main() {
	flag.Parse()
	runGRPCServer()
}

func runGRPCServer() {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	sugger := zapLogger.Sugar()

	grpc_zap.ReplaceGrpcLogger(zapLogger)

	zapOpts := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
	}

	grpcSrv := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zapLogger, zapOpts...),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(zapLogger, zapOpts...),
		),
	)
	lstn, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		sugger.Errorf("listen on %s failed, err %s", *grpcAddr, err)
		return
	}

	compassSrv := compass.NewCompassServer(compass.Config{
		TillerAddr:   *tillerAddr,
		RegistryAddr: *registryAddr,
	})
	if err := compassSrv.Start(); err != nil {
		sugger.Errorf("start compass server failed, err %s", err)
		return
	}
	defer compassSrv.Shutdown()

	tiller.RegisterReleaseServiceServer(grpcSrv, compassSrv)
	srvErrCh := make(chan error)
	go func() {
		sugger.Infof("compass server start to serve")
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
		sugger.Infof("compass server stoped with err %s", err)
		os.Exit(1)
	} else {
		sugger.Info("compass server stoped")
	}
}
