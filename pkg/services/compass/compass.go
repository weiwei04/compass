package compass

import (
	"time"

	"go.uber.org/zap"

	"google.golang.org/grpc"

	compassapi "github.com/weiwei04/compass/pkg/api/services/compass"
	"github.com/weiwei04/compass/pkg/chart"
	tillerapi "k8s.io/helm/pkg/proto/hapi/services"
)

type Config struct {
	TillerAddr   string
	RegistryAddr string
}

type CompassServer struct {
	config   Config
	logger   *zap.Logger
	registry chart.Store
	conn     *grpc.ClientConn
	tiller   tillerapi.ReleaseServiceClient
}

func NewCompassServer(config Config) *CompassServer {
	return &CompassServer{
		config: config,
	}
}

var _ compassapi.CompassServiceServer = &CompassServer{}

var _ tillerapi.ReleaseServiceServer = &CompassServer{}

func (s *CompassServer) Start() error {
	var err error
	if s.logger, err = zap.NewProduction(); err != nil {
		return err
	}
	s.registry, err = chart.NewHelmRegistryStore(s.config.RegistryAddr)
	if err != nil {
		s.logger.Error("NewHelmRegistryStore failed",
			zap.Error(err))
		return err
	}

	opts := []grpc.DialOption{
		grpc.WithTimeout(5 * time.Second),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	}
	s.conn, err = grpc.Dial(s.config.TillerAddr, opts...)
	if err != nil {
		s.logger.Error("dial tiller failed",
			zap.String("tillerAddr", s.config.TillerAddr),
			zap.Error(err))
		return err
	}
	s.tiller = tillerapi.NewReleaseServiceClient(s.conn)
	return nil
}

func (s *CompassServer) Shutdown() {
	s.conn.Close()
	s.logger.Sync()
}
