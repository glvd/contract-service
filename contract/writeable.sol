pragma solidity ^0.5.0;

import "https://github.com/OpenZeppelin/openzeppelin-solidity/blob/master/contracts/ownership/Ownable.sol";

/*
   @dev this contract is changed from Ownable
*/
contract Writeable is Ownable {
    address[] private _writer;

    event WritershipIncreased(uint indexed idx, address indexed newWriter);
    event WritershipDecreased(uint indexed idx, address indexed oldWriter);
    constructor () internal {
        address msgSender = _msgSender();
        _writer.push(msgSender);
        emit WritershipIncreased(_writer.length, msgSender);
    }

    /**
     * @dev Returns the address[] of the current writer.
     */
    function writer() public view returns (address[] memory) {
        return _writer;
    }

    function remove(uint _index) public onlyOwner{
        require(_index < _writer.length, "Writeable: index overflow");
        emit WritershipDecreased(_index, _writer[_index]);
        _writer[_index] = _writer[_writer.length-1];
        _writer.length--;
    }

    modifier onlyWriter() {
        require(isWriter(), "Writeable: caller is not the writer");
        _;
    }

    function isWriter() public view returns (bool) {
        for (uint i = 0 ; i< _writer.length;i++) {
            if (_msgSender() == _writer[i]){
                return true;
            }
        }
        return false;
    }

    function renounceWritership() public onlyOwner {
        delete _writer;
        emit WritershipDecreased(_writer.length, address(0));
    }

    function decreaseWritership(address oldWriter) public onlyOwner {
       _decreaseWritership(oldWriter);
    }
    
    function _decreaseWritership(address oldWriter) internal {
        require(_writer.length > 0, "Writerable: old writer has no data");
        for (uint i = 0; i < _writer.length; i++){
            if (_writer[i] == oldWriter){
                remove(i);
                return;
            }
        }
        return;
    }
    
    function increasedWritership(address newWriter) public onlyOwner {
        _increaseWritership(newWriter);
    }

    function _increaseWritership(address newWriter) internal {
        require(newWriter != address(0), "Writerable: new writer is the zero address");
        _writer.push(newWriter);
         emit WritershipIncreased(_writer.length, newWriter);
    }
}