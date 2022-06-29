include ../../docker.mk

CHAINID?=ibc0
CHAINDIR?=${HOME}/config/.simd/data
GRPCPORT?=9090

.PHONY: build
build:
	. ./scripts/setup_bls_build_flags.sh && \
	go build -o ./build/simd ./simapp/simd

# TODO remove build args such as PAT, USERNAME
.PHONY: docker-image
docker-image:
	$(DOCKER_BUILD) \
		--build-arg CHAINID=ibc0 \
		--build-arg BLS_TAG=v0.0.6 \
		--tag="$(DOCKER_REPO)tendermint-chain:$(DOCKER_TAG)" .
