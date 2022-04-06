.PHONY: compile-contracts deploy-contracts-shard0 deploy-contracts-shard1
compile-contracts:
	make -C contract compile

deploy-contracts-shard0:
	make -C contract deploy-shard0

deploy-contracts-shard1:
	make -C contract deploy-shard1

.PHONY: build-relayer
build-relayer:
	make -C relayer harmony-libs && \
	make -C relayer build
