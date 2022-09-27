#!/bin/bash

if [ x"$NETWORK" == x ]; then
  echo "Error: NETWORK must be populated"
  exit -1
fi

SEED_mainnet="seed-ctz.solidwallet.io:7100"
SEED_lisbon="seed-lisbon.solidwallet.io:7100"

start_chain() {
  while true; do
    RES=$(goloop system info 2>&1)
    if [ "$?" == "0" ]; then
      break
    fi
    sleep 1
  done
  echo $RES

  CID=$(goloop chain ls | jq .[0].cid)
  if [ "$CID" == "null" ]; then
    # set some runtime configs
    goloop system config rpcRosetta true
    goloop system config eeInstances 6

    # join chain
    GENESIS=/app/goloop-conf/${NETWORK,,}/icon_genesis.zip
    SEED=$(eval "echo \${SEED_${NETWORK,,}}")
    CID=$(goloop chain join \
        --platform icon \
        --channel icon_dex \
        --genesis ${GENESIS} \
        --tx_timeout 60000 \
        --node_cache small \
        --normal_tx_pool 10000 \
        --db_type rocksdb \
        --role 0 \
        --seed ${SEED})
  fi
  CID=$(eval echo $CID)
  echo "CID=$CID"
  goloop chain start $CID
}

echo "[*] START rosetta-icon server"
rosetta-icon run &

# start chain in backgound
start_chain &

echo "[*] START goloop server"
exec goloop server start
