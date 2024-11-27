package oasys

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

var deployments13 = []*deployment{
	{
		// commit: https://github.com/oasysgames/oasys-governance-contract/blob/v1.0.0/contracts/EVMAccessControl.sol
		contract: transactionFilter,
		code:     mustDecodeCode(`0x608060405234801561001057600080fd5b50600436106101365760003560e01c806390d3ed88116100b2578063d547741f11610081578063ec87621c11610066578063ec87621c146102df578063f00cab4314610306578063ff616c8e1461032657600080fd5b8063d547741f146102b9578063e969f4e1146102cc57600080fd5b806390d3ed881461025457806391d1485414610267578063a217fddf1461029e578063bd7f53d6146102a657600080fd5b806336568abe1161010957806354fd4d50116100ee57806354fd4d50146101ef57806373f25c381461022e5780637723a3eb1461024157600080fd5b806336568abe146101bc578063506921d0146101cf57600080fd5b806301ffc9a71461013b578063248a9ca31461016357806329853b86146101945780632f2ff15d146101a9575b600080fd5b61014e610149366004610fab565b610339565b60405190151581526020015b60405180910390f35b610186610171366004610fed565b60009081526020819052604090206001015490565b60405190815260200161015a565b6101a76101a2366004611022565b6103d2565b005b6101a76101b7366004611055565b610441565b6101a76101ca366004611055565b61046b565b6101e26101dd366004611078565b6104fc565b60405161015a91906110a2565b604080518082018252600581527f312e302e300000000000000000000000000000000000000000000000000000006020820152905161015a9190611113565b61014e61023c366004611146565b610511565b61014e61024f366004611146565b610535565b6101e2610262366004611078565b610559565b61014e610275366004611055565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b610186600081565b6101a76102b4366004611146565b610567565b6101a76102c7366004611055565b6105d4565b6101a76102da366004611022565b6105f9565b6101867f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0881565b61030e600181565b6040516001600160a01b03909116815260200161015a565b6101a7610334366004611146565b610668565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806103cc57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b086103fc816106d5565b610408600184846106e2565b6040516001600160a01b038416907f4a475457a2e45f0e02800bf274aede959c4e648bd208c5af470590d2e93d073290600090a2505050565b60008281526020819052604090206001015461045c816106d5565b61046683836108d4565b505050565b6001600160a01b03811633146104ee5760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201527f20726f6c657320666f722073656c66000000000000000000000000000000000060648201526084015b60405180910390fd5b6104f88282610972565b5050565b606061050a600284846109f1565b9392505050565b6001600160a01b0380821660009081526001602052604081205490911615156103cc565b6001600160a01b0380821660009081526002602052604081205490911615156103cc565b606061050a600184846109f1565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08610591816106d5565b61059c600183610ac4565b6040516001600160a01b038316907fca573f65e3d72a9f457e37bb6357cbb8f4bb06010cf200e958854309620342e990600090a25050565b6000828152602081905260409020600101546105ef816106d5565b6104668383610972565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08610623816106d5565b61062f600284846106e2565b6040516001600160a01b038416907f054279c30632ea9b7d314b136883c9448638eb556adc3cffb4dfac65024f416b90600090a2505050565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08610692816106d5565b61069d600283610ac4565b6040516001600160a01b038316907f2757c5d518b1e4ab9d6e4c8276711fd6bf73fbf7005ba27f059e1b5b6d56623190600090a25050565b6106df8133610c8b565b50565b6001600160a01b0382166107385760405162461bcd60e51b815260206004820152601160248201527f4541433a2061646472206973207a65726f00000000000000000000000000000060448201526064016104e5565b6000196001600160a01b038316016107925760405162461bcd60e51b815260206004820152601560248201527f4541433a20616464722069732073656e74696e656c000000000000000000000060448201526064016104e5565b6001600160a01b03828116600090815260208590526040902054166107f95760405162461bcd60e51b815260206004820152600e60248201527f4541433a206e6f7420666f756e6400000000000000000000000000000000000060448201526064016104e5565b6001600160a01b038116610814576108118383610cfe565b90505b6001600160a01b038181166000908152602085905260409020548116908316146108805760405162461bcd60e51b815260206004820181905260248201527f4541433a2070726576206164647265737320646f6573206e6f74206d6174636860448201526064016104e5565b6001600160a01b03918216600081815260209490945260408085208054938516865290852080549390941673ffffffffffffffffffffffffffffffffffffffff199384161790935590925280549091169055565b6000828152602081815260408083206001600160a01b038516845290915290205460ff166104f8576000828152602081815260408083206001600160a01b03851684529091529020805460ff1916600117905561092e3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6000828152602081815260408083206001600160a01b038516845290915290205460ff16156104f8576000828152602081815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b60606001600160a01b038316610a0657600192505b60008267ffffffffffffffff811115610a2157610a21611161565b604051908082528060200260200182016040528015610a4a578160200160208202803683370190505b50905060005b83811015610abb576001600160a01b0394851660009081526020879052604090205490941693600019850115610abb5784828281518110610a9357610a93611177565b6001600160a01b0390921660209283029190910190910152610ab4816111a3565b9050610a50565b50949350505050565b6001600160a01b038116610b1a5760405162461bcd60e51b815260206004820152601160248201527f4541433a2061646472206973207a65726f00000000000000000000000000000060448201526064016104e5565b306001600160a01b03821603610b725760405162461bcd60e51b815260206004820152601160248201527f4541433a20616464722069732073656c6600000000000000000000000000000060448201526064016104e5565b6000196001600160a01b03821601610bcc5760405162461bcd60e51b815260206004820152601560248201527f4541433a20616464722069732073656e74696e656c000000000000000000000060448201526064016104e5565b6001600160a01b038181166000908152602084905260409020541615610c345760405162461bcd60e51b815260206004820152601360248201527f4541433a20616c7265616479206578697374730000000000000000000000000060448201526064016104e5565b60016000818152602093909352604080842080546001600160a01b0394851680875292862080549590911673ffffffffffffffffffffffffffffffffffffffff199586161790559190935280549091169091179055565b6000828152602081815260408083206001600160a01b038516845290915290205460ff166104f857610cbc81610db8565b610cc7836020610dca565b604051602001610cd89291906111bc565b60408051601f198184030181529082905262461bcd60e51b82526104e591600401611113565b600060015b6001600160a01b0381811660009081526020869052604090205416600114610d70576001600160a01b03818116600090815260208690526040902054818516911603610d505790506103cc565b6001600160a01b0390811660009081526020859052604090205416610d03565b60405162461bcd60e51b815260206004820152601b60248201527f4541433a20707265762061646472657373206e6f7420666f756e64000000000060448201526064016104e5565b60606103cc6001600160a01b03831660145b60606000610dd983600261123d565b610de4906002611254565b67ffffffffffffffff811115610dfc57610dfc611161565b6040519080825280601f01601f191660200182016040528015610e26576020820181803683370190505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110610e5d57610e5d611177565b60200101906001600160f81b031916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610ea857610ea8611177565b60200101906001600160f81b031916908160001a9053506000610ecc84600261123d565b610ed7906001611254565b90505b6001811115610f5c577f303132333435363738396162636465660000000000000000000000000000000085600f1660108110610f1857610f18611177565b1a60f81b828281518110610f2e57610f2e611177565b60200101906001600160f81b031916908160001a90535060049490941c93610f5581611267565b9050610eda565b50831561050a5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e7460448201526064016104e5565b600060208284031215610fbd57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461050a57600080fd5b600060208284031215610fff57600080fd5b5035919050565b80356001600160a01b038116811461101d57600080fd5b919050565b6000806040838503121561103557600080fd5b61103e83611006565b915061104c60208401611006565b90509250929050565b6000806040838503121561106857600080fd5b8235915061104c60208401611006565b6000806040838503121561108b57600080fd5b61109483611006565b946020939093013593505050565b6020808252825182820181905260009190848201906040850190845b818110156110e35783516001600160a01b0316835292840192918401916001016110be565b50909695505050505050565b60005b8381101561110a5781810151838201526020016110f2565b50506000910152565b60208152600082518060208401526111328160408501602087016110ef565b601f01601f19169190910160400192915050565b60006020828403121561115857600080fd5b61050a82611006565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600182016111b5576111b561118d565b5060010190565b7f416363657373436f6e74726f6c3a206163636f756e74200000000000000000008152600083516111f48160178501602088016110ef565b7f206973206d697373696e6720726f6c652000000000000000000000000000000060179184019182015283516112318160288401602088016110ef565b01602801949350505050565b80820281158282048414176103cc576103cc61118d565b808201808211156103cc576103cc61118d565b6000816112765761127661118d565b50600019019056fea2646970667358221220a275e1c70f3d1b906ee360119f37a593a36f60d3d9c574672facac8ec4f93c5064736f6c63430008130033`),
		storage: storage{
			// mapping(bytes32 => RoleData) private _roles
			"0x00": &mapping{
				keyFn: func(key string) common.Hash {
					if key == "DEFAULT_ADMIN_ROLE" {
						return common.Hash{}
					}
					return common.BytesToHash(crypto.Keccak256([]byte(key)))
				},
				values: map[string]interface{}{
					"DEFAULT_ADMIN_ROLE": structvalue{
						// mapping(address => bool) members
						genesismap{
							params.OasysMainnetGenesisHash: &mapping{
								keyFn: addressKeyFn,
								values: map[string]interface{}{
									"0xe04EEaCb1f181cD23501f3De39014F4cbb7C1Cd8": "0x1",
								},
							},
							params.OasysTestnetGenesisHash: &mapping{
								keyFn: addressKeyFn,
								values: map[string]interface{}{
									"0xbe32b47A35C31d294B3c58d92ca83876DdC38776": "0x1",
								},
							},
							// For local development
							// GenesisHash: &mapping{
							// 	keyFn: addressKeyFn,
							// 	values: map[string]interface{}{
							// 		"0x75fBB5Bd6FDf076Dcaf55243e9E3f3c76F8b5640": "0x1",
							// 	},
							// },
						},
					},
					"MANAGER_ROLE": structvalue{
						// mapping(address => bool) members
						genesismap{
							params.OasysMainnetGenesisHash: &mapping{
								keyFn: addressKeyFn,
								values: map[string]interface{}{
									"0xc0bACBDA16Bb494d8C5be6DE84540465Fd83333E": "0x1",
								},
							},
							params.OasysTestnetGenesisHash: &mapping{
								keyFn: addressKeyFn,
								values: map[string]interface{}{
									"0xBb5a4FF43683a1281800A6Bc8a94365f055F444F": "0x1",
								},
							},
							// For local development
							// GenesisHash: &mapping{
							// 	keyFn: addressKeyFn,
							// 	values: map[string]interface{}{
							// 		"0x75fBB5Bd6FDf076Dcaf55243e9E3f3c76F8b5640": "0x1",
							// 	},
							// },
						},
					},
				},
			},
			// mapping(address => address) private _createAllowedList;
			"0x01": &mapping{
				keyFn: addressKeyFn,
				values: map[string]interface{}{
					"0x0000000000000000000000000000000000000001": "0x0000000000000000000000000000000000000001",
				},
			},
			// mapping(address => address) private _callDeniedList;
			"0x02": &mapping{
				keyFn: addressKeyFn,
				values: map[string]interface{}{
					"0x0000000000000000000000000000000000000001": "0x0000000000000000000000000000000000000001",
				},
			},
		},
	},
}
