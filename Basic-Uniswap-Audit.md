The BasicUniswap and also the PoolFactory contracts that are given for the manual audit may have high severity vulnerabilities,   
ignoring the lows such as lack of function visibility specifiers like public, external declarations for the functions etc. 

# 1. Lack of input validation (deemed to be High)  
Description: swapExactInput, swapExactOutput functions seems not to validate input addresses properly.  
Impact on the application: Malicious addresses could be passed, leading to unexpected behavior and potential fund loss.  
Suggested/Fix: Validate addresses to ensure they are not zero addresses and conform to expected formats.  

Existing code:
```  
function swapExactInput(
    IERC20 inputToken,
    uint256 inputAmount,
    IERC20 outputToken,
    uint256 minOutputAmount,
    uint64 deadline
)
    public
    revertIfZero(inputAmount)
    revertIfDeadlinePassed(deadline)
    returns (uint256 output)
{
    // Function logic
}  
```  

Code Fix can be:  
```  
function swapExactInput(
    IERC20 inputToken,
    uint256 inputAmount,
    IERC20 outputToken,
    uint256 minOutputAmount,
    uint64 deadline
)
    public
    revertIfZero(inputAmount)
    revertIfDeadlinePassed(deadline)
    returns (uint256 output)
{
    require(address(inputToken) != address(0), "Invalid input token address");
    require(address(outputToken) != address(0), "Invalid output token address");
}
```  
In the context of smart contracts, especially those written in Solidity for Ethereum-based applications, input validation is a critical practice 
to ensure that functions receive valid and expected data.  

```  
require(address(inputToken) != address(0), "Invalid input token address");
require(address(outputToken) != address(0), "Invalid output token address");
```  
```
require(address(inputToken) != address(0), "Invalid input token address");  
```   
This line checks if the inputToken address is not the zero address (0x0000000000000000000000000000000000000000). The zero address is a special address in Ethereum   
that represents a non-existent address or an invalid address. It is often used to signify the absence of a value or a null address. Passing the zero address as a   
token address would likely indicate a mistake or a malicious attempt to disrupt the contract’s logic.  
  
Similarly, the next line performs the same check for the outputToken address:
```  
require(address(outputToken) != address(0), "Invalid output token address");
```  

## Importance of Address Validation: Prevent Null Address Errors, Security enforcement, Data Integrity checks.

# 2.  Insecure Contract Initialization (High or Critical)  
Description: The constructor in BasicUniswap accepts address parameters without validation.  
Impact: An invalid or malicious WETH token address could be set, leading to improper functioning of the contract.  
Suggested Fix: Validate the wethToken and poolToken addresses in the constructor to ensure they are valid contract addresses.  

Existing code:  
```  
constructor(
    address poolToken,
    address wethToken,
    string memory liquidityTokenName,
    string memory liquidityTokenSymbol
) ERC20(liquidityTokenName, liquidityTokenSymbol) {
    i_wethToken = IERC20(wethToken);
    i_poolToken = IERC20(poolToken);
}
```  

Code Fix can be:  
```  
constructor(
    address poolToken,
    address wethToken,
    string memory liquidityTokenName,
    string memory liquidityTokenSymbol
) ERC20(liquidityTokenName, liquidityTokenSymbol) {
    require(isContract(poolToken), "Invalid pool token address");
    require(isContract(wethToken), "Invalid WETH token address");

    i_wethToken = IERC20(wethToken);
    i_poolToken = IERC20(poolToken);
}
```  
Constructor Parameters can be added like above, ERC20 can be initialized, address validation and state variable initialization can be done. 

## Security improvement, Error Prevention, Robustness adds an extra layer of protection and the contracts are safe to use.




