PKG_LIST := $(shell go list ./... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: run test build deploy race coverage coverhtml

test: ## Run unittests
	@go test -short ${PKG_LIST}

run :
	go run main.go

docker-build:
	docker build -t onboard-service .

docker-delete-container:
	docker rm onboard-service

docker-run:
	docker run --name onboard-service -p 8080:8080 -p 8081:8081 onboard-service

docker-stop:
	docker stop onboard-service

proto:
	protoc -I proto/ proto/messages.proto --go_out=plugins=grpc:proto

mock-message:
	@mockgen -source=./internal/messages/repository/repository.go -destination=./internal/mock/mock_messages.go -package=mock