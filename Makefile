LOCAL_BIN:=$(CURDIR)/bin
PROTOC_GEN_GO:=$(LOCAL_BIN)/protoc-gen-go
PROTOC_GEN_GO_GRPC:=$(LOCAL_BIN)/protoc-gen-go-grpc

API_PATH:=api/recognizer/v1
PKG_PATH:=pkg/api/recognizer/v1

.PHONY: gen
gen:
	mkdir -p $(PKG_PATH)
	protoc --proto_path $(API_PATH) \
	--go_out=$(PKG_PATH) --go_opt=paths=source_relative \
	--go-grpc_out=$(PKG_PATH) --go-grpc_opt=paths=source_relative \
	$(API_PATH)/voicekit.proto

.PHONY: clean
clean:
	rm -rf pkg/api/recognizer/v1/*.go

.PHONY: run-server
run-server:
	go run cmd/server/server.go

.PHONY: run-client
run-client:
	go run cmd/client/client.go