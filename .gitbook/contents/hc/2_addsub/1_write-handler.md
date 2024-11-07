# Write an Off-Chain Handler

The first step is to write our *off-chain handler*. This handler is responsible for receiving requests from the bundler along with their payloads and returning the appropriate responses.

## Next Example

Let's begin by allowing our handler to receive two numbers. It will perform both addition and subtraction on these numbers. If the result of the subtraction results in an underflow (i.e. if `b > a` in the expression `a-b`), the handler will respond with an underflow error. Create a `Python` file in which to copy all the following code blocks.

First, we'll need to import the Hybrid Compute SDK (viewable on Github [here](https://github.com/bobanetwork/aa-hc-sdk/tree/main/sdks)) and any other necessary libraries, depending on whether you develop in `Python` or `TypeScript`:

{% tabs %}
{% tab title="Python" %} 
```python
from web3 import Web3
from eth_abi import abi as ethabi
from jsonrpclib.SimpleJSONRPCServer import SimpleJSONRPCServer, SimpleJSONRPCRequestHandler
from hybrid_compute_sdk import HybridComputeSDK

# new sdk instance
sdk = HybridComputeSDK()
```

Initialize an `err_code` and a `resp` object with values in case of an exception:

```python
# define your server action
def offchain_add_sub(ver, sk, src_addr, src_nonce, oo_nonce, payload, *args):
    print("  -> offchain_add_sub handler called with ver={} subkey={} src_addr={} src_nonce={} oo_nonce={} payload={} extra_args={}".format(
        ver, sk, src_addr, src_nonce, oo_nonce, payload, args))
    err_code = 1
    resp = Web3.to_bytes(text="unknown error")
```

Next, we'll add a `try` block. Inside of it, let's make use of some of our imported library functions to parse and decode our binary data to an array of two `unit32`s:

```python 
    assert(ver == "0.2")
    try:
        req = sdk.parse_req(sk, src_addr, src_nonce, oo_nonce, payload)
        dec = ethabi.decode(['uint32', 'uint32'], req['reqBytes'])

    except Exception as e:
        print("DECODE FAILED", e)
```

With this decoded information, we can add and subtract the two numbers, re-encode the information (this time to an array of two `uint256`s), and check for an overflow. The entire `try` block should look like this:

```python
    assert(ver == "0.2")
    try:
        req = sdk.parse_req(sk, src_addr, src_nonce, oo_nonce, payload)
        dec = ethabi.decode(['uint32', 'uint32'], req['reqBytes'])

        if dec[0] >= dec[1]:
            s = dec[0] + dec[1]
            d = dec[0] - dec[1]
            resp = ethabi.encode(['uint256', 'uint256'], [s, d])
            err_code = 0
        else:
            print("offchain_add_sub underflow error", dec[0], dec[1])
            resp = Web3.to_bytes(text="underflow error")
    except Exception as e:
        print("DECODE FAILED", e)
```

Before we can return these objects, we need to transform them into a specific object. We can do so with the `sdk.gen_response()` function from our imported library. Putting the whole thing together, our whole `offchain_add_sub()` looks like this:

```python
# define your server action
def offchain_add_sub(ver, sk, src_addr, src_nonce, oo_nonce, payload, *args):
    print("  -> offchain_add_sub handler called with ver={} subkey={} src_addr={} src_nonce={} oo_nonce={} payload={} extra_args={}".format(
        ver, sk, src_addr, src_nonce, oo_nonce, payload, args))
    err_code = 1
    resp = Web3.to_bytes(text="unknown error")
    
    assert(ver == "0.2")
    try:
        req = sdk.parse_req(sk, src_addr, src_nonce, oo_nonce, payload)
        dec = ethabi.decode(['uint32', 'uint32'], req['reqBytes'])

        if dec[0] >= dec[1]:
            s = dec[0] + dec[1]
            d = dec[0] - dec[1]
            resp = ethabi.encode(['uint256', 'uint256'], [s, d])
            err_code = 0
        else:
            print("offchain_add_sub underflow error", dec[0], dec[1])
            resp = Web3.to_bytes(text="underflow error")
    except Exception as e:
        print("DECODE FAILED", e)

    return sdk.gen_response(req, err_code, resp)
```
{% endtab %}

{% tab title="Typescript" %}
```typescript
import {
    HybridComputeSDK,
    OffchainParameter,
    getParsedRequest,
    generateResponse,
    ServerActionResponse
} from '@bobanetwork/aa-hc-sdk-server';

import Web3 from "web3";
```

Next, set up a `try` block to handle parsing and decoding our binary data to an array of two `unit32`s:

```typescript
const action = (params: OffchainParameter): Promise<ServerActionResponse> => {
    const web3 = new Web3();
    const request = getParsedRequest(params);

    try {
        const decoded = web3.eth.abi.decodeParameters(
            ['uint32', 'uint32'],
            request.reqBytes
        );

    } catch (error: any) {
        return generateResponse(request, 1, web3.utils.asciiToHex(error.message));
    }
```

With this decoded information, we can add and subtract the two numbers, re-encode the information (this time to an array of two `uint256`s), and check for an overflow. The entire `try` block should look like this:

```typescript
    try {
        const decoded = web3.eth.abi.decodeParameters(
            ['uint32', 'uint32'],
            request.reqBytes
        );

        const num1 = Number(decoded[0]);
        const num2 = Number(decoded[1]);

        if (num1 >= num2) {
            const sum = num1 + num2;
            const difference = num1 - num2;
            const encodedResult = web3.eth.abi.encodeParameters(
                ['uint256', 'uint256'],
                [sum.toString(), difference.toString()]
            );
            return generateResponse(request, 0, encodedResult);
        } else {
            return generateResponse(
                request,
                1,
                web3.utils.asciiToHex("underflow error")
            );
        }
}
```

Putting it all together, our script looks like this:

```typescript
import {
    HybridComputeSDK,
    OffchainParameter,
    getParsedRequest,
    generateResponse,
    ServerActionResponse
} from '@bobanetwork/aa-hc-sdk-server';

import Web3 from "web3";

const action = (params: OffchainParameter): Promise<ServerActionResponse> => {
    const web3 = new Web3();
    const request = getParsedRequest(params);

    try {
        const decoded = web3.eth.abi.decodeParameters(
            ['uint32', 'uint32'],
            request.reqBytes
        );

        const num1 = Number(decoded[0]);
        const num2 = Number(decoded[1]);

        if (num1 >= num2) {
            const sum = num1 + num2;
            const difference = num1 - num2;
            const encodedResult = web3.eth.abi.encodeParameters(
                ['uint256', 'uint256'],
                [sum.toString(), difference.toString()]
            );
            return generateResponse(request, 0, encodedResult);
        } else {
            return generateResponse(
                request,
                1,
                web3.utils.asciiToHex("underflow error")
            );
        }
    } catch (error: any) {
        return generateResponse(request, 1, web3.utils.asciiToHex(error.message));
    }
}
```
{% endtab %}
{% endtabs %}

We've now successfully implemented a function which can receive a request from the bundler, perform some calculation with its payload, and return a response. Proceed to the next section to learn more about setting up a server to run this function.
