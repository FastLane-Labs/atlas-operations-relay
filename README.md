# Atlas Operations Relay

The Atlas operations relay is an infrastructure layer facilitating communication between parties involved in the Atlas ecosystem.

## Specifications

For APIs specifications and human readable description of the relay's tasks/duties, please see the [specifications](https://github.com/FastLane-Labs/atlas-operations-relay/tree/main/specs).

## Running the relay

### Configuration

Config parameters can be set either in the [config.json](https://github.com/FastLane-Labs/atlas-operations-relay/blob/main/config.json) file, or as CLI flags, and some via environment variables.
If a combination is used, the priority order is **ENV var** > **CLI param** > **config file**.

- **ETHEREUM RPC URL** (required)
    - *config.json*:  `network.rpc_url` (string)
    - *CLI flag*:     `--network.rpc_url`
    - *ENV var*:      `ETH_RPC_URL`
    - *comment*:      The relay needs access to an operational node, it makes view calls only
- **ATLAS CONTRACT ADDRESS** (required)
    - *config.json*:  `contracts.atlas` (string)
    - *CLI flag*:     `--contracts.atlas`
- **ATLAS VERIFICATION CONTRACT ADDRESS** (required)
    - *config.json*:  `contracts.atlasVerification` (string)
    - *CLI flag*:     `--contracts.atlasVerification`
- **ATLAS SIMULATOR CONTRACT ADDRESS** (required)
    - *config.json*:  `contracts.simulator` (string)
    - *CLI flag*:     `--contracts.simulator`
- **AUCTION DURATION** (optional)
    - *config.json*:  `relay.auction.duration` (integer)
    - *CLI flag*:     `--relay.auction.duration`
    - *default*:      `500`
    - *comment*:      in milliseconds, values <= 0 are ignored
- **MAX SOLUTIONS PER AUCTION**
    - *config.json*:  `relay.auction.max_solutions` (integer)
    - *CLI flag*:     `--relay.auction.max_solutions`
    - *default*:      `10`
    - *comment*:      values <= 0 are ignored
- **MAX GAS PER USER OPERATION** (optional)
    - *config.json*:  `relay.gas.max_per_user_operation` (integer)
    - *CLI flag*:     `--relay.gas.max_per_user_operation`
    - *default*:      1000000
    - *comment*:      values <= 0 are ignored
- **MAX GAS PER DAPP OPERATION** (optional)
    - *config.json*:  `relay.gas.max_per_dApp_operation` (integer)
    - *CLI flag*:     `--relay.gas.max_per_dApp_operation`
    - *default*:      1000000
    - *comment*:      values <= 0 are ignored

### Docker

Use [Dockerfile](https://github.com/FastLane-Labs/atlas-operations-relay/blob/main/Dockerfile) to build a Docker image, and run it with port `8080` exposed.
```
docker build -t ops_relay:latest .
docker run --name ops_relay -p 127.0.0.1:8080:8080 ops_relay
```

## Contributions

Contributions are welcomed, feel free to submit your pull requests.