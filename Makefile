HARMONY_GIT_OPTS?=-b v4.3.1
BLS_GIT_OPTS?=-b v0.0.6
MCL_GIT_OPTS?=
GIT_OPTS?=--depth 1
DEBUG_RETRY?=3

gopath=$(shell go env GOPATH)
harmony_root=${gopath}/src/github.com/harmony-one

.PHONY: setup-harmony
setup-harmony: clone-harmony build-harmony debug-harmony-retry

.PHONY: clone-harmony
clone-harmony:
	mkdir -p ${harmony_root}
	cd ${harmony_root} && \
	git clone https://github.com/harmony-one/mcl.git ${GIT_OPTS} ${MCL_GIT_OPTS} && \
	git clone https://github.com/harmony-one/bls.git ${GIT_OPTS} ${BLS_GIT_OPTS} && \
	git clone https://github.com/harmony-one/harmony.git ${GIT_OPTS} ${HARMONY_GIT_OPTS}

.PHONY: build-harmony
build-harmony:
	cd ${harmony_root}/harmony && \
	go mod tidy && \
	make

.PHONY: debug-harmony debug-harmony-retry debug-harmony-down
debug-harmony:
	cd ${harmony_root}/harmony && \
	make debug

# This is a workaround for when the http endpoint of shard 1 is different from the expected one.
debug-harmony-retry:
	@for i in $(shell seq ${DEBUG_RETRY}); do \
		make debug-harmony & \
		./scripts/check_shard.sh && break ; \
		make debug-harmony-down ; \
	done

debug-harmony-down:
	cd ${harmony_root}/harmony && \
	make debug-kill

.PHONY: compile-contracts deploy-contracts-shard0 deploy-contracts-shard1
compile-contracts:
	make -C contract compile

deploy-contracts-shard0:
	make -C contract deploy-shard0

deploy-contracts-shard1:
	make -C contract deploy-shard1

.PHONY: build-relayer
build-relayer:
	make -C relayer build
