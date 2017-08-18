package compass

import (
	"net/http"
	"time"

	tillerapi "k8s.io/helm/pkg/proto/hapi/services"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	context "golang.org/x/net/context"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	compassapi "github.com/weiwei04/compass/pkg/api/services/compass"
	"google.golang.org/grpc"
)

type GRPCHandler struct {
	*grpc.Server
}

func NewGRPCHandler(ctx context.Context, config Config) (GRPCHandler, error) {
	zapLogger, _ := zap.NewProduction()
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

	compassSrv := NewCompassServer(ctx, config)
	compassapi.RegisterCompassServiceServer(grpcSrv, compassSrv)
	tillerapi.RegisterReleaseServiceServer(grpcSrv, compassSrv)

	return GRPCHandler{grpcSrv}, nil
}

func (s GRPCHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.ServeHTTP(w, r)
}
