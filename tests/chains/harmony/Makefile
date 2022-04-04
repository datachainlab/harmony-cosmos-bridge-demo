include ../../docker.mk

HARMONY_TAG?=v4.3.1
BLS_TAG?=v0.0.6
NETWORK_ID?=2

# To check shard endpoints and restart Harmony network if necessary
RETRY ?= 5

SCRIPT_DIR=./scripts
CONTRACT_DIR=../../../contract
OUTPUT_ADDRESS_DIR=./output/addresses

.PHONY: docker-image
docker-image:
	make up-scaffold-with-check
	make deploy-contract
	make save-contract-address
	make docker-commit
	make down-scaffold

.PHONY: deploy-contract
deploy-contract:
	make wait-for-launch ATTEMPT=30 CONTAINER=harmony-scaffold
	cd ${CONTRACT_DIR} && npm run migrate

.PHONY: up-scaffold-with-check
# This starts harmony local network with verifying that
# each explorer node exists in the shard specified in config.
# If the verification fails, it restarts the network.
up-scaffold-with-check:
	@for i in $(shell seq ${RETRY}); do \
		make up-scaffold ; \
		make wait-for-launch ATTEMPT=30 CONTAINER=harmony-scaffold ; \
		./scripts/check_shard.sh && break ; \
		make down-scaffold ; \
	done

.PHONY: up-scaffold
up-scaffold:
	${DOCKER_COMPOSE} up --build -d

.PHONY: down-scaffold
down-scaffold:
	${DOCKER_COMPOSE} down --volumes --remove-orphans

.PHONY: save-contract-address
save-contract-address:
	mkdir -p ${OUTPUT_ADDRESS_DIR}
	${SCRIPT_DIR}/saveContractAddresses.sh ${NETWORK_ID} ${CONTRACT_DIR} ${OUTPUT_ADDRESS_DIR}

.PHONY: docker-commit
docker-commit:
	${SCRIPT_DIR}/commitImage.sh  ${DOCKER_REPO} ${DOCKER_TAG} harmony-chain harmony-scaffold ${OUTPUT_ADDRESS_DIR}
