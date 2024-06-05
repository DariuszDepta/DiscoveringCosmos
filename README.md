# Discovering [Cosmos](https://docs.cosmos.network)


## Prerequisities

- Install the newest version of [Go](https://go.dev/doc/install).

  Check installed Go version:

  ```shell
  go version
  ```

  Output:

  ```text
  go version go1.22.4 linux/amd64
  ```

- Install the newest version of [IgniteCLI](https://docs.ignite.com/welcome/install).
  
  Check installed Ignite CLI version:

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
  Your uname -a:                  [..]
  Your cwd:                       [..]
  Is on Gitpod:                   false
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

Re-run the chain, to make sure it works fine:

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
