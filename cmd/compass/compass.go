package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/weiwei04/compass/pkg/services/compass"

	"go.uber.org/zap"
)

var (
	grpcAddr = flag.String("listen", ":8910", "address:port to listen on")
	//enableTracing = flag.Bool("trace", false, "enable rpc tracing")
	tillerAddr   = flag.String("tiller", "127.0.0.1:44134", "tiller address, default: :44134")
	registryAddr = flag.String("registry", "http://helm-registry:8900", "registry address, default: http://helm-registry:8900")
)

func main() {
	flag.Parse()
	runServer()
}

func runServer() {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	sugger := zapLogger.Sugar()

	config := compass.Config{
		TillerAddr:   *tillerAddr,
		RegistryAddr: *registryAddr,
		ListenAt:     *grpcAddr,
	}
	sugger.Infof("tiller[%s] registry[%s] listen[%s]",
		config.TillerAddr, config.RegistryAddr, config.ListenAt)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcHandler, err := compass.NewGRPCHandler(ctx, config)
	if err != nil {
		return
	}

	restHandler, err := compass.NewRESTHandler(ctx, config)
	if err != nil {
		return
	}

	conn, err := net.Listen("tcp", fmt.Sprintf("%s", config.ListenAt))
	if err != nil {
		panic(err)
	}

	srv := http.Server{
		Addr: config.ListenAt,
		Handler: func(grpcHandler compass.GRPCHandler, restHandler compass.RESTHandler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
					grpcHandler.ServeHTTP(w, r)
				} else {
					restHandler.ServeHTTP(w, r)
				}
			})
		}(grpcHandler, restHandler),
	}

	srv.Serve(conn)
}
