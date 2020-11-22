DOCKER_REPOSITORY=maiaaraujo5
DOCKER_TAG=latest
APP_NAME_TAG=control-transaction:$(DOCKER_TAG)


build:
	go build -mod vendor -o ./dist/main cmd/main.go

test:
	go test ./internal...
integration_tests:
	go test -tags=integration ./it/...

run: build
	go run cmd/main.go

docker-build: build
	cp ./build/docker/Dockerfile .
	docker build -t $(DOCKER_REPOSITORY)/$(APP_NAME_TAG) .
	rm -rf Dockerfile

docker-run: docker-build
	docker run --net=host $(DOCKER_REPOSITORY)/$(APP_NAME_TAG)

docker-compose-run-dependencies:
	cp ./build/docker/docker-compose.yaml .
	cp ./build/docker/init.sql .
	docker-compose up -d
	rm -rf init.sql
	rm -rf docker-compose.yaml

docker-run-with-providers-dependencies: docker-build docker-compose-run-dependencies docker-run

