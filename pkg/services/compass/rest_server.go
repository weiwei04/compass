package compass

import (
	"net/http"
	"time"

	"github.com/golang/glog"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	compassapi "github.com/weiwei04/compass/pkg/api/services/compass"
	//"go.uber.org/zap"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

type restServer struct {
	config Config
}

func newRESTServer(ctx context.Context, config Config) restServer {
	return restServer{config: config}
}

func (s restServer) Serve(addr string) error {
	opts := []grpc.DialOption{
		grpc.WithTimeout(10 * time.Second),
		grpc.WithInsecure(),
	}

	var (
		conn *grpc.ClientConn
		err  error
	)
	for i := 0; i < 3; i++ {
		conn, err = grpc.Dial(s.config.RPCAddr, opts...)
		if err == nil {
			break
		}
	}
	if err != nil {
		return err
	}

	mux := runtime.NewServeMux()
	glog.Infof("Created Gateway ServeMux")
	err = compassapi.RegisterCompassServiceHandler(context.Background(), mux, conn)
	if err != nil {
		glog.Errorf("RegisterCompassServiceHandlerFromEndpoint:%s failed", s.config.RESTAddr)
		return err
	}
	glog.Infof("Gateway will serve at /")
	return http.ListenAndServe(s.config.RESTAddr, mux)
}
