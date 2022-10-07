# Rosetta ICON

`rosetta-icon` provides an implementation of the [Rosetta specification](https://github.com/coinbase/rosetta-specifications) for ICON in Go.

## How to Use

### System Requirements

```
CPU: minimum 4core, recommend 8core+
RAM: minimum 16GB, recommend 32GB+
DISK: minimum SSD 1.5TB, recommend SSD 2TB+ (for Mainnet sync)
```

### Development

 * Build a executable binary
   ```
   make build
   ```

 * Build a Docker image from the remote github repo
   ```
   make build-docker
   ```

 * Build a Docker image from the local context
   ```
   make build-local
   ```

### Run Docker

Assuming you have already built a Docker image called `rosetta-icon:latest` via above `build-docker` or `build-local` commands,
you can run these commands from the command line.

 * `Mainnet:Online`
   ```
   make run-mainnet-online
   ```

 * `Lisbon:Online` (Testnet)
   ```
   make run-lisbon-online
   ```

### Configuring the Environment Variables

#### Required Arguments

* **`MODE`**: determines if Rosetta can make outbound connections.
  - **Type:** `String`
  - **Options:** `ONLINE`, `OFFLINE`
  - **Default:** None


* **`NETWORK`**: the ICON network to launch or communicate with.
  - **Type:** `String`
  - **Options:** `MAINNET`, `LISBON` or `LOCALNET`
  - **Default:** None


* **`PORT`**: the port to use for Rosetta.
  - **Type:** `Integer`
  - **Options:** `8080`, any compatible port number
  - **Default:** None


#### Optional Arguments

* **`ENDPOINT`**: the endpoint for a running ICON node.
  - **Type:** `String`
  - **Options:** a node endpoint
  - **Default:** `http://localhost:9080`


### Testing with `rosetta-cli`

To validate `rosetta-icon`, [install `rosetta-cli`](https://github.com/coinbase/rosetta-cli#install)
and run one of the following commands:

* `NETWORK=lisbon make check_data`

  This command validates that the Data API implementation is correct using the ICON Lisbon (testnet) node.
  It also ensures that the implementation does not miss any balance-changing operations.

* `NETWORK=lisbon make check_construction`

  This command validates the Construction API implementation.
  It also verifies transaction construction, signing, and submissions to the Lisbon (testnet) network.

* `NETWORK=mainnet make check_data`

  This command validates that the Data API implementation is correct using the ICON Mainnet node.
  It also ensures that the implementation does not miss any balance-changing operations.


## License

This project is available under the [Apache License, Version 2.0](LICENSE).
