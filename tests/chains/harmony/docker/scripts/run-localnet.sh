#!/usr/bin/env bash
set -e

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
harmony_dir="$(go env GOPATH)/src/github.com/harmony-one/harmony"
# cf. https://github.com/harmony-one/harmony-test/blob/master/localnet/configs/localnet_deploy.config
localnet_config=$(realpath "$DIR/../configs/localnet_deploy.config")

function stop() {
  if [ "$KEEP" == "true" ]; then
    tail -f /dev/null
  fi
  kill_localnet
}

function kill_localnet() {
  pushd "$(pwd)"
  cd "$harmony_dir" && bash ./test/kill_node.sh
  popd
}

function setup() {
  if [ ! -d "$harmony_dir" ]; then
    echo "Test setup FAILED: Missing harmony directory at $harmony_dir"
    exit 1
  fi
  if [ ! -f "$localnet_config" ]; then
    echo "Test setup FAILED: Missing localnet deploy config at $localnet_config"
    exit 1
  fi
  kill_localnet
  error=0  # reset error/exit code
}

function build_and_start_localnet() {
  local localnet_log="$harmony_dir/localnet_deploy.log"
#  rm -rf "$harmony_dir/tmp_log*"
#  rm -rf "$harmony_dir/.dht*"
#  rm -f "$localnet_log"
#  rm -f "$harmony_dir/*.rlp"
  pushd "$(pwd)"
  cd "$harmony_dir"
  bash ./test/deploy.sh -e -B -D 60000 "$localnet_config" 2>&1 | tee "$localnet_log"
  popd
}

function wait_until_block() {
  local tries=200
  if [ -n "$2" ]; then
    tries=$2
  fi

  i=0
  local valid=false
  until $valid; do
    result=$(curl --silent --location --request POST "$API_ENDPOINT" \
      --header "Content-Type: application/json" \
      --data '{"jsonrpc":"2.0","method":"hmy_blockNumber","params":[],"id":1}' | jq '.result')
    if [[ "$result" < "\"$1\"" ]]; then
      echo "Waiting for block \"$1\", current block is $result..."
      if ((i > tries)); then
        echo "TRIES REACHED. EXITING."
        exit 1
      fi
      sleep 2
      i=$((i + 1))
    else
      valid=true
    fi
  done
}

function run() {
  echo -e "\n=== STARTING Harmony One localnet for Ganache ===\n"
  build_and_start_localnet || exit 1 &
  sleep 10
  wait_for_localnet_boot 100 # Timeout at ~300 seconds

  wait_until_block "0x6"
  echo "Initialization of localnet completed"
}

function wait_for_localnet_boot() {
  timeout=70
  if [ -n "$1" ]; then
    timeout=$1
  fi
  i=0
  until curl --silent --location --request POST "$API_ENDPOINT" \
    --header "Content-Type: application/json" \
    --data '{"jsonrpc":"2.0","method":"net_version","params":[],"id":1}' \
    --output /dev/null; do
    echo "Trying to connect to localnet..."
    if ((i > timeout)); then
      echo "TIMEOUT REACHED"
      exit 1
    fi
    sleep 3
    i=$((i + 1))
  done

  valid=false
  until $valid; do
    result=$(curl --silent --location --request POST "$API_ENDPOINT" \
      --header "Content-Type: application/json" \
      --data '{"jsonrpc":"2.0","method":"hmy_blockNumber","params":[],"id":1}' | jq '.result')
    if [ "$result" = "\"0x0\"" ]; then
      echo "Waiting for localnet to boot..."
      if ((i > timeout)); then
        echo "TIMEOUT REACHED"
        exit 1
      fi
      sleep 2
      i=$((i + 1))
    else
      valid=true
    fi
  done

  sleep 5  # Give some slack to ensure localnet is booted...
  echo "Localnet booted."
}

function wait_for_epoch() {
  wait_for_localnet_boot "$2"
  cur_epoch=0
  echo "Waiting for epoch $1..."
  until ((cur_epoch >= "$1")); do
    cur_epoch=$(curl --silent --location --request POST "$API_ENDPOINT" \
    --header "Content-Type: application/json" \
    --data '{"jsonrpc":"2.0","method":"hmyv2_latestHeader","params":[],"id":1}' | jq .result.epoch)
    if ((i > timeout)); then
      echo "TIMEOUT REACHED"
      exit 1
    fi
    sleep 3
    i=$((i + 1))
  done
}

trap stop SIGINT SIGTERM EXIT

KEEP=true

setup
run

exit "$error"
