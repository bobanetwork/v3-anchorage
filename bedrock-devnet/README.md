# bedrock-devnet

This is a utility for running a local Bedrock devnet. It is designed to replace the legacy Bash-based devnet runner as part of a progressive migration away from Bash automation.

The easiest way to invoke this script is to run `make devnet-up-deploy` from the root of this repository. Otherwise, to use this script run `python3 main.py --monorepo-dir=<path to the monorepo>`. You may need to set `PYTHONPATH` to this directory if you are invoking the script from somewhere other than `bedrock-devnet`.

Another way to spin the stack up locally is: `make devnet-hardhat-up`

## Deploy with AA contracts

Run from the root of the repository

```
make devnet-hardhat-up-with-aa
```
This deployment, for a local devnet, uses a sample create2Deployer (which does not exist in other networks), since other standard deploymnet proxies are not available on the local net. The local deployment also funds accounts on L2 for the deployment. Use the script without the --aa-setup flags to customize the script and run `python3 main.py --monorepo-dir=<path to the monorepo> --aa`. Customize eth_price_oracle address and the l2_boba_token addresses through env variables.

## Troubleshooting

### Genesis or rollup.json not generated
Under some circumstances the stack creates empty folders when the .devnet/ config files have not been generated yet. Due to the nature of docker volumes and docker cache you actively need to delete the volumes by itself to remove the faulty state.
