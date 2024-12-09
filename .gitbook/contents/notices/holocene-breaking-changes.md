# Prepare for Holocene Breaking Changes

The Holocene upgrade for Boba Sepolia and Boba BNB Testnet will be optimistically activated on **1736150400 Mon Jan 06 08:00:00 UTC 2025**.

## For Node Operators

Node operators are required to upgrade to Holocene before the activation date to avoid chain divergence. For Boba Testnets, the `op-node` release [v1.6.14](https://github.com/bobanetwork/boba/releases/tag/v1.6.14), the `op-erigon` release [v1.2.10](https://github.com/bobanetwork/op-erigon/releases/tag/v1.2.10) and the `op-geth` release [v1.101411.3](https://github.com/bobanetwork/op-geth/releases/tag/v1.101411.3) contain these changes

### Update to the Lastest Release

* op-node [v1.6.14](https://github.com/bobanetwork/boba/releases/tag/v1.6.14)
* op-erigon [v1.2.10](https://github.com/bobanetwork/op-erigon/releases/tag/v1.2.10)
* op-geth [v1.101411.3](https://github.com/bobanetwork/op-geth/releases/tag/v1.101411.3)

For `op-geth`, we have provided a new image built by Boba Team. The manually override is not needed anymore. Please use the new image `us-docker.pkg.dev/boba-392114/bobanetwork-tools-artifacts/images/op-geth:v1.101411.3`

### Verify Your Configuration

Make the following checks to verify that your node is properly configured.

- `op-node` , `op-erigon` and `op-geth` will log their configurations at startup
- Check that the Holocene time is set to `activation-timestamp` in the `op-node` startup logs
- Check that the Holocene time is set to `activation-timestamp` in the `op-erigon` startup logs
- Check that the Holocene time is set to `activation-timestamp` in the `op-geth` startup logs
