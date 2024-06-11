# Discovering CometBFT

Run the chain:

```shell
ignite chain serve
```

Output:

```text
  Blockchain is running
  
  👤 alice's account address: cosmos1secvjzt473ddvgtsv2lwrpe4r88hyrmamel2td
  👤 bob's account address: cosmos1k5wrjnxny7ymsymywau3av2pg4v4cg5u99pz9e
  
  🌍 Tendermint node: http://0.0.0.0:26657
  🌍 Blockchain API: http://0.0.0.0:1317
  🌍 Token faucet: http://0.0.0.0:4500
  
  ⋆ Data directory: ~/.disco
  ⋆ App binary: ~/go/bin/discod
  
  Press the 'q' key to stop serve
```

> Note that Tendermint node is available under this address: **http://0.0.0.0:26657**

### Get the node's health status information

```shell
curl -s -X GET "http://0.0.0.0:26657/health" | jq .
```

Output: 
```json
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {}
}
``` 

### Get node's status

```shell
curl -s -X GET "http://0.0.0.0:26657/status" | jq .
```

Output:
```json
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "node_info": {
      "protocol_version": {
        "p2p": "8",
        "block": "11",
        "app": "0"
      },
      "id": "b77488f9fa23715979fe03cfc19fd4f3b3750347",
      "listen_addr": "tcp://0.0.0.0:26656",
      "network": "disco",
      "version": "0.38.6",
      "channels": "40202122233038606100",
      "moniker": "mynode",
      "other": {
        "tx_index": "on",
        "rpc_address": "tcp://0.0.0.0:26657"
      }
    },
    "sync_info": {
      "latest_block_hash": "2512729BF836B206CB443BF3329A4130984B2D60E6198064A59EF1396AE6C1FD",
      "latest_app_hash": "7F06A859358907CDBD21053951516E0495C05F15DC2E68DCDCBCC65FDC6A85B6",
      "latest_block_height": "26083",
      "latest_block_time": "2024-06-11T14:07:57.12519517Z",
      "earliest_block_hash": "30FDA47FCC55758A1C98615379811C6BC4F4DBB2AE962C5E6F837D8E86ACD665",
      "earliest_app_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
      "earliest_block_height": "1",
      "earliest_block_time": "2024-06-05T14:26:30.900531975Z",
      "catching_up": false
    },
    "validator_info": {
      "address": "F96E556F296977D9B6F53E67F5390DEC151B20FF",
      "pub_key": {
        "type": "tendermint/PubKeyEd25519",
        "value": "rblUWQSomA8pUbGVINfs3ZEzL0oWzoG6+KAiutXJvuY="
      },
      "voting_power": "100"
    }
  }
}
```
