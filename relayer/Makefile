.PHONY: build
build:
	. ./scripts/setup_bls_build_flags.sh && \
	go build -o ./build/uly .

.PHONY: proto-gen
proto-gen:
	@echo "Generating Protobuf files"
	docker run -v $(CURDIR):/workspace --workdir /workspace tendermintdev/sdk-proto-gen sh ./scripts/protocgen.sh
