pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;
import "./dmessage.sol";
// import "./Ownable.sol";

contract DTag is Ownable{
    DMessage message;

    struct Tag {
      mapping(string => bool) flags;
      mapping(string => string[]) tags;   //tag标记=>id
    }
    
    constructor(address msgAddr) public {
        message = DMessage(msgAddr);
    }
    
    mapping(string => Tag) private mappingTags;   //tag名称=>tag
    
    function addTagIds(string memory _tag, string memory _sub,string[] memory _ids) public onlyOwner{
        require(mappingTags[_tag].flags[_sub] == false,"Tag: tag is exist");
        setTagIds(_tag,_sub,_ids);
    }
    
    function setTagIds(string memory _tag,string memory _sub,string[] memory _ids)public onlyOwner{
        require(message.checkAllMessage(_ids) == true,"Tag: message is not exist");
        mappingTags[_tag].flags[_sub] = true;
        mappingTags[_tag].tags[_sub] = _ids;
    }
    
    function replaceTagIds(string memory _tag,  string memory _sub,string[] memory _ids) public onlyOwner{
        require(mappingTags[_tag].flags[_sub] == true,"Tag: tag is not exist");
        require(message.checkAllMessage(_ids) == true,"Tag: message is not exist");
        mappingTags[_tag].tags[_sub] = _ids;
    }

    function _addTagId(string memory _tag,string memory _sub,string memory _id) private {
        mappingTags[_tag].flags[_sub] = true;
        mappingTags[_tag].tags[_sub].push(_id);
    }

    function addTagId(string memory _tag,string memory _sub, string memory _id)public onlyOwner{
        require(message.checkMessage(_id) == true,"Tag: message is not found");
        _addTagId(_tag,_sub,_id);
    }
    
    function checkTag(string memory _tag,string memory _sub)public view returns(bool){
        return mappingTags[_tag].flags[_sub];
    }
    
    function getTagIds(string memory _tag, string memory _sub) public view returns(string[] memory) {
        // require(mappingTags[tag].flags[sub] == true,"Tag: tag is not exist");
        return mappingTags[_tag].tags[_sub];
    }
    
    function checkTagIds(string memory _tag,string memory _sub)public view returns(uint,bool){
        return message.checkMessages(mappingTags[_tag].tags[_sub]);
    }
    
    function fixTagIds(string memory _tag,string memory _sub)public returns(uint _size){
        for (uint i = 0; i < mappingTags[_tag].tags[_sub].length; i++){
            if (!message.checkMessage(mappingTags[_tag].tags[_sub][i])){
                mappingTags[_tag].tags[_sub][i] = mappingTags[_tag].tags[_sub][mappingTags[_tag].tags[_sub].length-1];
                mappingTags[_tag].tags[_sub].length--;
            }
        }
        return mappingTags[_tag].tags[_sub].length;
    }
    
    function addTagMessage(string memory _tag,string memory _sub, string memory _id, string memory _message)public onlyOwner returns(bool) {
        message.addMessage(_id,_message);
        _addTagId(_tag,_sub,_id);
    }
    
    function getTagMessage(string memory _tag, string memory _sub) public view returns (string[] memory _value,uint _size) {
        _size = mappingTags[_tag].tags[_sub].length;
        _value = new string[](_size);
        for (uint i = 0; i < _size; i++){
            _value[i] = message.getMessage(mappingTags[_tag].tags[_sub][i]);
        }
        return (_value,_size);
    }
}