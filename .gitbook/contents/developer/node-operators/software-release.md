# Node Software Releases

This page provides a list of the necessary versions of node software and instructions on how to keep them updated.

Our latest releases, notes and changelogs can be found on Github. `op-node` releases can be found [here](https://github.com/bobanetwork/boba/tags) and `op-erigon` release can be found [here](https://github.com/bobanetwork/op-erigon/releases).

## Required Version by Network

These are the minimal required versions for the `op-node`, `op-erigon` and `op-geth` by network.

| Network          | op-node                                                      | op-erigon                                                    | op-geth                                                      |
| ---------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| Boba Mainnet | [v1.6.8](https://github.com/bobanetwork/boba/releases/tag/v1.6.8) | [v1.2.0](https://github.com/bobanetwork/op-erigon/releases/tag/v1.2.0) | [v1.101315.2](https://github.com/ethereum-optimism/op-geth/releases/tag/v1.101315.2) |
| Boba Sepolia | [v1.6.7](https://github.com/bobanetwork/boba/releases/tag/v1.6.7) | [v1.1.10](https://github.com/bobanetwork/op-erigon/releases/tag/v1.1.10) | [v1.101315.2](https://github.com/ethereum-optimism/op-geth/releases/tag/v1.101315.2) |
| Op Mainnet   | [v1.6.7](https://github.com/bobanetwork/boba/releases/tag/v1.6.7) | [v1.1.10](https://github.com/bobanetwork/op-erigon/releases/tag/v1.1.10) | [v1.101315.2](https://github.com/ethereum-optimism/op-geth/releases/tag/v1.101315.2) |
| Op Sepolia   | [v1.6.7](https://github.com/bobanetwork/boba/releases/tag/v1.6.7) | [v1.1.10](https://github.com/bobanetwork/op-erigon/releases/tag/v1.1.10) | [v1.101315.2](https://github.com/ethereum-optimism/op-geth/releases/tag/v1.101315.2) |
| Boba Bnb Testnet | [v1.6.7](https://github.com/bobanetwork/boba/releases/tag/v1.6.7) | [v1.1.10](https://github.com/bobanetwork/op-erigon/releases/tag/v1.1.10) | [v1.101315.2](https://github.com/ethereum-optimism/op-geth/releases/tag/v1.101315.2) |

## [op-node v1.6.8](https://github.com/bobanetwork/boba/releases/tag/v1.6.8)

**Description**

This is a mandatory release for node operators on Boba Mainnet networks. The **Fjord** protocol upgrades will activate on Tue Sep 10 07:00:00 UTC 2024 on Boba Mainnet networks.

The `op-node`  needs the flag `--plasma.enabled=false` to start it.

**Required Action**

Upgrade your `op-node` software and add  the `--plasma.enabled=false` to the configuration.

## [op-erigon v1.2.0](https://github.com/bobanetwork/op-erigon/releases/tag/v1.2.0)

**Description**

This is a mandatory release for node operators on Boba Sepolia and Boba BNB Testnet networks. The **Fjord** protocol upgrades will activate onTue Sep 10 07:00:00 UTC 2024 on Boba Sepolia and Boba BNB Testnet networks.

**Required Action**

Upgrade your `op-erigon` software.

## [op-geth v1.101315.2](https://github.com/ethereum-optimism/op-geth/releases/tag/v1.101315.2)

**Description**

This is a mandatory release for node operators on Op networks and Boba Sepolia and Boba BNB Testnet works.

**Required Action**

* Upgrade your `op-geth` software.
* Add `--override.fjord=1725951600` for the `Boba Mainnet` networks when you start it for the first time.
* Add `--override.fjord=1722297600` for the `Boba Sepolia` and `Boba BNB Testnet` networks when you start it for the first time.

## [op-node v1.6.7](https://github.com/bobanetwork/boba/releases/tag/v1.6.7)

**Description**

This is a mandatory release for node operators on Boba Sepolia and Boba BNB Testnet networks. The **Fjord** protocol upgrades will activate on Tue Jul 30 2024 00:00:00 UTC 2024 on Boba Sepolia and Boba BNB Testnet networks.

The `op-node`  needs the flag `--plasma.enabled=false` to start it.

**Required Action**

Upgrade your `op-node` software and add  the `--plasma.enabled=false` to the configuration.

## [op-erigon v1.1.10](https://github.com/bobanetwork/op-erigon/releases/tag/v1.1.10)

**Description**

This is a mandatory release for node operators on Boba Sepolia and Boba BNB Testnet networks. The **Fjord** protocol upgrades will activate on Tue Jul 30 2024 00:00:00 UTC 2024 on Boba Sepolia and Boba BNB Testnet networks.

**Required Action**

Upgrade your `op-erigon` software.
