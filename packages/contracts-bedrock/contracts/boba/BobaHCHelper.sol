//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

//import '@openzeppelin/contracts/access/Ownable.sol';

// import "./ITuringHelper.sol";
import { Burn } from "../libraries/Burn.sol";

contract BobaHCHelper /*is Ownable*/ {

  address immutable HelperAddr = 0x42000000000000000000000000000000000000Fd;


  event OffchainRandom(uint version, uint256 random);
  event DBG(bytes b);

  uint256 public NextSimpleRandom;
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

  function Ping(uint32 a) public returns (uint32) {
    BobaHCHelper Self = BobaHCHelper(HelperAddr);
    
    if (a==0) {
      return 1;
    } else {
      return 2 * Self.Ping(a - 1);
    }
  }
  
  
  // Return a random number which was previously deposited by the
  // offchain interface. Note that this method theoretically allows
  // a hostile Sequencer node to force a desired outcome. Use the
  // enhanced commit/revel method for applications which need
  // a stronger guarantee of randomness.
  //
  // Any application using random numbers should take care to ensure
  // that gas usage is independent of the random result, or else it
  // may be possible for users to cancel and resubmit transactions 
  // repeatedly until a desired outcome is obtained.
  
  function GetSimpleRandom() public returns (uint256) {
    require(NextSimpleRandom != 0, "Random number not generated");
    uint256 result = NextSimpleRandom;
    NextSimpleRandom = 0;
    emit OffchainRandom(0, result);
    return result;
  }
  
  function PutSimpleRandom(uint256 val) public {
  	//require (NextSimpleRandom == 0, "NextSimpleRandom already set");
	NextSimpleRandom = val;
  }
  
  function GetBounce(uint32 rType, string memory _url, bytes memory _payload)
    public returns (bytes memory) {
     emit DBG(_payload);
    bytes32 z = 0;
    return abi.encode(z);
}


  function GetResponse(uint32 rType, string memory _url, bytes memory _payload)
    public returns (bytes memory) {

    //require (msg.sender == address(this), "Turing:GetResponse:msg.sender != address(this)");
    //require (_payload.length > 0, "Turing:GetResponse:no payload");
    //require (rType == 2, string(GetErrorCode(rType))); // l2geth can pass values here to provide debug information

    // This is for legacy support, TBD whether or not it will be implemented.
    if (rType == 2) {
      require (_payload.length > 0, "Turing:GetResponse:no payload");
      return _payload;
    }
    
    bytes32 cacheKey = keccak256(abi.encodePacked(_url, msg.sender, _payload));
    require(OffchainResponses[cacheKey].length > 0, "Missing cache entry");
    bytes memory response = OffchainResponses[cacheKey];

    // We burn additional L2 gas here to reflect the cost of storing the Response data in the L1 rollup.
    // TODO - calculate (L1_Gas_per_byte * Length) * (L1_Gas_Price / L2_Gas_Price). 
    uint256 gasToBurn = response.length;
    Burn.gas (gasToBurn * 5000);
    
    delete(OffchainResponses[cacheKey]);
    return response;
  }
  
  // This is called by a special Offchain transaction inserted ahead of one which is anticipated
  // to call GetResponse. This will also transfer Boba tokens to pay for the offchain operation.
  // Note that the token payment is separate from the extra gas burned in GetResponse to cover 
  // L1 calldata storage.
  function PutResponse(bytes32 cacheKey, bytes memory _payload) public {
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
    // consume tokens

    // This function can also be used to clean up the state after a reverted GetResponse transaction.
    // In this case we do not need to charge any additional Boba here because that was paid in the PutResponse.    
    delete(OffchainResponses[cacheKey]);
  }
  
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
    
    bytes memory response = Self.GetResponse(1, url, request);
    
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

  function CallOffchain(bytes32 EK, bytes memory payload) public returns (bytes memory) {
    BobaHCHelper Self = BobaHCHelper(HelperAddr);
    string memory URL = Endpoints[EK].URL;
    require(bytes(URL).length > 0, "Endpoint not registered");

    return GetResponse(1, URL, payload);
  }

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
