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

    function getDeviceVersion(uint cid) public view returns(string memory,string memory){
        require(mappingDevices[cid] > 0, "Version: index version is not exist");
        return getVersion(mappingDevices[cid]);
    }

    function setDeviceVersion(uint _cid,string memory _ver,string memory _hash)public onlyOwner{
        mappingDevices[_cid]= totalVersion;
        setVersion(_cid,_ver,_hash);
    }

    function getVersion(uint idx)public view returns(string memory,string memory){
        return (mappingVersions[idx].version,mappingVersions[idx].hash);
    }

    function setVersion(uint cid,string memory ver,string memory hash)public onlyOwner{
        mappingVersions[totalVersion] = Version({cid:cid,version:ver,hash:hash});
        totalVersion++;
    }

    function replaceVersion(uint idx,uint cid,string memory ver,string memory hash)public onlyOwner{
        mappingVersions[idx] = Version({cid:cid,version:ver,hash:hash});
    }

    function removeNode(uint idx) public onlyOwner{
        require(idx < mappingNodes[last].count, "Node: delete node is not exist");
        uint limit = mappingNodes[last].count;
        mappingNodes[last].swarms[idx] = mappingNodes[last].swarms[limit-1];
        mappingNodes[last].count--;
    }
    
    function setNodeLast(uint idx) public onlyOwner{
        require(idx > last, "Node: last node is newer than now");
        last = idx;
    }

    function getNodeLast() public view returns(uint) {
        return last;
    }

    function storeNode(string memory swarm)public onlyOwner{
         mappingNodes[last].swarms[mappingNodes[last].count] = swarm;
         mappingNodes[last].count++;
    }

    function getNodeLastData()public view returns(string[] memory, uint){
        return getNode(last);
    }

    function getNode(uint idx) public view returns(string[] memory,uint){
        uint limit = mappingNodes[idx].count;
        string[] memory swarms = new string[](limit);
        for (uint i = 0 ; i < limit ; i++){
            swarms[i] = mappingNodes[idx].swarms[i];
        }
        return (swarms,limit);
    }
}