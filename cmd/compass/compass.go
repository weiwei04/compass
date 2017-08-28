package main

import (
	"flag"
	//"fmt"

	"github.com/weiwei04/compass/pkg/services/compass"
	//"go.uber.org/zap"
)

var (
	grpcAddr = flag.String("rpcAddr", "127.0.0.1:8910", "address:port to listen on")
	restAddr = flag.String("httpAddr", ":8911", "address:port to listen on")
	//enableTracing = flag.Bool("trace", false, "enable rpc tracing")
	tillerAddr   = flag.String("tiller", "127.0.0.1:44134", "tiller address, default: :44134")
	registryAddr = flag.String("registry", "http://helm-registry:8900", "registry address, default: http://helm-registry:8900")
	mock         = flag.Bool("mock", false, "enable mock, default false")
)

var count int

func main() {
	flag.Parse()
	runServer()
}

func runServer() {
	//logger, err := zap.NewProduction()
	//if err != nil {
	//	panic(fmt.Sprintf("New Logger failed, err:%s", err))
	//}
	config := compass.Config{
		TillerAddr:   *tillerAddr,
		RegistryAddr: *registryAddr,
		RPCAddr:      *grpcAddr,
		RESTAddr:     *restAddr,
		Mock:         *mock,
		//Logger:       logger,
	}
	srv := compass.NewServer(config)
	srv.Serve()
}
