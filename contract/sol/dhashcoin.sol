pragma solidity ^0.5.0;

import "./contracts/token/ERC20/ERC20.sol";
import "./contracts/token/ERC20/ERC20Mintable.sol";
import "./contracts/token/ERC20/ERC20Detailed.sol";

contract DHC is ERC20, ERC20Mintable,ERC20Detailed {
  constructor() ERC20Detailed("DHashCoin", "DHC",2) public {
  }
}