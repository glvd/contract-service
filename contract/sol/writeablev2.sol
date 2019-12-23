pragma solidity ^0.5.0;

import "./Ownable.sol";

/*
   @dev this contract is changed from Ownable
*/
contract Writeable is Ownable {
    mapping(address => bool) private _writer;
    mapping(uint => address) private _writers;
    uint private count;

    event WritershipIncreased(address indexed newWriter);
    event WritershipDecreased(address indexed oldWriter);
    constructor () internal {
        count = 0;
        emit WritershipIncreased(_msgSender());
    }

    function del(uint idx) private {
        if (idx >= count){
            return;
        }
        _writers[idx] = _writers[count-1];
        count--;
    }

    /**
     * @dev Returns the address[] of the writers.
     */
    function writers() public view returns (address[] memory _value){
        _value = new address[](count);
        for (uint i = 0;i<count;i++) {
            _value[i] = _writers[i];
        }
        return _value;
    }

    function recount() private {
        for (uint i = 0 ; i < count ;i++){
            if (_writer[_writers[i]] == false){
                del(i);
            }
        }
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
        _writer[oldWriter] = false;
        recount();
        emit WritershipDecreased(oldWriter);
    }

    function increasedWritership(address newWriter) public onlyOwner {
        _increaseWritership(newWriter);
    }

    function _increaseWritership(address newWriter) internal {
        require(newWriter != address(0),"Writerable: new writer is the zero address");
        require(_writer[newWriter] == false, "Writerable: new writer is exist");
        _writer[newWriter] = true;
        _writers[count] = newWriter;
        count++;
        emit WritershipIncreased(newWriter);
    }
}