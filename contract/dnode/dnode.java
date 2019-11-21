
// This file is an automatically generated Java binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package com.dhash.dnode;

import org.ethereum.geth.*;
import java.util.*;



public class Context {
	// ABI is the input ABI used to generate the binding from.
	public final static String ABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]";
	
	

	// Ethereum address where this contract is located at.
	public final Address Address;

	// Ethereum transaction in which this contract was deployed (if known!).
	public final Transaction Deployer;

	// Contract instance bound to a blockchain address.
	private final BoundContract Contract;

	// Creates a new instance of Context, bound to a specific deployed contract.
	public Context(Address address, EthereumClient client) throws Exception {
		this(Geth.bindContract(address, ABI, client));
	}

	

	
}


public class DNode {
	// ABI is the input ABI used to generate the binding from.
	public final static String ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cid\",\"type\":\"uint256\"}],\"name\":\"getDeviceVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodeLastData\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"removeNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"replaceVersion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_ver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"}],\"name\":\"setDeviceVersion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"setNodeLast\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"setVersion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"swarm\",\"type\":\"string\"}],\"name\":\"storeNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"truncateNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]";
	
		// DNodeFuncSigs maps the 4-byte function signature to its string representation.
		public final static Map<String, String> DNodeFuncSigs;
		static {
			Hashtable<String, String> temp = new Hashtable<String, String>();
			temp.put("ed6c8d0b", "getDeviceVersion(uint256)");
			temp.put("4f0f4aa9", "getNode(uint256)");
			temp.put("812b5e12", "getNodeLast()");
			temp.put("f428efb0", "getNodeLastData()");
			temp.put("b88da759", "getVersion(uint256)");
			temp.put("8f32d59b", "isOwner()");
			temp.put("8da5cb5b", "owner()");
			temp.put("ffd740df", "removeNode(uint256)");
			temp.put("715018a6", "renounceOwnership()");
			temp.put("342edfba", "replaceVersion(uint256,uint256,string,string)");
			temp.put("011fc71a", "setDeviceVersion(uint256,string,string)");
			temp.put("9c90028f", "setNodeLast(uint256)");
			temp.put("69cf7c9a", "setVersion(uint256,string,string)");
			temp.put("0a73ba53", "storeNode(string)");
			temp.put("f2fde38b", "transferOwnership(address)");
			temp.put("5bf1f19e", "truncateNode(uint256)");
			
			DNodeFuncSigs = Collections.unmodifiableMap(temp);
		}
	
	
	// BYTECODE is the compiled bytecode used for deploying new contracts.
	public final static String BYTECODE = "0x608060405260006100176001600160e01b0361006616565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006a565b3390565b6110a4806100796000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80638da5cb5b11610097578063ed6c8d0b11610066578063ed6c8d0b1461020b578063f2fde38b1461021e578063f428efb014610231578063ffd740df1461023957610100565b80638da5cb5b146101ad5780638f32d59b146101c25780639c90028f146101d7578063b88da759146101ea57610100565b80635bf1f19e116100d35780635bf1f19e1461016a57806369cf7c9a1461017d578063715018a614610190578063812b5e121461019857610100565b8063011fc71a146101055780630a73ba531461011a578063342edfba1461012d5780634f0f4aa914610140575b600080fd5b610118610113366004610b97565b61024c565b005b610118610128366004610b44565b61029b565b61011861013b366004610c14565b61030e565b61015361014e366004610b79565b61039a565b604051610161929190610ecf565b60405180910390f35b610118610178366004610b79565b6104bb565b61011861018b366004610b97565b6104f3565b61011861058a565b6101a06105f8565b6040516101619190610f72565b6101b56105ff565b6040516101619190610ec1565b6101ca61060e565b6040516101619190610eef565b6101186101e5366004610b79565b610632565b6101fd6101f8366004610b79565b61067c565b604051610161929190610efd565b6101fd610219366004610b79565b6107be565b61011861022c366004610b1e565b61080f565b61015361083f565b610118610247366004610b79565b610856565b61025461060e565b6102795760405162461bcd60e51b815260040161027090610f42565b60405180910390fd5b6005546000848152600160205260409020556102968383836104f3565b505050565b6102a361060e565b6102bf5760405162461bcd60e51b815260040161027090610f42565b600354600090815260026020908152604080832060018101548452825290912082516102ed928401906109a1565b50506003546000908152600260205260409020600190810180549091019055565b61031661060e565b6103325760405162461bcd60e51b815260040161027090610f42565b60408051606081018252848152602080820184815282840186905260008881526004835293909320825181559251805192939261037592600185019201906109a1565b50604082015180516103919160028401916020909101906109a1565b50505050505050565b600081815260026020908152604080832060010154815181815281840281019093019091526060929183908280156103e657816020015b60608152602001906001900390816103d15790505b50905060005b828110156104b1576000868152600260208181526040808420858552825292839020805484516001821615610100026000190190911693909304601f810183900483028401830190945283835291929083018282801561048d5780601f106104625761010080835404028352916020019161048d565b820191906000526020600020905b81548152906001019060200180831161047057829003601f168201915b505050505082828151811061049e57fe5b60209081029190910101526001016103ec565b5092509050915091565b6104c361060e565b6104df5760405162461bcd60e51b815260040161027090610f42565b600090815260026020526040812060010155565b6104fb61060e565b6105175760405162461bcd60e51b815260040161027090610f42565b60408051606081018252848152602080820184815282840186905260055460009081526004835293909320825181559251805192939261055d92600185019201906109a1565b50604082015180516105799160028401916020909101906109a1565b505060058054600101905550505050565b61059261060e565b6105ae5760405162461bcd60e51b815260040161027090610f42565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6003545b90565b6000546001600160a01b031690565b600080546001600160a01b031661062361091c565b6001600160a01b031614905090565b61063a61060e565b6106565760405162461bcd60e51b815260040161027090610f42565b60035481116106775760405162461bcd60e51b815260040161027090610f62565b600355565b600081815260046020908152604091829020600280820180548551600180831615610100026000190190921693909304601f8101869004860284018601909652858352606095869592949190910192918491908301828280156107205780601f106106f557610100808354040283529160200191610720565b820191906000526020600020905b81548152906001019060200180831161070357829003601f168201915b5050845460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959750869450925084019050828280156107ae5780601f10610783576101008083540402835291602001916107ae565b820191906000526020600020905b81548152906001019060200180831161079157829003601f168201915b5050505050905091509150915091565b60008181526001602052604090205460609081906107ee5760405162461bcd60e51b815260040161027090610f52565b6000838152600160205260409020546108069061067c565b91509150915091565b61081761060e565b6108335760405162461bcd60e51b815260040161027090610f42565b61083c81610920565b50565b6060600061084e60035461039a565b915091509091565b61085e61060e565b61087a5760405162461bcd60e51b815260040161027090610f42565b60035460009081526002602052604090206001015481106108ad5760405162461bcd60e51b815260040161027090610f32565b600354600090815260026020818152604080842060018082015460001980820188529290945282862087875292909520825493956108fb959194908116156101000290920190911604610a1f565b50506003546000908152600260205260409020600101805460001901905550565b3390565b6001600160a01b0381166109465760405162461bcd60e51b815260040161027090610f22565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106109e257805160ff1916838001178555610a0f565b82800160010185558215610a0f579182015b82811115610a0f5782518255916020019190600101906109f4565b50610a1b929150610a94565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610a585780548555610a0f565b82800160010185558215610a0f57600052602060002091601f016020900482015b82811115610a0f578254825591600101919060010190610a79565b6105fc91905b80821115610a1b5760008155600101610a9a565b8035610ab981611044565b92915050565b600082601f830112610ad057600080fd5b8135610ae3610ade82610fa7565b610f80565b91508082526020830160208301858383011115610aff57600080fd5b610b0a838284610ffe565b50505092915050565b8035610ab981611058565b600060208284031215610b3057600080fd5b6000610b3c8484610aae565b949350505050565b600060208284031215610b5657600080fd5b813567ffffffffffffffff811115610b6d57600080fd5b610b3c84828501610abf565b600060208284031215610b8b57600080fd5b6000610b3c8484610b13565b600080600060608486031215610bac57600080fd5b6000610bb88686610b13565b935050602084013567ffffffffffffffff811115610bd557600080fd5b610be186828701610abf565b925050604084013567ffffffffffffffff811115610bfe57600080fd5b610c0a86828701610abf565b9150509250925092565b60008060008060808587031215610c2a57600080fd5b6000610c368787610b13565b9450506020610c4787828801610b13565b935050604085013567ffffffffffffffff811115610c6457600080fd5b610c7087828801610abf565b925050606085013567ffffffffffffffff811115610c8d57600080fd5b610c9987828801610abf565b91505092959194509250565b6000610cb18383610d3e565b9392505050565b610cc181610fe2565b82525050565b6000610cd282610fd5565b610cdc8185610fd9565b935083602082028501610cee85610fcf565b8060005b85811015610d285784840389528151610d0b8582610ca5565b9450610d1683610fcf565b60209a909a0199925050600101610cf2565b5091979650505050505050565b610cc181610fed565b6000610d4982610fd5565b610d538185610fd9565b9350610d6381856020860161100a565b610d6c8161103a565b9093019392505050565b6000610d83602683610fd9565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181526564647265737360d01b602082015260400192915050565b6000610dcb601e83610fd9565b7f4e6f64653a2064656c657465206e6f6465206973206e6f742065786973740000815260200192915050565b6000610e04602083610fd9565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572815260200192915050565b6000610e3d602383610fd9565b7f56657273696f6e3a20696e6465782076657273696f6e206973206e6f742065788152621a5cdd60ea1b602082015260400192915050565b6000610e82602183610fd9565b7f4e6f64653a206c617374206e6f6465206973206e65776572207468616e206e6f8152607760f81b602082015260400192915050565b610cc1816105fc565b60208101610ab98284610cb8565b60408082528101610ee08185610cc7565b9050610cb16020830184610eb8565b60208101610ab98284610d35565b60408082528101610f0e8185610d3e565b90508181036020830152610b3c8184610d3e565b60208082528101610ab981610d76565b60208082528101610ab981610dbe565b60208082528101610ab981610df7565b60208082528101610ab981610e30565b60208082528101610ab981610e75565b60208101610ab98284610eb8565b60405181810167ffffffffffffffff81118282101715610f9f57600080fd5b604052919050565b600067ffffffffffffffff821115610fbe57600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b6000610ab982610ff2565b151590565b6001600160a01b031690565b82818337506000910152565b60005b8381101561102557818101518382015260200161100d565b83811115611034576000848401525b50505050565b601f01601f191690565b61104d81610fe2565b811461083c57600080fd5b61104d816105fc56fea365627a7a72315820a0193a6543d8a7d58ef7bef555559583036f248b079d9b00821d1863a1d5992d6c6578706572696d656e74616cf564736f6c634300050c0040";

	// deploy deploys a new Ethereum contract, binding an instance of DNode to it.
	public static DNode deploy(TransactOpts auth, EthereumClient client) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		String bytecode = BYTECODE;
		
		
		return new DNode(Geth.deployContract(auth, ABI, Geth.decodeFromHex(bytecode), client, args));
	}

	// Internal constructor used by contract deployment.
	private DNode(BoundContract deployment) {
		this.Address  = deployment.getAddress();
		this.Deployer = deployment.getDeployer();
		this.Contract = deployment;
	}
	

	// Ethereum address where this contract is located at.
	public final Address Address;

	// Ethereum transaction in which this contract was deployed (if known!).
	public final Transaction Deployer;

	// Contract instance bound to a blockchain address.
	private final BoundContract Contract;

	// Creates a new instance of DNode, bound to a specific deployed contract.
	public DNode(Address address, EthereumClient client) throws Exception {
		this(Geth.bindContract(address, ABI, client));
	}

	
	
	// GetDeviceVersionResults is the output of a call to getDeviceVersion.
	public class GetDeviceVersionResults {
		public String Return0;
		public String Return1;
		
	}
	

	// getDeviceVersion is a free data retrieval call binding the contract method 0xed6c8d0b.
	//
	// Solidity: function getDeviceVersion(uint256 cid) constant returns(string, string)
	public GetDeviceVersionResults getDeviceVersion(CallOpts opts, BigInt cid) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setBigInt(cid);args.set(0,arg0);
		

		Interfaces results = Geth.newInterfaces(2);
		Interface result0 = Geth.newInterface(); result0.setDefaultString(); results.set(0, result0);
		Interface result1 = Geth.newInterface(); result1.setDefaultString(); results.set(1, result1);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "getDeviceVersion", args);
		
			GetDeviceVersionResults result = new GetDeviceVersionResults();
			result.Return0 = results.get(0).getString();
			result.Return1 = results.get(1).getString();
			
			return result;
		
	}
	
	
	// GetNodeResults is the output of a call to getNode.
	public class GetNodeResults {
		public Strings Return0;
		public BigInt Return1;
		
	}
	

	// getNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
	//
	// Solidity: function getNode(uint256 idx) constant returns(string[], uint256)
	public GetNodeResults getNode(CallOpts opts, BigInt idx) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setBigInt(idx);args.set(0,arg0);
		

		Interfaces results = Geth.newInterfaces(2);
		Interface result0 = Geth.newInterface(); result0.setDefaultStrings(); results.set(0, result0);
		Interface result1 = Geth.newInterface(); result1.setDefaultBigInt(); results.set(1, result1);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "getNode", args);
		
			GetNodeResults result = new GetNodeResults();
			result.Return0 = results.get(0).getStrings();
			result.Return1 = results.get(1).getBigInt();
			
			return result;
		
	}
	
	

	// getNodeLast is a free data retrieval call binding the contract method 0x812b5e12.
	//
	// Solidity: function getNodeLast() constant returns(uint256)
	public BigInt getNodeLast(CallOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		

		Interfaces results = Geth.newInterfaces(1);
		Interface result0 = Geth.newInterface(); result0.setDefaultBigInt(); results.set(0, result0);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "getNodeLast", args);
		return results.get(0).getBigInt();
		
	}
	
	
	// GetNodeLastDataResults is the output of a call to getNodeLastData.
	public class GetNodeLastDataResults {
		public Strings Return0;
		public BigInt Return1;
		
	}
	

	// getNodeLastData is a free data retrieval call binding the contract method 0xf428efb0.
	//
	// Solidity: function getNodeLastData() constant returns(string[], uint256)
	public GetNodeLastDataResults getNodeLastData(CallOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		

		Interfaces results = Geth.newInterfaces(2);
		Interface result0 = Geth.newInterface(); result0.setDefaultStrings(); results.set(0, result0);
		Interface result1 = Geth.newInterface(); result1.setDefaultBigInt(); results.set(1, result1);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "getNodeLastData", args);
		
			GetNodeLastDataResults result = new GetNodeLastDataResults();
			result.Return0 = results.get(0).getStrings();
			result.Return1 = results.get(1).getBigInt();
			
			return result;
		
	}
	
	
	// GetVersionResults is the output of a call to getVersion.
	public class GetVersionResults {
		public String Return0;
		public String Return1;
		
	}
	

	// getVersion is a free data retrieval call binding the contract method 0xb88da759.
	//
	// Solidity: function getVersion(uint256 idx) constant returns(string, string)
	public GetVersionResults getVersion(CallOpts opts, BigInt idx) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setBigInt(idx);args.set(0,arg0);
		

		Interfaces results = Geth.newInterfaces(2);
		Interface result0 = Geth.newInterface(); result0.setDefaultString(); results.set(0, result0);
		Interface result1 = Geth.newInterface(); result1.setDefaultString(); results.set(1, result1);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "getVersion", args);
		
			GetVersionResults result = new GetVersionResults();
			result.Return0 = results.get(0).getString();
			result.Return1 = results.get(1).getString();
			
			return result;
		
	}
	
	

	// isOwner is a free data retrieval call binding the contract method 0x8f32d59b.
	//
	// Solidity: function isOwner() constant returns(bool)
	public boolean isOwner(CallOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		

		Interfaces results = Geth.newInterfaces(1);
		Interface result0 = Geth.newInterface(); result0.setDefaultBool(); results.set(0, result0);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "isOwner", args);
		return results.get(0).getBool();
		
	}
	
	

	// owner is a free data retrieval call binding the contract method 0x8da5cb5b.
	//
	// Solidity: function owner() constant returns(address)
	public Address owner(CallOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		

		Interfaces results = Geth.newInterfaces(1);
		Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "owner", args);
		return results.get(0).getAddress();
		
	}
	

	
	// removeNode is a paid mutator transaction binding the contract method 0xffd740df.
	//
	// Solidity: function removeNode(uint256 idx) returns()
	public Transaction removeNode(TransactOpts opts, BigInt idx) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setBigInt(idx);args.set(0,arg0);
		
		return this.Contract.transact(opts, "removeNode"	, args);
	}
	
	// renounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
	//
	// Solidity: function renounceOwnership() returns()
	public Transaction renounceOwnership(TransactOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		
		return this.Contract.transact(opts, "renounceOwnership"	, args);
	}
	
	// replaceVersion is a paid mutator transaction binding the contract method 0x342edfba.
	//
	// Solidity: function replaceVersion(uint256 idx, uint256 cid, string ver, string hash) returns()
	public Transaction replaceVersion(TransactOpts opts, BigInt idx, BigInt cid, String ver, String hash) throws Exception {
		Interfaces args = Geth.newInterfaces(4);
		Interface arg0 = Geth.newInterface();arg0.setBigInt(idx);args.set(0,arg0);
		Interface arg1 = Geth.newInterface();arg1.setBigInt(cid);args.set(1,arg1);
		Interface arg2 = Geth.newInterface();arg2.setString(ver);args.set(2,arg2);
		Interface arg3 = Geth.newInterface();arg3.setString(hash);args.set(3,arg3);
		
		return this.Contract.transact(opts, "replaceVersion"	, args);
	}
	
	// setDeviceVersion is a paid mutator transaction binding the contract method 0x011fc71a.
	//
	// Solidity: function setDeviceVersion(uint256 _cid, string _ver, string _hash) returns()
	public Transaction setDeviceVersion(TransactOpts opts, BigInt _cid, String _ver, String _hash) throws Exception {
		Interfaces args = Geth.newInterfaces(3);
		Interface arg0 = Geth.newInterface();arg0.setBigInt(_cid);args.set(0,arg0);
		Interface arg1 = Geth.newInterface();arg1.setString(_ver);args.set(1,arg1);
		Interface arg2 = Geth.newInterface();arg2.setString(_hash);args.set(2,arg2);
		
		return this.Contract.transact(opts, "setDeviceVersion"	, args);
	}
	
	// setNodeLast is a paid mutator transaction binding the contract method 0x9c90028f.
	//
	// Solidity: function setNodeLast(uint256 idx) returns()
	public Transaction setNodeLast(TransactOpts opts, BigInt idx) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setBigInt(idx);args.set(0,arg0);
		
		return this.Contract.transact(opts, "setNodeLast"	, args);
	}
	
	// setVersion is a paid mutator transaction binding the contract method 0x69cf7c9a.
	//
	// Solidity: function setVersion(uint256 cid, string ver, string hash) returns()
	public Transaction setVersion(TransactOpts opts, BigInt cid, String ver, String hash) throws Exception {
		Interfaces args = Geth.newInterfaces(3);
		Interface arg0 = Geth.newInterface();arg0.setBigInt(cid);args.set(0,arg0);
		Interface arg1 = Geth.newInterface();arg1.setString(ver);args.set(1,arg1);
		Interface arg2 = Geth.newInterface();arg2.setString(hash);args.set(2,arg2);
		
		return this.Contract.transact(opts, "setVersion"	, args);
	}
	
	// storeNode is a paid mutator transaction binding the contract method 0x0a73ba53.
	//
	// Solidity: function storeNode(string swarm) returns()
	public Transaction storeNode(TransactOpts opts, String swarm) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setString(swarm);args.set(0,arg0);
		
		return this.Contract.transact(opts, "storeNode"	, args);
	}
	
	// transferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
	//
	// Solidity: function transferOwnership(address newOwner) returns()
	public Transaction transferOwnership(TransactOpts opts, Address newOwner) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setAddress(newOwner);args.set(0,arg0);
		
		return this.Contract.transact(opts, "transferOwnership"	, args);
	}
	
	// truncateNode is a paid mutator transaction binding the contract method 0x5bf1f19e.
	//
	// Solidity: function truncateNode(uint256 idx) returns()
	public Transaction truncateNode(TransactOpts opts, BigInt idx) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setBigInt(idx);args.set(0,arg0);
		
		return this.Contract.transact(opts, "truncateNode"	, args);
	}
	
}


public class Ownable {
	// ABI is the input ABI used to generate the binding from.
	public final static String ABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]";
	
		// OwnableFuncSigs maps the 4-byte function signature to its string representation.
		public final static Map<String, String> OwnableFuncSigs;
		static {
			Hashtable<String, String> temp = new Hashtable<String, String>();
			temp.put("8f32d59b", "isOwner()");
			temp.put("8da5cb5b", "owner()");
			temp.put("715018a6", "renounceOwnership()");
			temp.put("f2fde38b", "transferOwnership(address)");
			
			OwnableFuncSigs = Collections.unmodifiableMap(temp);
		}
	
	

	// Ethereum address where this contract is located at.
	public final Address Address;

	// Ethereum transaction in which this contract was deployed (if known!).
	public final Transaction Deployer;

	// Contract instance bound to a blockchain address.
	private final BoundContract Contract;

	// Creates a new instance of Ownable, bound to a specific deployed contract.
	public Ownable(Address address, EthereumClient client) throws Exception {
		this(Geth.bindContract(address, ABI, client));
	}

	
	

	// isOwner is a free data retrieval call binding the contract method 0x8f32d59b.
	//
	// Solidity: function isOwner() constant returns(bool)
	public boolean isOwner(CallOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		

		Interfaces results = Geth.newInterfaces(1);
		Interface result0 = Geth.newInterface(); result0.setDefaultBool(); results.set(0, result0);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "isOwner", args);
		return results.get(0).getBool();
		
	}
	
	

	// owner is a free data retrieval call binding the contract method 0x8da5cb5b.
	//
	// Solidity: function owner() constant returns(address)
	public Address owner(CallOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		

		Interfaces results = Geth.newInterfaces(1);
		Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "owner", args);
		return results.get(0).getAddress();
		
	}
	

	
	// renounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
	//
	// Solidity: function renounceOwnership() returns()
	public Transaction renounceOwnership(TransactOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		
		return this.Contract.transact(opts, "renounceOwnership"	, args);
	}
	
	// transferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
	//
	// Solidity: function transferOwnership(address newOwner) returns()
	public Transaction transferOwnership(TransactOpts opts, Address newOwner) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setAddress(newOwner);args.set(0,arg0);
		
		return this.Contract.transact(opts, "transferOwnership"	, args);
	}
	
}

