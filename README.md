## Compass(a tiller frontend)

[![Build Status](http://reaper.qiniu.io/api/badges/weiwei04/compass/status.svg)](http://reaper.qiniu.io/weiwei04/compass)

#### Compass

TODO

#### Install Fusion Plugin

make bootstrap
make install-fusion

#### Compass REST API

install swagger tools

```bash
brew install go-swagger --with-goswagger
```

see rest api def

```bash
cd pkg/api/services/compass
goswagger serve --flavor=swagger compass.swagger.json
```
