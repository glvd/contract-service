pragma solidity ^0.5.0;

import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/ERC20.sol";
// import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/ERC20Mintable.sol";
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/ERC20Detailed.sol";
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/GSN/GSNRecipient.sol";
import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/ownership/Secondary.sol";

contract DHC is ERC20,ERC20Detailed,Secondary {
    uint256 private constant UINT256_MAX = 2**256 - 1;

    constructor(string memory name, string memory symbol, uint8 decimals) public ERC20Detailed(name, symbol, decimals) {
        // solhint-disable-previous-line no-empty-blocks
    }

    // The primary account (GSNRecipientERC20Fee) can mint tokens
    function mint(address account, uint256 amount) public onlyPrimary {
        _mint(account, amount);
    }

    // The primary account has 'infinite' allowance for all token holders
    function allowance(address owner, address spender) public view returns (uint256) {
        if (spender == primary()) {
            return UINT256_MAX;
        } else {
            return super.allowance(owner, spender);
        }
    }

    // Allowance for the primary account cannot be changed (it is always 'infinite')
    function _approve(address owner, address spender, uint256 value) internal {
        if (spender == primary()) {
            return;
        } else {
            super._approve(owner, spender, value);
        }
    }

    function transferFrom(address sender, address recipient, uint256 amount) public returns (bool) {
        if (recipient == primary()) {
            _transfer(sender, recipient, amount);
            return true;
        } else {
            return super.transferFrom(sender, recipient, amount);
        }
    }

}