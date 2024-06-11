# Discovering CometBFT

Based on: [CometBFT RPC](https://docs.cometbft.com/v1.0/rpc)

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

### Get genesis

```shell
curl -s -X GET "http://0.0.0.0:26657/genesis" | jq .
```

Output:

```json
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "genesis": {
      "genesis_time": "2024-06-05T14:26:30.900531975Z",
      "chain_id": "disco",
      "initial_height": "1",
      "consensus_params": {
        "block": {
          "max_bytes": "22020096",
          "max_gas": "-1"
        },
        "evidence": {
          "max_age_num_blocks": "100000",
          "max_age_duration": "172800000000000",
          "max_bytes": "1048576"
        },
        "validator": {
          "pub_key_types": [
            "ed25519"
          ]
        },
        "version": {
          "app": "0"
        },
        "abci": {
          "vote_extensions_enable_height": "0"
        }
      },
      "app_hash": "",
      "app_state": {
        "06-solomachine": null,
        "07-tendermint": null,
        "auth": {
          "params": {
            "max_memo_characters": "256",
            "tx_sig_limit": "7",
            "tx_size_cost_per_byte": "10",
            "sig_verify_cost_ed25519": "590",
            "sig_verify_cost_secp256k1": "1000"
          },
          "accounts": [
            {
              "@type": "/cosmos.auth.v1beta1.BaseAccount",
              "address": "cosmos1secvjzt473ddvgtsv2lwrpe4r88hyrmamel2td",
              "pub_key": null,
              "account_number": "0",
              "sequence": "0"
            },
            {
              "@type": "/cosmos.auth.v1beta1.BaseAccount",
              "address": "cosmos1k5wrjnxny7ymsymywau3av2pg4v4cg5u99pz9e",
              "pub_key": null,
              "account_number": "1",
              "sequence": "0"
            }
          ]
        },
        "authz": {
          "authorization": []
        },
        "bank": {
          "params": {
            "send_enabled": [],
            "default_send_enabled": true
          },
          "balances": [
            {
              "address": "cosmos1secvjzt473ddvgtsv2lwrpe4r88hyrmamel2td",
              "coins": [
                {
                  "denom": "stake",
                  "amount": "200000000"
                },
                {
                  "denom": "token",
                  "amount": "20000"
                }
              ]
            },
            {
              "address": "cosmos1k5wrjnxny7ymsymywau3av2pg4v4cg5u99pz9e",
              "coins": [
                {
                  "denom": "stake",
                  "amount": "100000000"
                },
                {
                  "denom": "token",
                  "amount": "10000"
                }
              ]
            }
          ],
          "supply": [
            {
              "denom": "stake",
              "amount": "300000000"
            },
            {
              "denom": "token",
              "amount": "30000"
            }
          ],
          "denom_metadata": [],
          "send_enabled": []
        },
        "capability": {
          "index": "1",
          "owners": []
        },
        "circuit": {
          "account_permissions": [],
          "disabled_type_urls": []
        },
        "consensus": null,
        "crisis": {
          "constant_fee": {
            "amount": "1000",
            "denom": "stake"
          }
        },
        "disco": {
          "params": {}
        },
        "distribution": {
          "delegator_starting_infos": [],
          "delegator_withdraw_infos": [],
          "fee_pool": {
            "community_pool": []
          },
          "outstanding_rewards": [],
          "params": {
            "base_proposer_reward": "0.000000000000000000",
            "bonus_proposer_reward": "0.000000000000000000",
            "community_tax": "0.020000000000000000",
            "withdraw_addr_enabled": true
          },
          "previous_proposer": "",
          "validator_accumulated_commissions": [],
          "validator_current_rewards": [],
          "validator_historical_rewards": [],
          "validator_slash_events": []
        },
        "evidence": {
          "evidence": []
        },
        "feegrant": {
          "allowances": []
        },
        "feeibc": {
          "fee_enabled_channels": [],
          "forward_relayers": [],
          "identified_fees": [],
          "registered_counterparty_payees": [],
          "registered_payees": []
        },
        "genutil": {
          "gen_txs": [
            {
              "body": {
                "messages": [
                  {
                    "@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
                    "description": {
                      "moniker": "mynode",
                      "identity": "",
                      "website": "",
                      "security_contact": "",
                      "details": ""
                    },
                    "commission": {
                      "rate": "0.100000000000000000",
                      "max_rate": "0.200000000000000000",
                      "max_change_rate": "0.010000000000000000"
                    },
                    "min_self_delegation": "1",
                    "delegator_address": "",
                    "validator_address": "cosmosvaloper1secvjzt473ddvgtsv2lwrpe4r88hyrma7dtl87",
                    "pubkey": {
                      "@type": "/cosmos.crypto.ed25519.PubKey",
                      "key": "rblUWQSomA8pUbGVINfs3ZEzL0oWzoG6+KAiutXJvuY="
                    },
                    "value": {
                      "denom": "stake",
                      "amount": "100000000"
                    }
                  }
                ],
                "memo": "b77488f9fa23715979fe03cfc19fd4f3b3750347@192.168.17.2:26656",
                "timeout_height": "0",
                "extension_options": [],
                "non_critical_extension_options": []
              },
              "auth_info": {
                "signer_infos": [
                  {
                    "public_key": {
                      "@type": "/cosmos.crypto.secp256k1.PubKey",
                      "key": "AlH82xkn2yPdPi493oMSC8YF7aUB3aLhnnQv4Hzg1yXT"
                    },
                    "mode_info": {
                      "single": {
                        "mode": "SIGN_MODE_DIRECT"
                      }
                    },
                    "sequence": "0"
                  }
                ],
                "fee": {
                  "amount": [],
                  "gas_limit": "200000",
                  "payer": "",
                  "granter": ""
                },
                "tip": null
              },
              "signatures": [
                "+7yCVWcJZvkl/0QtumRIjVrYb0KeJUNCm/VQ0Hm5Vh1vA8n1DqlvGPCqvplAPMJc/I0BvUC9rS5ggizze36W2Q=="
              ]
            }
          ]
        },
        "gov": {
          "constitution": "",
          "deposit_params": null,
          "deposits": [],
          "params": {
            "burn_proposal_deposit_prevote": false,
            "burn_vote_quorum": false,
            "burn_vote_veto": true,
            "expedited_min_deposit": [
              {
                "amount": "50000000",
                "denom": "stake"
              }
            ],
            "expedited_threshold": "0.667000000000000000",
            "expedited_voting_period": "86400s",
            "max_deposit_period": "172800s",
            "min_deposit": [
              {
                "amount": "10000000",
                "denom": "stake"
              }
            ],
            "min_deposit_ratio": "0.010000000000000000",
            "min_initial_deposit_ratio": "0.000000000000000000",
            "proposal_cancel_dest": "",
            "proposal_cancel_ratio": "0.500000000000000000",
            "quorum": "0.334000000000000000",
            "threshold": "0.500000000000000000",
            "veto_threshold": "0.334000000000000000",
            "voting_period": "172800s"
          },
          "proposals": [],
          "starting_proposal_id": "1",
          "tally_params": null,
          "votes": [],
          "voting_params": null
        },
        "group": {
          "group_members": [],
          "group_policies": [],
          "group_policy_seq": "0",
          "group_seq": "0",
          "groups": [],
          "proposal_seq": "0",
          "proposals": [],
          "votes": []
        },
        "ibc": {
          "channel_genesis": {
            "ack_sequences": [],
            "acknowledgements": [],
            "channels": [],
            "commitments": [],
            "next_channel_sequence": "0",
            "params": {
              "upgrade_timeout": {
                "height": {
                  "revision_height": "0",
                  "revision_number": "0"
                },
                "timestamp": "600000000000"
              }
            },
            "receipts": [],
            "recv_sequences": [],
            "send_sequences": []
          },
          "client_genesis": {
            "clients": [],
            "clients_consensus": [],
            "clients_metadata": [],
            "create_localhost": false,
            "next_client_sequence": "0",
            "params": {
              "allowed_clients": [
                "*"
              ]
            }
          },
          "connection_genesis": {
            "client_connection_paths": [],
            "connections": [],
            "next_connection_sequence": "0",
            "params": {
              "max_expected_time_per_block": "30000000000"
            }
          }
        },
        "interchainaccounts": {
          "controller_genesis_state": {
            "active_channels": [],
            "interchain_accounts": [],
            "params": {
              "controller_enabled": true
            },
            "ports": []
          },
          "host_genesis_state": {
            "active_channels": [],
            "interchain_accounts": [],
            "params": {
              "allow_messages": [
                "*"
              ],
              "host_enabled": true
            },
            "port": "icahost"
          }
        },
        "mint": {
          "minter": {
            "annual_provisions": "0.000000000000000000",
            "inflation": "0.130000000000000000"
          },
          "params": {
            "blocks_per_year": "6311520",
            "goal_bonded": "0.670000000000000000",
            "inflation_max": "0.200000000000000000",
            "inflation_min": "0.070000000000000000",
            "inflation_rate_change": "0.130000000000000000",
            "mint_denom": "stake"
          }
        },
        "nft": {
          "classes": [],
          "entries": []
        },
        "params": null,
        "runtime": null,
        "slashing": {
          "missed_blocks": [],
          "params": {
            "downtime_jail_duration": "600s",
            "min_signed_per_window": "0.500000000000000000",
            "signed_blocks_window": "100",
            "slash_fraction_double_sign": "0.050000000000000000",
            "slash_fraction_downtime": "0.010000000000000000"
          },
          "signing_infos": []
        },
        "staking": {
          "delegations": [],
          "exported": false,
          "last_total_power": "0",
          "last_validator_powers": [],
          "params": {
            "bond_denom": "stake",
            "historical_entries": 10000,
            "max_entries": 7,
            "max_validators": 100,
            "min_commission_rate": "0.000000000000000000",
            "unbonding_time": "1814400s"
          },
          "redelegations": [],
          "unbonding_delegations": [],
          "validators": []
        },
        "transfer": {
          "denom_traces": [],
          "params": {
            "receive_enabled": true,
            "send_enabled": true
          },
          "port_id": "transfer",
          "total_escrowed": []
        },
        "upgrade": {},
        "vesting": {}
      }
    }
  }
}
```
