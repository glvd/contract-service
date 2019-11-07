pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;
import "./writeable.sol";

contract DMessage is Writeable{
    struct Message {
       string id;         //self id
       string content;    //jsoncontent
       string version;    //jsonversion
    }
    
    mapping(string => bool) private mappingMessageFlags; //id => flag
    mapping(string => Message) private mappingMessages;  //id=>message
    mapping(uint256 => string) private mappingIds;
    uint256 private count;
    
    constructor() public {
        count = 0;
    }
    
    /*
     * @dev 添加一个值到数组
     * @param _value string, 要传入的数值
     */
    function addMsgId(string memory _value) private {
        mappingIds[count] = _value;
        count++;
        // msgIds.push(_value);
    }

    /*
     * @dev 获取数组的长度
     * @return length uint,返回数组的长度
     */
    function getMsgLength()public view returns (uint256) {
        return count;
        // return msgIds.length;
    }

    /*
     * @dev 更新数组的值
     * @param _index uint, 指定的索引
     * @param _value string, 要修改的值
     */
    function updateMsgId(uint _index, string memory  _value) private {
        mappingIds[_index] = _value;
    }

    /*
     * @dev 获取指定数组索引的值
     * @param _index uint, 索引值
     * @return _value uint, 返回结果
     */
    function getMsgId(uint _index) public view returns (string memory _value) {
        // require(_index < count, "Message ID: index is overflow");
        return mappingIds[_index];
        // return msgIds[_index];
    }

    /*
     * @dev 删除指定数组索引的值
     * @param _index uint, 索引值
     */
    function delMsgId(uint _index) private {
        if (_index >= count){
          return;  
        }
        mappingIds[_index] = mappingIds[count-1];
        // delete mappingIds[count-1];
        count--;
    }

    /*
     * @dev 取得所有的值
     * @return msgIds []string,返回数组
     */
    function getMsgIds() public view returns (string[] memory) {
        string[] memory ids = new string[](count);
        for (uint256 i = 0; i< count; i++) {
            ids[i] = mappingIds[i];
        }
        return ids;
    }

    function addMessage(string memory id, string memory content,string memory version)public onlyWriter returns(bool) {
        require(mappingMessageFlags[id] == false, "Message: add message is exist");
        Message memory newMessage = Message({id:id,content:content,version:version});
        mappingMessages[id] = newMessage;
        mappingMessageFlags[id] = true;
        addMsgId(id);
        return true;
    }

    function updateMessage(string memory id, string memory content,string memory version)public onlyWriter returns(bool) {
       require(mappingMessageFlags[id] == true, "Message: update message is not exist");
       Message memory newMessage = Message({id:id,content:content,version:version});
       mappingMessages[id] = newMessage;
       return true;
    }

    function checkMessage(string memory id)public view returns(bool) {
        return mappingMessageFlags[id];
    }

    function checkAllMessage(string[] memory ids)public view returns(bool){
         for (uint i = 0; i< ids.length;i++){
            if (!checkMessage(ids[i])){
                return false;
            }
        }
        return true;
    }

    function checkMessages(string[] memory ids)public view returns(uint,bool){
        for (uint i = 0; i< ids.length;i++){
            if (!checkMessage(ids[i])){
                return (i,false);
            }
        }
        return (0,true);
    }

    function getMessage(string memory id) public view returns (Message memory) {
    //   require(mappingMessageFlags[id] == true, "Message: get message is not exist");
       return mappingMessages[id];
    }
    
    function getMessages(uint  start,uint limit)public view returns (Message[] memory _value,uint _size){
        // require(start < count, "Message: start length is bigger than length");
        if ((start + limit) >= count){
            limit = count - start;
        }
        _value = new Message[](limit);

        for (uint i = 0 ; i < limit ; i++){
            _value[i] = getMessage(getMsgId(start+i));
        }
        
        return (_value,limit);
    }
    
    function recount() private {
        for (uint i = 0 ; i < getMsgLength() ;i++){
            if (!checkMessage(getMsgId(i))){
                delMsgId(i);
            }
        }
    }

    function delMessage(string memory id)public onlyOwner returns(Message memory,bool){
        require(mappingMessageFlags[id] == true, "Message: delete message is not exist");
        mappingMessageFlags[id]=false;
        recount();
        return (mappingMessages[id],true);
    }
    
}
