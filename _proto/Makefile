space := $(empty) $(empty)
comma := ,
empty :=

import_path = github.com/weiwei04/pkg/proto/compass

dst 	= ../pkg/proto/compass/services
target  = go
plugins = grpc

compass_ias    = $(subst $(space),$(comma),$(addsuffix =$(import_path)/$(version_pkg),$(addprefix M,$(version_pbs))))
compass_pbs    = $(sort $(wildcard compass/services/*.proto))
compass_pkg    = services

google_deps	 = Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any

.PHONY: all
all: services

.PHONY: services
services:
	PATH=../bin:$(PATH) protoc -I../vendor/k8s.io/helm/_proto/ -Icompass/services --$(target)_out=plugins=$(plugins),$(google_deps):$(dst) $(compass_pbs)

.PHONY: clean
clean:
	@rm -rf $(dst)/* 2>/dev/null
