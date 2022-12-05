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

