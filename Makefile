NOW = $(shell date +%s)

default: gen-go
	go build .
	
gen-go:
	mkdir -p generated/go
	protoc -I/usr/local/include -I. --go_out=plugins=grpc,paths=source_relative:generated/go proto/*.proto

install: default
	bash -c 'eval $$(minikube docker-env) && docker build -t ffs:latest .'

test: 
	go test ./...

deps:
	brew install protobuf go kubernetes-helm minikube

cluster:
	minikube start --cpus 4 --memory 8096 --disk-size=40g --kubernetes-version=v1.15.5 --docker-opt="default-ulimit=nofile=102400:102400"
