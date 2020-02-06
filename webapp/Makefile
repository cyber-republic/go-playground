IMAGE_NAME        ?= cyberrepublic/go-playground-webapp
IMAGE_TAG         ?= 1.0.0
IMAGE_VERSION     ?=
CONTAINER_NAME    ?= go-playground-webapp

.PHONY: vendor static

build: clean vendor fmt
	go build -o bin/webapp *.go

run: static
	go run *.go

static:
	$(MAKE) -C static

vendor:
	go mod vendor

docker-build: vendor
	docker build -t $(IMAGE_NAME):latest .

docker-run:
	(docker container stop $(CONTAINER_NAME) || true) && (docker container rm -f $(CONTAINER_NAME) || true) && \
	docker run \
		-d \
		-p 0.0.0.0:3000:80 \
		--name $(CONTAINER_NAME) \
		-t $(IMAGE_NAME):latest

docker-push:
	docker tag $(IMAGE_NAME):latest $(IMAGE_NAME):$(IMAGE_TAG) && \
	docker push $(IMAGE_NAME):latest && \
	docker push $(IMAGE_NAME):$(IMAGE_TAG)
