## Compass(a tiller frontend)

[![Build Status](http://drone.ke-cs.dev.qiniu.io/api/badges/weiwei04/compass/status.svg)](http://drone.ke-cs.dev.qiniu.io/weiwei04/compass)

#### Compass

TL,RD

#### Helm-Registry-Plugin


a helm plugin to push,list,inspect chart stored on [helm-registry](http://github.com/caicloud/helm-registry)

##### Install

```makefile
make build
cp -r _plugin $(HELM_HOME)/plugins/helm-registry-plugin
```

##### Usage
```shell
export HELM_REGISTRY_ADDR=http://xxxx
helm registry
```
