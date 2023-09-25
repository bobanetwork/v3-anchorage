This directory contains a test script and example contract which demonstrate
V3 Hybrid Compute.

First, bring up a local stack using "make devnet-hardhat-up" in the top-level
directory. In this directory, run "yarn" once for setup and then "yarn build"
to compile the demo contract. To run the test script, "yarn test:local".

The test script runs its own "offchain" webserver on port 1234, providing
functions which add two numbers or which generate a string with a specified
word count. The test script registers an Endpoint with the HCHelper contract,
deploys and registers the HCDemo contract along with a prepaid credit balance,
and then calls various functions in the contract. At the end of the test,
the endpoint is unregistered.
