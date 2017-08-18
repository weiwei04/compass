package compass

import (
	"time"

	context "golang.org/x/net/context"

	"go.uber.org/zap"

	"google.golang.org/grpc"

	compassapi "github.com/weiwei04/compass/pkg/api/services/compass"
	"github.com/weiwei04/compass/pkg/chart"
	tillerapi "k8s.io/helm/pkg/proto/hapi/services"
)

var _ compassapi.CompassServiceServer = &CompassServer{}

var _ tillerapi.ReleaseServiceServer = &CompassServer{}

type CompassServer struct {
	config   Config
	logger   *zap.Logger
	registry chart.Store
	conn     *grpc.ClientConn
	tiller   tillerapi.ReleaseServiceClient
}

func NewCompassServer(ctx context.Context, config Config) *CompassServer {
	s := &CompassServer{
		config: config,
	}

	var err error

	// init logger
	if s.logger, err = zap.NewProduction(); err != nil {
		return nil
	}

	// init conn to helm-registry
	if s.registry, err = chart.NewHelmRegistryStore(s.config.RegistryAddr); err != nil {
		s.logger.Error("NewHelmRegistryStore failed",
			zap.Error(err))
		return nil
	}

	// init conn to tiller service
	opts := []grpc.DialOption{
		grpc.WithTimeout(5 * time.Second),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	}
	if s.conn, err = grpc.Dial(s.config.TillerAddr, opts...); err != nil {
		s.logger.Error("dial tiller failed",
			zap.String("tillerAddr", s.config.TillerAddr),
			zap.Error(err))
		return nil
	}
	go func() {
		<-ctx.Done()
		if cerr := s.conn.Close(); cerr != nil {
			s.logger.Error("Failed to close conn to",
				zap.String("tiller", s.config.TillerAddr),
				zap.Error(cerr))
		} else {
			s.logger.Info("Closed conn to",
				zap.String("tiller", s.config.TillerAddr))
		}
		s.logger.Sync()
	}()
	s.tiller = tillerapi.NewReleaseServiceClient(s.conn)

	return s
}
