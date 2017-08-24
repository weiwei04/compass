package compass

import (
	"go.uber.org/zap"
)

type Config struct {
	TillerAddr   string
	RegistryAddr string
	RPCAddr      string
	RESTAddr     string
	Mock         bool
	Logger       *zap.Logger
}
