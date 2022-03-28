include ../../docker.mk

.PHONY: build
build:
	go build -o ./build/simd ./simapp/simd

.PHONY: docker-image
docker-image:
	$(DOCKER_BUILD) \
		--build-arg CHAINID=ibc0 \
		--tag $(DOCKER_REPO)tendermint-chain:$(DOCKER_TAG) .
