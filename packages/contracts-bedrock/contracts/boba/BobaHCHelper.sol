//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

//import '@openzeppelin/contracts/access/Ownable.sol';

// import "./ITuringHelper.sol";
import { Burn } from "../libraries/Burn.sol";

contract BobaHCHelper /*is Ownable*/ {

  address immutable HelperAddr = 0x42000000000000000000000000000000000000Fd;
  address immutable OffchainCaller = 0xdEAddEadDeaDDEaDDeadDeAddeadDEaddeaD9901;

  event OffchainRandom(uint version, uint256 random);
  event DBG(bytes b);

  mapping(bytes32 => bytes) OffchainResponses;

/*
  // This protects your own credits for this helper contract
  mapping(address => bool) public permittedCaller;

  event AddPermittedCaller(address _callerAddress);
  event RemovePermittedCaller(address _callerAddress);
  event CheckPermittedCaller(address _callerAddress, bool permitted);
  event OffchainResponse(uint version, bytes responseData);
  event Offchain42(uint version, uint256 random);


  modifier onlyPermittedCaller() {
    require(
      permittedCaller[msg.sender],
      'Invalid Caller Address'
    );
    _;
  }
*/


  constructor () {
   }

  // -------------------------------------------------------------------------
  
  /* Internal functions which provide the interface to core/vm/hybrid_compute.go. GetResponse()
    is called from other functions in this contract while PutResponse() and ExpireResponse() are
    called from Offchain transactions inserted by the sequencer node.
    
    The first time that GetResponse() is called, it will revert on a "missing cache entry". On a
    successful interaction, the offchain interactions will take place in the background and then
    the user transaction will be re-run. This second time it will find a cache entry and return
    the result. On a failure, the user transaction will be re-run a second time without any special
    logic, so that the "missing cache entry" is passed back to the caller.
    
    TODO - This could be re-written to provide an in-band error response by adding a success/failure
    flag to the cache entry.
  */
  
  // Try to get a previously-populated response
  function GetResponse(uint32 _rType, address _caller, string memory _url, bytes memory _payload)
    public returns (bytes memory) {

    // This function may only be called from the Helper contract
    require (msg.sender == address(this), "Invalid GetResponse() caller");

    bytes32 cacheKey = keccak256(abi.encodePacked(_rType, _caller, _url,  _payload));

    // This specific revert string triggers the offchain mechanism. Do not change.
    require(OffchainResponses[cacheKey].length > 0, "GetResponse: Missing cache entry");

    bytes memory response = OffchainResponses[cacheKey];

    // We burn additional L2 gas here to reflect the cost of storing the Response data in the L1 rollup.
    // TODO - calculate (L1_Gas_per_byte * Length) * (L1_Gas_Price / L2_Gas_Price). 
    uint256 gasToBurn = response.length;
    Burn.gas (gasToBurn * 5000);
    
    delete(OffchainResponses[cacheKey]);
    return response;
  }
  
  // Populate the cache ahead of an anticipated GetResponse(), charging Boba tokens.
  function PutResponse(bytes32 cacheKey, bytes memory _payload) public {
    // Reserved address for Offchain transactions (core/types/offchain_tx.go)
    require(msg.sender == OffchainCaller, "Invalid PutResponse() caller");
    // Eventually this should just overwrite
    require(OffchainResponses[cacheKey].length == 0, "DEBUG - Already exists");
    OffchainResponses[cacheKey] = _payload;
  }

  // This function will transfer a Boba token payment but will not store the _payload.
  // The sequencer will call this in the event that a user has called eth_estimateGas() to
  // populate the cache but has not submitted a real transaction within a specified time
  // interval. Since the offchain work has already been performed, this is how we bill for
  // it in the absence of a regular transaction.
  // Because there's no way to burn gas for the L1 calldata storage in this scenario, we
  // instead charge more Boba tokens than would be collected on a completed transaction.
  function ExpiredResponse(bytes32 cacheKey /*... + caller information */ ) public {
    // FIXME - consume tokens

    // Reserved address for Offchain transactions (core/types/offchain_tx.go)
    require(msg.sender == OffchainCaller, "Invalid PutResponse() caller");

    // This function can also be used to clean up the state after a reverted GetResponse transaction.
    // In this case we do not need to charge any additional Boba here because that was paid in the PutResponse.    
    delete(OffchainResponses[cacheKey]);
  }

  // -------------------------------------------------------------------------

  /* These functions are called to manage endpoints */

  // This is called to register an offchain endpoint. Registration requires
  // that endpoint to respond to an RPC call from the Helper contract, to
  // prove that the endpoint is under the control of the person attempting to
  // register it. The response to that challenge must be a bytes32 value which
  // is the keccak256 pre-image of "auth".
  //
  // This function also takes a parameter for initial funding. This must be
  // enough to cover the credit cost of the registration itself but may be
  // larger. The caller must have previously authorized the ERC20 token transfer
  // to the helper contract's address.
  //
  // Upon successful registration, the provided URL will be stored in a map
  // indexed by its Endpoint Key (EK) which is keccak256(abi.encodePacked(url)).
  // If a previous registration existed for that EK it will be overwritten.
  // 
  // The address which calls RegisterEndpoint() is registered as its owner, 
  // and only that owner will be permitted to unregister it later.
  
  struct EndpointEntry {
  	address Owner;
	string URL;
	mapping (address => bool) PermittedCallers;
  }

  mapping (bytes32 => EndpointEntry) public Endpoints;
  
  // Emitted on every attempt to register an endpoint, including failures
  // which progressed far enough to send an offchain request. Boba credits
  // are consumed per attempt, not per success.
  event EndpointRegistered(bool Success, bytes32 EK, address Owner, string URL);

  uint256 RegCredits = 100;

  function RegisterEndpoint(string memory url, uint256 credits, bytes32 auth)
    public returns (bool) {
    
    BobaHCHelper Self = BobaHCHelper(HelperAddr);
    bytes32 EK = keccak256(abi.encodePacked(url));
    bool success;

    require(credits >= RegCredits, "Insufficient registration credit");
    // FIXME - process the payment and require success
    
    // Future registration protocols may perform a 2-way exchange, but
    // for now we only care about authenticating the endpoint to the Helper contract.
    bytes memory request = abi.encodeWithSignature("RegisterV1()");
    
    bytes memory response = Self.GetResponse(1, address(this), url, request);
    
    if (response.length == 32 && keccak256(response) == auth) {
      // Success
      
      // FIXME - if overwriting an old entry, emit an Unregistration and deal with credits.
      
      Endpoints[EK].Owner = msg.sender;
      Endpoints[EK].URL = url;
      success = true;
    }

    emit EndpointRegistered(success, EK, msg.sender, url);
    return success;
  }

  function UnregisterEndpoint(bytes32 EK) public {
    require(Endpoints[EK].Owner == msg.sender, "Only owner may unregister");
    // FIXME refund any leftover credits to Owner
    delete(Endpoints[EK]);
  }
  
  // Add a user contract's address to the list of permitted callers for the 
  // specified previously-registered endpoint.
  // TODO - credit limits?
  function AddPermittedCaller(bytes32 _EK, address _callerAddress) public {
    require(Endpoints[_EK].Owner != address(0), "Endpoint is not registered");
    require(Endpoints[_EK].Owner == msg.sender, "Not the Endpoint owner");
    
    Endpoints[_EK].PermittedCallers[_callerAddress] = true;
  }
  
  // Remove a permited caller 
  function RemovePermittedCaller(bytes32 _EK, address _callerAddress) public {
    require(Endpoints[_EK].Owner != address(0), "Endpoint is not registered");
    require(Endpoints[_EK].Owner == msg.sender, "Not the Endpoint owner");
    
    Endpoints[_EK].PermittedCallers[_callerAddress] = false;
  }
  
  function CheckPermittedCaller(bytes32 _EK, address _callerAddress) public view returns (bool) {
    require(Endpoints[_EK].Owner != address(0), "Endpoint is not registered");
    return(Endpoints[_EK].PermittedCallers[_callerAddress]);
  }

  // -------------------------------------------------------------------------

  /* The following functions are called by users */

  function CallOffchain(bytes32 EK, bytes memory payload) public returns (bytes memory) {
    BobaHCHelper Self = BobaHCHelper(HelperAddr);
    string memory URL = Endpoints[EK].URL;
    require(bytes(URL).length > 0, "Endpoint not registered");

    require(Endpoints[EK].PermittedCallers[msg.sender], "Caller is not permitted");
    return Self.GetResponse(1, msg.sender, URL, payload);
  }

  function SimpleRandom() public returns (uint256) {
    BobaHCHelper Self = BobaHCHelper(HelperAddr);

    bytes memory response = Self.GetResponse(65537, msg.sender, "", "");
    require(response.length == 32, "Bad response length");
    
    uint256 result = abi.decode(response,(uint256));
    emit OffchainRandom(0, result);
    return result;
  }

  // -------------------------------------------------------------------------


  /* Legacy support */

/*
  function addPermittedCaller(address _callerAddress)
    public onlyOwner {
      permittedCaller[_callerAddress] = true;
      emit AddPermittedCaller(_callerAddress);
  }

  function removePermittedCaller(address _callerAddress)
    public onlyOwner {
      permittedCaller[_callerAddress] = false;
      emit RemovePermittedCaller(_callerAddress);
  }

  function checkPermittedCaller(address _callerAddress)
    public returns (bool) {
      bool permitted = permittedCaller[_callerAddress];
      emit CheckPermittedCaller(_callerAddress, permitted);
      return permitted;
  }

  function GetErrorCode(uint32 rType)
    internal view returns (string memory) {
      if(rType ==  1) return "TURING: Geth intercept failure";
      if(rType == 10) return "TURING: Incorrect input state";
      if(rType == 11) return "TURING: Calldata too short";
      if(rType == 12) return "TURING: URL >64 bytes";
      if(rType == 13) return "TURING: Server error";
      if(rType == 14) return "TURING: Could not decode server response";
      if(rType == 15) return "TURING: Could not create rpc client";
      if(rType == 16) return "TURING: RNG failure";
      if(rType == 17) return "TURING: API Response >322 chars";
      if(rType == 18) return "TURING: API Response >160 bytes";
      if(rType == 19) return "TURING: Insufficient credit";
      if(rType == 20) return "TURING: Missing cache entry";
  }

  / * This is the interface to the off-chain mechanism. Although
     marked as "public", it is only to be called by TuringCall()
     or TuringTX().
     The _payload parameter is overloaded to represent either the
     request parameters or the off-chain response, with the rType
     parameter indicating which is which.
     When called as a request (rType == 1), it starts the offchain call,
     which, if all all goes well, results in the rType changing to 2.
     This response is then passed back to the caller.
  * /
  function GetResponse(uint32 rType, string memory _url, bytes memory _payload)
    public returns (bytes memory) {

    require (msg.sender == address(this), "Turing:GetResponse:msg.sender != address(this)");
    require (_payload.length > 0, "Turing:GetResponse:no payload");
    require (rType == 2, string(GetErrorCode(rType))); // l2geth can pass values here to provide debug information
    return _payload;
  }

  function GetRandom(uint32 rType, uint256 _random)
    public returns (uint256) {

    require (msg.sender == address(this), "Turing:GetResponse:msg.sender != address(this)");
    require (rType == 2, string(GetErrorCode(rType)));
    return _random;
  }

  function Get42(uint32 rType, uint256 _random)
    public returns (uint256) {

    require (msg.sender == address(this), "Turing:GetResponse:msg.sender != address(this)");
    require (rType == 2, string(GetErrorCode(rType)));
    return _random;
  }

  / * Called from the external contract. It takes an api endponit URL
     and an abi-encoded request payload. The URL and the list of allowed
     methods are supplied when the contract is created. In the future
     some of this registration might be moved into l2geth, allowing for
     security measures such as TLS client certificates. A configurable timeout
     could also be added.

     Logs the offchain response so that a future verifier or fraud prover
     can replay the transaction and ensure that it results in the same state
     root as during the initial execution. Note - a future version might
     need to include a timestamp and/or more details about the
     offchain interaction.
  * /
  function TuringTx(string memory _url, bytes memory _payload)
    public onlyPermittedCaller override returns (bytes memory) {
      require (_payload.length > 0, "Turing:TuringTx:no payload");

      / * Initiate the request. This can't be a local function call
         because that would stay inside the EVM and not give l2geth
         a place to intercept and re-write the call.
      * /
      bytes memory response = Self.GetResponse(1, _url, _payload);
      emit OffchainResponse(0x01, response);
      return response;
  }

  function TuringRandom()
    public onlyPermittedCaller returns (uint256) {

      uint256 response = Self.GetRandom(1, 0);
      emit OffchainRandom(0x01, response);
      return response;
  }

  function Turing42()
    public onlyPermittedCaller returns (uint256) {

      uint256 response = Self.Get42(2, 42);
      emit Offchain42(0x01, response);
      return response;
  }

    // ERC165 check interface
    function supportsInterface(bytes4 _interfaceId) public pure returns (bool) {
        bytes4 firstSupportedInterface = bytes4(keccak256("supportsInterface(bytes4)")); // ERC165
        bytes4 secondSupportedInterface = ITuringHelper.TuringTx.selector;
        return _interfaceId == firstSupportedInterface || _interfaceId == secondSupportedInterface;
    }
*/
}
