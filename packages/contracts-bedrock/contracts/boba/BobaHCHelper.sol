//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.15;

//import '@openzeppelin/contracts/access/Ownable.sol';
//import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
//import "@openzeppelin/contracts/utils/introspection/ERC165Checker.sol";

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
	//string URL;
	//uint256 CreditBalance;
	mapping (address => bool) PermittedCallers;
  }

  mapping (bytes32 => EndpointEntry) public Endpoints;
  mapping(bytes32 => uint32) OffchainResponseCode;
  mapping(bytes32 => bytes)  OffchainResponseData;


  // events 
  event OffchainRandom(uint version, uint256 random);
  event DBG(bytes b);
  event AddBalanceTo(address, bytes32, uint256);
  //event WithdrawRevenue(address sender, uint256 withdrawAmount);


  constructor () {
   }

  // -------------------------------------------------------------------------


  /* Internal functions to compute a cache key and retrieve a result. Each
     method computes a cache key based on its parameters, then calls a common
     function to retrieve the response. If that reverts then the sequencer
     will extract the request parameters and perform the corresponding
     operation, writing the result into the cache and then re-running the
     original transaction which will now find the result it is expecting.

     ResponseCode is a packed structure of:
       - Version (internal consistency check)
       - User result code (0 = success, other = error)
       - Extra L2 Gas to burn
       - Boba token cost
     The L2 gas is provided as a mechanism to offset L1 storage costs. If the
     L1 cost is applied directly then this field should be set to 0.
  */
  
  function GetResponse(bytes32 _cacheKey) private returns (bool, bytes memory) {
  
    
    bytes memory cachedResponse = OffchainResponse[_cacheKey];
    require(cachedResponse.length > 0, "GetResponse: Missing cache entry");
    uint256 responseInfo;
    uint32 responseCode;
    bytes memory responseData;
    (responseInfo, responseData) = abi.decode(cachedResponse,(uint256,bytes));
    
    if (responseCode == 0) {
      return (true, responseData);
    } else {
      // Error code lookup?
      return (false, responseData);
    }
  }

  // Populate the cache ahead of an anticipated GetResponse(), charging Boba tokens.
  function PutResponse(bytes32 cacheKey, uint32 _code, bytes memory _payload) public {
    // Reserved address for Offchain transactions (core/types/offchain_tx.go)
    require(msg.sender == OffchainCaller, "Invalid PutResponse() caller");
    // Eventually this should just overwrite
    require(OffchainResponseData[cacheKey].length == 0, "DEBUG - Already exists");
    OffchainResponseCode[cacheKey] = _code;
    OffchainResponseData[cacheKey] = _payload;
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
    delete(OffchainResponseCode[cacheKey]);
    delete(OffchainResponseData[cacheKey]);
  }


  // ===================================================================

  /* Internal functions which provide the interface to core/vm/hybrid_compute.go. GetResponse()
    is called from other functions in this contract while PutResponse() and ExpireResponse() are
    called from Offchain transactions inserted by the sequencer node.
    
    The first time that GetResponse() is called, it will revert on a "missing cache entry". On a
    successful interaction, the offchain interactions will take place in the background and then
    the user transaction will be re-run. This second time it will find a cache entry and return
    the result. On a failure, the user transaction will be re-run a second time without any special
    logic, so that the "missing cache entry" is passed back to the caller.
  */
  
  // Try to get a previously-populated response
  function GetResponseV3(address _caller, string memory _url, bytes memory _payload)
    public returns (bool, bytes memory) {
    bool success;
    bytes memory responseData;

    require (msg.sender == address(this), "Invalid GetResponse() caller");

    bytes32 cacheKey = keccak256(abi.encodePacked(_caller, _url,  _payload));

    uint32 code = OffchainResponseCode[cacheKey];
    responseData = OffchainResponseData[cacheKey];

    if (code == 0) {
      // This specific revert string triggers the offchain mechanism. Do not change.
      require(responseData.length > 0, "GetResponse: Missing cache entry");
      success = true;

      // Placeholder to account for extra L1 storage costs
      // TODO - calculate (L1_Gas_per_byte * Length) * (L1_Gas_Price / L2_Gas_Price),
      // and merge with other L1 Fee logic
      uint256 gasToBurn = responseData.length;
      Burn.gas (gasToBurn * 5000);
    } else {
      string memory errMsg = "HybridCompute Error";  // TODO - look up a string based on the code.
      responseData = abi.encodeWithSignature("Error(string)", errMsg);
    }

    delete(OffchainResponseData[cacheKey]);
    delete(OffchainResponseCode[cacheKey]);
    return (success, responseData);
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
  
  
  // Emitted on every attempt to register an endpoint, including failures
  // which progressed far enough to send an offchain request. Boba credits
  // are consumed per attempt, not per success.
  event EndpointRegistered(bool Success, bytes32 EK, address Owner, string URL);

 function RegisterEndpoint(string memory url, uint256 credits, bytes32 auth)
    public returns (bool) {
    
    BobaHCHelper Self = BobaHCHelper(HelperAddr);
    bytes32 EK = keccak256(abi.encodePacked(url));
    bool success;
    
    require(credits >= RegCost, "Insufficient registration credit");
    // Transfer token to this contract, crediting any leftover balance to the endpoint

    // Future registration protocols may perform a 2-way exchange, but
    // for now we only care about authenticating the endpoint to the Helper contract.
    bytes memory request = abi.encodeWithSignature("RegisterV1()");
    bytes memory response; 

    (success, response) = Self.GetResponse(1, address(this), url, request);

    if (success && response.length == 32 && keccak256(response) == auth) {
      // Success

      address oldOwner = Endpoints[EK].Owner;
      uint256 oldBalance = Endpoints[EK].CreditBalance;
      
      // Process payment from new owner
      IERC20(hcToken).safeTransferFrom(msg.sender, address(this), credits);
      Endpoints[EK].CreditBalance += (credits - RegCost);
      ownerRevenue += RegCost;

      // Refund previous owner in case of overwrite
      if (oldOwner != address(0) && oldBalance > 0) {
        IERC20(hcToken).safeTransfer(oldOwner, oldBalance);
      }

      Endpoints[EK].Owner = msg.sender;
      Endpoints[EK].URL = url;
      success = true;
    } else {
      success = false;
    }

    emit EndpointRegistered(success, EK, msg.sender, url);
    return success;
  }

  function UnregisterEndpoint(bytes32 EK) public {
    require(Endpoints[EK].Owner == msg.sender, "Only owner may unregister");
    
    uint256 leftoverBalance = Endpoints[EK].CreditBalance;
    delete(Endpoints[EK]);
    IERC20(hcToken).safeTransfer(msg.sender, leftoverBalance);
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

  // Replaces BobaTuringCredit functionality
  function addBalanceTo(bytes32 _EK, uint256 _addBalanceAmount)
      public
     // onlyInitialized
  {
      require(Endpoints[_EK].Owner != address(0), "Endpoint is not registered");
      require(Endpoints[_EK].Owner == msg.sender, "Not the Endpoint owner");

      require(_addBalanceAmount != 0, "Invalid amount");
      /*
      require(Address.isContract(_helperContractAddress), "Address is EOA");
      require(
  	  ERC165Checker.supportsInterface(_helperContractAddress, 0x2f7adf43),
  	  "Invalid Helper Contract"
      );
      */

      // Transfer token to this contract
      IERC20(hcToken).safeTransferFrom(msg.sender, address(this), _addBalanceAmount);

      Endpoints[_EK].CreditBalance += _addBalanceAmount;

      emit AddBalanceTo(msg.sender, _EK, _addBalanceAmount);
  } 
  

  // -------------------------------------------------------------------------

  /* The following functions are called by users */

  function CallOffchain(bytes32 EK, bytes memory payload) public returns (bool, bytes memory) {
    BobaHCHelper Self = BobaHCHelper(HelperAddr);
    string memory URL = Endpoints[EK].URL;
    require(bytes(URL).length > 0, "Endpoint not registered");

    require(Endpoints[EK].PermittedCallers[msg.sender], "Caller is not permitted");
    
    require(Endpoints[EK].CreditBalance >= CallCost, "Insufficient credit balance");
    
    Endpoints[EK].CreditBalance -= CallCost;
    ownerRevenue += CallCost;

    bool success;
    bytes memory returnData;
    
    (success, returnData) = Self.GetResponse(1, msg.sender, URL, payload);
    return (success, returnData);
  }

  function SimpleRandom() public returns (uint256) {
    BobaHCHelper Self = BobaHCHelper(HelperAddr);
    bool success;
    bytes memory response;
    
    // This demonstrates a different option for billing. As random numbers are a
    // built-in function, there is no Endpoint and therefore no credit balance to use.
    // Instead we process a payment directly from the caller.
    //
    // An alternative would be to define non-URL Endpoints and give them credits.
    
    require(IERC20(hcToken).allowance(msg.sender, address(this)) >= RNGCost, "Insufficient credit authorization");
    IERC20(hcToken).safeTransferFrom(msg.sender, address(this), RNGCost);
    ownerRevenue += RNGCost;

    (success, response) = Self.GetResponse(65537, msg.sender, "", "");
    require(success, "Operation failed");
    require(response.length == 32, "Bad response length");
    
    uint256 result = abi.decode(response,(uint256));
    emit OffchainRandom(0, result);
    return result;
  }

  // Secure sequential RNG. Clients contribute their own randomness which is XOR-ed with
  // a server-side secret to generate the final answer, ensuring that neither side is able
  // to force a desired outcome. Each side provides a hash committment of its secret value
  // in the first transaction, revealing the secret in a later block (TBD whether it's
  // the next block or whether a greater confirmation depth will be required).
  //
  // "session" is an arbitrary identifier allowing a client to manage multiple ongoing
  // random sequences. The server combines this key with msg.sender to generate its internal
  // key for maaintaining state. Server-side states will eventually expire and may be lost
  // for other reasons. Callers must implement error handling.
  // 
  // A "nextHash" of bytes32(0) will terminate a sequence, freeing server state.
  // A "myNum" of 0 with a non-zero nextHash initiates a new session, overwriting 
  // any previous server state. 
  
  function NextRandom(bytes32 session, bytes32 nextHash, uint256 myNum) public returns (bytes32, uint256) {
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
/*
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


}
