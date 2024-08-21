# Deploy the Smart Contract

With a few adjustments to the [`deploy-local.py` script](https://github.com/bobanetwork/rundler-hc/blob/boba-develop/hybrid-compute/deploy-local.py), we can easily deploy our newly created smart contract.

First, we need to configure the script to connect to our L1 and L2 networks:

```python
l1 = Web3(Web3.HTTPProvider("http://127.0.0.1:8545"))
assert(l1.is_connected)
l1.middleware_onion.inject(geth_poa_middleware, layer = 0)

l2 = Web3(Web3.HTTPProvider("http://127.0.0.1:9545"))
assert(l2.is_connected)
l2.middleware_onion.inject(geth_poa_middleware, layer = 0)
```

Next, we load the contract using the `loadContract()` function:

```python
TC = loadContract(w3, "TestCounter", path_prefix + "test/TestCounter.sol")
```

Here `path_prefix` indicates where the contract is located.

After loading the contract, we can deploy it in this manner:

```solidity
vm.startBroadcast(deployerPrivateKey);

// These contracts are one example each
ret[0] = address(new AuctionFactory(ha1Addr));
ret[1] = address(new TestCaptcha(ha1Addr));
ret[2] = address(new TestCounter(ha1Addr));
ret[3] = address(new RainfallInsurance(ha1Addr));
ret[4] = address(new SportsBetting(ha1Addr));

vm.stopBroadcast();
```

{% hint style="info" %}
In the above example we deploy with Foundry, but feel free to deploy with whatever tooling you wish.
{% endhint %}

## Additional Examples

The documentation above was specifically written for the addition of two numbers. The `hybrid-compute/` folder contains more examples that can be used and experimented with. Let's integrate them into our `server-loop`:

```python
  def server_loop():
    server = SimpleJSONRPCServer(
        ('0.0.0.0', PORT),
        requestHandler=RequestHandler
    )

    // Add Sub
    server.register_function(offchain_addsub2, selector("addsub2(uint32,uint32)"))  # 97e0d7ba

    // Ramble
    server.register_function(offchain_ramble,  selector("ramble(uint256,bool)"))

    // CheckKyc
    server.register_function(offchain_checkkyc, selector("checkkyc(string)"))

    // getPrice
    server.register_function(offchain_getprice, selector("getprice(string)"))
```

Note: each of these `register_function()` statements represents a whitelisting for an off-chain RPC server.

You've now reached the end of this tutorial! These examples aim to demonstrate the different functions and best practices you should keep in mind as you develop your own efficient, flexible smart contracts with Hybrid Compute.
