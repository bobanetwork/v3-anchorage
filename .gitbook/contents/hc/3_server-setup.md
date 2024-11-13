# Set Up a Server

For any smart contract, with or without an off-chain handler, the next step is to set up a server to run it. In this example, we will create a simple `JSON-RPC` server using the Hybrid Compute SDK. Add the lines to your "addsub" script:

{% tabs %}
{% tab title="Python" %} 
```python
def server_loop():
    # prepare the server, add your action handler start the server
    sdk.create_json_rpc_server_instance()
    sdk.add_server_action("addsub2(uint32,uint32)", offchain_add_sub)
    sdk.serve_forever()

server_loop()
```

{% endtab %}

{% tab title="Typescript" %}
```typescript
/**  use the HC SDK to create a server, add a rpc method and start the server */
const hybridCompute = new HybridComputeSDK()
    .createJsonRpcServerInstance()
    .addServerAction('addsub(uint32,uint32)', action)
    .listenAt(1234);

console.log(`Started successfully: ${hybridCompute.isServerHealthy()}`)
```
{% endtab %}
{% endtabs %}

## Why is that?

The bundler sends us a `JSON-RPC v2` request, containing a `method` identifier. This identifier is the function selector of the desired off-chain method, along with several standard parameters (the names of which are subject to change) and a `payload` which contains the ABI-encoded request data:

``` JSON
{
  "jsonrpc":"2.0",
  "id":0,
  "method":"97e0d7ba",
  "params":{
    "ver":"0.3",
    "sk":"f27fb73f63cd38cee89c48053fe8bb3248ddb7a98ce9f45b9176d017df47d9ce",
    "src_addr":"b43a2532e87583351b9024d6a6d0ba7acfa54446",
    "src_nonce":"0000000000000000000000000000000000000000000003eb0000000000000003",
    "oo_nonce":"0xb43a2532e87583351b9024d6a6d0ba7acfa544460000000000000003",
    "payload":"00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000001"
  }
}
```

With the server set up, proceed to the following section to learn how to write our own smart contract.
