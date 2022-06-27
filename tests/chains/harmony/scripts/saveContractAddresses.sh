#!/usr/bin/env bash

set -eu

if [ $# -ne 3 ];
    then echo "illegal number of parameters"
    exit 1
fi

NETWORK_ID=$1
CONTRACT_DIR=$2
OUTPUT_DIR=$3
jq -r ".networks | .[\"${NETWORK_ID}\"].address" < ${CONTRACT_DIR}/build/contracts/IBCHost.json > ${OUTPUT_DIR}/IBCHost
jq -r ".networks | .[\"${NETWORK_ID}\"].address" < ${CONTRACT_DIR}/build/contracts/IBCHandler.json > ${OUTPUT_DIR}/IBCHandler
jq -r ".networks | .[\"${NETWORK_ID}\"].address" < ${CONTRACT_DIR}/build/contracts/SimpleToken.json > ${OUTPUT_DIR}/SimpleToken
jq -r ".networks | .[\"${NETWORK_ID}\"].address" < ${CONTRACT_DIR}/build/contracts/ICS20Bank.json > ${OUTPUT_DIR}/ICS20Bank
jq -r ".networks | .[\"${NETWORK_ID}\"].address" < ${CONTRACT_DIR}/build/contracts/ICS20TransferBank.json > ${OUTPUT_DIR}/ICS20TransferBank
