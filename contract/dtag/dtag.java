
// This file is an automatically generated Java binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package com.dhash.dtag;

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


public class DMessage {
	// ABI is the input ABI used to generate the binding from.
	public final static String ABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldWriter\",\"type\":\"address\"}],\"name\":\"WritershipDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newWriter\",\"type\":\"address\"}],\"name\":\"WritershipIncreased\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"addMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"ids\",\"type\":\"string[]\"}],\"name\":\"checkAllMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"checkMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"ids\",\"type\":\"string[]\"}],\"name\":\"checkMessages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldWriter\",\"type\":\"address\"}],\"name\":\"decreaseWritership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"delMessage\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"ids\",\"type\":\"string[]\"}],\"name\":\"getIdsMessages\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"_value\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"getMessage\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getMessages\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"_value\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getMsgId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_value\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMsgIds\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMsgLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWriter\",\"type\":\"address\"}],\"name\":\"increasedWritership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isWriter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceWritership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"updateMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"writer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"writers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_value\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]";
	
		// DMessageFuncSigs maps the 4-byte function signature to its string representation.
		public final static Map<String, String> DMessageFuncSigs;
		static {
			Hashtable<String, String> temp = new Hashtable<String, String>();
			temp.put("631e0c69", "addMessage(string,string)");
			temp.put("16313e92", "checkAllMessage(string[])");
			temp.put("1873ec68", "checkMessage(string)");
			temp.put("834daf0c", "checkMessages(string[])");
			temp.put("9f217bdf", "decreaseWritership(address)");
			temp.put("2309bc96", "delMessage(string)");
			temp.put("6c4ef994", "getIdsMessages(string[])");
			temp.put("0cc4e8d8", "getMessage(string)");
			temp.put("4be8362d", "getMessages(uint256,uint256)");
			temp.put("95cb2638", "getMsgId(uint256)");
			temp.put("68bfc43c", "getMsgIds()");
			temp.put("085c00e1", "getMsgLength()");
			temp.put("14b4d411", "increasedWritership(address)");
			temp.put("8f32d59b", "isOwner()");
			temp.put("4d6ee9fd", "isWriter()");
			temp.put("8da5cb5b", "owner()");
			temp.put("715018a6", "renounceOwnership()");
			temp.put("4075e83e", "renounceWritership()");
			temp.put("f2fde38b", "transferOwnership(address)");
			temp.put("75a0f3a4", "updateMessage(string,string)");
			temp.put("453a2abc", "writer()");
			temp.put("8b2f5369", "writers()");
			
			DMessageFuncSigs = Collections.unmodifiableMap(temp);
		}
	
	
	// BYTECODE is the compiled bytecode used for deploying new contracts.
	public final static String BYTECODE = "0x608060405234801561001057600080fd5b5060006100246001600160e01b036100e216565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35060006100816001600160e01b036100e216565b6001600160a01b0381166000818152600160208190526040808320805460ff191690921790915560038290555192935090917f8f2adc2e2f2723d375f8bb1359cb1027f713945395ca36c42b83dd94c5bdc7f29190a25060006007556100e6565b3390565b6118f5806100f56000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c806368bfc43c116100b85780638b2f53691161007c5780638b2f5369146102ab5780638da5cb5b146102c05780638f32d59b146102c857806395cb2638146102d05780639f217bdf146102e3578063f2fde38b146102f657610142565b806368bfc43c146102475780636c4ef9941461025c578063715018a61461026f57806375a0f3a414610277578063834daf0c1461028a57610142565b80632309bc961161010a5780632309bc96146101cd5780634075e83e146101ee578063453a2abc146101f65780634be8362d1461020b5780634d6ee9fd1461022c578063631e0c691461023457610142565b8063085c00e1146101475780630cc4e8d81461016557806314b4d4111461018557806316313e921461019a5780631873ec68146101ba575b600080fd5b61014f610309565b60405161015c9190611787565b60405180910390f35b6101786101733660046111a3565b610310565b60405161015c91906116c6565b610198610193366004611148565b6103be565b005b6101ad6101a836600461116e565b6103f7565b60405161015c91906116b8565b6101ad6101c83660046111a3565b61043f565b6101e06101db3660046111a3565b61046b565b60405161015c9291906116d7565b6101986105d7565b6101fe610670565b60405161015c9190611668565b61021e61021936600461125f565b610699565b60405161015c929190611698565b6101ad610728565b6101ad6102423660046111d8565b610758565b61024f610832565b60405161015c9190611687565b61021e61026a36600461116e565b610938565b6101986109c3565b6101ad6102853660046111d8565b610a31565b61029d61029836600461116e565b610ad7565b60405161015c929190611795565b6102b3610b18565b60405161015c9190611676565b6101fe610ba3565b6101ad610bb2565b6101786102de366004611241565b610bd6565b6101986102f1366004611148565b610c40565b610198610304366004611148565b610c6d565b6007545b90565b6060600582604051610322919061165c565b9081526040805160209281900383018120805460026001821615610100026000190190911604601f810185900485028301850190935282825290929091908301828280156103b15780601f10610386576101008083540402835291602001916103b1565b820191906000526020600020905b81548152906001019060200180831161039457829003601f168201915b505050505090505b919050565b6103c6610bb2565b6103eb5760405162461bcd60e51b81526004016103e290611767565b60405180910390fd5b6103f481610c9a565b50565b6000805b82518110156104365761042083828151811061041357fe5b602002602001015161043f565b61042e5760009150506103b9565b6001016103fb565b50600192915050565b6000600482604051610451919061165c565b9081526040519081900360200190205460ff169050919050565b60606000610477610bb2565b6104935760405162461bcd60e51b81526004016103e290611767565b6004836040516104a3919061165c565b9081526040519081900360200190205460ff1615156001146104d75760405162461bcd60e51b81526004016103e290611757565b60006004846040516104e9919061165c565b908152604051908190036020019020805491151560ff19909216919091179055610511610d6b565b600583604051610521919061165c565b90815260200160405180910390206001818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105c65780601f1061059b576101008083540402835291602001916105c6565b820191906000526020600020905b8154815290600101906020018083116105a957829003601f168201915b50505050509150915091505b915091565b6105df610bb2565b6105fb5760405162461bcd60e51b81526004016103e290611767565b600060016000610609610d9e565b6001600160a01b031681526020810191909152604001600020805460ff1916911515919091179055610639610d9e565b6001600160a01b03167fb24c36b70cc67df9ea22f053c15c4a41e7170d7be0637d2637a3445a266ac6af60405160405180910390a2565b600061067a610728565b1515600114156106935761068c610d9e565b905061030d565b50600090565b60606000600754838501106106b057836007540392505b826040519080825280602002602001820160405280156106e457816020015b60608152602001906001900390816106cf5790505b50915060005b8381101561071f57610700610173828701610bd6565b83828151811061070c57fe5b60209081029190910101526001016106ea565b50909391925050565b600060016000610736610d9e565b6001600160a01b0316815260208101919091526040016000205460ff16905090565b6000610762610728565b61077e5760405162461bcd60e51b81526004016103e290611747565b60048360405161078e919061165c565b9081526040519081900360200190205460ff16156107be5760405162461bcd60e51b81526004016103e290611707565b816005846040516107cf919061165c565b908152602001604051809103902090805190602001906107f0929190610f6e565b506001600484604051610803919061165c565b908152604051908190036020019020805491151560ff1990921691909117905561043683610da2565b92915050565b60608060075460405190808252806020026020018201604052801561086b57816020015b60608152602001906001900390816108565790505b50905060005b6007548110156109325760008181526006602090815260409182902080548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452909183018282801561090e5780601f106108e35761010080835404028352916020019161090e565b820191906000526020600020905b8154815290600101906020018083116108f157829003601f168201915b505050505082828151811061091f57fe5b6020908102919091010152600101610871565b50905090565b60606000825190508060405190808252806020026020018201604052801561097457816020015b606081526020019060019003908161095f5790505b50915060005b818110156109bd5761099e84828151811061099157fe5b6020026020010151610310565b8382815181106109aa57fe5b602090810291909101015260010161097a565b50915091565b6109cb610bb2565b6109e75760405162461bcd60e51b81526004016103e290611767565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000610a3b610728565b610a575760405162461bcd60e51b81526004016103e290611747565b600483604051610a67919061165c565b9081526040519081900360200190205460ff161515600114610a9b5760405162461bcd60e51b81526004016103e290611777565b81600584604051610aac919061165c565b90815260200160405180910390209080519060200190610acd929190610f6e565b5060019392505050565b600080805b8351811015610b0b57610af484828151811061041357fe5b610b03579150600090506105d2565b600101610adc565b5060009360019350915050565b6060600354604051908082528060200260200182016040528015610b46578160200160208202803883390190505b50905060005b600354811015610b9f5760008181526002602052604090205482516001600160a01b0390911690839083908110610b7f57fe5b6001600160a01b0390921660209283029190910190910152600101610b4c565b5090565b6000546001600160a01b031690565b600080546001600160a01b0316610bc7610d9e565b6001600160a01b031614905090565b60008181526006602090815260409182902080548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452606093928301828280156103b15780601f10610386576101008083540402835291602001916103b1565b610c48610bb2565b610c645760405162461bcd60e51b81526004016103e290611767565b6103f481610dd1565b610c75610bb2565b610c915760405162461bcd60e51b81526004016103e290611767565b6103f481610e94565b6001600160a01b038116610cc05760405162461bcd60e51b81526004016103e290611737565b6001600160a01b03811660009081526001602052604090205460ff1615610cf95760405162461bcd60e51b81526004016103e290611727565b6001600160a01b0381166000818152600160208181526040808420805460ff191684179055600380548552600290925280842080546001600160a01b0319168617905581549092019055517f8f2adc2e2f2723d375f8bb1359cb1027f713945395ca36c42b83dd94c5bdc7f29190a250565b60005b610d76610309565b8110156103f457610d896101c882610bd6565b610d9657610d9681610f15565b600101610d6e565b3390565b60075460009081526006602090815260409091208251610dc492840190610f6e565b5050600780546001019055565b6001600160a01b03811660009081526001602081905260409091205460ff16151514610e0f5760405162461bcd60e51b81526004016103e2906116f7565b6001600160a01b038116610e355760405162461bcd60e51b81526004016103e290611737565b6001600160a01b0381166000908152600160205260409020805460ff19169055610e5d610d6b565b6040516001600160a01b038216907fb24c36b70cc67df9ea22f053c15c4a41e7170d7be0637d2637a3445a266ac6af90600090a250565b6001600160a01b038116610eba5760405162461bcd60e51b81526004016103e290611717565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6007548110610f23576103f4565b60075460001990810160009081526006602052604080822084835291208154610f6093919291600261010060018416150290910190911604610fe8565b506007805460001901905550565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610faf57805160ff1916838001178555610fdc565b82800160010185558215610fdc579182015b82811115610fdc578251825591602001919060010190610fc1565b50610b9f92915061105d565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106110215780548555610fdc565b82800160010185558215610fdc57600052602060002091601f016020900482015b82811115610fdc578254825591600101919060010190611042565b61030d91905b80821115610b9f5760008155600101611063565b803561082c81611895565b600082601f83011261109357600080fd5b81356110a66110a1826117d7565b6117b0565b81815260209384019390925082018360005b838110156110e457813586016110ce88826110ee565b84525060209283019291909101906001016110b8565b5050505092915050565b600082601f8301126110ff57600080fd5b813561110d6110a1826117f8565b9150808252602083016020830185838301111561112957600080fd5b61113483828461184f565b50505092915050565b803561082c816118a9565b60006020828403121561115a57600080fd5b60006111668484611077565b949350505050565b60006020828403121561118057600080fd5b813567ffffffffffffffff81111561119757600080fd5b61116684828501611082565b6000602082840312156111b557600080fd5b813567ffffffffffffffff8111156111cc57600080fd5b611166848285016110ee565b600080604083850312156111eb57600080fd5b823567ffffffffffffffff81111561120257600080fd5b61120e858286016110ee565b925050602083013567ffffffffffffffff81111561122b57600080fd5b611237858286016110ee565b9150509250929050565b60006020828403121561125357600080fd5b6000611166848461113d565b6000806040838503121561127257600080fd5b600061127e858561113d565b92505060206112378582860161113d565b600061129b83836112b6565b505060200190565b60006112af8383611395565b9392505050565b6112bf81611833565b82525050565b60006112d082611826565b6112da818561182a565b93506112e583611820565b8060005b838110156113135781516112fd888261128f565b975061130883611820565b9250506001016112e9565b509495945050505050565b600061132982611826565b611333818561182a565b93508360208202850161134585611820565b8060005b8581101561137f578484038952815161136285826112a3565b945061136d83611820565b60209a909a0199925050600101611349565b5091979650505050505050565b6112bf8161183e565b60006113a082611826565b6113aa818561182a565b93506113ba81856020860161185b565b6113c38161188b565b9093019392505050565b60006113d882611826565b6113e281856103b9565b93506113f281856020860161185b565b9290920192915050565b600061140960258361182a565b7f57726974657261626c653a206f6c642077726974657220776173206e6f7420778152643934ba32b960d91b602082015260400192915050565b6000611450601d8361182a565b7f4d6573736167653a20616464206d657373616765206973206578697374000000815260200192915050565b600061148960268361182a565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181526564647265737360d01b602082015260400192915050565b60006114d1601f8361182a565b7f57726974657261626c653a206e65772077726974657220697320657869737400815260200192915050565b600061150a602a8361182a565b7f57726974657261626c653a206e65772077726974657220697320746865207a65815269726f206164647265737360b01b602082015260400192915050565b600061155660238361182a565b7f577269746561626c653a2063616c6c6572206973206e6f7420746865207772698152623a32b960e91b602082015260400192915050565b600061159b60248361182a565b7f4d6573736167653a2064656c657465206d657373616765206973206e6f7420658152631e1a5cdd60e21b602082015260400192915050565b60006115e160208361182a565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572815260200192915050565b600061161a60248361182a565b7f4d6573736167653a20757064617465206d657373616765206973206e6f7420658152631e1a5cdd60e21b602082015260400192915050565b6112bf8161030d565b60006112af82846113cd565b6020810161082c82846112b6565b602080825281016112af81846112c5565b602080825281016112af818461131e565b604080825281016116a9818561131e565b90506112af6020830184611653565b6020810161082c828461138c565b602080825281016112af8184611395565b604080825281016116e88185611395565b90506112af602083018461138c565b6020808252810161082c816113fc565b6020808252810161082c81611443565b6020808252810161082c8161147c565b6020808252810161082c816114c4565b6020808252810161082c816114fd565b6020808252810161082c81611549565b6020808252810161082c8161158e565b6020808252810161082c816115d4565b6020808252810161082c8161160d565b6020810161082c8284611653565b604081016117a38285611653565b6112af602083018461138c565b60405181810167ffffffffffffffff811182821017156117cf57600080fd5b604052919050565b600067ffffffffffffffff8211156117ee57600080fd5b5060209081020190565b600067ffffffffffffffff82111561180f57600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b600061082c82611843565b151590565b6001600160a01b031690565b82818337506000910152565b60005b8381101561187657818101518382015260200161185e565b83811115611885576000848401525b50505050565b601f01601f191690565b61189e81611833565b81146103f457600080fd5b61189e8161030d56fea365627a7a723158202a24e1f5f0cf350cb56a28b69f9a6d1414df00cc6274a90f9b8a3b1ea7b527286c6578706572696d656e74616cf564736f6c634300050c0040";

	// deploy deploys a new Ethereum contract, binding an instance of DMessage to it.
	public static DMessage deploy(TransactOpts auth, EthereumClient client) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		String bytecode = BYTECODE;
		
		
		return new DMessage(Geth.deployContract(auth, ABI, Geth.decodeFromHex(bytecode), client, args));
	}

	// Internal constructor used by contract deployment.
	private DMessage(BoundContract deployment) {
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

	// Creates a new instance of DMessage, bound to a specific deployed contract.
	public DMessage(Address address, EthereumClient client) throws Exception {
		this(Geth.bindContract(address, ABI, client));
	}

	
	

	// checkAllMessage is a free data retrieval call binding the contract method 0x16313e92.
	//
	// Solidity: function checkAllMessage(string[] ids) constant returns(bool)
	public boolean checkAllMessage(CallOpts opts, Strings ids) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setStrings(ids);args.set(0,arg0);
		

		Interfaces results = Geth.newInterfaces(1);
		Interface result0 = Geth.newInterface(); result0.setDefaultBool(); results.set(0, result0);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "checkAllMessage", args);
		return results.get(0).getBool();
		
	}
	
	

	// checkMessage is a free data retrieval call binding the contract method 0x1873ec68.
	//
	// Solidity: function checkMessage(string id) constant returns(bool)
	public boolean checkMessage(CallOpts opts, String id) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setString(id);args.set(0,arg0);
		

		Interfaces results = Geth.newInterfaces(1);
		Interface result0 = Geth.newInterface(); result0.setDefaultBool(); results.set(0, result0);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "checkMessage", args);
		return results.get(0).getBool();
		
	}
	
	
	// CheckMessagesResults is the output of a call to checkMessages.
	public class CheckMessagesResults {
		public BigInt Return0;
		public boolean Return1;
		
	}
	

	// checkMessages is a free data retrieval call binding the contract method 0x834daf0c.
	//
	// Solidity: function checkMessages(string[] ids) constant returns(uint256, bool)
	public CheckMessagesResults checkMessages(CallOpts opts, Strings ids) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setStrings(ids);args.set(0,arg0);
		

		Interfaces results = Geth.newInterfaces(2);
		Interface result0 = Geth.newInterface(); result0.setDefaultBigInt(); results.set(0, result0);
		Interface result1 = Geth.newInterface(); result1.setDefaultBool(); results.set(1, result1);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "checkMessages", args);
		
			CheckMessagesResults result = new CheckMessagesResults();
			result.Return0 = results.get(0).getBigInt();
			result.Return1 = results.get(1).getBool();
			
			return result;
		
	}
	
	
	// GetIdsMessagesResults is the output of a call to getIdsMessages.
	public class GetIdsMessagesResults {
		public Strings Value;
		public BigInt Size;
		
	}
	

	// getIdsMessages is a free data retrieval call binding the contract method 0x6c4ef994.
	//
	// Solidity: function getIdsMessages(string[] ids) constant returns(string[] _value, uint256 _size)
	public GetIdsMessagesResults getIdsMessages(CallOpts opts, Strings ids) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setStrings(ids);args.set(0,arg0);
		

		Interfaces results = Geth.newInterfaces(2);
		Interface result0 = Geth.newInterface(); result0.setDefaultStrings(); results.set(0, result0);
		Interface result1 = Geth.newInterface(); result1.setDefaultBigInt(); results.set(1, result1);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "getIdsMessages", args);
		
			GetIdsMessagesResults result = new GetIdsMessagesResults();
			result.Value = results.get(0).getStrings();
			result.Size = results.get(1).getBigInt();
			
			return result;
		
	}
	
	

	// getMessage is a free data retrieval call binding the contract method 0x0cc4e8d8.
	//
	// Solidity: function getMessage(string id) constant returns(string)
	public String getMessage(CallOpts opts, String id) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setString(id);args.set(0,arg0);
		

		Interfaces results = Geth.newInterfaces(1);
		Interface result0 = Geth.newInterface(); result0.setDefaultString(); results.set(0, result0);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "getMessage", args);
		return results.get(0).getString();
		
	}
	
	
	// GetMessagesResults is the output of a call to getMessages.
	public class GetMessagesResults {
		public Strings Value;
		public BigInt Size;
		
	}
	

	// getMessages is a free data retrieval call binding the contract method 0x4be8362d.
	//
	// Solidity: function getMessages(uint256 start, uint256 limit) constant returns(string[] _value, uint256 _size)
	public GetMessagesResults getMessages(CallOpts opts, BigInt start, BigInt limit) throws Exception {
		Interfaces args = Geth.newInterfaces(2);
		Interface arg0 = Geth.newInterface();arg0.setBigInt(start);args.set(0,arg0);
		Interface arg1 = Geth.newInterface();arg1.setBigInt(limit);args.set(1,arg1);
		

		Interfaces results = Geth.newInterfaces(2);
		Interface result0 = Geth.newInterface(); result0.setDefaultStrings(); results.set(0, result0);
		Interface result1 = Geth.newInterface(); result1.setDefaultBigInt(); results.set(1, result1);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "getMessages", args);
		
			GetMessagesResults result = new GetMessagesResults();
			result.Value = results.get(0).getStrings();
			result.Size = results.get(1).getBigInt();
			
			return result;
		
	}
	
	

	// getMsgId is a free data retrieval call binding the contract method 0x95cb2638.
	//
	// Solidity: function getMsgId(uint256 _index) constant returns(string _value)
	public String getMsgId(CallOpts opts, BigInt _index) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setBigInt(_index);args.set(0,arg0);
		

		Interfaces results = Geth.newInterfaces(1);
		Interface result0 = Geth.newInterface(); result0.setDefaultString(); results.set(0, result0);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "getMsgId", args);
		return results.get(0).getString();
		
	}
	
	

	// getMsgIds is a free data retrieval call binding the contract method 0x68bfc43c.
	//
	// Solidity: function getMsgIds() constant returns(string[])
	public Strings getMsgIds(CallOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		

		Interfaces results = Geth.newInterfaces(1);
		Interface result0 = Geth.newInterface(); result0.setDefaultStrings(); results.set(0, result0);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "getMsgIds", args);
		return results.get(0).getStrings();
		
	}
	
	

	// getMsgLength is a free data retrieval call binding the contract method 0x085c00e1.
	//
	// Solidity: function getMsgLength() constant returns(uint256)
	public BigInt getMsgLength(CallOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		

		Interfaces results = Geth.newInterfaces(1);
		Interface result0 = Geth.newInterface(); result0.setDefaultBigInt(); results.set(0, result0);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "getMsgLength", args);
		return results.get(0).getBigInt();
		
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
	
	

	// isWriter is a free data retrieval call binding the contract method 0x4d6ee9fd.
	//
	// Solidity: function isWriter() constant returns(bool)
	public boolean isWriter(CallOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		

		Interfaces results = Geth.newInterfaces(1);
		Interface result0 = Geth.newInterface(); result0.setDefaultBool(); results.set(0, result0);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "isWriter", args);
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
	
	

	// writer is a free data retrieval call binding the contract method 0x453a2abc.
	//
	// Solidity: function writer() constant returns(address)
	public Address writer(CallOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		

		Interfaces results = Geth.newInterfaces(1);
		Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "writer", args);
		return results.get(0).getAddress();
		
	}
	
	

	// writers is a free data retrieval call binding the contract method 0x8b2f5369.
	//
	// Solidity: function writers() constant returns(address[] _value)
	public Addresses writers(CallOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		

		Interfaces results = Geth.newInterfaces(1);
		Interface result0 = Geth.newInterface(); result0.setDefaultAddresses(); results.set(0, result0);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "writers", args);
		return results.get(0).getAddresses();
		
	}
	

	
	// addMessage is a paid mutator transaction binding the contract method 0x631e0c69.
	//
	// Solidity: function addMessage(string id, string message) returns(bool)
	public Transaction addMessage(TransactOpts opts, String id, String message) throws Exception {
		Interfaces args = Geth.newInterfaces(2);
		Interface arg0 = Geth.newInterface();arg0.setString(id);args.set(0,arg0);
		Interface arg1 = Geth.newInterface();arg1.setString(message);args.set(1,arg1);
		
		return this.Contract.transact(opts, "addMessage"	, args);
	}
	
	// decreaseWritership is a paid mutator transaction binding the contract method 0x9f217bdf.
	//
	// Solidity: function decreaseWritership(address oldWriter) returns()
	public Transaction decreaseWritership(TransactOpts opts, Address oldWriter) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setAddress(oldWriter);args.set(0,arg0);
		
		return this.Contract.transact(opts, "decreaseWritership"	, args);
	}
	
	// delMessage is a paid mutator transaction binding the contract method 0x2309bc96.
	//
	// Solidity: function delMessage(string id) returns(string, bool)
	public Transaction delMessage(TransactOpts opts, String id) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setString(id);args.set(0,arg0);
		
		return this.Contract.transact(opts, "delMessage"	, args);
	}
	
	// increasedWritership is a paid mutator transaction binding the contract method 0x14b4d411.
	//
	// Solidity: function increasedWritership(address newWriter) returns()
	public Transaction increasedWritership(TransactOpts opts, Address newWriter) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setAddress(newWriter);args.set(0,arg0);
		
		return this.Contract.transact(opts, "increasedWritership"	, args);
	}
	
	// renounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
	//
	// Solidity: function renounceOwnership() returns()
	public Transaction renounceOwnership(TransactOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		
		return this.Contract.transact(opts, "renounceOwnership"	, args);
	}
	
	// renounceWritership is a paid mutator transaction binding the contract method 0x4075e83e.
	//
	// Solidity: function renounceWritership() returns()
	public Transaction renounceWritership(TransactOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		
		return this.Contract.transact(opts, "renounceWritership"	, args);
	}
	
	// transferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
	//
	// Solidity: function transferOwnership(address newOwner) returns()
	public Transaction transferOwnership(TransactOpts opts, Address newOwner) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setAddress(newOwner);args.set(0,arg0);
		
		return this.Contract.transact(opts, "transferOwnership"	, args);
	}
	
	// updateMessage is a paid mutator transaction binding the contract method 0x75a0f3a4.
	//
	// Solidity: function updateMessage(string id, string message) returns(bool)
	public Transaction updateMessage(TransactOpts opts, String id, String message) throws Exception {
		Interfaces args = Geth.newInterfaces(2);
		Interface arg0 = Geth.newInterface();arg0.setString(id);args.set(0,arg0);
		Interface arg1 = Geth.newInterface();arg1.setString(message);args.set(1,arg1);
		
		return this.Contract.transact(opts, "updateMessage"	, args);
	}
	
}


public class DTag {
	// ABI is the input ABI used to generate the binding from.
	public final static String ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_sub\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"addTagId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_sub\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_ids\",\"type\":\"string[]\"}],\"name\":\"addTagIds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_sub\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"}],\"name\":\"addTagMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_sub\",\"type\":\"string\"}],\"name\":\"checkTag\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_sub\",\"type\":\"string\"}],\"name\":\"checkTagIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_sub\",\"type\":\"string\"}],\"name\":\"fixTagIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_sub\",\"type\":\"string\"}],\"name\":\"getTagIds\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_sub\",\"type\":\"string\"}],\"name\":\"getTagMessage\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"_value\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_sub\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_ids\",\"type\":\"string[]\"}],\"name\":\"replaceTagIds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_sub\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_ids\",\"type\":\"string[]\"}],\"name\":\"setTagIds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]";
	
		// DTagFuncSigs maps the 4-byte function signature to its string representation.
		public final static Map<String, String> DTagFuncSigs;
		static {
			Hashtable<String, String> temp = new Hashtable<String, String>();
			temp.put("16017c65", "addTagId(string,string,string)");
			temp.put("e314604e", "addTagIds(string,string,string[])");
			temp.put("31c2c83f", "addTagMessage(string,string,string,string)");
			temp.put("16896051", "checkTag(string,string)");
			temp.put("aa79bd66", "checkTagIds(string,string)");
			temp.put("44dfc0e1", "fixTagIds(string,string)");
			temp.put("8bb125e9", "getTagIds(string,string)");
			temp.put("76e74f92", "getTagMessage(string,string)");
			temp.put("8f32d59b", "isOwner()");
			temp.put("8da5cb5b", "owner()");
			temp.put("715018a6", "renounceOwnership()");
			temp.put("e748442e", "replaceTagIds(string,string,string[])");
			temp.put("8d540dfa", "setTagIds(string,string,string[])");
			temp.put("f2fde38b", "transferOwnership(address)");
			
			DTagFuncSigs = Collections.unmodifiableMap(temp);
		}
	
	
	// BYTECODE is the compiled bytecode used for deploying new contracts.
	public final static String BYTECODE = "0x60806040523480156200001157600080fd5b5060405162001c9738038062001c978339810160408190526200003491620000d0565b6000620000496001600160e01b03620000b916565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600180546001600160a01b0319166001600160a01b039290921691909117905562000125565b3390565b8051620000ca816200010b565b92915050565b600060208284031215620000e357600080fd5b6000620000f18484620000bd565b949350505050565b60006001600160a01b038216620000ca565b6200011681620000f9565b81146200012257600080fd5b50565b611b6280620001356000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80638d540dfa1161008c578063aa79bd6611610066578063aa79bd66146101d9578063e314604e146101fa578063e748442e1461020d578063f2fde38b14610220576100ea565b80638d540dfa146101a95780638da5cb5b146101bc5780638f32d59b146101d1576100ea565b806344dfc0e1116100c857806344dfc0e114610140578063715018a61461016057806376e74f92146101685780638bb125e914610189576100ea565b806316017c65146100ef578063168960511461010457806331c2c83f1461012d575b600080fd5b6101026100fd366004611403565b610233565b005b61011761011236600461130b565b610311565b6040516101249190611922565b60405180910390f35b61011761013b36600461148a565b61035a565b61015361014e36600461130b565b610418565b60405161012491906119d7565b6101026106e2565b61017b61017636600461130b565b610750565b6040516101249291906118f1565b61019c61019736600461130b565b6108d8565b60405161012491906118e0565b6101026101b7366004611372565b6109ee565b6101c4610b57565b60405161012491906118d2565b610117610b67565b6101ec6101e736600461130b565b610b8b565b6040516101249291906119e5565b610102610208366004611372565b610c5b565b61010261021b366004611372565b610ce5565b61010261022e36600461129b565b610e1a565b61023b610b67565b6102605760405162461bcd60e51b8152600401610257906119b7565b60405180910390fd5b60015460405163030e7d8d60e31b81526001600160a01b0390911690631873ec6890610290908490600401611930565b60206040518083038186803b1580156102a857600080fd5b505afa1580156102bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506102e091908101906112b9565b15156001146103015760405162461bcd60e51b8152600401610257906119a7565b61030c838383610e4a565b505050565b600060028360405161032391906118c6565b9081526040519081900360200181209061033e9084906118c6565b9081526040519081900360200190205460ff1690505b92915050565b6000610364610b67565b6103805760405162461bcd60e51b8152600401610257906119b7565b60015460405163631e0c6960e01b81526001600160a01b039091169063631e0c69906103b29086908690600401611941565b602060405180830381600087803b1580156103cc57600080fd5b505af11580156103e0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061040491908101906112b9565b50610410858585610e4a565b949350505050565b6000805b60028460405161042c91906118c6565b90815260200160405180910390206001018360405161044b91906118c6565b9081526040519081900360200190205481101561069a576001546040516001600160a01b0390911690631873ec68906002906104889088906118c6565b9081526020016040518091039020600101856040516104a791906118c6565b908152602001604051809103902083815481106104c057fe5b906000526020600020016040518263ffffffff1660e01b81526004016104e69190611966565b60206040518083038186803b1580156104fe57600080fd5b505afa158015610512573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061053691908101906112b9565b6106925760028460405161054a91906118c6565b90815260200160405180910390206001018360405161056991906118c6565b9081526020016040518091039020600160028660405161058991906118c6565b9081526020016040518091039020600101856040516105a891906118c6565b90815260200160405180910390208054905003815481106105c557fe5b906000526020600020016002856040516105df91906118c6565b9081526020016040518091039020600101846040516105fe91906118c6565b9081526020016040518091039020828154811061061757fe5b906000526020600020019080546001816001161561010002031660029004610640929190610f88565b5060028460405161065191906118c6565b90815260200160405180910390206001018360405161067091906118c6565b90815260405190819003602001902080549061069090600019830161100d565b505b60010161041c565b506002836040516106ab91906118c6565b9081526020016040518091039020600101826040516106ca91906118c6565b90815260405190819003602001902054905092915050565b6106ea610b67565b6107065760405162461bcd60e51b8152600401610257906119b7565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6060600060028460405161076491906118c6565b90815260200160405180910390206001018360405161078391906118c6565b9081526020016040518091039020805490509050806040519080825280602002602001820160405280156107cb57816020015b60608152602001906001900390816107b65790505b50915060005b818110156108d0576001546040516001600160a01b0390911690630cc4e8d8906002906107ff9089906118c6565b90815260200160405180910390206001018660405161081e91906118c6565b9081526020016040518091039020838154811061083757fe5b906000526020600020016040518263ffffffff1660e01b815260040161085d9190611966565b60006040518083038186803b15801561087557600080fd5b505afa158015610889573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526108b191908101906112d7565b8382815181106108bd57fe5b60209081029190910101526001016107d1565b509250929050565b60606002836040516108ea91906118c6565b90815260200160405180910390206001018260405161090991906118c6565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156109e25760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156109ce5780601f106109a3576101008083540402835291602001916109ce565b820191906000526020600020905b8154815290600101906020018083116109b157829003601f168201915b505050505081526020019060010190610937565b50505050905092915050565b6109f6610b67565b610a125760405162461bcd60e51b8152600401610257906119b7565b600154604051630b189f4960e11b81526001600160a01b03909116906316313e9290610a429084906004016118e0565b60206040518083038186803b158015610a5a57600080fd5b505afa158015610a6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610a9291908101906112b9565b1515600114610ab35760405162461bcd60e51b8152600401610257906119c7565b6001600284604051610ac591906118c6565b90815260405190819003602001812090610ae09085906118c6565b908152604051908190036020018120805492151560ff19909316929092179091558190600290610b119086906118c6565b908152602001604051809103902060010183604051610b3091906118c6565b90815260200160405180910390209080519060200190610b51929190611031565b50505050565b6000546001600160a01b03165b90565b600080546001600160a01b0316610b7c610f03565b6001600160a01b031614905090565b60015460405160009182916001600160a01b039091169063834daf0c90600290610bb69088906118c6565b908152602001604051809103902060010185604051610bd591906118c6565b9081526040519081900360200181206001600160e01b031960e084901b168252610c0191600401611911565b604080518083038186803b158015610c1857600080fd5b505afa158015610c2c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610c509190810190611546565b915091509250929050565b610c63610b67565b610c7f5760405162461bcd60e51b8152600401610257906119b7565b600283604051610c8f91906118c6565b90815260405190819003602001812090610caa9084906118c6565b9081526040519081900360200190205460ff1615610cda5760405162461bcd60e51b815260040161025790611987565b61030c8383836109ee565b610ced610b67565b610d095760405162461bcd60e51b8152600401610257906119b7565b600283604051610d1991906118c6565b90815260405190819003602001812090610d349084906118c6565b9081526040519081900360200190205460ff161515600114610d685760405162461bcd60e51b815260040161025790611997565b600154604051630b189f4960e11b81526001600160a01b03909116906316313e9290610d989084906004016118e0565b60206040518083038186803b158015610db057600080fd5b505afa158015610dc4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610de891908101906112b9565b1515600114610e095760405162461bcd60e51b8152600401610257906119c7565b80600284604051610b1191906118c6565b610e22610b67565b610e3e5760405162461bcd60e51b8152600401610257906119b7565b610e4781610f07565b50565b6001600284604051610e5c91906118c6565b90815260405190819003602001812090610e779085906118c6565b908152604051908190036020018120805492151560ff1990931692909217909155600290610ea69085906118c6565b908152602001604051809103902060010182604051610ec591906118c6565b9081526040516020918190038201902080546001810180835560009283529183902084519293610efc93919092019185019061108a565b5050505050565b3390565b6001600160a01b038116610f2d5760405162461bcd60e51b815260040161025790611977565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610fc15780548555610ffd565b82800160010185558215610ffd57600052602060002091601f016020900482015b82811115610ffd578254825591600101919060010190610fe2565b506110099291506110f8565b5090565b81548183558181111561030c5760008381526020902061030c918101908301611112565b82805482825590600052602060002090810192821561107e579160200282015b8281111561107e578251805161106e91849160209091019061108a565b5091602001919060010190611051565b50611009929150611112565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106110cb57805160ff1916838001178555610ffd565b82800160010185558215610ffd579182015b82811115610ffd5782518255916020019190600101906110dd565b610b6491905b8082111561100957600081556001016110fe565b610b6491905b8082111561100957600061112c8282611135565b50600101611118565b50805460018160011615610100020316600290046000825580601f1061115b5750610e47565b601f016020900490600052602060002090810190610e4791906110f8565b803561035481611af9565b600082601f83011261119557600080fd5b81356111a86111a382611a26565b611a00565b81815260209384019390925082018360005b838110156111e657813586016111d088826111fb565b84525060209283019291909101906001016111ba565b5050505092915050565b805161035481611b0d565b600082601f83011261120c57600080fd5b813561121a6111a382611a46565b9150808252602083016020830185838301111561123657600080fd5b611241838284611ab7565b50505092915050565b600082601f83011261125b57600080fd5b81516112696111a382611a46565b9150808252602083016020830185838301111561128557600080fd5b611241838284611ac3565b805161035481611b16565b6000602082840312156112ad57600080fd5b60006104108484611179565b6000602082840312156112cb57600080fd5b600061041084846111f0565b6000602082840312156112e957600080fd5b81516001600160401b038111156112ff57600080fd5b6104108482850161124a565b6000806040838503121561131e57600080fd5b82356001600160401b0381111561133457600080fd5b611340858286016111fb565b92505060208301356001600160401b0381111561135c57600080fd5b611368858286016111fb565b9150509250929050565b60008060006060848603121561138757600080fd5b83356001600160401b0381111561139d57600080fd5b6113a9868287016111fb565b93505060208401356001600160401b038111156113c557600080fd5b6113d1868287016111fb565b92505060408401356001600160401b038111156113ed57600080fd5b6113f986828701611184565b9150509250925092565b60008060006060848603121561141857600080fd5b83356001600160401b0381111561142e57600080fd5b61143a868287016111fb565b93505060208401356001600160401b0381111561145657600080fd5b611462868287016111fb565b92505060408401356001600160401b0381111561147e57600080fd5b6113f9868287016111fb565b600080600080608085870312156114a057600080fd5b84356001600160401b038111156114b657600080fd5b6114c2878288016111fb565b94505060208501356001600160401b038111156114de57600080fd5b6114ea878288016111fb565b93505060408501356001600160401b0381111561150657600080fd5b611512878288016111fb565b92505060608501356001600160401b0381111561152e57600080fd5b61153a878288016111fb565b91505092959194509250565b6000806040838503121561155957600080fd5b60006115658585611290565b9250506020611368858286016111f0565b6000611582838361167b565b9392505050565b600061158283836116e2565b61159e81611a9b565b82525050565b60006115af82611a7f565b6115b98185611a8d565b9350836020820285016115cb85611a6d565b8060005b8581101561160557848403895281516115e88582611576565b94506115f383611a6d565b60209a909a01999250506001016115cf565b5091979650505050505050565b600061161d82611a83565b6116278185611a8d565b93508360208202850161163985611a73565b8060005b85811015611605578484038952816116558582611589565b945061166083611a87565b60209a909a019992505060010161163d565b61159e81611aa6565b600061168682611a7f565b6116908185611a8d565b93506116a0818560208601611ac3565b6116a981611aef565b9093019392505050565b60006116be82611a7f565b6116c88185611a96565b93506116d8818560208601611ac3565b9290920192915050565b6000815460018116600081146116ff576001811461172557611764565b607f60028304166117108187611a8d565b60ff1984168152955050602085019250611764565b600282046117338187611a8d565b955061173e85611a73565b60005b8281101561175d57815488820152600190910190602001611741565b8701945050505b505092915050565b6000611779602683611a8d565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181526564647265737360d01b602082015260400192915050565b60006117c1601183611a8d565b70151859ce881d1859c81a5cc8195e1a5cdd607a1b815260200192915050565b60006117ee601583611a8d565b74151859ce881d1859c81a5cc81b9bdd08195e1a5cdd605a1b815260200192915050565b600061181f601983611a8d565b7f5461673a206d657373616765206973206e6f7420666f756e6400000000000000815260200192915050565b6000611858602083611a8d565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572815260200192915050565b6000611891601983611a8d565b7f5461673a206d657373616765206973206e6f7420657869737400000000000000815260200192915050565b61159e81610b64565b600061158282846116b3565b602081016103548284611595565b6020808252810161158281846115a4565b6040808252810161190281856115a4565b905061158260208301846118bd565b602080825281016115828184611612565b602081016103548284611672565b60208082528101611582818461167b565b60408082528101611952818561167b565b90508181036020830152610410818461167b565b6020808252810161158281846116e2565b602080825281016103548161176c565b60208082528101610354816117b4565b60208082528101610354816117e1565b6020808252810161035481611812565b602080825281016103548161184b565b6020808252810161035481611884565b6020810161035482846118bd565b604081016119f382856118bd565b6115826020830184611672565b6040518181016001600160401b0381118282101715611a1e57600080fd5b604052919050565b60006001600160401b03821115611a3c57600080fd5b5060209081020190565b60006001600160401b03821115611a5c57600080fd5b506020601f91909101601f19160190565b60200190565b60009081526020902090565b5190565b5490565b60010190565b90815260200190565b919050565b600061035482611aab565b151590565b6001600160a01b031690565b82818337506000910152565b60005b83811015611ade578181015183820152602001611ac6565b83811115610b515750506000910152565b601f01601f191690565b611b0281611a9b565b8114610e4757600080fd5b611b0281611aa6565b611b0281610b6456fea365627a7a723158204d6c105a79521e4c3f72bdac6813431cbea1f03c5cbb4ee311ca3987b03945426c6578706572696d656e74616cf564736f6c634300050c0040";

	// deploy deploys a new Ethereum contract, binding an instance of DTag to it.
	public static DTag deploy(TransactOpts auth, EthereumClient client, Address msgAddr) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		String bytecode = BYTECODE;
		
		Interface arg0 = Geth.newInterface();arg0.setAddress(msgAddr);args.set(0,arg0);
		
		return new DTag(Geth.deployContract(auth, ABI, Geth.decodeFromHex(bytecode), client, args));
	}

	// Internal constructor used by contract deployment.
	private DTag(BoundContract deployment) {
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

	// Creates a new instance of DTag, bound to a specific deployed contract.
	public DTag(Address address, EthereumClient client) throws Exception {
		this(Geth.bindContract(address, ABI, client));
	}

	
	

	// checkTag is a free data retrieval call binding the contract method 0x16896051.
	//
	// Solidity: function checkTag(string _tag, string _sub) constant returns(bool)
	public boolean checkTag(CallOpts opts, String _tag, String _sub) throws Exception {
		Interfaces args = Geth.newInterfaces(2);
		Interface arg0 = Geth.newInterface();arg0.setString(_tag);args.set(0,arg0);
		Interface arg1 = Geth.newInterface();arg1.setString(_sub);args.set(1,arg1);
		

		Interfaces results = Geth.newInterfaces(1);
		Interface result0 = Geth.newInterface(); result0.setDefaultBool(); results.set(0, result0);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "checkTag", args);
		return results.get(0).getBool();
		
	}
	
	
	// CheckTagIdsResults is the output of a call to checkTagIds.
	public class CheckTagIdsResults {
		public BigInt Return0;
		public boolean Return1;
		
	}
	

	// checkTagIds is a free data retrieval call binding the contract method 0xaa79bd66.
	//
	// Solidity: function checkTagIds(string _tag, string _sub) constant returns(uint256, bool)
	public CheckTagIdsResults checkTagIds(CallOpts opts, String _tag, String _sub) throws Exception {
		Interfaces args = Geth.newInterfaces(2);
		Interface arg0 = Geth.newInterface();arg0.setString(_tag);args.set(0,arg0);
		Interface arg1 = Geth.newInterface();arg1.setString(_sub);args.set(1,arg1);
		

		Interfaces results = Geth.newInterfaces(2);
		Interface result0 = Geth.newInterface(); result0.setDefaultBigInt(); results.set(0, result0);
		Interface result1 = Geth.newInterface(); result1.setDefaultBool(); results.set(1, result1);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "checkTagIds", args);
		
			CheckTagIdsResults result = new CheckTagIdsResults();
			result.Return0 = results.get(0).getBigInt();
			result.Return1 = results.get(1).getBool();
			
			return result;
		
	}
	
	

	// getTagIds is a free data retrieval call binding the contract method 0x8bb125e9.
	//
	// Solidity: function getTagIds(string _tag, string _sub) constant returns(string[])
	public Strings getTagIds(CallOpts opts, String _tag, String _sub) throws Exception {
		Interfaces args = Geth.newInterfaces(2);
		Interface arg0 = Geth.newInterface();arg0.setString(_tag);args.set(0,arg0);
		Interface arg1 = Geth.newInterface();arg1.setString(_sub);args.set(1,arg1);
		

		Interfaces results = Geth.newInterfaces(1);
		Interface result0 = Geth.newInterface(); result0.setDefaultStrings(); results.set(0, result0);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "getTagIds", args);
		return results.get(0).getStrings();
		
	}
	
	
	// GetTagMessageResults is the output of a call to getTagMessage.
	public class GetTagMessageResults {
		public Strings Value;
		public BigInt Size;
		
	}
	

	// getTagMessage is a free data retrieval call binding the contract method 0x76e74f92.
	//
	// Solidity: function getTagMessage(string _tag, string _sub) constant returns(string[] _value, uint256 _size)
	public GetTagMessageResults getTagMessage(CallOpts opts, String _tag, String _sub) throws Exception {
		Interfaces args = Geth.newInterfaces(2);
		Interface arg0 = Geth.newInterface();arg0.setString(_tag);args.set(0,arg0);
		Interface arg1 = Geth.newInterface();arg1.setString(_sub);args.set(1,arg1);
		

		Interfaces results = Geth.newInterfaces(2);
		Interface result0 = Geth.newInterface(); result0.setDefaultStrings(); results.set(0, result0);
		Interface result1 = Geth.newInterface(); result1.setDefaultBigInt(); results.set(1, result1);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "getTagMessage", args);
		
			GetTagMessageResults result = new GetTagMessageResults();
			result.Value = results.get(0).getStrings();
			result.Size = results.get(1).getBigInt();
			
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
	

	
	// addTagId is a paid mutator transaction binding the contract method 0x16017c65.
	//
	// Solidity: function addTagId(string _tag, string _sub, string _id) returns()
	public Transaction addTagId(TransactOpts opts, String _tag, String _sub, String _id) throws Exception {
		Interfaces args = Geth.newInterfaces(3);
		Interface arg0 = Geth.newInterface();arg0.setString(_tag);args.set(0,arg0);
		Interface arg1 = Geth.newInterface();arg1.setString(_sub);args.set(1,arg1);
		Interface arg2 = Geth.newInterface();arg2.setString(_id);args.set(2,arg2);
		
		return this.Contract.transact(opts, "addTagId"	, args);
	}
	
	// addTagIds is a paid mutator transaction binding the contract method 0xe314604e.
	//
	// Solidity: function addTagIds(string _tag, string _sub, string[] _ids) returns()
	public Transaction addTagIds(TransactOpts opts, String _tag, String _sub, Strings _ids) throws Exception {
		Interfaces args = Geth.newInterfaces(3);
		Interface arg0 = Geth.newInterface();arg0.setString(_tag);args.set(0,arg0);
		Interface arg1 = Geth.newInterface();arg1.setString(_sub);args.set(1,arg1);
		Interface arg2 = Geth.newInterface();arg2.setStrings(_ids);args.set(2,arg2);
		
		return this.Contract.transact(opts, "addTagIds"	, args);
	}
	
	// addTagMessage is a paid mutator transaction binding the contract method 0x31c2c83f.
	//
	// Solidity: function addTagMessage(string _tag, string _sub, string _id, string _message) returns(bool)
	public Transaction addTagMessage(TransactOpts opts, String _tag, String _sub, String _id, String _message) throws Exception {
		Interfaces args = Geth.newInterfaces(4);
		Interface arg0 = Geth.newInterface();arg0.setString(_tag);args.set(0,arg0);
		Interface arg1 = Geth.newInterface();arg1.setString(_sub);args.set(1,arg1);
		Interface arg2 = Geth.newInterface();arg2.setString(_id);args.set(2,arg2);
		Interface arg3 = Geth.newInterface();arg3.setString(_message);args.set(3,arg3);
		
		return this.Contract.transact(opts, "addTagMessage"	, args);
	}
	
	// fixTagIds is a paid mutator transaction binding the contract method 0x44dfc0e1.
	//
	// Solidity: function fixTagIds(string _tag, string _sub) returns(uint256 _size)
	public Transaction fixTagIds(TransactOpts opts, String _tag, String _sub) throws Exception {
		Interfaces args = Geth.newInterfaces(2);
		Interface arg0 = Geth.newInterface();arg0.setString(_tag);args.set(0,arg0);
		Interface arg1 = Geth.newInterface();arg1.setString(_sub);args.set(1,arg1);
		
		return this.Contract.transact(opts, "fixTagIds"	, args);
	}
	
	// renounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
	//
	// Solidity: function renounceOwnership() returns()
	public Transaction renounceOwnership(TransactOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		
		return this.Contract.transact(opts, "renounceOwnership"	, args);
	}
	
	// replaceTagIds is a paid mutator transaction binding the contract method 0xe748442e.
	//
	// Solidity: function replaceTagIds(string _tag, string _sub, string[] _ids) returns()
	public Transaction replaceTagIds(TransactOpts opts, String _tag, String _sub, Strings _ids) throws Exception {
		Interfaces args = Geth.newInterfaces(3);
		Interface arg0 = Geth.newInterface();arg0.setString(_tag);args.set(0,arg0);
		Interface arg1 = Geth.newInterface();arg1.setString(_sub);args.set(1,arg1);
		Interface arg2 = Geth.newInterface();arg2.setStrings(_ids);args.set(2,arg2);
		
		return this.Contract.transact(opts, "replaceTagIds"	, args);
	}
	
	// setTagIds is a paid mutator transaction binding the contract method 0x8d540dfa.
	//
	// Solidity: function setTagIds(string _tag, string _sub, string[] _ids) returns()
	public Transaction setTagIds(TransactOpts opts, String _tag, String _sub, Strings _ids) throws Exception {
		Interfaces args = Geth.newInterfaces(3);
		Interface arg0 = Geth.newInterface();arg0.setString(_tag);args.set(0,arg0);
		Interface arg1 = Geth.newInterface();arg1.setString(_sub);args.set(1,arg1);
		Interface arg2 = Geth.newInterface();arg2.setStrings(_ids);args.set(2,arg2);
		
		return this.Contract.transact(opts, "setTagIds"	, args);
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


public class Writeable {
	// ABI is the input ABI used to generate the binding from.
	public final static String ABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldWriter\",\"type\":\"address\"}],\"name\":\"WritershipDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newWriter\",\"type\":\"address\"}],\"name\":\"WritershipIncreased\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldWriter\",\"type\":\"address\"}],\"name\":\"decreaseWritership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWriter\",\"type\":\"address\"}],\"name\":\"increasedWritership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isWriter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceWritership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"writer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"writers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_value\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]";
	
		// WriteableFuncSigs maps the 4-byte function signature to its string representation.
		public final static Map<String, String> WriteableFuncSigs;
		static {
			Hashtable<String, String> temp = new Hashtable<String, String>();
			temp.put("9f217bdf", "decreaseWritership(address)");
			temp.put("14b4d411", "increasedWritership(address)");
			temp.put("8f32d59b", "isOwner()");
			temp.put("4d6ee9fd", "isWriter()");
			temp.put("8da5cb5b", "owner()");
			temp.put("715018a6", "renounceOwnership()");
			temp.put("4075e83e", "renounceWritership()");
			temp.put("f2fde38b", "transferOwnership(address)");
			temp.put("453a2abc", "writer()");
			temp.put("8b2f5369", "writers()");
			
			WriteableFuncSigs = Collections.unmodifiableMap(temp);
		}
	
	

	// Ethereum address where this contract is located at.
	public final Address Address;

	// Ethereum transaction in which this contract was deployed (if known!).
	public final Transaction Deployer;

	// Contract instance bound to a blockchain address.
	private final BoundContract Contract;

	// Creates a new instance of Writeable, bound to a specific deployed contract.
	public Writeable(Address address, EthereumClient client) throws Exception {
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
	
	

	// isWriter is a free data retrieval call binding the contract method 0x4d6ee9fd.
	//
	// Solidity: function isWriter() constant returns(bool)
	public boolean isWriter(CallOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		

		Interfaces results = Geth.newInterfaces(1);
		Interface result0 = Geth.newInterface(); result0.setDefaultBool(); results.set(0, result0);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "isWriter", args);
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
	
	

	// writer is a free data retrieval call binding the contract method 0x453a2abc.
	//
	// Solidity: function writer() constant returns(address)
	public Address writer(CallOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		

		Interfaces results = Geth.newInterfaces(1);
		Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "writer", args);
		return results.get(0).getAddress();
		
	}
	
	

	// writers is a free data retrieval call binding the contract method 0x8b2f5369.
	//
	// Solidity: function writers() constant returns(address[] _value)
	public Addresses writers(CallOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		

		Interfaces results = Geth.newInterfaces(1);
		Interface result0 = Geth.newInterface(); result0.setDefaultAddresses(); results.set(0, result0);
		

		if (opts == null) {
			opts = Geth.newCallOpts();
		}
		this.Contract.call(opts, results, "writers", args);
		return results.get(0).getAddresses();
		
	}
	

	
	// decreaseWritership is a paid mutator transaction binding the contract method 0x9f217bdf.
	//
	// Solidity: function decreaseWritership(address oldWriter) returns()
	public Transaction decreaseWritership(TransactOpts opts, Address oldWriter) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setAddress(oldWriter);args.set(0,arg0);
		
		return this.Contract.transact(opts, "decreaseWritership"	, args);
	}
	
	// increasedWritership is a paid mutator transaction binding the contract method 0x14b4d411.
	//
	// Solidity: function increasedWritership(address newWriter) returns()
	public Transaction increasedWritership(TransactOpts opts, Address newWriter) throws Exception {
		Interfaces args = Geth.newInterfaces(1);
		Interface arg0 = Geth.newInterface();arg0.setAddress(newWriter);args.set(0,arg0);
		
		return this.Contract.transact(opts, "increasedWritership"	, args);
	}
	
	// renounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
	//
	// Solidity: function renounceOwnership() returns()
	public Transaction renounceOwnership(TransactOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		
		return this.Contract.transact(opts, "renounceOwnership"	, args);
	}
	
	// renounceWritership is a paid mutator transaction binding the contract method 0x4075e83e.
	//
	// Solidity: function renounceWritership() returns()
	public Transaction renounceWritership(TransactOpts opts) throws Exception {
		Interfaces args = Geth.newInterfaces(0);
		
		return this.Contract.transact(opts, "renounceWritership"	, args);
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

