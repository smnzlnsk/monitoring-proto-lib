.PHONY: generate clean

PROTO_DIR := proto
GO_OUT_DIR := gen/go
PYTHON_OUT_DIR := gen/python

generate: generate-go generate-python

generate-go:
	@mkdir -p $(GO_OUT_DIR)
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(GO_OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GO_OUT_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(shell find $(PROTO_DIR) -name '*.proto')

generate-python:
	@mkdir -p $(PYTHON_OUT_DIR)
	python3 -m grpc_tools.protoc \
		--proto_path=$(PROTO_DIR) \
		--python_out=$(PYTHON_OUT_DIR) \
		--grpc_python_out=$(PYTHON_OUT_DIR) \
		--pyi_out=$(PYTHON_OUT_DIR) \
		$(shell find $(PROTO_DIR) -name '*.proto')

clean:
	rm -rf $(GO_OUT_DIR)
	rm -rf $(PYTHON_OUT_DIR)
