.PHONY: setup generate generate-go generate-python lint

# Setup development environment
setup:
	# Install Python dependencies
	pip install -r requirements.txt
	# Install Go dependencies
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	# Check if protoc is installed, if not suggest installation
	@which protoc > /dev/null || (echo "Please install protoc: sudo apt-get install -y protobuf-compiler (Ubuntu) or brew install protobuf (MacOS)")
	# Check if buf is installed, if not suggest installation
	@which buf > /dev/null || (echo "Please install buf: https://buf.build/docs/installation")

generate: generate-go generate-python

generate-go:
	buf generate --template buf.gen.yaml

generate-python:
	buf generate --template buf.gen.yaml

lint:
	buf lint
