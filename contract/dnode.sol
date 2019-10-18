pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;
import "https://github.com/OpenZeppelin/openzeppelin-solidity/blob/master/contracts/ownership/Ownable.sol";

contract DNode is Ownable{

    struct Node {
      mapping(uint => string) swarms;
      uint count;
    }

    mapping(string => Node)  mappingNodes;
    string private last;
    constructor() public {
        mappingNodes[last] = Node(0);
    }

    function setLast(string memory idx) public onlyOwner{
        last = idx;
        mappingNodes[last] = Node(0);
    }

    function getLast() public view returns(string memory) {
        return last;
    }

    function remove(uint idx) public onlyOwner{
        require(idx < mappingNodes[last].count, "Node: delete node is not exist");
        uint limit = mappingNodes[last].count;
        mappingNodes[last].swarms[idx] = mappingNodes[last].swarms[limit-1];
        mappingNodes[last].count--;
    }

    function store(string memory swarm)public onlyOwner{
         mappingNodes[last].swarms[mappingNodes[last].count] = swarm;
         mappingNodes[last].count++;
    }

    function getLastNode()public view returns(string[] memory, uint){
        return getNode(last);
    }

    function getNode(string memory idx) public view returns(string[] memory,uint){
        uint limit = mappingNodes[idx].count;
        string[] memory swarms = new string[](limit);
        for (uint i = 0 ; i < limit ; i++){
            swarms[i] = mappingNodes[idx].swarms[i];
        }
        return (swarms,limit);
    }
}