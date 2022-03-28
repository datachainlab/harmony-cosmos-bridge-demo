include ../../docker.mk

HARMONY_TAG?=v4.3.1

.PHONY: docker-image
docker-image:
	$(DOCKER_BUILD) \
		--build-arg HARMONY_TAG=$(HARMONY_TAG) \
		--tag $(DOCKER_REPO)harmony-chain:$(DOCKER_TAG) .
