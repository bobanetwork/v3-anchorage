//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

interface BobaHCHelper {
  function SimpleRandom() external returns(uint256);
  function GetRandom(uint32, uint256) external returns(uint256);
  function CallOffchain(string calldata, bytes calldata) external returns (bool, bytes memory);
  function SequentialRandom(bytes32, bytes32, uint256) external returns (bytes32, uint256);
  function GetLegacyResponse(uint32, string memory, bytes memory) external returns (bytes memory);
}

contract HCDemo {
  BobaHCHelper Helper;
  HCDemo Self;

  event CoinFlip(bool isHeads, uint256 raw);
  event MathResponse(uint32 Result);
  event StringLength (uint256 Length);
  event HashDbg(bytes32 indexed);

  constructor (address helper)  {
    Helper = BobaHCHelper(address(helper));
    Self = HCDemo(address(this));
  }

  // Call an offchain RPC endpoint to calculate (a+5)
  function Add5(uint32 a) public returns (uint32) {
    bytes memory response;
    bool success;

    (success, response) = Helper.CallOffchain(
      "http://192.168.4.2:1234/hc", 
      abi.encodeWithSignature("add5(uint32)",a)
      );

    require(success);

    uint32 b = abi.decode(response,(uint32));
    emit MathResponse(b);
    return b;
  }

  // This generates a string containing the word "chicken" repeated the
  // specified number of times. The function counts the length of the
  // response string and emits an event showing that length. The purpose
  // of this is to test that longer offchain responses consume enough
  // additional gas to compensate for L1 data storage.
  function Chicken(uint32 a) public returns (uint256) {

    bool success;
    bytes memory response;

    (success, response) = Helper.CallOffchain(
    	"http://192.168.4.2:1234/hc",
	 abi.encodeWithSignature("chicken(uint32)",a)
	 );

    require(success);

    bytes memory s = abi.decode(response,(bytes));
    uint256 len = s.length;
    emit StringLength(len);
    return len;
  }

  // Simple RNG example. Note that it may be possible for callers to predict or influence the result.
  function FlipCoin() public {
    uint256 r = Helper.GetRandom(1,0);
    bool isHeads = (r % 2 == 1);
    emit CoinFlip(isHeads, r);
  }

  // This is called multiple times to generate a sequence of random numbers
  // using a commit/reveal scheme. The client "secrets" are a sequential counter.
  // The sequence is terminated by supplying a 0 hash when revealing the final result.
  function SeqRandom(uint32 i) public returns (bytes32, uint256) {
    require(i >= 1 && i <= 3, "Invalid index");

    bytes32 responseHash;
    bytes32 clientHash;
    uint256 responseNum;
    uint256 clientNum;
    bytes32 seq = 0x1000000000000000000000000000000000000000000000000000000000000001;

    if (i==1) {
      clientNum = 0;
      clientHash = keccak256(abi.encodePacked(uint256(1)));
    } else if(i==2) {
      clientNum = 1;
      clientHash = keccak256(abi.encodePacked(uint256(2)));
    } else if(i==3) {
      clientNum = 2;
      clientHash = bytes32(0);
    }

    (responseHash, responseNum) = Helper.SequentialRandom(seq, clientHash, clientNum);

    return (responseHash, responseNum);
  }

  // Original Turing interface in which the server must process a length parameter
  function LegacyV1() public returns (uint32) {
    // Compute (a^2 + b^2)
    uint32 a = 3;
    uint32 b = 4;
    uint32 s;
    bytes memory request = abi.encode(a, b);
    bytes memory response = Helper.GetLegacyResponse(1, "http://192.168.4.2:1234/SoS_v1", request);
    s = abi.decode(response,(uint32));
    emit MathResponse(s);
    return s;
  }

  // The legacy V2 API removes the need to insert a length parameter, and allows larger responses (not tested here)
  function LegacyV2() public returns (uint32) {
    // Compute (a^2 + b^2)
    uint32 a = 4;
    uint32 b = 3;
    uint32 s;
    bytes memory request = abi.encode(a, b);
    bytes memory response = Helper.GetLegacyResponse(0x02000001, "http://192.168.4.2:1234/SoS_v2", request);
    s = abi.decode(response,(uint32));
    emit MathResponse(s);
    return s;
  }

  // ERC165 check interface, needed to register with legacy credit contract
  function supportsInterface(bytes4 _interfaceId) public pure returns (bool) {
      bytes4 firstSupportedInterface = bytes4(keccak256("supportsInterface(bytes4)")); // ERC165
      bytes4 secondSupportedInterface = bytes4(keccak256("TuringTx(string,bytes)"));
      return _interfaceId == firstSupportedInterface || _interfaceId == secondSupportedInterface;
  }

}

