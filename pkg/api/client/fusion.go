package client

type Config struct {
	ReleaseAddr  string
	RegistryAddr string
	Logger       Logger
}

type Fusion interface {
	Release() Release
	Registry() Registry
}

func NewFusion(config Config) Fusion {
	return &fusion{
		logger:   config.Logger,
		release:  NewReleaseClient(config.ReleaseAddr),
		registry: NewHelmRegistryClient(config.RegistryAddr),
	}
}

type fusion struct {
	logger   Logger
	release  Release
	registry Registry
}

func (f *fusion) Release() Release {
	return f.release
}

func (f *fusion) Registry() Registry {
	return f.registry
}
