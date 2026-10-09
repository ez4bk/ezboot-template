override APP_NS := ezboot
override APP_NAME := $(shell basename "$(PWD)")
override KUBE_NAMESPACE := platform

override BUILD_TIME := $(shell date '+%y%m%d')
override GIT_REVISION := $(shell git rev-parse --short HEAD)
override KUBE_CONTEXT := $(shell kubectl config current-context)

ifdef DOCKER_HOST
override APP_GOPATH := /home/.remote/go
override BUF_CACHE := /home/.remote/buf
else
override APP_GOPATH := $(shell go env GOPATH)
override BUF_CACHE := $(HOME)/.cache/buf
endif

override DOCKER_REGISTRY ?= ghcr.io/ez4bk/$(APP_NS)
override RT_GOOGLE_WIRE_IMAGE := ghcr.io/ez4bk/devops/google-wire
override RT_GO_COMPILER_IMAGE := ghcr.io/ez4bk/devops/go-compiler
override RT_PROTOBUF_GENERATOR := ghcr.io/ez4bk/devops/protoc-buf

override APP_FULL_NAME := $(patsubst -%,%,$(APP_NS)-$(patsubst $(APP_NS)-%,%,$(APP_NAME)))
override APP_IMAGE_NAME := $(DOCKER_REGISTRY)/$(APP_FULL_NAME)
override APP_IMAGE_TAG := $(BUILD_TIME)-$(GIT_REVISION)

.DEFAULT_GOAL: help


##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk command is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


##@ Development

.PHONY: api
api: ## Generate api
	@docker pull $(RT_PROTOBUF_GENERATOR)
	@$(eval VOLUME_NAME=$(shell bash -c 'openssl rand -hex 6'))
	@docker volume create $(VOLUME_NAME)
	@docker container create --name $(VOLUME_NAME) -v $(VOLUME_NAME):/data busybox
	@docker cp ./api $(VOLUME_NAME):/data/api
	@docker cp ./buf.gen.yaml $(VOLUME_NAME):/data
	@docker cp ./buf.lock $(VOLUME_NAME):/data
	@docker cp ./buf.yaml $(VOLUME_NAME):/data
	docker run -ti --rm --entrypoint="/bin/bash" -v $(VOLUME_NAME):/data -v $(BUF_CACHE):/root/.cache/buf \
		$(RT_PROTOBUF_GENERATOR) -ec 'cd /data && buf dep update && buf build && buf generate'
	@docker cp $(VOLUME_NAME):/data/api .
	@docker rm $(VOLUME_NAME)
	@docker volume rm $(VOLUME_NAME)

.PHONY: wire
wire:
	@docker pull $(RT_GOOGLE_WIRE_IMAGE)
	@$(eval VOLUME_NAME=$(shell bash -c 'openssl rand -hex 6'))
	@docker volume create $(VOLUME_NAME)
	@docker container create --name $(VOLUME_NAME) -v $(VOLUME_NAME):/workspace busybox
	@docker cp . $(VOLUME_NAME):/workspace
	docker run -ti --rm -v $(VOLUME_NAME):/workspace -v $(APP_GOPATH):/go/pkg \
		$(RT_GOOGLE_WIRE_IMAGE) wire ./...
	@docker cp $(VOLUME_NAME):/workspace/cmd .
	@docker rm $(VOLUME_NAME)
	@docker volume rm $(VOLUME_NAME)

.PHONY: app
app:
	docker build -t $(APP_IMAGE_NAME):$(APP_IMAGE_TAG) .

.PHONY: push
push:
	docker push $(APP_IMAGE_NAME):$(APP_IMAGE_TAG)

.PHONY: deploy
deploy: app push
	@echo "Deploy by helm using context \033[0;31m$(KUBE_CONTEXT)\033[0m, waiting 5s to confirm"
	@sleep 5
	helm upgrade $(APP_NAME) deployments/myna \
		-n $(KUBE_NAMESPACE) \
		--set 'image.repository=$(APP_IMAGE_NAME)' \
		--set 'image.tag=$(APP_IMAGE_TAG)' \
		--install
