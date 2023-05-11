//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.15;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
// import "./ITuringHelper.sol";
import { Burn } from "../libraries/Burn.sol";

contract BobaHCHelper /*is Ownable*/ {
  //using SafeMath for uint256;
  using SafeERC20 for IERC20;

  address immutable HelperAddr = 0x42000000000000000000000000000000000000Fd;
  address immutable OffchainCaller = 0xdEAddEadDeaDDEaDDeadDeAddeadDEaddeaD9901;
  address immutable hcToken = 0x42000000000000000000000000000000000000FE;	// Boba token

  // Cost (in hcToken) for various operations
  uint256 constant RegCost = 100;  // Register an endpoint
  uint256 constant CallCost = 50;  // Call off-chain
  uint256 constant RNGCost = 30;   // Use a random-number function

  // Collected balance
  uint256 public ownerRevenue;

  mapping(address => uint256) public prepaidCredit;
  mapping(address => uint256) public pendingCharge;

  struct EndpointEntry {
  	address Owner;
	mapping (address => bool) PermittedCallers;
  }
  mapping (bytes32 => EndpointEntry) public Endpoints;

  mapping(bytes32 => bytes) OffchainResponses;


  // events 
  event DBG(bytes b);

  constructor () {
   }


  function CallOffchain(string calldata _url, bytes calldata _payload) public returns (bool, bytes memory) {
    uint32  method = 0xc1fd7e46; // "CallOffchain(string,bytes)" signature
    bytes32 EK = keccak256(abi.encodePacked(_url));

    if (msg.sender != address(this)) {
      require (Endpoints[EK].Owner != address(0), "Endpoint not registered");
      require (Endpoints[EK].PermittedCallers[msg.sender], "Caller is not permitted");
    }

    bytes32 key = keccak256(abi.encodePacked(method, msg.sender, abi.encodePacked(_url), _payload));

    return GetResponse(key, msg.sender, CallCost);
  }

  function SimpleRandom() public returns (uint256) {
    uint32 method = 0xc6e4b0d3; // "SimpleRandom()" signature
    bool success;
    bytes memory responsedata;
    bytes32 key = keccak256(abi.encodePacked(method, msg.sender));
    (success, responsedata) = GetResponse(key, msg.sender, RNGCost);
    require(success && (responsedata.length == 32), "Random number generation failed");

    uint256 result = abi.decode(responsedata,(uint256));
    
    return result;
  }

  function SequentialRandom(bytes32 session, bytes32 nextHash, uint256 myNum) public returns (bytes32, uint256) {
    uint32  method = 0x32be428f; // "SequentialRandom(bytes32,bytes32,uint256)" signature

    bool success;
    bytes memory responsedata;
    bytes32 key = keccak256(abi.encodePacked(method, msg.sender, session, nextHash, myNum));
    (success, responsedata) = GetResponse(key, msg.sender, RNGCost);

    require(success && (responsedata.length == 64), "Random number generation failed");
    
    return abi.decode(responsedata,(bytes32,uint256));
  }

  // -------------------------------------------------------------------------

  function RegisterEndpoint(string calldata _url, bytes32 _auth)
    public returns (bool) {
    
    BobaHCHelper Self = BobaHCHelper(HelperAddr);
    bytes32 EK = keccak256(abi.encodePacked(_url));

    IERC20(hcToken).safeTransferFrom(msg.sender, address(this), RegCost);
    ownerRevenue += RegCost;

    bytes memory request = abi.encodeWithSignature("RegisterV1()");
    bytes memory response; 
    bool success;

    (success, response) = Self.CallOffchain(_url, request);

    if (success && response.length == 32 && keccak256(response) == _auth) {
      Endpoints[EK].Owner = msg.sender;
    } else {
      success = false;
    }

    return success;
  }

  function AddPermittedCaller(string calldata _url, address _callerAddress) public {
    bytes32 EK = keccak256(abi.encodePacked(_url));
    require(Endpoints[EK].Owner != address(0), "Endpoint is not registered");
    require(Endpoints[EK].Owner == msg.sender, "Not the Endpoint owner");

    Endpoints[EK].PermittedCallers[_callerAddress] = true;
  }

  function RemovePermittedCaller(string calldata _url, address _callerAddress) public {
    bytes32 EK = keccak256(abi.encodePacked(_url));
    require(Endpoints[EK].Owner != address(0), "Endpoint is not registered");
    require(Endpoints[EK].Owner == msg.sender, "Not the Endpoint owner");

    Endpoints[EK].PermittedCallers[_callerAddress] = false;
  }

  function CheckPermittedCaller(string calldata _url, address _callerAddress) public view returns (bool) {
    bytes32 EK = keccak256(abi.encodePacked(_url));
    require(Endpoints[EK].Owner != address(0), "Endpoint is not registered");
    return(Endpoints[EK].PermittedCallers[_callerAddress]);
  }

  function AddCredit(address _caller, uint256 _amount) public {
    require(_amount > 0, "Invalid amount");

    IERC20(hcToken).safeTransferFrom(msg.sender, address(this), _amount);
    prepaidCredit[_caller] += _amount; 
  }

  // -------------------------------------------------------------------------

  function GetResponse(bytes32 _cacheKey, address _caller, uint256 _cost) internal returns (bool, bytes memory) {
    if (_caller != address(this)) {
      uint256 creditCheck = _cost + pendingCharge[_caller];
      require (prepaidCredit[_caller] >= creditCheck, "Insufficient credit");

      prepaidCredit[_caller] -= _cost;
      ownerRevenue += _cost;
      // TODO - could extend this to also charge credits which would be paid to the endpoint owner.
    }

    bytes memory cachedResponse = OffchainResponses[_cacheKey];
    require(cachedResponse.length > 0, "GetResponse: Missing cache entry"); // Trigger string; don't edit.

    delete(OffchainResponses[_cacheKey]);
    return abi.decode(cachedResponse,(bool, bytes));
  }

  function PutResponse(bytes32 _cacheKey, bool _success, bytes calldata _returndata) public {
    require(msg.sender == OffchainCaller, "Invalid PutResponse() caller");

    // Eventually this should just overwrite
    require(OffchainResponses[_cacheKey].length == 0, "DEBUG - Already exists");
    OffchainResponses[_cacheKey] = abi.encode(_success, _returndata);
  }

  
}
