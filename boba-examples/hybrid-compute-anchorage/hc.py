import os,sys
from web3 import Web3
import threading
import signal
import time
from random import *
import queue
import requests,json
from web3.gas_strategies.time_based import fast_gas_price_strategy
from web3.middleware import geth_poa_middleware
from web3.logs import STRICT, IGNORE, DISCARD, WARN
import logging
import base64

import rlp
from multiprocessing import Process
from jsonrpclib.SimpleJSONRPCServer import SimpleJSONRPCServer,SimpleJSONRPCRequestHandler

from eth_abi import abi as ethabi

# ----------------------------------------------------------
# Methods provided by the local webserver

# Expect an ABI-encoded uint32 'x'. Return an ABI-encoded 'x+5'
def add5(req):
  print("  -> ADD5 handler called with",req)

  try:
    reqBytes = Web3.toBytes(hexstr=req)
    dec = ethabi.decode_abi(['uint32'], reqBytes)
  except Exception as e:
    print("DECODE FAILED", e)

  answer = int(dec[0]) + 5
  print("  -> ADD5 calculated", answer)

  try:
    e1 = ethabi.encode_abi(['uint32'], [answer])
    enc = Web3.toHex(e1)
  except Exception as e:
    print("ENCODE FAILED", e)
  print("  -> ADD5 responding", enc)

  return enc

# Expect an ABI-encoded uint32 'x'. Return that many "chicken"s
# cf. https://improbable.com/airchives/paperair/volume12/v12i5/chicken-12-5.pdf
def chicken(req):
  print("  -> Chicken handler called with",req)

  try:
    reqBytes = Web3.toBytes(hexstr=req)
    dec = ethabi.decode_abi(['uint32'], reqBytes)
  except Exception as e:
    print("DECODE FAILED", e)

  answer = "chicken "*int(dec[0])

  print("  -> Chicken response length", len(answer))

  try:
    answer = Web3.toBytes(text=answer)
    e1 = ethabi.encode_abi(['bytes'], [answer])
    enc = Web3.toHex(e1)
  except Exception as e:
    print("ENCODE FAILED", e)

  return enc

def sumSquares_v1(req):
  print("  -> SumSquares_v1 handler called with",req)

  try:
    reqBytes = Web3.toBytes(hexstr=req)
    dec = ethabi.decode_abi(['uint32','uint32','uint32'], reqBytes)
  except Exception as e:
    print("DECODE FAILED", e)

  # dec[0] is a length parameter, ignored for this test
  answer = int(dec[1]*dec[1]) + int(dec[2]*dec[2])
  print("  -> SumSquares calculated", answer)

  try:
    e1 = ethabi.encode_abi(['uint32','uint32'], [32, answer])
    enc = Web3.toHex(e1)
  except Exception as e:
    print("ENCODE FAILED", e)
  print("  -> SumSquares responding", enc)

  return enc

# Expect a ABI-encoded uint32 'x' and 'y'. Return (x^2 + y^2)
def sumSquares_v2(req):
  print("  -> SumSquares_v2 handler called with",req)

  try:
    reqBytes = Web3.toBytes(hexstr=req)
    dec = ethabi.decode_abi(['uint32','uint32'], reqBytes)
  except Exception as e:
    print("DECODE FAILED", e)

  answer = int(dec[0]*dec[0]) + int(dec[1]*dec[1])
  print("  -> SumSquares calculated", answer)

  try:
    e1 = ethabi.encode_abi(['uint32'], [answer])
    enc = Web3.toHex(e1)
  except Exception as e:
    print("ENCODE FAILED", e)
  print("  -> SumSquares responding", enc)

  return enc

# ----------------------------------------------------------
# Main code for the local webserver

def offchain(*args):
  print("  -> offchain handler called with", args)
  print("a0",args[0])
  selector = args[0][:10]
  rest = args[0][10:]
  print("sel",selector)

  if selector == "0x4dde17c3":
    return add5(rest)
  elif selector == "0xf69de8e1":
    return chicken(rest)
  elif selector == "0x61f334d3":  # Register endpoint
    return "0x0000000000000000000000000000000000000000000000000000000000000000"
  elif selector == "0x00000000": # A hack to handle legacy requests using the same server process.
    if len(args[0]) == 130:
      return sumSquares_v2(args[0])
    else:
      return sumSquares_v1(args[0])
  else:
    print("Unknown", args[0][0:10])

  return ""

def ocReg(*args):
  print(" *** ocReg", args)

  return "0x0000000000000000000000000000000000000000000000000000000000000000"

class RequestHandler(SimpleJSONRPCRequestHandler):
  rpc_paths = ('/', '/hc', '/SoS_v1', '/SoS_v2')

def server_loop(contractAddr):
  print ("Registering method", contractAddr)
  server = SimpleJSONRPCServer(('192.168.4.2', 1234), requestHandler=RequestHandler)
  server.register_function(offchain, contractAddr)
  # Registration events are sent from the HCHelper contract, which is predeployed
  # to the 0x4200...03E9 address.
  server.register_function(ocReg, "0x42000000000000000000000000000000000003E9")
  server.serve_forever()

# ----------------------------------------------------------
# RPC endpoints and contracts

with open("../../.devnet/addresses.json") as f:
  config = json.loads(f.read())

w3 = Web3(Web3.HTTPProvider("http://127.0.0.1:8545"))
assert (w3.isConnected)
w3.middleware_onion.inject(geth_poa_middleware, layer=0)

l2 = Web3(Web3.HTTPProvider("http://127.0.0.1:9545"))
assert (l2.isConnected)
l2.middleware_onion.inject(geth_poa_middleware, layer=0)

# Test account created for local Boba devnet (Need to fix the deployment to fund this one on L1;
# for now using a shared test account
#addr=Web3.toChecksumAddress("0xb0bA04c08d8f1471bcA20C12a64DcCa17B01d96f")
#key="c9776e5eb09b348dfde140019e21142503d3c2a5c6d2019d0b30f5099ff2c8dd"
addr=Web3.toChecksumAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
key="ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

# Hardhat test account which is Owner of the legacy TuringCredit contract (needed to update price)
legacyOwner = Web3.toChecksumAddress("0xa0Ee7A142d267C1f36714E4a8F75612F20a79720")
legacyOwnerKey = 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6

with open("../../packages/contracts-bedrock/forge-artifacts/OptimismMintableERC20.sol/OptimismMintableERC20.json") as f:
  abi = json.loads(f.read())['abi']
boba_l2 = l2.eth.contract(address=Web3.toChecksumAddress("0x4200000000000000000000000000000000000023"), abi=abi)

with open("../../packages/contracts-bedrock/forge-artifacts/BobaHCHelper.sol/BobaHCHelper.json") as f:
  abi = json.loads(f.read())['abi']
hc = l2.eth.contract(address=Web3.toChecksumAddress("0x42000000000000000000000000000000000003E9"), abi=abi)

with open("../../packages/contracts-bedrock/forge-artifacts/BobaTuringCredit.sol/BobaTuringCredit.0.8.15.json") as f:
  abi = json.loads(f.read())['abi']
legacyCredit = l2.eth.contract(address=Web3.toChecksumAddress("0x42000000000000000000000000000000000003e8"), abi=abi)

with open("./artifacts/contracts/HCDemo.sol/HCDemo.json") as f:
  demoJson = json.loads(f.read())

with open("../../packages/contracts-bedrock/forge-artifacts/L1StandardBridge.sol/L1StandardBridge.json") as f:
  abi = json.loads(f.read())['abi']
l1sb = w3.eth.contract(address=config['L1StandardBridgeProxy'], abi=abi)

with open("../../packages/contracts-bedrock/forge-artifacts/BOBA.sol/BOBA.json") as f:
  abi = json.loads(f.read())['abi']
boba_l1 = w3.eth.contract(address=config['BOBA'], abi=abi)

# ----------------------------------------------------------
# Utility functions

def signAndSend(tx, label, gasEstimate):
  balStart = l2.eth.getBalance(addr)
  signed_tx = l2.eth.account.sign_transaction(tx, key)
  ret = l2.eth.send_raw_transaction(signed_tx.rawTransaction)
  print("Submitted {} tx, id={}".format(label,Web3.toHex(ret)))
  rcpt = l2.eth.wait_for_transaction_receipt(ret, poll_latency=0.75)
  balEnd = l2.eth.getBalance(addr)
  if gasEstimate:
    print("Got receipt in block {} status {}, gasUsed {}/{} ({} leftover)".format(
      rcpt.blockNumber,
      rcpt.status,
      rcpt.gasUsed,
      gasEstimate,
      gasEstimate - rcpt.gasUsed))
  else:
    print("Got receipt in block {} status {}, gasUsed {}".format(rcpt.blockNumber, rcpt.status, rcpt.gasUsed))
  totFee = balStart - balEnd
  l1Fee = Web3.toInt(hexstr=rcpt.l1Fee)
  l2Fee = rcpt.gasUsed * rcpt.effectiveGasPrice

  extraStr = ""
  if totFee > (l1Fee + l2Fee):
    extraStr = "extra {}".format(Web3.fromWei(totFee - l1Fee - l2Fee, 'gwei'))
  elif totFee < (l1Fee + l2Fee):
    extraStr = "***UNDERPAID*** {}".format(Web3.fromWei((l1Fee + l2Fee) - totFee, 'gwei'))
  print("Fee paid (gwei): {} total, {} l1, {} l2 {}".format(
    Web3.fromWei(totFee, 'gwei'),
    Web3.fromWei(l1Fee, 'gwei'),
    Web3.fromWei(l2Fee, 'gwei'),
    extraStr))
  if rcpt.status != 1:
    print("RECEIPT:", rcpt)
  assert(rcpt.status == 1)
  if gasEstimate:
    assert (gasEstimate >= rcpt.gasUsed)
  return rcpt

# =========================================
# Setup - ensure accounts are funded and approved; deploy test contracts

if boba_l1.functions.allowance(addr,l1sb.address).call() == 0:
  print("Approving bridge contract")
  tx = boba_l1.functions.approve(l1sb.address, Web3.toInt(hexstr="0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")).buildTransaction({
      'nonce': w3.eth.get_transaction_count(addr),
      'from':addr,
      'chainId': 900
     })
  signed_txn =w3.eth.account.sign_transaction(tx, key)
  ret = w3.eth.send_raw_transaction(signed_txn.rawTransaction)
  rcpt = w3.eth.wait_for_transaction_receipt(ret)
  assert(rcpt.status == 1)
  print("Approval done")

balCheck = l2.eth.getBalance(addr)
balCheck2 = boba_l2.functions.balanceOf(addr).call()
print("Starting balances", Web3.fromWei(balCheck,'ether'), "l2_eth", Web3.fromWei(balCheck2,'ether'), "l2_boba")

if balCheck == 0:
  tx = {
      'nonce': w3.eth.get_transaction_count(addr),
      'from':addr,
      'to':config['OptimismPortalProxy'],
      'gas':210000,
      'chainId': 900,
      'value': Web3.toWei(10,'ether')
  }
  if w3.eth.gasPrice > 1000000:
    tx['gasPrice'] = w3.eth.gasPrice
  else:
    tx['gasPrice'] = Web3.toWei(1, 'gwei')

  signed_txn =w3.eth.account.sign_transaction(tx, key)
  ret = w3.eth.send_raw_transaction(signed_txn.rawTransaction)
  print("\nSubmitted ETH TX", Web3.toHex(ret))
  rcpt = w3.eth.wait_for_transaction_receipt(ret)
  print("Got ETH receipt in block", rcpt.blockNumber, "status", rcpt.status, "gasPrice", rcpt.effectiveGasPrice)
  assert(rcpt.status == 1)

if balCheck2 == 0:
  tx = l1sb.functions.depositERC20(
     boba_l1.address,
     boba_l2.address,
     Web3.toWei(10,'ether'),
     4000000,
     "").buildTransaction({
       'nonce': w3.eth.get_transaction_count(addr),
       'from':addr,
       'chainId': 900
     })
  tx['gas'] = int(w3.eth.estimate_gas(tx) * 1.5)

  signed_txn =w3.eth.account.sign_transaction(tx, key)
  ret2 = w3.eth.send_raw_transaction(signed_txn.rawTransaction)
  print("Submitted BOBA TX", Web3.toHex(ret2))

  rcpt = w3.eth.wait_for_transaction_receipt(ret2)
  print("Got BOBA receipt in block", rcpt.blockNumber, "status", rcpt.status, "gasPrice", rcpt.effectiveGasPrice)
  assert(rcpt.status == 1)

while balCheck == 0 or balCheck2 == 0:  # Could check for low balance and top up, but this works for simple tests
  print("Waiting...")
  time.sleep(2)
  balCheck = l2.eth.getBalance(addr)
  balCheck2 = boba_l2.functions.balanceOf(addr).call()

approveTx = boba_l2.functions.approve(hc.address, Web3.toWei(1000000,'ether')).buildTransaction({
       'nonce': l2.eth.get_transaction_count(addr),
       'from':addr,
       'chainId': 901,
       'gas':1000000,
       'maxFeePerGas':Web3.toWei(10, 'gwei'),
})
signAndSend(approveTx,"ApproveToken", None)

approveTx = boba_l2.functions.approve(legacyCredit.address, Web3.toWei(1000000,'ether')).buildTransaction({
       'nonce': l2.eth.get_transaction_count(addr),
       'from':addr,
       'chainId': 901,
       'gas':1000000,
       'maxFeePerGas':Web3.toWei(10, 'gwei'),
})
signAndSend(approveTx,"ApproveToken (legacy)", None)

print("Setting legacy TuringPrice")
# This has to be done by the owner of the contact.
legacyPriceTx = legacyCredit.functions.updateTuringPrice(1001).buildTransaction({
       'nonce': l2.eth.get_transaction_count(legacyOwner),
       'from':legacyOwner,
       'chainId': 901,
       'gas':1000000,
       'maxFeePerGas':Web3.toWei(10, 'gwei'),
})
signed_tx = l2.eth.account.sign_transaction(legacyPriceTx, legacyOwnerKey)
ret = l2.eth.send_raw_transaction(signed_tx.rawTransaction)
rcpt = l2.eth.wait_for_transaction_receipt(ret, poll_latency=0.75)
assert(rcpt.status == 1)

demoFactory = l2.eth.contract(abi=demoJson['abi'], bytecode=demoJson['bytecode'])
deployTx = demoFactory.constructor(hc.address).buildTransaction({
  'from':addr,
  'nonce':l2.eth.get_transaction_count(addr),
  'chainId': 901,
  'gas':1000000,
  'maxFeePerGas':Web3.toWei(10, 'gwei'),
})
rcpt = signAndSend(deployTx, "deploy", None)

demo = l2.eth.contract(abi=demoJson['abi'],address=rcpt.contractAddress)
cAddr = rcpt.contractAddress
print("Demo contract deployed to", cAddr)

serverProc = Process(target=server_loop,args=(cAddr,))
serverProc.start()
print("Server started")

T0 = time.time()
print("ownerRevenue", hc.functions.ownerRevenue().call(), legacyCredit.functions.ownerRevenue().call())

URL = "http://192.168.4.2:1234/hc"
RegHash = "0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563"
tx = hc.functions.RegisterEndpoint(URL, RegHash).buildTransaction({
       'nonce': l2.eth.get_transaction_count(addr),
       'from':addr,
       'chainId': 901,
       'gas':1222333,
       'maxFeePerGas':Web3.toWei(10, 'gwei'),
})
rcpt = signAndSend(tx, "Register", None)
ans = hc.events.EndpointRegistered().processReceipt(rcpt,errors=DISCARD)
print("Registration", ans)

print("ownerRevenue", hc.functions.ownerRevenue().call(), legacyCredit.functions.ownerRevenue().call())

# Register our Demo contract as a permitted caller of the endpoint
tx = hc.functions.AddPermittedCaller(URL, cAddr).buildTransaction({
       'nonce': l2.eth.get_transaction_count(addr),
       'from':addr,
       'chainId': 901,
       'gas':1000000,
       'maxFeePerGas':Web3.toWei(10, 'gwei'),
})
signAndSend(tx, "AddPermittedCaller", None)

# Deposit credits for Demo contract
tx = hc.functions.AddCredit(cAddr, 100000).buildTransaction({
       'nonce': l2.eth.get_transaction_count(addr),
       'from':addr,
       'chainId': 901,
       'gas':1000000,
       'maxFeePerGas':Web3.toWei(10, 'gwei'),
})
signAndSend(tx, "AddCredit", None)

tx = legacyCredit.functions.addBalanceTo(100000, cAddr).buildTransaction({
       'nonce': l2.eth.get_transaction_count(addr),
       'from':addr,
       'chainId': 901,
       'gas':1000000,
       'maxFeePerGas':Web3.toWei(10, 'gwei'),
})
signAndSend(tx, "AddCredit(legacy)", None)

# =========================================
# Tests start here

print("\nGetSimpleRandom starting")
tx = demo.functions.FlipCoin().buildTransaction({
       'nonce': l2.eth.get_transaction_count(addr),
       'from':addr,
       'chainId': 901,
       'gas':1000000,
       'maxFeePerGas':Web3.toWei(10, 'gwei'),
})
eg = l2.eth.estimate_gas(tx)
print("GasEstimate", eg)
rcpt = signAndSend(tx, "FlipCoin", eg)

ans = demo.events.CoinFlip().processReceipt(rcpt,errors=DISCARD)
print("CoinFlip event", ans[0].args)

print("\nSequentialRandom starting")
time.sleep(2)

tx = demo.functions.SeqRandom(1).buildTransaction({
       'nonce': l2.eth.get_transaction_count(addr),
       'from':addr,
       'chainId': 901,
       'gas':1000000,
       'maxFeePerGas':Web3.toWei(10, 'gwei'),
})
eg = l2.eth.estimate_gas(tx)
print("GasEstimate", eg)
rcpt = signAndSend(tx, "SeqRandom(1)", eg)
time.sleep(2)

tx = demo.functions.SeqRandom(2).buildTransaction({
       'nonce': l2.eth.get_transaction_count(addr),
       'from':addr,
       'chainId': 901,
       'gas':1000000,
       'maxFeePerGas':Web3.toWei(10, 'gwei'),
})
eg = l2.eth.estimate_gas(tx)
print("GasEstimate", eg)
rcpt = signAndSend(tx, "SeqRandom(2)", eg)
time.sleep(2)

tx = demo.functions.SeqRandom(2).buildTransaction({
       'nonce': l2.eth.get_transaction_count(addr),
       'from':addr,
       'chainId': 901,
       'gas':1000000,
       'maxFeePerGas':Web3.toWei(10, 'gwei'),
})
eg = l2.eth.estimate_gas(tx)
print("GasEstimate", eg)
rcpt = signAndSend(tx, "SeqRandom(3)", eg)

print("\nAdd5 (2) starting")
tx = demo.functions.Add5(2).buildTransaction({
       'nonce': l2.eth.get_transaction_count(addr),
       'from':addr,
       'chainId': 901,
       'gas':1000000,
       'maxFeePerGas':Web3.toWei(10, 'gwei'),
})

eg = l2.eth.estimate_gas(tx)
print("GasEstimate", eg)

rcpt = signAndSend(tx, "Add5", eg)
ans = demo.events.MathResponse().processReceipt(rcpt,errors=DISCARD)
print("MathResponse event: ", ans[0].args)

print("\nAdd5 (10) starting, no eth.estimate_gas")
tx = demo.functions.Add5(10).buildTransaction({
       'nonce': l2.eth.get_transaction_count(addr),
       'from':addr,
       'chainId': 901,
       'gas':1000000,
       'maxFeePerGas':Web3.toWei(10, 'gwei'),
})

rcpt = signAndSend(tx, "Add5", None)
ans = demo.events.MathResponse().processReceipt(rcpt,errors=DISCARD)
print("MathResponse event: ", ans[0].args)

print("ownerRevenue", hc.functions.ownerRevenue().call(), legacyCredit.functions.ownerRevenue().call())

print("\nChicken(1) starting")
tx = demo.functions.Chicken(1).buildTransaction({
       'nonce': l2.eth.get_transaction_count(addr),
       'from':addr,
       'chainId': 901,
       'gas':1000000,
       'maxFeePerGas':Web3.toWei(10, 'gwei'),
})
eg = l2.eth.estimate_gas(tx)
print("GasEstimate", eg)
rcpt = signAndSend(tx, "Chicken", eg)

ans = demo.events.StringLength().processReceipt(rcpt,errors=DISCARD)
print("StringLength event: ", ans[0].args)

# Large requests may expose bugs in L1 fee calculation
print("\nChicken(250) starting")
tx = demo.functions.Chicken(250).buildTransaction({
       'nonce': l2.eth.get_transaction_count(addr),
       'from':addr,
       'chainId': 901,
       'maxFeePerGas':Web3.toWei(10, 'gwei'),
})
eg = l2.eth.estimate_gas(tx)
print("GasEstimate", eg)
tx['gas'] = int(eg * 1.0)  # Can adjust for debugging
rcpt = signAndSend(tx, "Chicken", eg)

ans = demo.events.StringLength().processReceipt(rcpt,errors=DISCARD)
print("StringLength event: ", ans[0].args)

# ----------------------------------------------------------

print("\nLegacyV1 starting")
tx = demo.functions.LegacyV1().buildTransaction({
       'nonce': l2.eth.get_transaction_count(addr),
       'from':addr,
       'chainId': 901,
       'gas':1000000,
       'maxFeePerGas':Web3.toWei(10, 'gwei'),
})
eg = l2.eth.estimate_gas(tx)
print("GasEstimate", eg)
rcpt = signAndSend(tx, "LegacyV1", eg)

ans = demo.events.MathResponse().processReceipt(rcpt,errors=DISCARD)
print("MathResponse event: ", ans[0].args)

print("\nLegacyV2 starting")
tx = demo.functions.LegacyV2().buildTransaction({
       'nonce': l2.eth.get_transaction_count(addr),
       'from':addr,
       'chainId': 901,
       'gas':1000000,
       'maxFeePerGas':Web3.toWei(10, 'gwei'),
})
eg = l2.eth.estimate_gas(tx)
print("GasEstimate", eg)
rcpt = signAndSend(tx, "LegacyV2", eg)

ans = demo.events.MathResponse().processReceipt(rcpt,errors=DISCARD)
print("MathResponse event: ", ans[0].args)

# ----------------------------------------------------------
# Cleanup and shutdown

print("\nCleanup starting")

tx = hc.functions.RemovePermittedCaller(URL, cAddr).buildTransaction({
       'nonce': l2.eth.get_transaction_count(addr),
       'from':addr,
       'chainId': 901,
       'gas':1000000,
       'maxFeePerGas':Web3.toWei(10, 'gwei'),
})
signAndSend(tx, "RemovePermittedCaller", None)

tx = hc.functions.UnregisterEndpoint(URL).buildTransaction({
       'nonce': l2.eth.get_transaction_count(addr),
       'from':addr,
       'chainId': 901,
       'gas':1000000,
       'maxFeePerGas':Web3.toWei(10, 'gwei'),
})
signAndSend(tx, "Unregister", None)

print("ownerRevenue", hc.functions.ownerRevenue().call(), "new,", legacyCredit.functions.ownerRevenue().call(), "legacy")
serverProc.kill()

