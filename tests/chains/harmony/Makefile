include ../../docker.mk

HARMONY_TAG?=v4.3.1
BLS_TAG?=v0.0.6

.PHONY: docker-image
docker-image:
	$(DOCKER_BUILD) \
		--build-arg HARMONY_TAG=$(HARMONY_TAG) \
		--build-arg BLS_TAG=$(BLS_TAG) \
		--tag $(DOCKER_REPO)harmony-chain:$(DOCKER_TAG) .
