import argparse
import logging
import os
import subprocess
import json
import socket
import calendar
import datetime
import time
import shutil
import http.client
import gzip
from multiprocessing import Process, Queue
import concurrent.futures
from collections import namedtuple

pjoin = os.path.join

parser = argparse.ArgumentParser(description='Bedrock devnet AA setup')
parser.add_argument('--monorepo-dir', help='Directory of the monorepo', default=os.getcwd())
parser.add_argument('--test', help='Tests the deployment, must already be deployed', type=bool, action=argparse.BooleanOptionalAction)

log = logging.getLogger()

class Bunch:
    def __init__(self, **kwds):
        self.__dict__.update(kwds)

def main():
    args = parser.parse_args()

    monorepo_dir = os.path.abspath(args.monorepo_dir)
    devnet_dir = pjoin(monorepo_dir, '.devnet')
    ops_bedrock_dir = pjoin(monorepo_dir, 'ops-bedrock')

    paths = Bunch(
        mono_repo_dir=monorepo_dir,
        devnet_dir=devnet_dir,
        ops_bedrock_dir=ops_bedrock_dir,
        addresses_json_path=pjoin(devnet_dir, 'addresses.json'),
        aa_addresses_json_path=pjoin(devnet_dir, 'aa_addresses.json'),
    )

    if args.test:
      log.info('Testing aa deployment')
      aa_test()
      return

    aa_deploy(paths)

def aa_deploy(paths):
    wait_up(9545)
    wait_for_rpc_server('127.0.0.1:9545')
    wait_up(8545)
    wait_for_rpc_server('127.0.0.1:8545')

    addresses = read_json(paths.addresses_json_path)
    L1CrossDomainMessengerAddress = addresses['L1CrossDomainMessengerProxy']
    L1StandardBridgeAddress = addresses['L1StandardBridgeProxy']

    docker_env = {
        'PWD': paths.ops_bedrock_dir,
        'L1_CROSS_DOMAIN_MESSENGER_ADDRESS': L1CrossDomainMessengerAddress,
        'L1_STANDARD_BRIDGE_ADDRESS': L1StandardBridgeAddress,
    }
    run_command(['docker', 'compose', '-f', 'docker-compose.yml', '-f', 'docker-compose-side.yml', 'build', 'aa_deployer', 'bundler'], cwd=paths.ops_bedrock_dir, env=docker_env)

    if os.path.exists(paths.aa_addresses_json_path):
        log.info('AA already deployed')
    else:
        log.info('Deploying AA contracts')
        run_command(['docker', 'compose', '-f', 'docker-compose.yml', '-f', 'docker-compose-side.yml', 'up', 'aa_deployer'], cwd=paths.ops_bedrock_dir, env=docker_env)

    log.info('Bring up bundler')
    aa_addresses = read_json(paths.aa_addresses_json_path)
    AddressManagerAddress = aa_addresses['L2_AddressManager']
    docker_env['ADDRESS_MANAGER_ADDRESS'] = AddressManagerAddress
    run_command(['docker', 'compose', '-f', 'docker-compose.yml', '-f', 'docker-compose-side.yml', 'up', '-d', 'bundler'], cwd=paths.ops_bedrock_dir, env=docker_env)

    # Fin.
    log.info('AA ready.')

def aa_test():
    wait_up(3000)
    wait_for_rpc_server('127.0.0.1:3000/rpc')

def run_command(args, check=True, shell=False, cwd=None, env=None, timeout=None):
    env = env if env else {}
    return subprocess.run(
        args,
        check=check,
        shell=shell,
        env={
            **os.environ,
            **env
        },
        cwd=cwd,
        timeout=timeout
    )

def wait_up(port, retries=10, wait_secs=1):
    for i in range(0, retries):
        log.info(f'Trying 127.0.0.1:{port}')
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        try:
            s.connect(('127.0.0.1', int(port)))
            s.shutdown(2)
            log.info(f'Connected 127.0.0.1:{port}')
            return True
        except Exception:
            time.sleep(wait_secs)

    raise Exception(f'Timed out waiting for port {port}.')

def wait_for_rpc_server(url):
    log.info(f'Waiting for RPC server at {url}')

    headers = {'Content-type': 'application/json'}
    body = '{"id":1, "jsonrpc":"2.0", "method": "eth_chainId", "params":[]}'

    while True:
        try:
            conn = http.client.HTTPConnection(url)
            conn.request('POST', '/', body, headers)
            response = conn.getresponse()
            if response.status < 300:
                log.info(f'RPC server at {url} ready')
                return
        except Exception as e:
            log.info(f'Waiting for RPC server at {url}')
            time.sleep(1)
        finally:
            if conn:
                conn.close()

def read_json(path):
    with open(path, 'r') as f:
        return json.load(f)
