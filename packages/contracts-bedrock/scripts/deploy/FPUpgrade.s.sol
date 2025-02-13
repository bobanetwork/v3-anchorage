// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// Testing
import { console2 as console } from "forge-std/console2.sol";
import { stdJson } from "forge-std/StdJson.sol";
import { AlphabetVM } from "test/mocks/AlphabetVM.sol";
import { EIP1967Helper } from "test/mocks/EIP1967Helper.sol";

// Scripts
import { Deployer } from "scripts/deploy/Deployer.sol";
import { DeployUtils } from "scripts/libraries/DeployUtils.sol";
import { ChainAssertions } from "scripts/deploy/ChainAssertions.sol";

// Contracts
import { Chains } from "scripts/libraries/Chains.sol";
import { Config } from "scripts/libraries/Config.sol";
import { Proxy } from "src/universal/Proxy.sol";
import { ProxyAdmin } from "src/universal/ProxyAdmin.sol";
import { SuperchainConfig } from "src/L1/SuperchainConfig.sol";
import { OptimismPortal2 } from "src/L1/OptimismPortal2.sol";
import { DisputeGameFactory } from "src/dispute/DisputeGameFactory.sol";
import { FaultDisputeGame } from "src/dispute/FaultDisputeGame.sol";
import { DelayedWETH } from "src/dispute/DelayedWETH.sol";
import { AnchorStateRegistry } from "src/dispute/AnchorStateRegistry.sol";
import { PreimageOracle } from "src/cannon/PreimageOracle.sol";
import { MIPS } from "src/cannon/MIPS.sol";

// Libraries
import { Types } from "scripts/libraries/Types.sol";
import { Duration } from "src/dispute/lib/LibUDT.sol";
import { GameType, Claim, GameTypes, OutputRoot, Hash } from "src/dispute/lib/Types.sol";

// Interfaces
import { IDisputeGameFactory } from "interfaces/dispute/IDisputeGameFactory.sol";
import { IDisputeGame } from "interfaces/dispute/IDisputeGame.sol";
import { IFaultDisputeGame } from "interfaces/dispute/IFaultDisputeGame.sol";
import { IDelayedWETH } from "interfaces/dispute/IDelayedWETH.sol";
import { IAnchorStateRegistry } from "interfaces/dispute/IAnchorStateRegistry.sol";
import { IPreimageOracle } from "interfaces/cannon/IPreimageOracle.sol";
import { ISuperchainConfig } from "interfaces/L1/ISuperchainConfig.sol";
import { IBigStepper } from "interfaces/dispute/IBigStepper.sol";
import { IPreimageOracle } from "interfaces/cannon/IPreimageOracle.sol";

/// @title Deploy
/// @notice Script used to upgrade the FP contracts
contract Deploy is Deployer {
    using stdJson for string;

    /// @notice FaultDisputeGameParams is a struct that contains the parameters necessary to call
    ///         the function _setFaultGameImplementation. This struct exists because the EVM needs
    ///         to finally adopt PUSHN and get rid of stack too deep once and for all.
    ///         Someday we will look back and laugh about stack too deep, today is not that day.
    struct FaultDisputeGameParams {
        AnchorStateRegistry anchorStateRegistry;
        DelayedWETH weth;
        GameType gameType;
        Claim absolutePrestate;
        IBigStepper faultVm;
        uint256 maxGameDepth;
    }

    ////////////////////////////////////////////////////////////////
    //                         Helper                            //
    ////////////////////////////////////////////////////////////////

    /// @notice read proxyd address from the hardhat deployment files
    function readProxyAddress(string memory _contractName) internal view returns (address proxyAddress_) {
        string memory _deploymentPath = vm.envOr("DEPLOYMENT_PATH", string(""));
        require(bytes(_deploymentPath).length > 0, "Deploy: must set DEPLOYMENT_PATH to deployment files");
        string memory _contractJson = vm.readFile(_deploymentPath);
        bytes memory _contractAddress = stdJson.parseRaw(_contractJson, string(abi.encodePacked(".", _contractName)));
        proxyAddress_ = bytesToAddress(_contractAddress);
    }

    /// @notice Convert bytes to address
    function bytesToAddress(bytes memory _bys) private pure returns (address addr_) {
        assembly {
            addr_ := mload(add(_bys, 32))
        }
    }

    ////////////////////////////////////////////////////////////////
    //                        Modifiers                           //
    ////////////////////////////////////////////////////////////////

    /// @notice Modifier that wraps a function in broadcasting.
    modifier broadcast() {
        vm.startBroadcast(msg.sender);
        _;
        vm.stopBroadcast();
    }

    ////////////////////////////////////////////////////////////////
    //                        Accessors                           //
    ////////////////////////////////////////////////////////////////

    /// @notice The create2 salt used for deployment of the contract implementations.
    ///         Using this helps to reduce config across networks as the implementation
    ///         addresses will be the same across networks when deployed with create2.
    function _implSalt() internal view returns (bytes32) {
        return keccak256(bytes(Config.implSalt()));
    }

    /// @notice Returns the proxy addresses. If a proxy is not found, it will have address(0).
    function _proxies() internal view returns (Types.ContractSet memory proxies_) {
        proxies_ = Types.ContractSet({
            L1CrossDomainMessenger: artifacts.mustGetAddress("L1CrossDomainMessengerProxy"),
            L1StandardBridge: artifacts.mustGetAddress("L1StandardBridgeProxy"),
            L2OutputOracle: artifacts.mustGetAddress("L2OutputOracleProxy"),
            DisputeGameFactory: artifacts.mustGetAddress("DisputeGameFactoryProxy"),
            PermissionedDelayedWETH: artifacts.getAddress("PermissionedDelayedWETHProxy"),
            DelayedWETH: artifacts.mustGetAddress("DelayedWETHProxy"),
            AnchorStateRegistry: artifacts.mustGetAddress("AnchorStateRegistryProxy"),
            OptimismMintableERC20Factory: artifacts.mustGetAddress("OptimismMintableERC20FactoryProxy"),
            OptimismPortal: artifacts.mustGetAddress("OptimismPortalProxy"),
            SystemConfig: artifacts.mustGetAddress("SystemConfigProxy"),
            L1ERC721Bridge: artifacts.mustGetAddress("L1ERC721BridgeProxy"),
            ProtocolVersions: artifacts.mustGetAddress("ProtocolVersionsProxy"),
            SuperchainConfig: artifacts.mustGetAddress("SuperchainConfigProxy")
        });
    }

    /// @notice Returns the proxy addresses, not reverting if any are unset.
    function _proxiesUnstrict() internal view returns (Types.ContractSet memory proxies_) {
        proxies_ = Types.ContractSet({
            L1CrossDomainMessenger: artifacts.getAddress("L1CrossDomainMessengerProxy"),
            L1StandardBridge: artifacts.getAddress("L1StandardBridgeProxy"),
            L2OutputOracle: artifacts.getAddress("L2OutputOracleProxy"),
            DisputeGameFactory: artifacts.getAddress("DisputeGameFactoryProxy"),
            DelayedWETH: artifacts.getAddress("DelayedWETHProxy"),
            PermissionedDelayedWETH: artifacts.getAddress("PermissionedDelayedWETHProxy"),
            AnchorStateRegistry: artifacts.getAddress("AnchorStateRegistryProxy"),
            OptimismMintableERC20Factory: artifacts.getAddress("OptimismMintableERC20FactoryProxy"),
            OptimismPortal: artifacts.getAddress("OptimismPortalProxy"),
            SystemConfig: artifacts.getAddress("SystemConfigProxy"),
            L1ERC721Bridge: artifacts.getAddress("L1ERC721BridgeProxy"),
            ProtocolVersions: artifacts.getAddress("ProtocolVersionsProxy"),
            SuperchainConfig: artifacts.getAddress("SuperchainConfigProxy")
        });
    }

    ////////////////////////////////////////////////////////////////
    //            State Changing Helper Functions                 //
    ////////////////////////////////////////////////////////////////

    /// @notice Call to the Proxy's upgrade and call method
    function _upgradeToAndCall(address payable _proxy, address _implementation, bytes memory _innerCallData) internal {
        Proxy(_proxy).upgradeToAndCall(_implementation, _innerCallData);
    }

    ////////////////////////////////////////////////////////////////
    //                    SetUp and Run                           //
    ////////////////////////////////////////////////////////////////

    /// @notice Deploy all FP contracts.
    function run() public {
        console.log("Upgrading the protocol to support FP");
        _run();
    }

    /// @notice Internal function containing the deploy logic.
    function _run() internal virtual {
        console.log("Start of L1 Deploy!");
        deployProxies();
        deployImplementations();
        initializeImplementations();

        setAlphabetFaultGameImplementation();
        setFastFaultGameImplementation();
        setCannonFaultGameImplementation();

        transferERC1967Proxy("DisputeGameFactoryProxy");
        transferERC1967Proxy("DelayedWETHProxy");
        transferERC1967Proxy("AnchorStateRegistryProxy");

        console.log("PLEASE UPGRADE OptimismPortal2 MANUALLY");
        console.log("PLEASE UPDATE DGF ADDRESS IN SystemConfigProxy MANUALLY");
    }

    /// @notice Deploy all of the proxies
    function deployProxies() public {
        console.log("Deploying proxies");

        // Both the DisputeGameFactory and L2OutputOracle proxies are deployed regardles of whether FPAC is enabled
        // to prevent a nastier refactor to the deploy scripts. In the future, the L2OutputOracle will be removed. If
        // fault proofs are not enabled, the DisputeGameFactory proxy will be unused.
        deployERC1967ProxyWithOwner("DisputeGameFactoryProxy", msg.sender);
        deployERC1967ProxyWithOwner("DelayedWETHProxy", msg.sender);
        deployERC1967ProxyWithOwner("AnchorStateRegistryProxy", msg.sender);
    }

    /// @notice Deploy all of the implementations
    function deployImplementations() public {
        console.log("Deploying implementations");

        // Fault proofs
        deployOptimismPortal2();
        deployDisputeGameFactory();
        deployDelayedWETH();
        deployPreimageOracle();
        deployMips();
        // Fix later
        // deployAnchorStateRegistry();
    }

    /// @notice Initialize all of the implementations
    function initializeImplementations() public {
        console.log("Initializing implementations");
        // This has to be manually triggered!
        // because we have already initialized the contracts in the previous deployment
        // initializeOptimismPortal2();

        initializeDisputeGameFactory();
        initializeDelayedWETH();
        // Fix later
        // initializeAnchorStateRegistry();
    }

    /// @notice Transfer the ERC1967Proxy ownership to the ProxyAdmin
    function transferERC1967Proxy(string memory _name) public {
        address proxyAdmin = readProxyAddress("ProxyAdmin");
        require(proxyAdmin != address(0), "Deploy: ProxyAdmin address not found");
        EIP1967Helper.setAdmin(artifacts.mustGetAddress(_name), proxyAdmin);

        // verify the admin
        require(EIP1967Helper.getAdmin(artifacts.mustGetAddress(_name)) == proxyAdmin, "Deploy: ProxyAdmin transfer failed");
        console.log("Transferred %s ownership to ProxyAdmin at %s", _name, proxyAdmin);
    }

    /// @notice Deploys an ERC1967Proxy contract with a specified owner.
    /// @param _name The name of the proxy contract to be deployed.
    /// @param _proxyOwner The address of the owner of the proxy contract.
    /// @return addr_ The address of the deployed proxy contract.
    function deployERC1967ProxyWithOwner(
        string memory _name,
        address _proxyOwner
    )
        public
        broadcast
        returns (address addr_)
    {
        console.log(string.concat("Deploying ERC1967 proxy for ", _name));
        Proxy proxy = new Proxy({ _admin: _proxyOwner });

        require(EIP1967Helper.getAdmin(address(proxy)) == _proxyOwner);

        artifacts.save(_name, address(proxy));
        console.log("   at %s", address(proxy));
        addr_ = address(proxy);
    }

    /// @notice Deploy the OptimismPortal2
    function deployOptimismPortal2() public broadcast returns (address addr_) {
        console.log("Deploying OptimismPortal2 implementation");

        // Could also verify this inside DeployConfig but doing it here is a bit more reliable.
        require(
            uint32(cfg.respectedGameType()) == cfg.respectedGameType(), "Deploy: respectedGameType must fit into uint32"
        );

        OptimismPortal2 portal = new OptimismPortal2({
            _proofMaturityDelaySeconds: cfg.proofMaturityDelaySeconds(),
            _disputeGameFinalityDelaySeconds: cfg.disputeGameFinalityDelaySeconds()
        });

        artifacts.save("OptimismPortal2", address(portal));
        console.log("OptimismPortal2 deployed at %s", address(portal));

        // Override the `OptimismPortal2` contract to the deployed implementation. This is necessary
        // to check the `OptimismPortal2` implementation alongside dependent contracts, which
        // are always proxies.
        Types.ContractSet memory contracts = _proxiesUnstrict();
        contracts.OptimismPortal = address(portal);
        ChainAssertions.checkOptimismPortal2({ _contracts: contracts, _cfg: cfg, _isProxy: false });

        addr_ = address(portal);
    }

    /// @notice Deploy the DisputeGameFactory
    function deployDisputeGameFactory() public broadcast returns (address addr_) {
        console.log("Deploying DisputeGameFactory implementation");
        DisputeGameFactory factory = new DisputeGameFactory();
        artifacts.save("DisputeGameFactory", address(factory));
        console.log("DisputeGameFactory deployed at %s", address(factory));

        // Override the `DisputeGameFactory` contract to the deployed implementation. This is necessary to check the
        // `DisputeGameFactory` implementation alongside dependent contracts, which are always proxies.
        Types.ContractSet memory contracts = _proxiesUnstrict();
        contracts.DisputeGameFactory = address(factory);
        ChainAssertions.checkDisputeGameFactory({ _contracts: contracts, _expectedOwner: address(0), _isProxy: false });

        addr_ = address(factory);
    }

    function deployDelayedWETH() public broadcast returns (address addr_) {
        console.log("Deploying DelayedWETH implementation");
        DelayedWETH weth = new DelayedWETH(cfg.faultGameWithdrawalDelay());
        artifacts.save("DelayedWETH", address(weth));
        console.log("DelayedWETH deployed at %s", address(weth));

        // Override the `DelayedWETH` contract to the deployed implementation. This is necessary
        // to check the `DelayedWETH` implementation alongside dependent contracts, which are
        // always proxies.
        Types.ContractSet memory contracts = _proxiesUnstrict();
        contracts.DelayedWETH = address(weth);
        ChainAssertions.checkDelayedWETH({
            _contracts: contracts,
            _cfg: cfg,
            _isProxy: false,
            _expectedOwner: address(0)
        });

        addr_ = address(weth);
    }

    /// @notice Deploy the PreimageOracle
    function deployPreimageOracle() public broadcast returns (address addr_) {
        console.log("Deploying PreimageOracle implementation");
        PreimageOracle preimageOracle = new PreimageOracle({
            _minProposalSize: cfg.preimageOracleMinProposalSize(),
            _challengePeriod: cfg.preimageOracleChallengePeriod()
        });
        artifacts.save("PreimageOracle", address(preimageOracle));
        console.log("PreimageOracle deployed at %s", address(preimageOracle));

        addr_ = address(preimageOracle);
    }

    /// @notice Deploy Mips
    function deployMips() public broadcast returns (address addr_) {
        console.log("Deploying Mips implementation");
        MIPS mips = new MIPS(IPreimageOracle(artifacts.mustGetAddress("PreimageOracle")));
        artifacts.save("Mips", address(mips));
        console.log("MIPS deployed at %s", address(mips));

        addr_ = address(mips);
    }

    // Fix later
    // /// @notice Deploy the AnchorStateRegistry
    // function deployAnchorStateRegistry() public broadcast returns (address addr_) {
    //     console.log("Deploying AnchorStateRegistry implementation");
    //     AnchorStateRegistry anchorStateRegistry =
    //         new AnchorStateRegistry(IDisputeGameFactory(artifacts.mustGetAddress("DisputeGameFactoryProxy")));
    //     artifacts.save("AnchorStateRegistry", address(anchorStateRegistry));
    //     console.log("AnchorStateRegistry deployed at %s", address(anchorStateRegistry));

    //     addr_ = address(anchorStateRegistry);
    // }

    /// @notice Initialize the DisputeGameFactory
    function initializeDisputeGameFactory() public broadcast {
        console.log("Upgrading and initializing DisputeGameFactory proxy");
        address disputeGameFactoryProxy = artifacts.mustGetAddress("DisputeGameFactoryProxy");
        address disputeGameFactory = artifacts.mustGetAddress("DisputeGameFactory");

        _upgradeToAndCall({
            _proxy: payable(disputeGameFactoryProxy),
            _implementation: disputeGameFactory,
            _innerCallData: abi.encodeCall(DisputeGameFactory.initialize, (msg.sender))
        });

        string memory version = DisputeGameFactory(disputeGameFactoryProxy).version();
        console.log("DisputeGameFactory version: %s", version);

        ChainAssertions.checkDisputeGameFactory({
            _contracts: _proxiesUnstrict(),
            _expectedOwner: msg.sender,
            _isProxy: true
        });
    }

    function initializeDelayedWETH() public broadcast {
        console.log("Upgrading and initializing DelayedWETH proxy");
        address delayedWETHProxy = artifacts.mustGetAddress("DelayedWETHProxy");
        address delayedWETH = artifacts.mustGetAddress("DelayedWETH");
        address superchainConfigProxy = readProxyAddress("SuperchainConfigProxy");
        require(superchainConfigProxy != address(0), "Deploy: SuperchainConfigProxy address not found");

        // Override superchainconfig address
        Types.ContractSet memory contracts = _proxiesUnstrict();
        contracts.SuperchainConfig = address(superchainConfigProxy);

        _upgradeToAndCall({
            _proxy: payable(delayedWETHProxy),
            _implementation: delayedWETH,
            _innerCallData: abi.encodeCall(DelayedWETH.initialize, (msg.sender, ISuperchainConfig(superchainConfigProxy)))
        });

        string memory version = DelayedWETH(payable(delayedWETHProxy)).version();
        console.log("DelayedWETH version: %s", version);

        ChainAssertions.checkDelayedWETH({ _contracts: contracts, _cfg: cfg, _isProxy: true, _expectedOwner: msg.sender });
    }

    // Fix later
    // function initializeAnchorStateRegistry() public broadcast {
    //     console.log("Upgrading and initializing AnchorStateRegistry proxy");
    //     address anchorStateRegistryProxy = artifacts.mustGetAddress("AnchorStateRegistryProxy");
    //     address anchorStateRegistry = artifacts.mustGetAddress("AnchorStateRegistry");
    //     address superchainConfigProxy = readProxyAddress("SuperchainConfigProxy");
    //     require(superchainConfigProxy != address(0), "Deploy: SuperchainConfigProxy address not found");

    //     AnchorStateRegistry.StartingAnchorRoot[] memory roots = new AnchorStateRegistry.StartingAnchorRoot[](4);
    //     roots[0] = AnchorStateRegistry.StartingAnchorRoot({
    //         gameType: GameTypes.CANNON,
    //         outputRoot: OutputRoot({
    //             root: Hash.wrap(cfg.faultGameGenesisOutputRoot()),
    //             l2BlockNumber: cfg.faultGameGenesisBlock()
    //         })
    //     });
    //     roots[1] = AnchorStateRegistry.StartingAnchorRoot({
    //         gameType: GameTypes.PERMISSIONED_CANNON,
    //         outputRoot: OutputRoot({
    //             root: Hash.wrap(cfg.faultGameGenesisOutputRoot()),
    //             l2BlockNumber: cfg.faultGameGenesisBlock()
    //         })
    //     });
    //     roots[2] = AnchorStateRegistry.StartingAnchorRoot({
    //         gameType: GameTypes.ALPHABET,
    //         outputRoot: OutputRoot({
    //             root: Hash.wrap(cfg.faultGameGenesisOutputRoot()),
    //             l2BlockNumber: cfg.faultGameGenesisBlock()
    //         })
    //     });
    //     roots[3] = AnchorStateRegistry.StartingAnchorRoot({
    //         gameType: GameTypes.ASTERISC,
    //         outputRoot: OutputRoot({
    //             root: Hash.wrap(cfg.faultGameGenesisOutputRoot()),
    //             l2BlockNumber: cfg.faultGameGenesisBlock()
    //         })
    //     });

    //     _upgradeToAndCall({
    //         _proxy: payable(anchorStateRegistryProxy),
    //         _implementation: anchorStateRegistry,
    //         _innerCallData: abi.encodeCall(
    //             AnchorStateRegistry.initialize, (roots, ISuperchainConfig(superchainConfigProxy))
    //         )
    //     });

    //     string memory version = AnchorStateRegistry(payable(anchorStateRegistryProxy)).version();
    //     console.log("AnchorStateRegistry version: %s", version);
    // }

    /// @notice Sets the implementation for the `CANNON` game type in the `DisputeGameFactory`
    function setCannonFaultGameImplementation() public broadcast {
        console.log("Setting Cannon FaultDisputeGame implementation");
        IDisputeGameFactory factory = IDisputeGameFactory(artifacts.mustGetAddress("DisputeGameFactoryProxy"));
        IDelayedWETH weth = IDelayedWETH(artifacts.mustGetAddress("DelayedWETHProxy"));

        // Set the Cannon FaultDisputeGame implementation in the factory.
        _setFaultGameImplementation({
            _factory: factory,
            _params: IFaultDisputeGame.GameConstructorParams({
                gameType: GameTypes.CANNON,
                absolutePrestate: loadMipsAbsolutePrestate(),
                maxGameDepth: cfg.faultGameMaxDepth(),
                splitDepth: cfg.faultGameSplitDepth(),
                clockExtension: Duration.wrap(uint64(cfg.faultGameClockExtension())),
                maxClockDuration: Duration.wrap(uint64(cfg.faultGameMaxClockDuration())),
                vm: IBigStepper(artifacts.mustGetAddress("Mips")),
                weth: weth,
                anchorStateRegistry: IAnchorStateRegistry(artifacts.mustGetAddress("AnchorStateRegistryProxy")),
                l2ChainId: cfg.l2ChainID()
            })
        });
    }

    /// @notice Sets the implementation for the `ALPHABET` game type in the `DisputeGameFactory`
    function setAlphabetFaultGameImplementation() public broadcast {
        console.log("Setting Alphabet FaultDisputeGame implementation");
        IDisputeGameFactory factory = IDisputeGameFactory(artifacts.mustGetAddress("DisputeGameFactoryProxy"));
        IDelayedWETH weth = IDelayedWETH(artifacts.mustGetAddress("DelayedWETHProxy"));

        Claim outputAbsolutePrestate = Claim.wrap(bytes32(cfg.faultGameAbsolutePrestate()));
        _setFaultGameImplementation({
            _factory: factory,
            _params: IFaultDisputeGame.GameConstructorParams({
                gameType: GameTypes.ALPHABET,
                absolutePrestate: outputAbsolutePrestate,
                // The max depth for the alphabet trace is always 3. Add 1 because split depth is fully inclusive.
                maxGameDepth: cfg.faultGameSplitDepth() + 3 + 1,
                splitDepth: cfg.faultGameSplitDepth(),
                clockExtension: Duration.wrap(uint64(cfg.faultGameClockExtension())),
                maxClockDuration: Duration.wrap(uint64(cfg.faultGameMaxClockDuration())),
                vm: IBigStepper(new AlphabetVM(outputAbsolutePrestate, IPreimageOracle(artifacts.mustGetAddress("PreimageOracle")))),
                weth: weth,
                anchorStateRegistry: IAnchorStateRegistry(artifacts.mustGetAddress("AnchorStateRegistryProxy")),
                l2ChainId: cfg.l2ChainID()
            })
        });
    }

    /// @notice Sets the implementation for the `ALPHABET` game type in the `DisputeGameFactory`
    function setFastFaultGameImplementation() public broadcast {
        console.log("Setting Fast FaultDisputeGame implementation");
        IDisputeGameFactory factory = IDisputeGameFactory(artifacts.mustGetAddress("DisputeGameFactoryProxy"));
        IDelayedWETH weth = IDelayedWETH(artifacts.mustGetAddress("DelayedWETHProxy"));

        Claim outputAbsolutePrestate = Claim.wrap(bytes32(cfg.faultGameAbsolutePrestate()));
        IPreimageOracle fastOracle = IPreimageOracle(
            DeployUtils.create2AndSave({
                _save: artifacts,
                _salt: _implSalt(),
                _name: "PreimageOracle",
                _nick: "FastPreimageOracle",
                _args: DeployUtils.encodeConstructor(
                    abi.encodeCall(IPreimageOracle.__constructor__, (cfg.preimageOracleMinProposalSize(), 0))
                )
            })
        );
        _setFaultGameImplementation({
            _factory: factory,
            _params: IFaultDisputeGame.GameConstructorParams({
                gameType: GameTypes.FAST,
                absolutePrestate: outputAbsolutePrestate,
                // The max depth for the alphabet trace is always 3. Add 1 because split depth is fully inclusive.
                maxGameDepth: cfg.faultGameSplitDepth() + 3 + 1,
                splitDepth: cfg.faultGameSplitDepth(),
                clockExtension: Duration.wrap(uint64(cfg.faultGameClockExtension())),
                maxClockDuration: Duration.wrap(0), // Resolvable immediately
                vm: IBigStepper(new AlphabetVM(outputAbsolutePrestate, fastOracle)),
                weth: weth,
                anchorStateRegistry: IAnchorStateRegistry(artifacts.mustGetAddress("AnchorStateRegistryProxy")),
                l2ChainId: cfg.l2ChainID()
            })
        });
    }

    /// @notice Sets the implementation for the given fault game type in the `DisputeGameFactory`.
    function _setFaultGameImplementation(
        IDisputeGameFactory _factory,
        IFaultDisputeGame.GameConstructorParams memory _params
    )
        internal
    {
        if (address(_factory.gameImpls(_params.gameType)) != address(0)) {
            console.log(
                "[WARN] DisputeGameFactoryProxy: `FaultDisputeGame` implementation already set for game type: %s",
                vm.toString(GameType.unwrap(_params.gameType))
            );
            return;
        }

        uint32 rawGameType = GameType.unwrap(_params.gameType);
        require(
            rawGameType != GameTypes.PERMISSIONED_CANNON.raw(), "Deploy: Permissioned Game should be deployed by OPCM"
        );

        _factory.setImplementation(
            _params.gameType,
            IDisputeGame(
                DeployUtils.create2AndSave({
                    _save: artifacts,
                    _salt: _implSalt(),
                    _name: "FaultDisputeGame",
                    _nick: string.concat("FaultDisputeGame_", vm.toString(rawGameType)),
                    _args: DeployUtils.encodeConstructor(abi.encodeCall(IFaultDisputeGame.__constructor__, (_params)))
                })
            )
        );

        string memory gameTypeString;
        if (rawGameType == GameTypes.CANNON.raw()) {
            gameTypeString = "Cannon";
        } else if (rawGameType == GameTypes.ALPHABET.raw()) {
            gameTypeString = "Alphabet";
        } else {
            gameTypeString = "Unknown";
        }

        console.log(
            "DisputeGameFactoryProxy: set `FaultDisputeGame` implementation (Backend: %s | GameType: %s)",
            gameTypeString,
            vm.toString(rawGameType)
        );
    }

    /// @notice Custom error for missing Cannon prestate dump
    error CannonPrestateDumpNotFound();

    /// @notice Loads the mips absolute prestate from the prestate-proof for devnets otherwise
    ///         from the config.
    function loadMipsAbsolutePrestate() internal returns (Claim mipsAbsolutePrestate_) {
        if (block.chainid == Chains.LocalDevnet || block.chainid == Chains.GethDevnet) {
            // Fetch the absolute prestate dump
            string memory filePath = string.concat(vm.projectRoot(), "/../../op-program/bin/prestate-proof.json");
            string[] memory commands = new string[](3);
            commands[0] = "bash";
            commands[1] = "-c";
            commands[2] = string.concat("[[ -f ", filePath, " ]] && echo \"present\"");
            if (vm.ffi(commands).length == 0) {
                revert CannonPrestateDumpNotFound();
            }
            commands[2] = string.concat("cat ", filePath, " | jq -r .pre");
            mipsAbsolutePrestate_ = Claim.wrap(abi.decode(vm.ffi(commands), (bytes32)));
            console.log(
                "[Cannon Dispute Game] Using devnet MIPS Absolute prestate: %s",
                vm.toString(Claim.unwrap(mipsAbsolutePrestate_))
            );
        } else {
            console.log(
                "[Cannon Dispute Game] Using absolute prestate from config: %x", cfg.faultGameAbsolutePrestate()
            );
            mipsAbsolutePrestate_ = Claim.wrap(bytes32(cfg.faultGameAbsolutePrestate()));
        }
    }
}
