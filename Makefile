include .env

setup: build start

deploy: update_grpc build start

build:	
	@echo "Build app"; \
	cd ${SERVER_PATH}; \
	go build

start:
	@echo "Run app"; \
	cd ${SERVER_PATH}; \
	./server
	
update_grpc:
	@echo "Update grpc"; \
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/gophkeeper.proto

run_client:
	@echo "Build client"; \
	cd ${CLIENT_PATH}; \
	go build
	@echo "Run client"; \
	cd ${CLIENT_PATH}; \
	./client