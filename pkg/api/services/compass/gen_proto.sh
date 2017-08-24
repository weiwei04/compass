#!/bin/bash
protoc --go_out=plugins=grpc:. *.proto -I. -I../../../../vendor/k8s.io/helm/_proto -I../../../../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
protoc --grpc-gateway_out=logtostderr=true:. *.proto -I. -I../../../../vendor/k8s.io/helm/_proto -I../../../../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
protoc --swagger_out=logtostderr=true:. *.proto -I. -I../../../../vendor/k8s.io/helm/_proto -I../../../../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
#goswagger generate client -f compass.swagger.json
