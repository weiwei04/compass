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
	if config.Logger == nil {
		config.Logger = &logger{}
	}
	return &fusion{
		logger:   config.Logger,
		release:  NewReleaseClient(config.ReleaseAddr, config.Logger),
		registry: NewHelmRegistryClient(config.RegistryAddr, config.Logger),
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
