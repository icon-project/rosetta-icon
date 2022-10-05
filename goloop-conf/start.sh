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

    NETKEY=${NETWORK,,}
    CONF_DIR=/app/goloop-conf/${NETKEY}
    if [ $NETKEY == "mainnet" ]; then
      cp ${CONF_DIR}/icon_governance.zip /app/ || exit -1
    fi

    # join chain
    GENESIS=${CONF_DIR}/icon_genesis.zip
    SEED=$(eval "echo \${SEED_${NETKEY}}")
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
  if [ $NETKEY == "mainnet" ]; then
    if [ ! -f ${GOLOOP_NODE_DIR}/1/block_v1_proof.bin ]; then
      cp ${CONF_DIR}/block_v1_proof.bin ${GOLOOP_NODE_DIR}/1/ || exit -1
    fi
  fi
  goloop chain start $CID
}

echo "[*] START rosetta-icon server"
rosetta-icon run &

# start chain in backgound
start_chain &

echo "[*] START goloop server"
exec goloop server start
