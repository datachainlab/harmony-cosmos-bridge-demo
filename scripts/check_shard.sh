#!/usr/bin/env bash
set -e

expected_shard0_url=${LOCAL_SHARD_0_URL}
expected_shard1_url=${LOCAL_SHARD_1_URL}
max_shard=1

function check() {
  if [ "$#" -ne 2 ]; then
    echo "Illegal number of parameters"
    exit 1
  fi

  local url=$1
  local shard=$2
  if [ $shard -gt $max_shard ]; then
    echo "invalid shard arg"
    exit 1
  fi

  local result=$(curl --silent --location --request POST "$url" \
    --header "Content-Type: application/json" \
    --data '{"jsonrpc":"2.0","method":"hmyv2_getShardingStructure","params":[],"id":1}' | jq ".result[$shard].current")
  if [ "$result" == "true" ]; then
    echo "$url is on shardID: $shard"
  else
    echo "$url is not on shardID: $shard"
    exit 1
  fi
}

function wait_for_localnet_boot() {
  local timeout=70
  if [ -n "$1" ]; then
    timeout=$1
  fi
  local i=0
  until curl --silent --location --request POST "$expected_shard0_url" \
    --header "Content-Type: application/json" \
    --data '{"jsonrpc":"2.0","method":"net_version","params":[],"id":1}' >/dev/null; do
    echo "Trying to connect to localnet..."
    if ((i > timeout)); then
      echo "TIMEOUT REACHED"
      exit 1
    fi
    sleep 3
    i=$((i + 1))
  done

  local valid=false
  until $valid; do
    local result=$(curl --silent --location --request POST "$expected_shard0_url" \
      --header "Content-Type: application/json" \
      --data '{"jsonrpc":"2.0","method":"hmy_blockNumber","params":[],"id":1}' | jq '.result')
    if [ "$result" = "\"0x0\"" ]; then
      echo "Waiting for localnet to boot..."
      if ((i > timeout)); then
        echo "TIMEOUT REACHED"
        exit 1
      fi
      sleep 3
      i=$((i + 1))
    else
      valid=true
    fi
  done
}

wait_for_localnet_boot
check $expected_shard0_url 0
check $expected_shard1_url 1

exit 0
