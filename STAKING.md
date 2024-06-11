# Staking

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

Discover validators using CLI:

```shell
discod query staking validators
```

Output:

```text
pagination:
  total: "1"
validators:
- commission:
    commission_rates:
      max_change_rate: "10000000000000000"
      max_rate: "200000000000000000"
      rate: "100000000000000000"
    update_time: "2024-06-05T14:26:30.900531975Z"
  consensus_pubkey:
    type: tendermint/PubKeyEd25519
    value: rblUWQSomA8pUbGVINfs3ZEzL0oWzoG6+KAiutXJvuY=
  delegator_shares: "100000000000000000000000000"
  description:
    moniker: mynode
  min_self_delegation: "1"
  operator_address: cosmosvaloper1secvjzt473ddvgtsv2lwrpe4r88hyrma7dtl87
  status: 3
  tokens: "100000000"
  unbonding_time: "1970-01-01T00:00:00Z"
```

Discover validators using REST:

```shell
curl http://0.0.0.0:1317/cosmos/staking/v1beta1/validators | jq .
```

Output:

```json
{
  "validators": [
    {
      "operator_address": "cosmosvaloper1secvjzt473ddvgtsv2lwrpe4r88hyrma7dtl87",
      "consensus_pubkey": {
        "@type": "/cosmos.crypto.ed25519.PubKey",
        "key": "rblUWQSomA8pUbGVINfs3ZEzL0oWzoG6+KAiutXJvuY="
      },
      "jailed": false,
      "status": "BOND_STATUS_BONDED",
      "tokens": "100000000",
      "delegator_shares": "100000000.000000000000000000",
      "description": {
        "moniker": "mynode",
        "identity": "",
        "website": "",
        "security_contact": "",
        "details": ""
      },
      "unbonding_height": "0",
      "unbonding_time": "1970-01-01T00:00:00Z",
      "commission": {
        "commission_rates": {
          "rate": "0.100000000000000000",
          "max_rate": "0.200000000000000000",
          "max_change_rate": "0.010000000000000000"
        },
        "update_time": "2024-06-05T14:26:30.900531975Z"
      },
      "min_self_delegation": "1",
      "unbonding_on_hold_ref_count": "0",
      "unbonding_ids": []
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

## [REST endpoints](https://docs.cosmos.network/v0.50/build/modules/staking#rest)

```text
/cosmos/staking/v1beta1/delegations/{delegatorAddr}
/cosmos/staking/v1beta1/delegators/{delegatorAddr}/redelegations
/cosmos/staking/v1beta1/delegators/{delegatorAddr}/unbonding_delegations
/cosmos/staking/v1beta1/delegators/{delegatorAddr}/validators
/cosmos/staking/v1beta1/delegators/{delegatorAddr}/validators/{validatorAddr}
/cosmos/staking/v1beta1/historical_info/{height}
/cosmos/staking/v1beta1/params
/cosmos/staking/v1beta1/pool
/cosmos/staking/v1beta1/validators
/cosmos/staking/v1beta1/validators/{validatorAddr}
/cosmos/staking/v1beta1/validators/{validatorAddr}/delegations
/cosmos/staking/v1beta1/validators/{validatorAddr}/delegations/{delegatorAddr}
/cosmos/staking/v1beta1/validators/{validatorAddr}/delegations/{delegatorAddr}/unbonding_delegation
/cosmos/staking/v1beta1/validators/{validatorAddr}/unbonding_delegations
```
