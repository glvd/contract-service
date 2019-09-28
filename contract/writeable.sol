pragma solidity ^0.5.0;

import "https://github.com/OpenZeppelin/openzeppelin-solidity/blob/master/contracts/ownership/Ownable.sol";

/*
   @dev this contract is changed from Ownable
*/
contract Writeable is Ownable {
    mapping(address => bool) private _writer;

    event WritershipIncreased(address indexed newWriter);
    event WritershipDecreased(address indexed oldWriter);
    constructor () internal {
        address msgSender = _msgSender();
        _writer[msgSender] = true;
        emit WritershipIncreased(msgSender);
    }

    /**
     * @dev Returns the address of the current writer.
     */
    function writer() public view returns (address) {
        if (isWriter() == true){
            return _msgSender();
        }
        return address(0);
    }

    /**
     * @dev Throws if called by any account other than the writer.
    */
    modifier onlyWriter() {
        require(isWriter(), "Writeable: caller is not the writer");
        _;
    }
    
    /**
     * @dev Returns true if the caller is the current writer.
     */
    function isWriter() public view returns (bool) {
        return _writer[_msgSender()];
    }

    /**
     * @dev Leaves the contract without owner. It will not be possible to call
     * `onlyOwner` functions anymore. Can only be called by the current writer.
     *
     * NOTE: Renouncing ownership will leave the contract without an owner,
     * thereby removing any functionality that is only available to the owner.
     */
    function renounceWritership() public onlyOwner {
        _writer[_msgSender()] = false;
        emit WritershipDecreased(_msgSender());
    }
    
    /**
     * @dev Decrease writership of the contract to a new account (`oldWriter`).
     * Can only be called by the current owner.
     */
    function decreaseWritership(address oldWriter) public onlyOwner {
        _decreaseWritership(oldWriter);
    }

    /**
     * @dev Decrease writership of the contract to a new account (`oldWriter`).
     */
    function _decreaseWritership(address oldWriter) internal {
        require(_writer[oldWriter] == true, "Writerable: old writer was not writer");
        _writer[oldWriter]=false;
        emit WritershipDecreased(oldWriter);
    }
    
    /**
     * @dev Increase writership of the contract to a new account (`newWriter`).
     * Can only be called by the current owner.
     */    
    function increaseWritership(address newWriter) public onlyOwner {
        _increaseWritership(newWriter);
    }
    
    /**
     * @dev Increase writership of the contract to a new account (`newWriter`).
     */
    function _increaseWritership(address newWriter) internal {
        require(newWriter != address(0), "Writerable: new writer is the zero address");
        _writer[newWriter]=true;
        emit WritershipIncreased(newWriter);
    }
}