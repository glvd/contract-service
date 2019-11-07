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
    
    function addTagIds(string memory tag, string memory sub,string[] memory ids) public onlyOwner{
        require(mappingTags[tag].flags[sub] == false,"Tag: tag is exist");
        setTagIds(tag,sub,ids);
    }
    
    function setTagIds(string memory tag,string memory sub,string[] memory ids)public onlyOwner{
        require(message.checkAllMessage(ids) == true,"Tag: message is not exist");
        mappingTags[tag].flags[sub] = true;
        mappingTags[tag].tags[sub] = ids;
    }
    
    function replaceTagIds(string memory tag,  string memory sub,string[] memory ids) public onlyOwner{
        require(mappingTags[tag].flags[sub] == true,"Tag: tag is not exist");
        require(message.checkAllMessage(ids) == true,"Tag: message is not exist");
        mappingTags[tag].tags[sub] = ids;
    }

    function _addTagId(string memory tag,string memory sub,string memory id) private {
        mappingTags[tag].flags[sub] = true;
        mappingTags[tag].tags[sub].push(id);
    }

    function addTagId(string memory tag,string memory sub, string memory id)public onlyOwner{
        require(message.checkMessage(id) == true,"Tag: message is not found");
        _addTagId(tag,sub,id);
    }
    
    function checkTag(string memory tag,string memory sub)public view returns(bool){
        return mappingTags[tag].flags[sub];
    }
    
    function getTagIds(string memory tag, string memory sub) public view returns(string[] memory) {
        // require(mappingTags[tag].flags[sub] == true,"Tag: tag is not exist");
        return mappingTags[tag].tags[sub];
    }
    
    function checkTagIds(string memory tag,string memory sub)public view returns(uint,bool){
        return message.checkMessages(mappingTags[tag].tags[sub]);
    }
    
    function fixTagIds(string memory tag,string memory sub)public returns(uint _size){
        for (uint i = 0; i < mappingTags[tag].tags[sub].length; i++){
            if (!message.checkMessage(mappingTags[tag].tags[sub][i])){
                mappingTags[tag].tags[sub][i] = mappingTags[tag].tags[sub][mappingTags[tag].tags[sub].length-1];
                mappingTags[tag].tags[sub].length--;
            }
        }
        return mappingTags[tag].tags[sub].length;
    }
    
    function addTagMessage(string memory tag,string memory sub, string memory id, string memory content,string memory version)public onlyOwner returns(bool) {
        message.addMessage(id,content,version);
        _addTagId(tag,sub,id);
    }
    
    function getTagMessage(string memory tag, string memory sub) public view returns (DMessage.Message[] memory _value,uint _size) {
        _size = mappingTags[tag].tags[sub].length;
        _value = new DMessage.Message[](_size);
        for (uint i = 0; i < _size; i++){
            _value[i] = message.getMessage(mappingTags[tag].tags[sub][i]);
        }
        return (_value,_size);
    }
}