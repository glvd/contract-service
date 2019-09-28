pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;
import "https://github.com/OpenZeppelin/openzeppelin-solidity/blob/master/contracts/ownership/Ownable.sol";

contract DMessage is Ownable{
    struct Message {
       string id;         //self id
       string content;    //jsoncontent
       string version;    //jsonversion
    }
    string[] private msgIds;
    mapping(string => bool) private mappingMessageFlags; //id => flag
    mapping(string => Message) private mappingMessages;  //id=>message
    
        /*
     * @dev 添加一个值到数组
     * @param _value string, 要传入的数值
     */
    function addMsgId(string memory _value) private {
        msgIds.push(_value);
    }

    /*
     * @dev 获取数组的长度
     * @return length uint,返回数组的长度
     */
    function getMsgLength()public view returns (uint256) {
        return msgIds.length;
    }

    /*
     * @dev 更新数组的值
     * @param _index uint, 指定的索引
     * @param _value string, 要修改的值
     */
    function updateMsgId(uint _index, string memory  _value) private {
        msgIds[_index] = _value;
    }

    /*
     * @dev 获取指定数组索引的值
     * @param _index uint, 索引值
     * @return _value uint, 返回结果
     */
    function getMsgId(uint _index) public view returns (string memory _value) {
        require(_index >= msgIds.length, "Message ID: index is overflow");
        return msgIds[_index];
    }

    /*
     * @dev 删除指定数组索引的值
     * @param _index uint, 索引值
     */
    function delMsgId(uint _index) private {
        if (_index >= msgIds.length){
          return;  
        } 
        msgIds[_index] = msgIds[msgIds.length-1];
        msgIds.length--;
    }

    /*
     * @dev 取得所有的值
     * @return msgIds []string,返回数组
     */
    function getMsgIds() public view returns (string[] memory _value) {
        return msgIds;
    }

    function addMessage(Message memory msg)public onlyOwner returns(bool) {
        require(mappingMessageFlags[msg.id] == true, "Message: add message is exist");
        mappingMessages[msg.id] = msg;
        mappingMessageFlags[msg.id] = true;
        addMsgId(msg.id);
        return true;
    }

    function updateMessage(Message memory msg)public onlyOwner returns(bool) {
       require(mappingMessageFlags[msg.id] == false, "Message: update message is not found");
       if (!mappingMessageFlags[msg.id]){
           return false;
       } 
       mappingMessages[msg.id] = msg;
       return true;
    }

    function checkMessage(string memory id)public view returns(bool) {
        return mappingMessageFlags[id];
    }

    function getMessage(string memory id) public view returns (Message memory) {
       require(mappingMessageFlags[id] == false, "Message: get message is not found");
       return mappingMessages[id];
    }
    
    function getMessages(uint  start,uint limit)public view returns (Message[] memory,uint){
        require(start > msgIds.length, "Message: start length is bigger than length");
        if ((start + limit) > msgIds.length){
            limit = msgIds.length - start;
        }
         Message[] memory msg = new Message[](limit);

        for (uint i = 0 ; i < limit ; i++){
            msg[i] = getMessage(getMsgId(start+i));
        }
        
        return (msg,limit);
    }
    
    function recount() private {
        for (uint i = 0 ; i < getMsgLength() ;i++){
            if (!mappingMessageFlags[getMsgId(i)]){
                delMsgId(i);
            }
        }
    }

    function delMessage(string memory id)public onlyOwner returns(Message memory,bool){
        require(mappingMessageFlags[id] == false, "Message: delete message is not found");
        mappingMessageFlags[id]=false;
        recount();
        return (mappingMessages[id],true);
    }
    
}