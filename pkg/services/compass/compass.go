package compass

import (

	//"log"
	"time"

	"google.golang.org/grpc"

	tillerapi "k8s.io/helm/pkg/proto/hapi/services"
	//"k8s.io/helm/pkg/version"
	compassapi "github.com/weiwei04/compass/pkg/api/services/compass"
)

type CompassServer struct {
	tillerAddr string
	conn       *grpc.ClientConn
	tiller     tillerapi.ReleaseServiceClient
}

func NewCompassServer(tillerAddr string) *CompassServer {
	return &CompassServer{tillerAddr: tillerAddr}
}

var _ compassapi.CompassServiceServer = &CompassServer{}

var _ tillerapi.ReleaseServiceServer = &CompassServer{}

func (s *CompassServer) Start() error {
	opts := []grpc.DialOption{
		grpc.WithTimeout(5 * time.Second),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	}
	var err error
	s.conn, err = grpc.Dial(s.tillerAddr, opts...)
	if err != nil {
		return err
	}
	s.tiller = tillerapi.NewReleaseServiceClient(s.conn)
	return nil
}

func (s *CompassServer) Shutdown() {
	s.conn.Close()
}
