.PHONY: run test proto

run :
	go run main.go

docker-build:
	docker build -t onboard-service .

docker-delete-container:
	docker rm onboard-service

docker-run:
	docker run --name onboard-service -p 8080:8080 onboard-service

proto:
	protoc -I proto/ proto/messages.proto --go_out=plugins=grpc:proto