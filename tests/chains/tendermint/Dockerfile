FROM golang:1.16-buster AS scaffold

ARG BLS_TAG
ARG LIB_OUTPUT=/root/build/lib/

SHELL ["/bin/bash", "-c"]
RUN apt-get update \
    && apt-get install -y \
        unzip libgmp-dev libssl-dev curl git jq make gcc g++ bash sudo python3 python3-pip

WORKDIR $GOPATH/src/github.com/harmony-one
RUN git clone https://github.com/harmony-one/bls.git -b ${BLS_TAG} --depth 1 \
    && git clone https://github.com/harmony-one/mcl.git --depth 1 \
    && make -C mcl -j8 \
	&& make -C bls BLS_SWAP_G=1 -j8 \
    && mkdir -p ${LIB_OUTPUT} \
    && cp mcl/lib/* ${LIB_OUTPUT} \
    && cp bls/lib/* ${LIB_OUTPUT}

WORKDIR /root
COPY ./ ./

# TODO remove
# for private repo
ARG USERNAME
ARG PAT
ARG GOPRIVATE=github.com/datachainlab/ibc-harmony-client
RUN echo "machine github.com login ${USERNAME} password ${PAT}" > /root/.netrc \
    && chmod 600 /root/.netrc

RUN source ./scripts/setup_bls_build_flags.sh \
    && go build -mod=readonly -o ./build/simd ./simapp/simd

FROM debian:buster AS initializer
RUN apt-get -y update && apt-get -y install \
    jq libgmp-dev libssl-dev
WORKDIR /root

COPY --from=scaffold /root/build/simd /usr/bin/simd
COPY --from=scaffold /root/build/lib/* /root/lib/
COPY ./scripts ./scripts

ARG CHAINID
ARG CHAINDIR=./data
ARG RPCPORT=26657
ARG P2PPORT=26656
ARG PROFPORT=6060
ARG GRPCPORT=9090
ENV LD_LIBRARY_PATH=/root/lib/

RUN ./scripts/tm-chain simd $CHAINID $CHAINDIR $RPCPORT $P2PPORT $PROFPORT $GRPCPORT

FROM debian:buster
RUN apt-get -y update && apt-get -y install \
    libgmp-dev libssl-dev \
    && apt-get -y clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /root

ARG CHAINID
ARG CHAINDIR=./data
ARG RPCPORT=26657
ARG P2PPORT=26656
ARG PROFPORT=6060
ARG GRPCPORT=9090

ENV CHAINID=$CHAINID
ENV CHAINDIR=$CHAINDIR
ENV GRPCPORT=$GRPCPORT
ENV LD_LIBRARY_PATH=/root/lib/

COPY --from=scaffold /root/build/simd /usr/bin/simd
COPY --from=initializer /root/$CHAINDIR /root/$CHAINDIR
COPY --from=initializer /root/lib /root/lib
COPY ./scripts/entrypoint.sh /root/entrypoint.sh

EXPOSE $RPCPORT $P2PPORT $PROFPORT $GRPCPORT

ENTRYPOINT ["./entrypoint.sh"]
