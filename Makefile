DOCKER_REPOSITORY=maiaaraujo5
DOCKER_TAG=latest
APP_NAME_TAG=control-transaction:$(DOCKER_TAG)


build-go:
	go mod vendor
	go build -mod vendor -o ./dist/main cmd/main.go

test:
	go test ./internal...

integration_tests: docker-compose-run-dependencies-test
	go test -tags=integration ./it/...

run: build-go
	go run cmd/main.go

docker-build: build-go
	cp ./build/docker/Dockerfile .
	docker build -t $(DOCKER_REPOSITORY)/$(APP_NAME_TAG) .
	rm -rf Dockerfile

docker-run: docker-rm-postgres-container docker-build
	docker run --net=host $(DOCKER_REPOSITORY)/$(APP_NAME_TAG)

docker-compose-run-dependencies: docker-rm-postgres-container
	cp ./build/docker/development/docker-compose.yaml .
	cp ./build/docker/init.sql .
	docker-compose up -d
	rm -rf init.sql
	rm -rf docker-compose.yaml

docker-compose-run-dependencies-test: docker-rm-postgres-container
	cp ./build/docker/test/docker-compose.yaml .
	docker-compose up -d
	rm -rf docker-compose.yaml

docker-run-with-providers-dependencies: docker-rm-postgres-container docker-build docker-compose-run-dependencies docker-run

docker-rm-postgres-container:
	-docker rm pismo-postgres --force
	-docker rm pismo-postgres-test --force
