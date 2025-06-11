# Используется bin в текущей директории для установки плагинов protoc
LOCAL_BIN:=$(CURDIR)/bin

run:
	go run ./cmd/loms

build:
	go build -o ./bin/loms ./cmd/loms

.PHONY: genproto
genproto: .proto-generate
	go mod tidy

.PHONY: .proto-generate
.proto-generate: .bin-proto .vendor-proto  .order-api-generate .stock-api-generate .merge-swagger

# https://github.com/grpc-ecosystem/grpc-gateway?tab=readme-ov-file
# https://grpc.io/docs/languages/go/quickstart/
.PHONY: .bin-proto
.bin-proto:
	$(info Installing binary dependencies...)
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6 && \
    GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.26.3 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.26.3 && \
	GOBIN=$(LOCAL_BIN) go install github.com/g3co/go-swagger-merger@v0.3.0

.vendor-proto: .vendor-rm  vendor-proto/google/protobuf vendor-proto/buf/validate vendor-proto/google/api vendor-proto/protoc-gen-openapiv2/options
	go mod tidy

.PHONY: .vendor-rm
.vendor-rm:
	rm -rf vendor-proto

# Устанавливается proto описания google/protobuf
vendor-proto/google/protobuf:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/protocolbuffers/protobuf vendor-proto/protobuf &&\
	cd vendor-proto/protobuf &&\
	git sparse-checkout set --no-cone src/google/protobuf &&\
	git checkout
	mkdir -p vendor-proto/google
	mv vendor-proto/protobuf/src/google/protobuf vendor-proto/google
	rm -rf vendor-proto/protobuf

# Устанавливается proto описания buf/validate для protovalidate
vendor-proto/buf/validate:
	git clone -b main --single-branch --depth=1 --filter=tree:0 \
		https://github.com/bufbuild/protovalidate vendor-proto/tmp && \
		cd vendor-proto/tmp && \
		git sparse-checkout set --no-cone proto/protovalidate/buf/validate &&\
		git checkout
		mkdir -p vendor-proto/buf
		mv vendor-proto/tmp/proto/protovalidate/buf vendor-proto/
		rm -rf vendor-proto/tmp

# Устанавливается proto описания google/googleapis
vendor-proto/google/api:
	git clone -b master --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/googleapis/googleapis vendor-proto/googleapis && \
 	cd vendor-proto/googleapis && \
	git sparse-checkout set --no-cone google/api && \
	git checkout
	mkdir -p  vendor-proto/google
	mv vendor-proto/googleapis/google/api vendor-proto/google
	rm -rf vendor-proto/googleapis

# Устанавливается proto описания protoc-gen-openapiv2/options
vendor-proto/protoc-gen-openapiv2/options:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/grpc-ecosystem/grpc-gateway vendor-proto/grpc-ecosystem && \
 	cd vendor-proto/grpc-ecosystem && \
	git sparse-checkout set --no-cone protoc-gen-openapiv2/options && \
	git checkout
	mkdir -p vendor-proto/protoc-gen-openapiv2
	mv vendor-proto/grpc-ecosystem/protoc-gen-openapiv2/options vendor-proto/protoc-gen-openapiv2
	rm -rf vendor-proto/grpc-ecosystem

ORDER_API_PROTO_PATH:=api/order
.PHONY: .order-api-generate
.order-api-generate:
	rm -rf pkg/${ORDER_API_PROTO_PATH}
	mkdir -p pkg/${ORDER_API_PROTO_PATH}
	protoc \
	-I ${ORDER_API_PROTO_PATH} \
	-I vendor-proto \
	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
	--go_out pkg/${ORDER_API_PROTO_PATH} \
	--go_opt paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
	--go-grpc_out pkg/${ORDER_API_PROTO_PATH} \
	--go-grpc_opt paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=$(LOCAL_BIN)/protoc-gen-grpc-gateway \
	--grpc-gateway_out pkg/${ORDER_API_PROTO_PATH} \
	--grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true \
	--plugin=protoc-gen-openapiv2=$(LOCAL_BIN)/protoc-gen-openapiv2 \
    --openapiv2_out pkg/${ORDER_API_PROTO_PATH} \
    --openapiv2_opt logtostderr=true \
	${ORDER_API_PROTO_PATH}/*.proto

STOCK_API_PROTO_PATH:=api/stock
.PHONY: .stock-api-generate
.stock-api-generate:
	rm -rf pkg/${STOCK_API_PROTO_PATH}
	mkdir -p pkg/${STOCK_API_PROTO_PATH}
	protoc \
	-I ${STOCK_API_PROTO_PATH} \
	-I vendor-proto \
	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
	--go_out pkg/${STOCK_API_PROTO_PATH} \
	--go_opt paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
	--go-grpc_out pkg/${STOCK_API_PROTO_PATH} \
	--go-grpc_opt paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=$(LOCAL_BIN)/protoc-gen-grpc-gateway \
	--grpc-gateway_out pkg/${STOCK_API_PROTO_PATH} \
	--grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true \
	--plugin=protoc-gen-openapiv2=$(LOCAL_BIN)/protoc-gen-openapiv2 \
    --openapiv2_out pkg/${STOCK_API_PROTO_PATH} \
    --openapiv2_opt logtostderr=true \
	${STOCK_API_PROTO_PATH}/*.proto

.PHONY: .merge-swagger
.merge-swagger:
	rm -rf pkg/swagger
	mkdir -p pkg/swagger
	$(LOCAL_BIN)/go-swagger-merger \
	-o pkg/swagger/swagger.json \
	pkg/${ORDER_API_PROTO_PATH}/order.swagger.json \
	pkg/${STOCK_API_PROTO_PATH}/stock.swagger.json

.PHONY: generate-apis
generate-apis: .stock-api-generate .order-api-generate

.PHONY: .bin-mock
.bin-mock:
	$(info Installing mockery...)
	GOBIN=$(LOCAL_BIN) go install github.com/vektra/mockery/v3@v3.2.5

.PHONY: mocks
mocks:
	$(info Generate mocks...)
	$(LOCAL_BIN)/mockery

.PHONY: test
test:
	$(info Run tests...)
	go test -v -race




GOOSE_DRIVER=postgres
GOOSE_DBSTRING_1="host=localhost port=54321 user=postgres password=postgres dbname=loms sslmode=disable"
GOOSE_DBSTRING_2="host=localhost port=54322 user=postgres password=postgres dbname=loms sslmode=disable"
GOOSE_MIGRATION_DIR=./migrations

.PHONY: .bin-goose
.bin-goose:
	$(info Installing goose...)
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

.create-migration:
	goose -s -dir "./migrations" create create_table_stock sql

migrate-up: .bin-goose
	$(LOCAL_BIN)/goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING_1) up

migrate-down: .bin-goose
	$(LOCAL_BIN)/goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING_1) down

migrate-status: .bin-goose
	$(LOCAL_BIN)/goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING_1) status
	$(LOCAL_BIN)/goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING_2) status

.PHONY: .bin-sqlc
.bin-sqlc:
	$(info Installing goose...)
	GOBIN=$(LOCAL_BIN) go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

gensqlc: .bin-sqlc
	$(LOCAL_BIN)/sqlc generate