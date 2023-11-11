#!python
#
# Copyright 2021-2023 mmontour@enya.ai. Internal use only; all rights reserved
#
# Basic fraud checker. Reads state roots from SCC and checks them against
# actual L2 values, reporting mismatches. Note - code was copied from
# stress_tester.py and contains a lot of irrelevant cruft to be cleaned up
# later.

import os
from web3 import Web3
import inspect
import time
import json
from ratelimiter import RateLimiter
import logging
from web3.middleware import geth_poa_middleware
from web3.logs import STRICT, IGNORE, DISCARD, WARN
from threading import Thread, Lock

# pip install jsonrpclib-pelix
from jsonrpclib.SimpleJSONRPCServer import PooledJSONRPCServer
from jsonrpclib.threadpool import ThreadPool

logging.basicConfig(format='%(levelname)s %(asctime)s %(message)s', datefmt='%Y%m%dT%H%M%S')
logger = logging.getLogger('fraud-detector')
logger.setLevel(logging.DEBUG)
#logger.addHandler(logging.StreamHandler())

try:
  logger.debug ("L1_NODE_WEB3_URL = {}".format(os.environ['L1_NODE_WEB3_URL']))
  logger.debug ("L2_NODE_WEB3_URL = {}".format(os.environ['L2_NODE_WEB3_URL']))
  logger.debug ("VERIFIER_WEB3_URL = {}".format(os.environ['VERIFIER_WEB3_URL']))
  logger.debug ("RATE_LIMITER_MAX_CALLS = {}".format(os.environ['RATE_LIMITER_MAX_CALLS']))
  logger.debug ("RATE_LIMITER_PERIOD = {}".format(os.environ['RATE_LIMITER_PERIOD']))
  logger.debug ("L2_CHECK_INTERVAL = {}".format(os.environ['L2_CHECK_INTERVAL']))
  logger.debug ("L2OO_CONTRACT_ADDRESS = {}".format(os.environ['L2OO_CONTRACT_ADDRESS']))
except:
  logger.error("Missing required environment variable(s)")
  exit(1)

MAX_CALLS = int(os.environ['RATE_LIMITER_MAX_CALLS'])
PERIOD = int(os.environ['RATE_LIMITER_PERIOD'])
MAX_ATTEMPTS = 3

# To reduce load on the mainnet L2 when many users are running fraud detectors, the L2
# block is only queried once per <interval> seconds. The primary concern with respect
# to fraud is whether or not the output roots published to L1 correspond to the ones
# derived locally by the Verifier
l2_check_interval = int(os.environ['L2_CHECK_INTERVAL'])
last_l2check = 0

Matched = {
  'Index':0, # Highest good block, and its corresponding state root
  'Root':"0x",
  'Time':time.time(), # Local timestamp
  'is_ok':True # Set false once a mismatch is found. This disables checkpointing
}
matchedLock = Lock()

def getRPCRateLimiter():
  @RateLimiter(max_calls=MAX_CALLS, period=PERIOD)
  def rateLimiter(method, *args, **kwargs):
      attempts = 0
      while True:
        try:
          return method(*args, **kwargs)
        except Exception as e:
          if attempts >= 3:
            logger.error("An error occured in fraud-detector: %s", str(e))
            return str(e)
          else:
            logger.error("An error occured in fraud-detector: %s \n %s", str(inspect.getouterframes(inspect.currentframe())[-2].frame), str(e))
            logger.info("Will try again...")
            attempts += 1
            time.sleep(3)
  return rateLimiter

def status(*args):
  global Matched
  matchedLock.acquire()
  status = {
    'matchedBlock':Matched['Index'],
    'matchedRoot':Matched['Root'],
    'timestamp':Matched['Time'],
    'isOK':Matched['is_ok']
  }
  matchedLock.release()
  return status

def do_checkpoint():
  global Matched
  global checkpoint_index

  matchedLock.acquire()
  isOK = Matched['is_ok']
  checkpoint_index = Matched['Index']
  matchedLock.release()

  try:
    if isOK:
      with open("./db/checkpoint.dat", "w") as f:
        f.write(json.dumps([checkpoint_index]))
  except:
    logger.error("Error writing checkpoint.dat")     
# #########################

checkpoint_index = -1
rpc = [None]*4
rpcRateLimters = [getRPCRateLimiter() if i > 0 else None for i in range(4)]

try:
 with open("./db/checkpoint.dat", "r") as f:
    checkpoint_index = json.loads(f.read())[0]
    logger.info("Starting at checkpoint index " + str(checkpoint_index))
except:
  pass

while True:
  try:
    rpc[1] = Web3(Web3.HTTPProvider(os.environ['L1_NODE_WEB3_URL']))
    assert (rpcRateLimters[1](rpc[1].is_connected))
    break
  except:
    logger.info ("Waiting for L1...")
    time.sleep(10)

rpcRateLimters[1](rpc[1].middleware_onion.inject, geth_poa_middleware, layer=0)
logger.debug("Connected to L1_NODE_WEB3_URL")

while True:
  try:
    rpc[2] = Web3(Web3.HTTPProvider(os.environ['L2_NODE_WEB3_URL']))
    assert (rpcRateLimters[2](rpc[2].is_connected))
    break
  except:
    logger.info ("Waiting for L2...")
    time.sleep(10)

rpcRateLimters[2](rpc[2].middleware_onion.inject, geth_poa_middleware, layer=0)
logger.debug("Connected to L2_NODE_WEB3_URL")

while True:
  try:
    rpc[3] = Web3(Web3.HTTPProvider(os.environ['VERIFIER_WEB3_URL']))
    assert (rpcRateLimters[3](rpc[3].is_connected))
    break
  except:
    logger.info ("Waiting for verifier...")
    time.sleep(10)

rpcRateLimters[3](rpc[3].middleware_onion.inject, geth_poa_middleware, layer=0)
logger.debug("Connected to VERIFIER_WEB3_URL")

def loadContract(rateLimiter, rpc, addr, abiPath):
  with open(abiPath) as f:
    abi = json.loads(f.read())['abi']
  c = rateLimiter(rpc.eth.contract, address=addr, abi=abi)
  return c

l2oo = loadContract(rpcRateLimters[1], rpc[1], os.environ['L2OO_CONTRACT_ADDRESS'], './contracts/L2OutputOracle.json')

# Check that these are consistent
startingBlockNumber = l2oo.functions.startingBlockNumber().call()
startingTimestamp = l2oo.functions.startingTimestamp().call()
b = rpc[2].eth.get_block(startingBlockNumber)
assert(b.timestamp == startingTimestamp)

def checkBlocks(bn):
  b_seq = rpcRateLimters[2](rpc[2].eth.get_block, bn)
  b_vfy = rpc[3].eth.get_block(bn)
  
  if b_seq.hash != b_vfy.hash:
    logger.warn("L2 blockhash mismatch at {}, sequencer {} verifier {}".format(bn, Web3.to_hex(b_seq.hash), Web3.to_hex(b_vft.hash)))
    return False
  if b_seq.stateRoot != b_vfy.stateRoot:
    logger.warn("L2 stateRoot mismatch at {}, sequencer {} verifier {}".format(bn, Web3.to_hex(b_seq.stateRoot), Web3.to_hex(b_vft.stateRoot)))
    return False
  return True

def checkProposals(fromIdx, toIdx, checkL2):
  global l2_check_interval,last_l2check
  match = ""

  for i in range(fromIdx, toIdx+1):
    op = rpcRateLimters[1](l2oo.functions.getL2Output(i).call)
    bn = op[2]

    while rpc[3].eth.block_number < bn:
      logger.debug("Waiting for verifier to catch up, block number {} / {}".format(rpc[3].eth.block_number, bn))
      time.sleep(10)
    b = rpc[3].eth.get_block(bn)
    proof = rpc[3].eth.get_proof("0x4200000000000000000000000000000000000016", [], bn)
    mpsr = proof.storageHash
    sr = b.stateRoot
    bh = b.hash

    orp = "0x0000000000000000000000000000000000000000000000000000000000000000" + Web3.to_hex(sr)[2:] + Web3.to_hex(mpsr)[2:] + Web3.to_hex(bh)[2:]
    calcRoot = Web3.keccak(hexstr=orp)
    now = time.time()
    if(calcRoot != op[0]):
      match = "***MISMATCH*** Output Root"
    elif checkL2 and (now > last_l2check + l2_check_interval):
      if not checkBlocks(bn):
        match = "***MISMATCH*** Block Hashes"
      last_l2check = now

    log_str = "{} {} {} {} {}".format(i, bn, Web3.to_hex(op[0]), Web3.to_hex(calcRoot), match)
    matchedLock.acquire()
    if match == "" and Matched['is_ok']:
      Matched['Index'] = i
      Matched['Root'] = Web3.to_hex(op[0])
      Matched['Time'] = time.time()
      logger.info(log_str)
    else:
      Matched['is_ok'] = False
      logger.warning(log_str)
    matchedLock.release()

# ###############################

def server_loop():
  notify_pool = ThreadPool(max_threads=10, min_threads=0)
  request_pool = ThreadPool(max_threads=50, min_threads=10)
  notify_pool.start()
  request_pool.start()

  ws = PooledJSONRPCServer(('', 8555), thread_pool=request_pool)
  ws.set_notification_pool(notify_pool)
  ws.register_function(status)

  try:
    ws.serve_forever()
  finally:
    request_pool.stop()
    notify_pool.stop()
    ws.set_notification_pool(None)

def fpLoop():
  toIdx = rpcRateLimters[1](l2oo.functions.latestOutputIndex().call)
  while toIdx < checkpoint_index:
    logger.warn("Waiting for index to advance beyond checkpoint {} / {}".format(toIdx, checkpoint_index))
    time.sleep(10)
    toIdx = rpcRateLimters[1](l2oo.functions.latestOutputIndex().call)
  if toIdx > checkpoint_index:
    checkProposals(checkpoint_index + 1, toIdx, False) # L2 not checked when catching up on old proposals, to reduce RPC traffic
    do_checkpoint()
  lastIdx = toIdx

  logger.debug("Reached latest index {}, waiting for it to advance".format(checkpoint_index))
  while True:
    time.sleep(10)
    toIdx = rpcRateLimters[1](l2oo.functions.latestOutputIndex().call)
    if toIdx < lastIdx:
      logger.warn("latestOutputIndex has decreased, was {} now {}".format(lastIdx, toIdx))
    elif toIdx > lastIdx:
      checkProposals(lastIdx + 1, toIdx, True)
      do_checkpoint() # These update infrequently enough that we can checkpoint each time
    lastIdx = toIdx
  logger.info("Exiting")

server_thread = Thread(target=server_loop,args=())
server_thread.start()

fpLoop()
