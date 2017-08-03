#!/bin/bash
protoc --go_out=plugins=grpc:. *.proto -I. -I../../../../vendor/k8s.io/helm/_proto
