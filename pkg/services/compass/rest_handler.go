package compass

import (
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	compassapi "github.com/weiwei04/compass/pkg/api/services/compass"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

type RESTHandler struct {
	*http.ServeMux
}

func NewRESTHandler(ctx context.Context, config Config) (RESTHandler, error) {
	var (
		err     error
		handler RESTHandler
	)
	opts := []grpc.DialOption{
		grpc.WithTimeout(5 * time.Second),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	}

	gwmux := runtime.NewServeMux()
	err = compassapi.RegisterCompassServiceHandlerFromEndpoint(ctx, gwmux, "SOME_ENDPOINT", opts)
	if err != nil {
		return handler, err
	}
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	return RESTHandler{mux}, err
}

func (h RESTHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.ServeHTTP(w, r)
}
