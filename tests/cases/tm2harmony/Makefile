include ../../docker.mk

RETRY ?= 5

.PHONY: network
network:
	make network-tendermint
	make network-harmony-with-check
	make wait-for-launch ATTEMPT=30 CONTAINER=harmony-chain

.PHONY: network-tendermint
network-tendermint:
	TAG=${DOCKER_TAG} $(DOCKER_COMPOSE) \
		-f ../docker-compose-test.yaml \
		up -d \
		tendermint-chain

.PHONY: network-harmony-with-check
# This starts harmony localnet with verifying that
# each explorer node exists in the shard specified in config.
# If the verification fails, it restarts the localnet.
network-harmony-with-check:
	@for i in $(shell seq ${RETRY}); do \
		make network-harmony & \
		../../scripts/check_shard.sh && break ; \
		make network-harmony-down ; \
	done

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
