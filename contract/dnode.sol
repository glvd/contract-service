pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;
import "https://github.com/OpenZeppelin/openzeppelin-solidity/blob/master/contracts/ownership/Ownable.sol";

contract DNode is Ownable{

    struct Node {
      mapping(uint => string) swarms;
      uint count;
    }

    struct Version{
        string hash;
        string version;
    }

    mapping(uint => Node)  mappingNodes;
    uint private last;

    mapping(uint => Version) mappingVersions;
    uint private lastVersion;

    function setLast(uint idx) public onlyOwner{
        require(idx > last, "Node: last node is newer than now");
        last = idx;
    }

    function getLast() public view returns(uint) {
        return last;
    }

    function getLastVersion() public view returns(string memory,string memory){
        require(lastVersion > 0, "Version: version index is not exist");
        return getVersion(lastVersion-1);
    }

    function getVersion(uint idx)public view returns(string memory,string memory){
        return (mappingVersions[idx].version,mappingVersions[idx].hash);
    }

    function setVersion(uint idx,string memory ver,string memory hash)public{
        require(idx < lastVersion, "Version: version index is not exist");
        mappingVersions[idx] = Version({version:ver,hash:hash});
    }

    function setLastVersion(string memory ver,string memory hash)public{
        mappingVersions[lastVersion] = Version({version:ver,hash:hash});
        lastVersion++;
    }

    function remove(uint idx) public onlyOwner{
        require(idx < mappingNodes[last].count, "Node: delete node is not exist");
        uint limit = mappingNodes[last].count;
        mappingNodes[last].swarms[idx]= mappingNodes[last].swarms[limit-1];
        mappingNodes[last].count--;
    }

    function store(string memory swarm)public onlyOwner{
         mappingNodes[last].swarms[mappingNodes[last].count]= swarm;
         mappingNodes[last].count++;
    }

    function getLastNode()public view returns(string[] memory, uint){
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