# Discovering [Cosmos](https://docs.cosmos.network)

## Prerequisities

### Install the newest version of [Go](https://go.dev/doc/install)

Check installed `Go` version:

```shell
go version
```

Output:

```text
go version go1.22.4 linux/amd64
```

### Install the newest version of [IgniteCLI](https://docs.ignite.com/welcome/install)
  
Check installed `Ignite CLI` version:

```shell
ignite version
```

Output:

```text
Ignite CLI version:             v28.4.0
Ignite CLI build date:          2024-05-15T13:42:13Z
Ignite CLI source hash:         83ee9ba5f81f2d2104ed91808f2cb72719a23e41
Ignite CLI config version:      v1
Cosmos SDK version:             v0.50.6
Your OS:                        linux
Your arch:                      amd64
Your Node.js version:           v22.0.0
Your go version:                go version go1.22.4 linux/amd64
Your uname -a:                  [..omitted]
Your cwd:                       [..omitted]
Is on Gitpod:                   false
```

### Install [jq](https://jqlang.github.io/jq/download)

Check installed `jq` version:

```shell
jq --version
```

Output:

```text
jq-1.7.1
```

### Install [curl](https://curl.se)

Check installed `curl` version:

```shell
curl --version
```

Output:

```text
curl 8.6.0 (x86_64-redhat-linux-gnu) [..omitted]
Release-Date: 2024-01-31
Protocols: [..omitted]
Features: [..omitted]
```

## Create a chain

Scaffold a chain named `disco`:

```shell
ignite scaffold chain disco
```
 Output:
 
```text
⭐️ Successfully created a new blockchain 'disco'.
👉 Get started with the following commands:

 % cd disco
 % ignite chain serve

Documentation: https://docs.ignite.com
```

Follow the instructions above:

```shell
cd disco
```

```shell
ignite chain serve
```

Output:

```text
Blockchain is running
  
  ✔ Added account alice with address cosmos1secvjzt473ddvgtsv2lwrpe4r88hyrmamel2td and mnemonic:
  science pony machine viable trash bitter average off bind remove hungry click rigid fork select clown theme cable episode black actress van rifle length
  
  ✔ Added account bob with address cosmos1k5wrjnxny7ymsymywau3av2pg4v4cg5u99pz9e and mnemonic:
  pond puppy margin cotton govern genre argue trial clown paper salute pride wedding defense steel share ketchup pause vacuum flee trophy solid quit evoke
  
  🌍 Tendermint node: http://0.0.0.0:26657
  🌍 Blockchain API: http://0.0.0.0:1317
  🌍 Token faucet: http://0.0.0.0:4500
  
  ⋆ Data directory: ~/.disco
  ⋆ App binary: ~/go/bin/discod
  
  Press the 'q' key to stop serve
```

Stop the chain by pressing `q`:

```text
  💿 Genesis state saved in ~/.ignite/local-chains/disco/exported_genesis.json
  
  𝓲 Stopped
```

Re-run the chain, to make sure it works:

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

## Discover blockchain API

Start the chain:

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

Blockchain exposes its API on this address:

```text
  🌍 Blockchain API: http://0.0.0.0:1317
```

The blockchain API documentation is available [**here**](https://docs.cosmos.network/api).

### Exploring REST API endpoints

Open a new terminal window and try the following commands.

#### Get all existing [Accounts](https://docs.cosmos.network/api#tag/Query/operation/Accounts)

```shell
curl http://0.0.0.0:1317/cosmos/auth/v1beta1/accounts | jq .
```

Output:

```json
{
  "accounts": [
    {
      "@type": "/cosmos.auth.v1beta1.ModuleAccount",
      "base_account": {
        "address": "cosmos1fl48vsnmsdzcv85q5d2q4z5ajdha8yu34mf0eh",
        "pub_key": null,
        "account_number": "4",
        "sequence": "0"
      },
      "name": "bonded_tokens_pool",
      "permissions": [
        "burner",
        "staking"
      ]
    },
    {
      "@type": "/cosmos.auth.v1beta1.ModuleAccount",
      "base_account": {
        "address": "cosmos1tygms3xhhs3yv487phx3dw4a95jn7t7lpm470r",
        "pub_key": null,
        "account_number": "5",
        "sequence": "0"
      },
      "name": "not_bonded_tokens_pool",
      "permissions": [
        "burner",
        "staking"
      ]
    },
    {
      "@type": "/cosmos.auth.v1beta1.ModuleAccount",
      "base_account": {
        "address": "cosmos10d07y265gmmuvt4z0w9aw880jnsr700j6zn9kn",
        "pub_key": null,
        "account_number": "6",
        "sequence": "0"
      },
      "name": "gov",
      "permissions": [
        "burner"
      ]
    },
    {
      "@type": "/cosmos.auth.v1beta1.BaseAccount",
      "address": "cosmos1secvjzt473ddvgtsv2lwrpe4r88hyrmamel2td",
      "pub_key": {
        "@type": "/cosmos.crypto.secp256k1.PubKey",
        "key": "AlH82xkn2yPdPi493oMSC8YF7aUB3aLhnnQv4Hzg1yXT"
      },
      "account_number": "0",
      "sequence": "1"
    },
    {
      "@type": "/cosmos.auth.v1beta1.ModuleAccount",
      "base_account": {
        "address": "cosmos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd88lyufl",
        "pub_key": null,
        "account_number": "3",
        "sequence": "0"
      },
      "name": "distribution",
      "permissions": []
    },
    {
      "@type": "/cosmos.auth.v1beta1.BaseAccount",
      "address": "cosmos1k5wrjnxny7ymsymywau3av2pg4v4cg5u99pz9e",
      "pub_key": null,
      "account_number": "1",
      "sequence": "0"
    },
    {
      "@type": "/cosmos.auth.v1beta1.ModuleAccount",
      "base_account": {
        "address": "cosmos1m3h30wlvsf8llruxtpukdvsy0km2kum8g38c8q",
        "pub_key": null,
        "account_number": "7",
        "sequence": "0"
      },
      "name": "mint",
      "permissions": [
        "minter"
      ]
    },
    {
      "@type": "/cosmos.auth.v1beta1.ModuleAccount",
      "base_account": {
        "address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
        "pub_key": null,
        "account_number": "2",
        "sequence": "0"
      },
      "name": "fee_collector",
      "permissions": []
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "8"
  }
}
```

There are few addresses of modules returned:

| Module name            | Module address                                  |
|------------------------|-------------------------------------------------|
| fee_collector          | `cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta` |
| distribution           | `cosmos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd88lyufl` |
| bonded_tokens_pool     | `cosmos1fl48vsnmsdzcv85q5d2q4z5ajdha8yu34mf0eh` | 
| not_bonded_tokens_pool | `cosmos1tygms3xhhs3yv487phx3dw4a95jn7t7lpm470r` |
| gov                    | `cosmos10d07y265gmmuvt4z0w9aw880jnsr700j6zn9kn` |
| mint                   | `cosmos1m3h30wlvsf8llruxtpukdvsy0km2kum8g38c8q` |

And two base accounts: 

| Account number | Base account address                            |
|----------------|-------------------------------------------------|
| 0 (alice)      | `cosmos1secvjzt473ddvgtsv2lwrpe4r88hyrmamel2td` |
| 1 (bob)        | `cosmos1k5wrjnxny7ymsymywau3av2pg4v4cg5u99pz9e` |

> Take a look at the addresses displayed after starting the chain (the same as above). Labels `alice` and `bob` are just helpers
> to avoid copying and pasting the whole long addresses in commands.
> ```text
> 👤 alice's account address: cosmos1secvjzt473ddvgtsv2lwrpe4r88hyrmamel2td
> 👤 bob's account address: cosmos1k5wrjnxny7ymsymywau3av2pg4v4cg5u99pz9e
> ```

> [!NOTE]  
> Why `alice` has a public key set
> 
> ```json
> {
>   "@type": "/cosmos.auth.v1beta1.BaseAccount",
>   "address": "cosmos1secvjzt473ddvgtsv2lwrpe4r88hyrmamel2td",
>   "pub_key": {
>     "@type": "/cosmos.crypto.secp256k1.PubKey",
>     "key": "AlH82xkn2yPdPi493oMSC8YF7aUB3aLhnnQv4Hzg1yXT"
>   },
>   "account_number": "0",
>   "sequence": "1"
> }
> ```
> and `bob` does not?
> 
> ```json
> {
>   "@type": "/cosmos.auth.v1beta1.BaseAccount",
>   "address": "cosmos1k5wrjnxny7ymsymywau3av2pg4v4cg5u99pz9e",
>   "pub_key": null,
>   "account_number": "1",
>   "sequence": "0"
> }
> ```

#### Get [AccountInfo](https://docs.cosmos.network/api#tag/Query/operation/AccountInfo)

```shell
curl http://0.0.0.0:1317/cosmos/auth/v1beta1/account_info/cosmos1secvjzt473ddvgtsv2lwrpe4r88hyrmamel2td | jq .
```

Output:

```json
{
  "info": {
    "address": "cosmos1secvjzt473ddvgtsv2lwrpe4r88hyrmamel2td",
    "pub_key": {
      "@type": "/cosmos.crypto.secp256k1.PubKey",
      "key": "AlH82xkn2yPdPi493oMSC8YF7aUB3aLhnnQv4Hzg1yXT"
    },
    "account_number": "0",
    "sequence": "1"
  }
}
```

```shell
curl http://0.0.0.0:1317/cosmos/auth/v1beta1/account_info/cosmos1k5wrjnxny7ymsymywau3av2pg4v4cg5u99pz9e | jq .
```

Output:

```json
{
  "info": {
    "address": "cosmos1k5wrjnxny7ymsymywau3av2pg4v4cg5u99pz9e",
    "pub_key": null,
    "account_number": "1",
    "sequence": "0"
  }
}
```
