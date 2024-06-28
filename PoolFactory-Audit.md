# Vulnerability: Dependency on External Contracts without Validations  
The PoolManager contract relies on external contracts, specifically BasicUniswap and IERC20, without validating their addresses.   
This poses a significant risk because if an incorrect address is provided, the contract could interact with a non-existent or   
malicious contract, leading to unexpected behavior, contract failure, or security vulnerabilities.  

## Risk analysis
Incorrect Contract Address: If an incorrect address is provided, the contract may interact with a non-ERC20 token, causing function calls to fail.  
Malicious Contract: If a malicious contract is provided, it could potentially exploit the PoolManager contract, leading to loss of funds or data.  
Unexpected Behavior: Even if the address points to a contract, if it's not the expected type, it could cause the PoolManager contract to behave in unexpected ways, compromising its functionality and security.  

## Suggested Fix
Implement checks to ensure the provided token address is a valid ERC20 token contract. Additionally, validate the created   
BasicUniswap contract to ensure it conforms to expected behavior.  

## Improvised CreatePool function can be seen as follows:
```  
pragma solidity 0.8.20;

import { BasicUniswap } from "./BasicUniswap.sol";
import { IERC20 } from "forge-std/interfaces/IERC20.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";

contract PoolManager is Ownable {
    error PoolManager__PoolAlreadyExists(address tokenAddress);
    error PoolManager__PoolDoesNotExist(address tokenAddress);

    /*//////////////////////////////////////////////////////////////
                            STATE VARIABLES
    //////////////////////////////////////////////////////////////*/
    mapping(address => address) private s_pools;
    mapping(address => address) private s_tokens;

    address private immutable i_wethToken;

    /*//////////////////////////////////////////////////////////////
                                 EVENTS
    //////////////////////////////////////////////////////////////*/
    event PoolCreated(address indexed tokenAddress, address indexed poolAddress);

    /*//////////////////////////////////////////////////////////////
                               FUNCTIONS
    //////////////////////////////////////////////////////////////*/
    constructor(address wethToken) {
        require(isContract(wethToken), "WETH token address is not a contract");
        i_wethToken = wethToken;
    }

    /*//////////////////////////////////////////////////////////////
                           EXTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/
    function createPool(address tokenAddress) external onlyOwner returns (address) {
        if (s_pools[tokenAddress] != address(0)) {
            revert PoolManager__PoolAlreadyExists(tokenAddress);
        }
        require(isContract(tokenAddress), "Provided address is not a contract");
        require(isERC20(tokenAddress), "Provided address is not an ERC20 token");

        string memory liquidityTokenName = string.concat("BASIC-Uniswap ", IERC20(tokenAddress).name());
        string memory liquidityTokenSymbol = string.concat("bu", IERC20(tokenAddress).name());
        BasicUniswap tPool = new BasicUniswap(tokenAddress, i_wethToken, liquidityTokenName, liquidityTokenSymbol);
        require(isContract(address(tPool)), "Failed to create pool contract");

        s_pools[tokenAddress] = address(tPool);
        s_tokens[address(tPool)] = tokenAddress;

        emit PoolCreated(tokenAddress, address(tPool));
        return address(tPool);
    }

    /*//////////////////////////////////////////////////////////////
                   EXTERNAL AND PUBLIC VIEW AND PURE
    //////////////////////////////////////////////////////////////*/
    function getPool(address tokenAddress) external view returns (address) {
        return s_pools[tokenAddress];
    }

    function getToken(address pool) external view returns (address) {
        return s_tokens[pool];
    }

    function getWethToken() external view returns (address) {
        return i_wethToken;
    }

    /*//////////////////////////////////////////////////////////////
                           INTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/
    function isContract(address addr) internal view returns (bool) {
        uint256 size;
        assembly {
            size := extcodesize(addr)
        }
        return size > 0;
    }

    function isERC20(address addr) internal view returns (bool) {
        try IERC20(addr).totalSupply() returns (uint256) {
            return true;
        } catch {
            return false;
        }
    }
}
``` 



