<p align="center">
  <a href="https://www.rosetta-api.org">
    <img width="90%" alt="Rosetta" src="https://www.rosetta-api.org/img/rosetta_header.png">
  </a>
</p>
<h3 align="center">
   Rosetta ICON
</h3>

## Overview
`rosetta-icon` provides a reference implementation of the [Rosetta specification](https://github.com/coinbase/rosetta-specifications) for ICON in Golang.

## System Requirements
`ICON Citizen Node` has been tested on an [AWS c5.2xlarge](https://aws.amazon.com/ec2/instance-types/c5/) instance. This instance type has 4 vCPU and 16 GB of RAM.

## Run in Docker Compose
```shell script
cd docker
docker-compose up
```

## Run Local
### pre-requirements
1. Run Citizen Node:  ref) https://www.icondev.io/docs/node
2. Build rosetta-icon: make
3. Run rosetta-icon with enviroment variables (ex TestNet)
    * ENDPOINT=http://localhost:9000
    * MODE=ONLINE # (ONLINE, OFFLINE)
    * NETWORK=TESTNET # (MAINNET, TESTNET, ZICON, DEVNET)
    * PORT=8080
    
## Run Local without Citizen Node
### pre-requirements
1. Build rosetta-icon: make
2. Run rosetta-icon with enviroment variables (ex TestNet)
    * ENDPOINT=https://testwallet.icon.foundation
    * MODE=ONLINE # (ONLINE, OFFLINE)
    * NETWORK=TESTNET # (MAINNET, TESTNET, ZICON, DEVNET)
    * PORT=8080
    
## Caution
* ICON Node Required Full DB.
    * ICON Node doesn't support light client.
* Recommend Docker-compose DB Snapshot(FASTEST_START: "yes")


## Docker Compose Configration
```
version: '3'
services:
  prep-node:
    image: 'iconloop/prep-node:dev'
    container_name: "prep-node"
    restart: "on-failure"
    environment:
      LOOPCHAIN_LOG_LEVEL: "SPAM" # Log Level
      ICON_LOG_LEVEL: "DEBUG" # Log Level
      DEFAULT_PATH: "/data/loopchain" # DB, SCORE Storage Path
      LOG_OUTPUT_TYPE: "file" # Log export (file, console)
      SERVICE: "testnet" # mainnet, testnet(을지로), zicon(파고다)
      IS_AUTOGEN_CERT: "true" # Auto Generate Cert
      FASTEST_START: "yes" # download DB Snapshot
      IS_COMPRESS_LOG: "true"
      USER_DEFINED_ENV: |
        .CHANNEL_OPTION.icon_dex.crep_root_hash=0xb7cc19da5bff37a2c4954e16473ab65610a9481f8f864d7ea587c65bff82402f|configure_json
    cap_add:
      - SYS_TIME
    volumes:
      - ./data/loopchain/mainnet:/data/loopchain/
      - ./cert:/prep_peer/cert
    ports:
      - '7100:7100' # GRPC
      - '9000:9000' # RPC
  rosetta:
    image: 'jinyyo/rosetta-test'
    container_name: "rosetta"
    restart: "on-failure"
    environment:
      ENDPOINT: 'http://localhost:9000'
      MODE: 'ONLINE'
      NETWORK: 'TESTNET'
    ports:
      - '8080:8080'
```

