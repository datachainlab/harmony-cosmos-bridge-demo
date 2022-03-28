FROM harmonyone/localnet-test

ENV API_ENDPOINT=

# ex: v4.3.1
ARG HARMONY_TAG

# Setup hmy with default accounts
# Main funding account
RUN hmy keys import-private-key 144109d9b1182b51233955c112f64a545bb70143539f161e936bb01f8b1e081d

# TODO: Remove the harmony one cloning when the harmonyone/localnet-test is updated.
WORKDIR $GOPATH/src/github.com/harmony-one

RUN git clone https://github.com/harmony-one/harmony.git -b $HARMONY_TAG --depth 1

# Build to fetch all dependencies for faster test builds
#RUN cd harmony && go get github.com/pborman/uuid && go get github.com/rjeczalik/notify \
#    && go get github.com/cespare/cp && go get github.com/libp2p/go-libp2p-crypto && go get github.com/kr/pretty \
#    && go get github.com/kr/text && go get gopkg.in/check.v1 && bash scripts/install_build_tools.sh && make

RUN cd harmony \
    && go mod tidy \
    && bash scripts/install_build_tools.sh \
    && make

WORKDIR $GOPATH/src/github.com/harmony-one/harmony-test/localnet

COPY scripts scripts

RUN chmod +x $GOPATH/src/github.com/harmony-one/harmony-test/localnet/scripts/run-localnet.sh
WORKDIR $GOPATH/src/github.com/harmony-one/harmony
ENTRYPOINT ["/go/src/github.com/harmony-one/harmony-test/localnet/scripts/run-localnet.sh"]
