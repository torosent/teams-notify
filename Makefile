
DOCKER_REGISTRY ?= "torosent"

.PHONY: build
build:
	mkdir -p bin/
	go build -o bin/teams-notify ./main.go

.PHONY: docker-build
docker-build:
	mkdir -p rootfs
	GOOS=linux GOARCH=amd64 go build -o rootfs/teams-notify ./main.go
	docker build -t $(DOCKER_REGISTRY)/teams-notify:latest .

.PHONY: docker-push
docker-push:
	docker push $(DOCKER_REGISTRY)/teams-notify

.PHONY: bootstrap
bootstrap:


