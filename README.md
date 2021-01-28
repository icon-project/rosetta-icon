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
    * MODE=ONLINE
    * NETWORK=TESTNET
    * PORT=8080
    
## Run Local without Citizen Node
### pre-requirements
1. Build rosetta-icon: make
2. Run rosetta-icon with enviroment variables (ex TestNet)
    * ENDPOINT=https://testwallet.icon.foundation
    * MODE=ONLINE
    * NETWORK=TESTNET
    * PORT=8080
    
## Caution
* ICON Node Required Full DB.
    * ICON Node doesn't support light client.
* Recommend Docker-compose DB Snapshot(FASTEST_START: "yes")