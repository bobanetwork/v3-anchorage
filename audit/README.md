# Custom Fee Token Audit

The op-stack from Optimism forms the basis of Boba Anchorage.  For
Boba-Ethereum, we are using unmodified versions of the contracts from op-stack
but for Boba-BNB we need to support the BOBA ERC20 token as the native gas
token and convert BNB Chain's native gas token BNB to and from a wrapped ERC20
version of the token on Boba BNB.

## Existing custom gas token support

Optimism provides an op-stack solution for supporting most of the basic custom
gas token functionality.  This includes bridging the L1 ERC20 token onto the
L2 as the native gas token and back.  It however does not permit bridging the
native L1 gas token into the L2 as a wrapped ERC20 token (and similarly does
not allow the reverse).

This change set implements this missing functionality on top of the existing
function.  It is intended to maintain feature compatibility with the existing
Boba-BNB bridging semantics.

## Audit Scope

We assume that the base set of contracts supplied by Optimism are sound and
rely on their existing audits for the many op-stack based chains which are
already deployed.  The changes from the upstream contracts are documented in
[custom-fee-token.diff](custom-fee-token.diff).  This is based on
a diff between the custom-fee-token branch
(c6b46d4ca9cdf7a58148d79ed1d3bdda01d31032) and the develop branch
(6c5577f30bf7405a0dc2e317696392eb88fdd911)

This diff can be reproduced from the root dir via:

```
git diff 6c5577f30bf7405a0dc2e317696392eb88fdd911..9cc0e3d8f0dbac286ea189bb86126284d1d86e44 -- \
   packages/contracts-bedrock/src/{L1,L2,libraries,universal} > custom-fee-token.diff
```

Note, this diff omits the deletion of the `WETH9.sol` contract from the
`src/boba` directory as it is not relevant to the bridging changes.  It also
omits the changes to the `deployment` directory, as they are largely error
message changes to decrease the size of the initial deployment and can be
reverted after.
