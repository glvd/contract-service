pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;
import "./dmessage.sol";
import "https://github.com/OpenZeppelin/openzeppelin-solidity/blob/master/contracts/ownership/Ownable.sol";

contract DTag is Ownable{
    DMessage message;

    struct Tag {
    //   bool flag;
      mapping(string => bool) flags;
      mapping(string => string[]) tags;   //tag标记=>id
    }
    
    constructor(address msgAddr) public {
        message = DMessage(msgAddr);
    }
    
    mapping(string => Tag) private mappingTags;   //tag名称=>tag
    
    function checkMessages(string[] memory ids)public view returns(uint,bool){
        for (uint i = 0 ;i<ids.length;i++){
            if (!message.checkMessage(ids[i])){
                return (i,false);
            }
        }
        return (ids.length,true);
    }
      
    function addTag(string memory tag, string memory sub,string[] memory ids) public onlyOwner{
        checkMessages(ids);
        mappingTags[tag].flags[sub] = true;
        mappingTags[tag].tags[sub] = ids;
    }
    
    function replaceTag(string memory tag,  string memory sub,string[] memory ids) public onlyOwner{
        require(mappingTags[tag].flags[sub] == true,"Tag: tag is not exist");
        checkMessages(ids);
        mappingTags[tag].tags[sub] = ids;
    }

    function addOrReplacTag(string memory tag, string memory sub, string[] memory ids)public onlyOwner{
        checkMessages(ids);
        mappingTags[tag].flags[sub] = true;
        mappingTags[tag].tags[sub] = ids;
    }

    function addTagId(string memory tag,string memory sub, string memory id)public onlyOwner{
        mappingTags[tag].flags[sub] = true;
        mappingTags[tag].tags[sub].push(id);
    }
    
    function getTagSub(string memory tag, string memory sub) public view returns(string[] memory) {
        require(mappingTags[tag].flags[sub] == true,"Tag: tag is not exist");
        return mappingTags[tag].tags[sub];
    }
    
    function addTagMessage(string memory tag,string memory sub, DMessage.Message memory msg)public onlyOwner returns(bool) {
        message.addMessage(msg);
        addTagId(tag,sub,msg.id);
    }
    
    // function getTagMessage(string memory tag, string memory sub) public returns (Message[] memory) {
    //   string[] list =   getTagSub(tag,sub);
    //   Message[] msg;
    //     for (uint i = 0; i < list.length; i++){
    //         msg.push(mappingMessages[list[i]]);
    //     }
    //     return msg;
    // }
}