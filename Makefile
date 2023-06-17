IMAGE ?= streamnative/pulsar-beat-output
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
GIT_HASH := $(shell git rev-parse HEAD)

.PHONY: build-image
build-image: ## Build docker image.  Default: IMAGE=streamnative/pulsar-beat-output
	docker build -t $(IMAGE):$(GIT_HASH) .

.PHONY: push
push: build-image ## Build and push docker image. Default: IMAGE=streamnative/pulsar-beat-output
	docker push $(IMAGE):$(GIT_HASH)


.PHONY: build-m1-image
build-m1-image: ## Build docker image.  Default: IMAGE=streamnative/pulsar-beat-output
	docker buildx build --platform linux/amd64 -t $(IMAGE):$(GIT_HASH) .

.PHONY: push-m1
push-m1: build-m1-image ## Build and push docker image. Default: IMAGE=streamnative/pulsar-beat-output
	docker push $(IMAGE):$(GIT_HASH)