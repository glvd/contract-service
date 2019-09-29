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

    modifier onlyWriter() {
        require(isWriter(), "Writeable: caller is not the writer");
        _;
    }

    function isWriter() public view returns (bool) {
        return _writer[_msgSender()];
    }

    function renounceWritership() public onlyOwner {
        _writer[_msgSender()] = false;
        emit WritershipDecreased(_msgSender());
    }

    function decreaseWritership(address oldWriter) public onlyOwner {
       _decreaseWritership(oldWriter);
    }
    
    function _decreaseWritership(address oldWriter) internal {
        require(_writer[oldWriter] == true, "Writerable: old writer was not writer");
        require(oldWriter != address(0),"Writerable: new writer is the zero address");
        _writer[oldWriter]=false;
        emit WritershipDecreased(oldWriter);
    }
    
    function increasedWritership(address newWriter) public onlyOwner {
        _increaseWritership(newWriter);
    }

    function _increaseWritership(address newWriter) internal {
        require(newWriter != address(0),"Writerable: new writer is the zero address");
        require(_writer[newWriter] == false, "Writerable: new writer is exist");
        _writer[newWriter]=true;
        emit WritershipIncreased(newWriter);
    }
}