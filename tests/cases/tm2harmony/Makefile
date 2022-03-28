include ../../docker.mk

RETRY ?= 5

.PHONY: network
network:
	TAG=${DOCKER_TAG} $(DOCKER_COMPOSE) \
		-f ../docker-compose-test.yaml \
		up -d \
		tendermint-chain

	@for i in $(shell seq ${RETRY}); do \
		make network-harmony & \
		../../scripts/check_shard.sh && break ; \
		make network-harmony-down ; \
	done
	make wait-for-launch ATTEMPT=30 CONTAINER=harmony-chain

.PHONY: network-harmony
network-harmony:
	@TAG=${DOCKER_TAG} $(DOCKER_COMPOSE) \
		-f ../docker-compose-test.yaml \
		up -d \
		harmony-chain

.PHONY: network-harmony-down
network-harmony-down:
	@TAG=${DOCKER_TAG} $(DOCKER_COMPOSE) \
		-f ../docker-compose-test.yaml \
		rm -fsv \
		harmony-chain

.PHONY: test
test:
	./scripts/fixture
	./scripts/init-rly
	./scripts/handshake
	#./scripts/test-tx

.PHONY: network-down
network-down:
	TAG=${DOCKER_TAG} $(DOCKER_COMPOSE) \
		-f ../docker-compose-test.yaml \
		down -v --remove-orphans
