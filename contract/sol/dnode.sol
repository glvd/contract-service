pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;
import "./Ownable.sol";

contract DNode is Ownable{

    struct Node {
      mapping(uint => string) swarms;
      uint count;
    }

    struct Version{
        uint   cid;//client id
        string hash;
        string version;
    }

    mapping(uint => uint) mappingDevices;

    mapping(uint => Node)  mappingNodes;
    uint private last;

    mapping(uint => Version) mappingVersions;
    uint private totalVersion;

    function setNodeLast(uint _idx) public onlyOwner{
        require(_idx > last, "Node: last node is newer than now");
        last = _idx;
    }

    function getNodeLast() public view returns(uint) {
        return last;
    }

    function getDeviceVersion(uint _cid) public view returns(Version memory _version){
        require(mappingDevices[_cid] > 0, "Version: index version is not exist");
        return getVersion(mappingDevices[_cid]);
    }

    function setDeviceVersion(uint _cid,string memory _ver,string memory _hash)public onlyOwner{
        mappingDevices[_cid]= totalVersion;
        setVersion(_cid,_ver,_hash);
    }

    function getVersion(uint _idx)public view returns(Version memory _version){
        return mappingVersions[_idx];
    }

    function setVersion(uint _cid,string memory _ver,string memory _hash)private onlyOwner{
        mappingVersions[totalVersion] = Version({cid:_cid,version:_ver,hash:_hash});
        totalVersion++;
    }

    function replaceVersion(uint _idx,uint _cid,string memory _ver,string memory _hash)public onlyOwner{
        mappingVersions[_idx] = Version({cid:_cid,version:_ver,hash:_hash});
    }

    function removeNode(uint _idx) public onlyOwner{
        require(_idx < mappingNodes[last].count, "Node: delete node is not exist");
        uint limit = mappingNodes[last].count;
        mappingNodes[last].swarms[_idx] = mappingNodes[last].swarms[limit-1];
        mappingNodes[last].count--;
    }

    function storeNode(string memory _swarm)public onlyOwner{
         mappingNodes[last].swarms[mappingNodes[last].count] = _swarm;
         mappingNodes[last].count++;
    }

    function getLastNode()public view returns(string[] memory, uint){
        return getNode(last);
    }

    function getNode(uint _idx) public view returns(string[] memory,uint){
        uint limit = mappingNodes[_idx].count;
        string[] memory swarms = new string[](limit);
        for (uint i = 0 ; i < limit ; i++){
            swarms[i] = mappingNodes[_idx].swarms[i];
        }
        return (swarms,limit);
    }
}