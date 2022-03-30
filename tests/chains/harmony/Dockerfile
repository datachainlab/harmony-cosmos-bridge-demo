FROM golang:1.16

ENV API_ENDPOINT=
# ex: v4.3.1
ARG HARMONY_TAG
ARG BLS_TAG

WORKDIR $GOPATH/src/github.com/harmony-one

SHELL ["/bin/bash", "-c"]

RUN apt update && apt upgrade -y && apt update -y && apt install unzip \
libgmp-dev libssl-dev curl git jq make gcc g++ bash sudo python3 python3-pip -y

RUN git clone https://github.com/harmony-one/harmony.git -b ${HARMONY_TAG} --depth 1 \
    && git clone https://github.com/harmony-one/bls.git -b ${BLS_TAG} --depth 1 \
    && git clone https://github.com/harmony-one/mcl.git --depth 1

# TODO
# Currently, TendermintLightClient exceeds MaxCodeSize
RUN cd harmony && \
    sed -i 's/24576/49152/' internal/params/protocol_params.go

# Build to fetch all dependencies for faster test builds
RUN cd harmony && go mod tidy && go get github.com/pborman/uuid && go get github.com/rjeczalik/notify \
    && go get github.com/cespare/cp && go get github.com/libp2p/go-libp2p-crypto && go get github.com/kr/pretty \
    && go get github.com/kr/text && go get gopkg.in/check.v1 && bash scripts/install_build_tools.sh && make \
    && make clean

# Install testing tools
RUN curl -L -o /go/bin/hmy https://harmony.one/hmycli \
    && chmod +x /go/bin/hmy \
    # Setup hmy with default accounts
    # Main funding account
    && /go/bin/hmy keys import-private-key 144109d9b1182b51233955c112f64a545bb70143539f161e936bb01f8b1e081d

# comment out until used
#WORKDIR $GOPATH/src/github.com/coinbase
#
#RUN git clone https://github.com/coinbase/rosetta-cli.git
#RUN cd rosetta-cli && make install

WORKDIR $GOPATH/src/github.com/harmony-one/harmony-test/localnet

COPY scripts scripts
COPY configs configs

RUN chmod +x $GOPATH/src/github.com/harmony-one/harmony-test/localnet/scripts/run-localnet.sh
WORKDIR $GOPATH/src/github.com/harmony-one/harmony
ENTRYPOINT ["/go/src/github.com/harmony-one/harmony-test/localnet/scripts/run-localnet.sh"]
