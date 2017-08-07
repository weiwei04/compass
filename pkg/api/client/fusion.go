package client

type Fusion interface {
	Connect() error
	Shutdown()
	Release() Release
	Registry() Registry
}

func NewFusion(registryAddr string, releaseAddr string) Fusion {
	return &fusion{
		release:  NewCompassReleaseClient(releaseAddr),
		registry: NewHelmRegistryClient(registryAddr),
	}
}

type fusion struct {
	release  Release
	registry Registry
}

func (f *fusion) Connect() error {
	var err error
	defer func() {
		if err == nil {
			return
		}
		f.release.Shutdown()
		f.registry.Shutdown()
	}()
	err = f.release.Connect()
	if err != nil {
		return err
	}
	err = f.registry.Connect()
	return err
}

func (f *fusion) Shutdown() {
	f.release.Shutdown()
	f.registry.Shutdown()
}

func (f *fusion) Release() Release {
	return f.release
}

func (f *fusion) Registry() Registry {
	return f.registry
}
