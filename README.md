<p align="center">
  <a href="https://www.rosetta-api.org">
    <img width="90%" alt="Rosetta" src="https://www.rosetta-api.org/img/rosetta_header.png">
  </a>
</p>
<h3 align="center">
   Rosetta ICON
</h3>

ROSETTA-ICON IS CONSIDERED [ALPHA SOFTWARE](https://en.wikipedia.org/wiki/Software_release_life_cycle#Alpha) USE AT YOUR OWN RISK.

## Overview
`rosetta-icon` provides a reference implementation of the [Rosetta specification](https://github.com/coinbase/rosetta-specifications) for ICON in Golang.

## System Requirements
`ICON Citizen Node` has been tested on an [AWS c5.2xlarge](https://aws.amazon.com/ec2/instance-types/c5/) instance. This instance type has 4 vCPU and 16 GB of RAM.

## Run Local
### pre-requirements
1. Run Citizen Node:  ref) https://icondev.io/icon-node/quickstart
2. Build rosetta-icon: make
3. Run rosetta-icon with enviroment variables (ex TestNet)
    * ENDPOINT=http://localhost:9000
    * MODE=ONLINE # (ONLINE, OFFLINE)
    * NETWORK=TESTNET # (MAINNET, TESTNET, DEVNET)
    * PORT=8080
    
## Caution
* ICON Node Required Full DB.
    * ICON Node doesn't support light client.
