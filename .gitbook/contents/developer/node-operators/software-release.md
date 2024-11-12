# Node Software Releases

This page provides a list of the necessary versions of node software and instructions on how to keep them updated.

Our latest releases, notes and changelogs can be found on Github. `op-node` releases can be found [here](https://github.com/bobanetwork/boba/tags) and `op-erigon` release can be found [here](https://github.com/bobanetwork/op-erigon/releases).

## Required Version by Network

These are the minimal required versions for the `op-node`, `op-erigon` and `op-geth` by network.

| Network          | op-node                                                      | op-erigon                                                    | op-geth                                                      |
| ---------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| Boba Mainnet | [v1.6.11](https://github.com/bobanetwork/boba/releases/tag/v1.6.11) | [v1.2.5](https://github.com/bobanetwork/op-erigon/releases/tag/v1.2.5) | [v1.101408.0](https://github.com/ethereum-optimism/op-geth/releases/tag/v1.101408.0) |
| Boba Sepolia | [v1.6.10](https://github.com/bobanetwork/boba/releases/tag/v1.6.10) | [v1.2.2](https://github.com/bobanetwork/op-erigon/releases/tag/v1.2.2) | [v1.101408.0](https://github.com/ethereum-optimism/op-geth/releases/tag/v1.101408.0) |
| Op Mainnet   | [v1.6.10](https://github.com/bobanetwork/boba/releases/tag/v1.6.10) | [v1.2.2](https://github.com/bobanetwork/op-erigon/releases/tag/v1.2.2) | [v1.101408.0](https://github.com/ethereum-optimism/op-geth/releases/tag/v1.101408.0) |
| Op Sepolia   | [v1.6.10](https://github.com/bobanetwork/boba/releases/tag/v1.6.10) | [v1.2.2](https://github.com/bobanetwork/op-erigon/releases/tag/v1.2.2) | [v1.101408.0](https://github.com/ethereum-optimism/op-geth/releases/tag/v1.101408.0) |
| Boba Bnb Testnet | [v1.6.10](https://github.com/bobanetwork/boba/releases/tag/v1.6.10) | [v1.2.2](https://github.com/bobanetwork/op-erigon/releases/tag/v1.2.2) | [v1.101408.0](https://github.com/ethereum-optimism/op-geth/releases/tag/v1.101408.0) |

## [op-erigon v1.2.5](https://github.com/bobanetwork/op-erigon/releases/tag/v1.2.5)

**Description**

This is a mandatory release for node operators on Boba Eth Mainnet. The **Granite** protocol upgrades will activate on Thu 24 Oct 07:00:00 UTC 2024 on Boba Eth Mainnet.

**Required Action**

Upgrade your `op-erigon` software.

## [op-node v1.6.10](https://github.com/bobanetwork/boba/releases/tag/v1.6.10)

**Description**

This is a mandatory release for node operators on Boba Sepolia and Boba BNB Testnet networks. The **Granite** protocol upgrades will activate on Mon 16 Sep 07:00:00 UTC 2024 on Boba Testnet networks.

**Required Action**

Upgrade your `op-node` software.

## [op-erigon v1.2.2](https://github.com/bobanetwork/op-erigon/releases/tag/v1.2.2)

**Description**

This is a mandatory release for node operators on Boba Sepolia and Boba BNB Testnet networks. The **Granite** protocol upgrades will activate on Mon 16 Sep 07:00:00 UTC 2024 on Boba Testnet networks.

**Required Action**

Upgrade your `op-erigon` software.

## [op-geth v1.101408.0](https://github.com/ethereum-optimism/op-geth/releases/tag/v1.101408.0)

**Description**

This is a mandatory release for node operators on on Boba Sepolia and Boba BNB Testnet networks.

**Required Action**

* Upgrade your `op-geth` software.
* Add `--override.granite=1729753200` for the `Boba Mainnet` network when you start it for the first time.
* Add `--override.granite=1726470000` for the `Boba Sepolia` and `Boba BNB Testnet` networks when you start it for the first time.
