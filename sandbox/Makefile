IMAGE_NAME        ?= cyberrepublic/go-playground-sandbox
IMAGE_TAG         ?= 2.0.0
IMAGE_VERSION     ?=
CONTAINER_NAME    ?= go-playground-sandbox

docker-build:
	docker build -f ./sandbox-basic.Dockerfile -t go-playground-sandbox-basic:latest . && \
	docker image rm -f $(IMAGE_NAME):latest && \
	docker build -f ./Dockerfile -t $(IMAGE_NAME):latest .

docker-run:
	(docker container stop $(CONTAINER_NAME) || true) && (docker container rm -f $(CONTAINER_NAME) || true) && \
	docker run \
		-d \
		-p 0.0.0.0:8080:8080 \
		--name $(CONTAINER_NAME) \
		-t $(IMAGE_NAME):latest

docker-push:
	docker tag $(IMAGE_NAME):latest $(IMAGE_NAME):$(IMAGE_TAG) && \
	docker push $(IMAGE_NAME):latest && \
	docker push $(IMAGE_NAME):$(IMAGE_TAG)