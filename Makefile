server:
	go run cmd/main.go

test:
	go test -v ./...

composeup:
	docker-compose up

composebuild:
	docker-compose up --build

composedown:
	docker-compose down

proto:
	rm -f /pkg/pb/*.go
	protoc --proto_path=proto --go_out=pkg/pb --go_opt=paths=source_relative \
    --go-grpc_out=pkg/pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

.PHONY: server test composeup composebuild composedown proto