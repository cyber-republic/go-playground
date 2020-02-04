IMAGE_NAME        ?= cyberrepublic/go-playground
IMAGE_TAG         ?= 1.0.0
IMAGE_VERSION     ?=
CONTAINER_NAME    ?= go-playground

docker-build:
	docker build -t $(IMAGE_NAME):latest .

docker-push:
	docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(IMAGE_NAME):$(IMAGE_TAG) && \
	docker push $(IMAGE_NAME):$(IMAGE_TAG) && \
	docker push $(IMAGE_NAME):$(IMAGE_TAG)

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
