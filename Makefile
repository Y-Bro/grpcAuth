protoc -I$@/${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. $@/${PROTO_DIR}/*.proto


proto:
	protoc -I ./cmd/proto --go_opt=module=github.com/Y-Bro/go-grpcAuth --go_out=. --go-grpc_opt=module=github.com/Y-Bro/go-grpcAuth --go-grpc_out=. ./cmd/proto/*.proto


help: ## Show this help
	@${HELP_CMD}

auth:
	go build -o ./bin/auth/server ./cmd/server