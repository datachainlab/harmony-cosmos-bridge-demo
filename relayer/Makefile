BLS_TAG?=v0.0.6
BLS_GIT_OPTS?=-b ${BLS_TAG}
MCL_GIT_OPTS?=
GIT_OPTS?=--depth 1
LIB_OUTPUT?=./build/

gopath=$(shell go env GOPATH)
harmony_root=${gopath}/src/github.com/harmony-one

.PHONY: harmony-libs
harmony-libs: clone-harmony-libs build-harmony-libs

# this is needed to build relayer
.PHONY: clone-harmony-lib
clone-harmony-libs:
	mkdir -p ${harmony_root}
	git clone ${GIT_OPTS} ${MCL_GIT_OPTS} -- https://github.com/harmony-one/mcl.git ${harmony_root}/mcl && \
	git clone ${GIT_OPTS} ${BLS_GIT_OPTS} -- https://github.com/harmony-one/bls.git ${harmony_root}/bls

.PHONY: build-harmony-lib
build-harmony-libs:
	make -C ${harmony_root}/mcl -j8 && \
	make -C ${harmony_root}/bls BLS_SWAP_G=1 -j8 && \
	mkdir -p ${LIB_OUTPUT} && \
	cp ${harmony_root}/mcl/lib/* ${LIB_OUTPUT} && \
	cp ${harmony_root}/bls/lib/* ${LIB_OUTPUT}

.PHONY: build
build:
	. ./scripts/setup_bls_build_flags.sh && \
	go build -o ./build/uly .

.PHONY: proto-gen
proto-gen:
	@echo "Generating Protobuf files"
	docker run -v $(CURDIR):/workspace --workdir /workspace tendermintdev/sdk-proto-gen sh ./scripts/protocgen.sh
