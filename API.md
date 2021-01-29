# `rosetta-icon` Restful API

Based on rosetta specification `v1.4.9`, rosetta-icon` provides following Restful APIs:

## Network

**/network/list**

*Get List of Available Networks*

Request:

```json
{
	"metadata": {}
}
```

Response:

Sample

```json
{
	"network_identifiers": [
		{
			"blockchain": "ICON",
			"network": "Testnet"
		}
	]
}
```

**/network/options**

*Get List of Available Networks*

Request:

```json
{
	"network_identifier": {
		"blockchain": "ICON",
		"network": "Testnet"
	},
	"metadata": {}
}
```

Response:

Sample

```json
{
	"version": {
		"rosetta_version": "1.4.0",
		"node_version": "1.8.0",
		"middleware_version": "0.0.1"
	},
	"allow": {
		"operation_statuses": [
			{
				"status": "SUCCESS",
				"successful": true
			},
			{
				"status": "FAIL",
				"successful": false
			}
		],
		"operation_types": [
			"GENESIS",
			"TRANSFER",
			"FEE",
			"BASE",
			"ISSUE",
			"BURN",
			"ICXTRANSFER",
			"CLAIM",
			"MESSAGE"
		],
		"errors": [
			{
				"code": 0,
				"message": "Endpoint not implemented",
				"retriable": false
			},
			{
				"code": 1,
				"message": "Endpoint unavailable offline",
				"retriable": false
			},
			{
				"code": 2,
				"message": "ICON Node is not ready",
				"retriable": true
			},
			{
				"code": 13,
				"message": "Wrong Block Hash",
				"retriable": true
			}
		],
		"historical_balance_lookup": false,
		"call_methods": null,
		"balance_exemptions": null
	}
}
```

**/network/status**

*Get Network Status*

Request:

```json
{
	"network_identifier": {
		"blockchain": "ICON",
		"network": "Testnet"
	},
	"metadata": {}
}
```

Response:

Sample

```json
{
	"current_block_identifier": {
		"index": 16516303,
		"hash": "0x180c1d0bebf0aefe3cd3e2a087a0fdd0baf386a41580a4daf29584356d21948a"
	},
	"current_block_timestamp": 1611839529778,
	"genesis_block_identifier": {
		"index": 0,
		"hash": "0xec470d5b046106b86a87eff612af8b1b0279312a4f14a70323828faad47e4364"
	},
	"peers": [
		{
			"peer_id": "hx8d94f443d4a0002ebd4b86634b4a8502a7d180d4",
			"metadata": {
				"address": "hx8d94f443d4a0002ebd4b86634b4a8502a7d180d4",
				"blockHeight": "0x117",
				"city": "Seoul",
				"country": "KOR",
				"delegated": "0x831231d8e247897b2c000",
				"details": "https://icon.jinwoo.com/json",
				"email": "banana@jinwoo.com",
				"grade": "0x0",
				"irep": "0xa968163f0a57b400000",
				"irepUpdateBlockHeight": "0x117",
				"lastGenerateBlockHeight": "0xfc04a6",
				"name": "akrt-peer01",
				"nodeAddress": "hx8d94f443d4a0002ebd4b86634b4a8502a7d180d4",
				"p2pEndpoint": "13.124.238.145:7100",
				"penalty": "0x0",
				"stake": "0x0",
				"status": "0x0",
				"totalBlocks": "0xf94c20",
				"txIndex": "0x0",
				"unvalidatedSequenceBlocks": "0x0",
				"validatedBlocks": "0xf9406a",
				"website": "https://icon.jinwoo.com"
			}
		},
		{
			"peer_id": "hx3814a13b7dad940b343e4ae4219cba1437552feb",
			"metadata": {
				"address": "hx3814a13b7dad940b343e4ae4219cba1437552feb",
				"blockHeight": "0x11f",
				"city": "Seoul",
				"country": "KOR",
				"delegated": "0x749dbac1046256b314000",
				"details": "https://icon.jinwoo.com/json",
				"email": "banana@jinwoo.com",
				"grade": "0x0",
				"irep": "0xa968163f0a57b400000",
				"irepUpdateBlockHeight": "0x11f",
				"lastGenerateBlockHeight": "0xfc04b0",
				"name": "aort-peer01",
				"nodeAddress": "hx3814a13b7dad940b343e4ae4219cba1437552feb",
				"p2pEndpoint": "34.208.246.112:7100",
				"penalty": "0x0",
				"stake": "0x0",
				"status": "0x0",
				"totalBlocks": "0xfb5b39",
				"txIndex": "0x0",
				"unvalidatedSequenceBlocks": "0x0",
				"validatedBlocks": "0xfb58d3",
				"website": "https://icon.jinwoo.com"
			}
		},
		{
			"peer_id": "hxf2152fad2960b17bedf60d4bc249b42d7f2a0ebc",
			"metadata": {
				"address": "hxf2152fad2960b17bedf60d4bc249b42d7f2a0ebc",
				"blockHeight": "0x113",
				"city": "Seoul",
				"country": "KOR",
				"delegated": "0x66172ff51b0118d214000",
				"details": "https://icon.foundation/prep/details.json",
				"email": "banana@jinwoo.com",
				"grade": "0x0",
				"irep": "0xa968163f0a57b400000",
				"irepUpdateBlockHeight": "0x113",
				"lastGenerateBlockHeight": "0xfc04ba",
				"name": "ajpt-peer01",
				"nodeAddress": "hxf2152fad2960b17bedf60d4bc249b42d7f2a0ebc",
				"p2pEndpoint": "13.115.169.117:7100",
				"penalty": "0x0",
				"stake": "0x100bd5616bee7d8000",
				"status": "0x0",
				"totalBlocks": "0xf8ca1f",
				"txIndex": "0x0",
				"unvalidatedSequenceBlocks": "0x0",
				"validatedBlocks": "0xf8be73",
				"website": "https://icon.foundation"
			}
		},
		{
			"peer_id": "hx459af36f56b9140407df21a856fcfd4ff44db648",
			"metadata": {
				"address": "hx459af36f56b9140407df21a856fcfd4ff44db648",
				"blockHeight": "0x121",
				"city": "Seoul",
				"country": "KOR",
				"delegated": "0x5af4d88304f7dfd3f4000",
				"details": "https://icon.jinwoo.com/json",
				"email": "banana@jinwoo.com",
				"grade": "0x0",
				"irep": "0xa968163f0a57b400000",
				"irepUpdateBlockHeight": "0x121",
				"lastGenerateBlockHeight": "0xfc04c4",
				"name": "aort-peer02",
				"nodeAddress": "hx459af36f56b9140407df21a856fcfd4ff44db648",
				"p2pEndpoint": "34.215.104.203:7100",
				"penalty": "0x0",
				"stake": "0x0",
				"status": "0x0",
				"totalBlocks": "0xfb5b39",
				"txIndex": "0x0",
				"unvalidatedSequenceBlocks": "0x0",
				"validatedBlocks": "0xfb58cf",
				"website": "https://icon.jinwoo.com"
			}
		},
		{
			"peer_id": "hxf4347d08be308bfc9f8428e1e536536983695b33",
			"metadata": {
				"address": "hxf4347d08be308bfc9f8428e1e536536983695b33",
				"blockHeight": "0x125",
				"city": "Seoul",
				"country": "KOR",
				"delegated": "0x1d43b90b9bb99ae748000",
				"details": "https://icon.jinwoo.com/json",
				"email": "banana@jinwoo.com",
				"grade": "0x0",
				"irep": "0xa968163f0a57b400000",
				"irepUpdateBlockHeight": "0x125",
				"lastGenerateBlockHeight": "0xfc04ce",
				"name": "avat-peer02",
				"nodeAddress": "hxf4347d08be308bfc9f8428e1e536536983695b33",
				"p2pEndpoint": "52.20.138.241:7100",
				"penalty": "0x0",
				"stake": "0x0",
				"status": "0x0",
				"totalBlocks": "0x41a9cd",
				"txIndex": "0x0",
				"unvalidatedSequenceBlocks": "0x0",
				"validatedBlocks": "0x41a7ef",
				"website": "https://icon.jinwoo.com"
			}
		},
		{
			"peer_id": "hx7f8783d8ca3d4923e0ec1dbe21ba5b6edf301b17",
			"metadata": {
				"address": "hx7f8783d8ca3d4923e0ec1dbe21ba5b6edf301b17",
				"blockHeight": "0x119",
				"city": "Seoul",
				"country": "KOR",
				"delegated": "0x48da18833225534a6c000",
				"details": "https://icon.jinwoo.com/json",
				"email": "banana@jinwoo.com",
				"grade": "0x0",
				"irep": "0xa968163f0a57b400000",
				"irepUpdateBlockHeight": "0x119",
				"lastGenerateBlockHeight": "0xfc0492",
				"name": "akrt-peer02",
				"nodeAddress": "hx7f8783d8ca3d4923e0ec1dbe21ba5b6edf301b17",
				"p2pEndpoint": "52.78.98.154:7100",
				"penalty": "0x0",
				"stake": "0x0",
				"status": "0x0",
				"totalBlocks": "0xfb5b39",
				"txIndex": "0x0",
				"unvalidatedSequenceBlocks": "0x0",
				"validatedBlocks": "0xfb5a09",
				"website": "https://icon.jinwoo.com"
			}
		},
		{
			"peer_id": "hx36ceb073381b7b5ce383ccefc361435e3ec54dc4",
			"metadata": {
				"address": "hx36ceb073381b7b5ce383ccefc361435e3ec54dc4",
				"blockHeight": "0x115",
				"city": "Seoul",
				"country": "KOR",
				"delegated": "0x2baf0c657c5c081c60000",
				"details": "https://icon.jinwoo.com/json",
				"email": "banana@jinwoo.com",
				"grade": "0x0",
				"irep": "0xa968163f0a57b400000",
				"irepUpdateBlockHeight": "0x115",
				"lastGenerateBlockHeight": "0xfc049c",
				"name": "ajpt-peer02",
				"nodeAddress": "hx36ceb073381b7b5ce383ccefc361435e3ec54dc4",
				"p2pEndpoint": "13.230.61.61:7100",
				"penalty": "0x0",
				"stake": "0x0",
				"status": "0x0",
				"totalBlocks": "0xfb5b39",
				"txIndex": "0x0",
				"unvalidatedSequenceBlocks": "0x0",
				"validatedBlocks": "0xfb599c",
				"website": "https://icon.jinwoo.com"
			}
		}
	]
}
```

### Account

**/account/balance**

*Get an Account Balance*

Request:

```json
{
	"network_identifier": {
		"blockchain": "ICON",
		"network": "Testnet"
	},
	"account_identifier": {
		"address": "hx2141bf8b6d2213d4d7204e2ddab92653dc245c5f",
		"sub_account": {
			"address": "empty"
		},
		"metadata": {}
	}
}
```

Response:

Sample

```json
{
	"block_identifier": {
		"index": 16516418,
		"hash": "0x4e898d096370f572bff77cd282ac937c09b69bba4b603e03322041866a3b9d89"
	},
	"balances": [
		{
			"value": "0",
			"currency": {
				"symbol": "ICX",
				"decimals": 18
			}
		}
	]
}
```

## Block

**/block**

*Get a Block*s

Request:

Using Index)

```json
{
	"network_identifier": {
		"blockchain": "ICON",
		"network": "Testnet"
	},
	"block_identifier": {
		"index": 1
	}
}
```

Using Hash

```json
{
	"network_identifier": {
		"blockchain": "ICON",
		"network": "Testnet"
	},
	"block_identifier": {
		"hash": "0x7a93f6e66063e2399e5fdce9c921bd5214e22e3d145e3523545a1f69a3b8d8c7"
	}
}
```



Response:

Sample

```json
{
	"block": {
		"block_identifier": {
			"index": 1,
			"hash": "0x7a93f6e66063e2399e5fdce9c921bd5214e22e3d145e3523545a1f69a3b8d8c7"
		},
		"parent_block_identifier": {
			"index": 0,
			"hash": "0xec470d5b046106b86a87eff612af8b1b0279312a4f14a70323828faad47e4364"
		},
		"timestamp": 1578533634950,
		"transactions": null,
		"metadata": {
			"leader": "hxf2152fad2960b17bedf60d4bc249b42d7f2a0ebc",
			"leaderVotesHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
			"logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			"nextLeader": "hxf2152fad2960b17bedf60d4bc249b42d7f2a0ebc",
			"nextRepsHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
			"prevVotesHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
			"receiptsHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
			"repsHash": "0x38ec404f0d0d90a9a8586eccf89e3e78de0d3c7580063b20823308e7f722cd12",
			"signature": "knkcRLajtp6q7DidyCsBSUHND2MMlbtj7HRGPFRi2KcbPp3jU1D6WAikzsKcR/yPs3V2X+h52J4tWBAmTosiXgE=",
			"stateHash": "0xa7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a",
			"transactionsHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
			"version": "0.3"
		}
	}
}
```

**/block/transaction**

*Get a Block Transaction - Payment*

Request:

```json
{
	"network_identifier": {
		"blockchain": "ICON",
		"network": "Testnet"
	},
	"block_identifier": {
		"index": 1582509,
		"hash": "4cc2adbb6fe5f14952b1a7043b0a3fb0a33016fe0de99d1bc2102f349e3cd3ad"
	},
	"transaction_identifier": {
		"hash": "0x2c89b69a75ce737ac61b76a6a86ffa233362ae7b05eabd54d067737e282b75c0"
	}
}
```

Response:

Sample
__Note__: The operation type is `base`.

```json
{
	"transaction": {
		"transaction_identifier": {
			"hash": "0x2c89b69a75ce737ac61b76a6a86ffa233362ae7b05eabd54d067737e282b75c0"
		},
		"operations": [
			{
				"operation_identifier": {
					"index": 0
				},
				"type": "BASE",
				"status": "SUCCESS"
			},
			{
				"operation_identifier": {
					"index": 1
				},
				"type": "ISSUE",
				"status": "SUCCESS",
				"account": {
					"address": "hx1000000000000000000000000000000000000000"
				},
				"amount": {
					"value": "3093140432098765383",
					"currency": {
						"symbol": "ICX",
						"decimals": 18
					}
				}
			}
		],
		"metadata": {
			"data": {
				"prep": {
					"irep": "0xa968163f0a57b400000",
					"rrep": "0x3fe",
					"totalDelegation": "0x2b2dc7ab9162daa1000000",
					"value": "0x2aed0ad5b7a52a47"
				},
				"result": {
					"coveredByFee": "0x0",
					"coveredByOverIssuedICX": "0x0",
					"issue": "0x2aed0ad5b7a52a47"
				}
			},
			"dataType": "base",
			"nid": null,
			"nonce": null,
			"signature": null,
			"timestamp": "0x59bda519b1c52",
			"version": "0x3"
		}
	}
}
```



# TODO

*Get a Block Transaction - Contract Deployment*

Request:

```json
{
    "network_identifier": {
        "blockchain": "zilliqa",
        "network": "mainnet"
    },
    "block_identifier": {
    	"index": 670379,
    	"hash": "e71a6d73ec69accb63cb77e67ce6bdde92e6de5a9f1981d8cd9f2f4630031a7b"
    },
    "transaction_identifier": {
    	"hash": "5a3662d689468b423f050824c93343b790a7295d44a4e0f5ebee119ecc18d065"
    }
}
```

Response:

Sample
__Note__: The operation type is `contract_deployment`.

```json
{
    "transaction": {
        "transaction_identifier": {
            "hash": "5a3662d689468b423f050824c93343b790a7295d44a4e0f5ebee119ecc18d065"
        },
        "operations": [
            {
                "operation_identifier": {
                    "index": 0
                },
                "type": "contract_deployment",
                "status": "SUCCESS",
                "account": {
                    "address": "zil1a35lxvh38y3u8xe7kzxfkgdhmctj387zs92llt",
                    "metadata": {
                        "base16": "ec69F332F13923C39B3eB08c9b21B7De17289FC2"
                    }
                },
                "amount": {
                    "value": "0",
                    "currency": {
                        "symbol": "ZIL",
                        "decimals": 12
                    }
                },
                "metadata": {
                    "code": "\nscilla_version 0\nimport BoolUtils\nlibrary ResolverLib\ntype RecordKeyValue =\n  | RecordKeyValue of String String\nlet nilMessage = Nil {Message}\nlet oneMsg =\n  fun(msg: Message) =>\n    Cons {Message} msg nilMessage\nlet eOwnerSet =\n  fun(address: ByStr20) =>\n    {_eventname: \"OwnerSet\"; address: address}\n(* @deprecated eRecordsSet is emitted instead (since 0.1.1) *)\nlet eRecordSet =\n  fun(key: String) =>\n  fun(value: String) =>\n    {_eventname: \"RecordSet\"; key: key; value: value}\n(* @deprecated eRecordsSet is emitted instead (since 0.1.1) *)\nlet eRecordUnset =\n  fun(key: String) =>\n    {_eventname: \"RecordUnset\"; key: key}\nlet eRecordsSet =\n  fun(registry: ByStr20) =>\n  fun(node: ByStr32) =>\n    {_eventname: \"RecordsSet\"; registry: registry; node: node}\nlet eError =\n  fun(message: String) =>\n    {_eventname: \"Error\"; message: message}\nlet emptyValue = \"\"\nlet mOnResolverConfigured =\n  fun(registry: ByStr20) =>\n  fun(node: ByStr32) =>\n    let m = {_tag: \"onResolverConfigured\"; _amount: Uint128 0; _recipient: registry; node: node} in\n      oneMsg m\nlet copyRecordsFromList =\n  fun (recordsMap: Map String String) =>\n  fun (recordsList: List RecordKeyValue) =>\n    let foldl = @list_foldl RecordKeyValue Map String String in\n      let iter =\n        fun (recordsMap: Map String String) =>\n        fun (el: RecordKeyValue) =>\n          match el with\n          | RecordKeyValue key val =>\n            let isEmpty = builtin eq val emptyValue in\n              match isEmpty with\n              | True => builtin remove recordsMap key\n              | False => builtin put recordsMap key val\n              end\n          end\n      in\n        foldl iter recordsMap recordsList\ncontract Resolver(\n  initialOwner: ByStr20,\n  registry: ByStr20,\n  node: ByStr32,\n  initialRecords: Map String String\n)\nfield vendor: String = \"UD\"\nfield version: String = \"0.1.1\"\nfield owner: ByStr20 = initialOwner\nfield records: Map String String = initialRecords\n(* Sets owner address *)\n(* @ensures a sender address is an owner of the contract *)\n(* @param address *)\n(* @emits OwnerSet if the operation was successful *)\n(* @emits Error if a sender address has no permission for the operation *)\ntransition setOwner(address: ByStr20)\n  currentOwner <- owner;\n  isOkSender = builtin eq currentOwner _sender;\n  match isOkSender with\n  | True =>\n    owner := address;\n    e = eOwnerSet address;\n    event e\n  | _ =>\n    e = let m = \"Sender not owner\" in eError m;\n    event e\n  end\nend\n(* Sets a key value pair *)\n(* @ensures a sender address is an owner of the contract *)\n(* @param key *)\n(* @param value *)\n(* @emits RecordSet if the operation was successful *)\n(* @emits Error if a sender address has no permission for the operation *)\n(* @sends onResolverConfigured to the registry *)\ntransition set(key: String, value: String)\n  currentOwner <- owner;\n  isOkSender = builtin eq currentOwner _sender;\n  match isOkSender with\n  | True =>\n    records[key] := value;\n    e = eRecordsSet registry node;\n    event e;\n    msgs = mOnResolverConfigured registry node;\n    send msgs\n  | _ =>\n    e = let m = \"Sender not owner\" in eError m;\n    event e\n  end\nend\n(* Remove a key from records map *)\n(* @ensures a sender address is an owner of the contract *)\n(* @param key *)\n(* @emits RecordUnset if the operation was successful *)\n(* @emits Error if a sender address has no permission for the operation *)\n(* @sends onResolverConfigured to the registry *)\ntransition unset(key: String)\n  keyExists <- exists records[key];\n  currentOwner <- owner;\n  isOk =\n    let isOkSender = builtin eq currentOwner _sender in\n      andb isOkSender keyExists;\n  match isOk with\n  | True =>\n    delete records[key];\n    e = eRecordsSet registry node;\n    event e;\n    msgs = mOnResolverConfigured registry node;\n    send msgs\n  | _ =>\n    e = let m = \"Sender not owner or key does not exist\" in\n      eError m;\n    event e\n  end\nend\n(* Set multiple keys to records map *)\n(* Removes records from the map if according passed value is empty *)\n(* @ensures a sender address is an owner of the contract *)\n(* @param newRecords *)\n(* @emits RecordsSet if the operation was successful *)\n(* @emits Error if a sender address has no permission for the operation *)\n(* @sends onResolverConfigured to the registry *)\ntransition setMulti(newRecords: List RecordKeyValue)\n  currentOwner <- owner;\n  isOkSender = builtin eq currentOwner _sender;\n  match isOkSender with\n  | True =>\n    oldRecords <- records;\n    newRecordsMap = copyRecordsFromList oldRecords newRecords;\n    records := newRecordsMap;\n    e = eRecordsSet registry node;\n    event e;\n    msgs = mOnResolverConfigured registry node;\n    send msgs\n  | _ =>\n    e = let m = \"Sender not owner\" in eError m;\n    event e\n  end\nend\n",
                    "data": "[{\"vname\":\"_scilla_version\",\"type\":\"Uint32\",\"value\":\"0\"},{\"vname\":\"initialOwner\",\"type\":\"ByStr20\",\"value\":\"0x4887fb6920a8ae50886543ee8aa504da6c9f83bf\"},{\"vname\":\"registry\",\"type\":\"ByStr20\",\"value\":\"0x9611c53be6d1b32058b2747bdececed7e1216793\"},{\"vname\":\"node\",\"type\":\"ByStr32\",\"value\":\"0xd72c3c6e1e3b1b1238b5ba82ff7afe688f542b1cdbfee692a912dd88b1d31f76\"},{\"vname\":\"initialRecords\",\"type\":\"Map String String\",\"value\":[{\"key\":\"ZIL\",\"val\":\"0x803637d03997e4c29729e9ce9e4bc41c0c867354\"}]}]",
                    "gasLimit": "10000",
                    "gasPrice": "1000000000",
                    "nonce": "2007",
                    "receipt": {
                        "accept": false,
                        "errors": null,
                        "exceptions": null,
                        "success": true,
                        "cumulative_gas": "6024",
                        "epoch_num": "670379",
                        "event_logs": null,
                        "transitions": null
                    },
                    "senderPubKey": "0x032E38FCA06A680FFE1BA40956ADA08CB94236FC985B4F7571D455408A0A27E1A2",
                    "signature": "0xB665902059A7519F9A8E118B87ACE4EFDC0FB434475617B19B94E38ABAB68AE8DC85650E781C52CBE281FA527E7556EE9BC593531743A3594BF25C4330EC0165",
                    "version": "65537"
                }
            }
        ]
    }
}
```

### Displaying Contract Calls Information for Block Transactions
A contract call can be defined in the either one of the following two forms:
1. An account has invoked a function in a contract
2. An account has invoked a function in a contract which further invokes another function in a different contract (a.k.a Chain Calling)

Depending on the functions invoked by the contract, a contract call may perform additional smart contract deposits to some accounts.
These smart contract deposits will be shown under the `operations []` idential to how typical payment transaction is displayed.
Additional metadata information related to the transaction such as the *contract address* and *gas amount* are displayed only at the __final operation block__ to reduce cluttering of metadata.

*Get a Block Transaction - Contract Call without Smart Contract Deposits*

Request:

```json
{
    "network_identifier": {
        "blockchain": "zilliqa",
        "network": "testnet"
    },
    "block_identifier": {
    	"index": 1558244,
    	"hash": "4f00a6059b22ebd73e6a60d77fbc20f65bfa3be3f5ae57712422699e3bb031ac"
    },
    "transaction_identifier": {
    	"hash": "ad8a8aa7c1aff0a59a3d56f9c9a72176c344e8a35bbd66e69b2bc7011b44e637"
    }
}
```

Response:

Sample
__Note__: The operation type is `contract_call`.

```json
{
    "transaction": {
        "transaction_identifier": {
            "hash": "ad8a8aa7c1aff0a59a3d56f9c9a72176c344e8a35bbd66e69b2bc7011b44e637"
        },
        "operations": [
            {
                "operation_identifier": {
                    "index": 0
                },
                "type": "contract_call",
                "status": "SUCCESS",
                "account": {
                    "address": "zil1ha4z3qu69uxr6h2m7v9ggcjt332cjupzp7c2ae",
                    "metadata": {
                        "base16": "Bf6a28839a2f0c3d5d5Bf30A84624B8c55897022"
                    }
                },
                "amount": {
                    "value": "0",
                    "currency": {
                        "symbol": "ZIL",
                        "decimals": 12
                    }
                },
                "metadata": {
                    "contractAddress": "c36087407e6474e038d7c316a620afe2a752ad0e",
                    "data": "{\"_tag\":\"SubmitHeaderBlock\",\"params\":[{\"vname\":\"new_hash\",\"type\":\"ByStr32\",\"value\":\"0x62c8b569f485f22878f8b31f6e159981be4ea78bdeb09062fbcdcbc4802deae2\"},{\"vname\":\"block\",\"type\":\"Uint64\",\"value\":\"178656\"}]}",
                    "gasLimit": "40000",
                    "gasPrice": "1000000000",
                    "nonce": "352",
                    "receipt": {
                        "accept": false,
                        "errors": null,
                        "exceptions": null,
                        "success": true,
                        "cumulative_gas": "841",
                        "epoch_num": "1558244",
                        "event_logs": [
                            {
                                "_eventname": "SubmitHashSuccess",
                                "address": "0xc36087407e6474e038d7c316a620afe2a752ad0e",
                                "params": [
                                    {
                                        "type": "ByStr32",
                                        "value": "0x62c8b569f485f22878f8b31f6e159981be4ea78bdeb09062fbcdcbc4802deae2",
                                        "vname": "hash"
                                    },
                                    {
                                        "type": "Int32",
                                        "value": "2",
                                        "vname": "code"
                                    }
                                ]
                            }
                        ],
                        "transitions": null
                    },
                    "senderPubKey": "0x025A5A6AFBB5797E44F29FEFA81B43EB3600C70F021B78ABCE7CF2D4D01D467AFF",
                    "signature": "0xA9FA2B79A0927B544528693D51BB7FCAD1E283146310CE3B12167EAA982AF69EB879942A8310D66F4D5E46655C930DFB4664861435F7CCC2E3DACA653A3966FF",
                    "version": "21823489"
                }
            }
        ]
    }
}
```

*Get a Block Transaction - Contract Call with Smart Contract Deposits (With Chain Calls)*

Request:

```json
{
    "network_identifier": {
        "blockchain": "zilliqa",
        "network": "testnet"
    },
    "block_identifier": {
    	"index": 1406004,
    	"hash": "84c14dc0685e01b3c7d06f2f2dd9198880b182a82d16ed62a67752560badc6b7"
    },
    "transaction_identifier": {
    	"hash": "17c46c252569a4f3fc41ae45fc6a898892b3f75dde11d517f8b7a037caf658e3"
    }
}
```

Response:


Sample
__Note__: The operation type is `contract_call`, follow by `contract_call_transfer` for subsequent operations with a smart contract deposits.


In the sample, the sequence of operations are as follows:
- Initiator `zil16ura3fhsf84h60s7w6xjy4u2wxel892n7sq5dp` -> Contract `zil135gsjk2wqxwecn00axm2s40ey6g6ne8668046h` (invokes a contract call to add funds)
- Contract `zil135gsjk2wqxwecn00axm2s40ey6g6ne8668046h` (`8d1109594e019d9c4defe9b6a855f92691a9e4fa`) -> Recipient `zil12n6h5gqhlpw87gtzlqe5sq5r7pq2spj8x2g8pe` (amount is deducted from contract balance and trasnferred to recipient)

```json
{
    "transaction": {
        "transaction_identifier": {
            "hash": "17c46c252569a4f3fc41ae45fc6a898892b3f75dde11d517f8b7a037caf658e3"
        },
        "operations": [
            {
                "operation_identifier": {
                    "index": 0
                },
                "type": "contract_call",
                "status": "SUCCESS",
                "account": {
                    "address": "zil16ura3fhsf84h60s7w6xjy4u2wxel892n7sq5dp",
                    "metadata": {
                        "base16": "d707D8a6F049Eb7d3E1e768D22578A71b3f39553"
                    }
                },
                "amount": {
                    "value": "0",
                    "currency": {
                        "symbol": "ZIL",
                        "decimals": 12
                    }
                }
            },
            {
                "operation_identifier": {
                    "index": 1
                },
                "related_operations": [
                    {
                        "index": 0
                    }
                ],
                "type": "contract_call_transfer",
                "status": "SUCCESS",
                "account": {
                    "address": "zil135gsjk2wqxwecn00axm2s40ey6g6ne8668046h",
                    "metadata": {
                        "base16": "8D1109594E019D9C4DEFe9b6A855F92691A9E4fA"
                    }
                },
                "amount": {
                    "value": "-123073860347289351",
                    "currency": {
                        "symbol": "ZIL",
                        "decimals": 12
                    }
                }
            },
            {
                "operation_identifier": {
                    "index": 2
                },
                "related_operations": [
                    {
                        "index": 1
                    }
                ],
                "type": "contract_call_transfer",
                "status": "SUCCESS",
                "account": {
                    "address": "zil12n6h5gqhlpw87gtzlqe5sq5r7pq2spj8x2g8pe",
                    "metadata": {
                        "base16": "54F57A2017F85c7F2162f833480283f040a80647"
                    }
                },
                "amount": {
                    "value": "123073860347289351",
                    "currency": {
                        "symbol": "ZIL",
                        "decimals": 12
                    }
                },
                "metadata": {
                    "contractAddress": "8d1109594e019d9c4defe9b6a855f92691a9e4fa",
                    "data": "{\"_tag\": \"AddFunds\", \"params\": []}",
                    "gasLimit": "10000",
                    "gasPrice": "1000000000",
                    "nonce": "8",
                    "receipt": {
                        "accept": false,
                        "errors": null,
                        "exceptions": null,
                        "success": true,
                        "cumulative_gas": "1402",
                        "epoch_num": "1406004",
                        "event_logs": [
                            {
                                "_eventname": "Verifier add funds",
                                "address": "0x54f57a2017f85c7f2162f833480283f040a80647",
                                "params": [
                                    {
                                        "type": "ByStr20",
                                        "value": "0xd707d8a6f049eb7d3e1e768d22578a71b3f39553",
                                        "vname": "verifier"
                                    }
                                ]
                            }
                        ],
                        "transitions": [
                            {
                                "accept": false,
                                "addr": "0x8d1109594e019d9c4defe9b6a855f92691a9e4fa",
                                "depth": 0,
                                "msg": {
                                    "_amount": "123073860347289351",
                                    "_recipient": "0x54f57a2017f85c7f2162f833480283f040a80647",
                                    "_tag": "AddFunds",
                                    "params": [
                                        {
                                            "vname": "initiator",
                                            "type": "ByStr20",
                                            "value": "0xd707d8a6f049eb7d3e1e768d22578a71b3f39553"
                                        }
                                    ]
                                }
                            }
                        ]
                    },
                    "senderPubKey": "0x02BCD59F13A3DF40DE7D6B901B10DA416D2EFDD41E9A3631D6673809D7F5B9C4EF",
                    "signature": "0xB0B321303D8CABDC3E1AD6B3ECD5CECD90A7D0A839C69C1C23A68CC0AFD283DCC9696EB503EBAE167C3DF3943A54E6EAA1D35D28D9F414FBA44109DAAAEF4F56",
                    "version": "21823489"
                }
            }
        ]
    }
}
```

## Construction

### Construction Flow
The construction flow is always in this sequence:
1. /construction/derive
2. /construction/preprocess
3. /construction/metadata
4. /construction/payloads
5. /construction/parse
6. /construction/combine
7. /construction/parse (to confirm correctness)
8. /construction/hash
9. /construction/submit

**/construction/combine**

*Create Network Transaction from Signature*

__Note__: Before calling `/combine`, please call `/payloads` to have the `unsigned_transaction`. Next, use [goZilliqa SDK](https://github.com/Zilliqa/gozilliqa-sdk) or other Zilliqa's SDKs to craft a transaction object and sign the transaction object; print out the __*signature*__ and __*transaction object*__ in __hexadecimals__. 

Refer to the `signRosettaTransaction.js` in the `examples` folder to craft and sign a transaction object.

Use them as request parameters as follows:

```
{
    ...,
    "unsigned_transaction": ... // from /payloads
    "signatures": [
        {
            "signing_payload": {
                "address": "string", // sender account address
                "hex_bytes": "string",  // signed transaction object in hexadecimals representation obtained after signing with goZilliqa SDK or other Zilliqa SDK
                "signature_type": "ecdsa"
            },
            "public_key": {
                "hex_bytes": "string", // sender public key
                "curve_type": "secp256k1"
            },
            "signature_type": "ecdsa",
            "hex_bytes": "string" // signature of the signed transaction object 
        }
    ]
}

```


Request:

```json
{
    "network_identifier": {
        "blockchain": "zilliqa",
        "network": "testnet"
    },
    "unsigned_transaction": "{\"amount\":2000000000000,\"code\":\"\",\"data\":\"\",\"gasLimit\":1,\"gasPrice\":2000000000,\"nonce\":187,\"pubKey\":\"02e44ef2c5c2031386faa6cafdf5f67318cc661871b0112a27458e65f37a35655e\",\"senderAddr\":\"zil1n8uafq4thhzlq5nj50p55al9jvamr3s45hm49r\",\"toAddr\":\"zil1f9uqwhwkq7fnzgh5x4djyzg4a7j3apx8dsnnc0\",\"version\":21823489}",
    "signatures": [
        {
            "signing_payload": {
                "account_identifier": {
                    "address": "zil1n8uafq4thhzlq5nj50p55al9jvamr3s45hm49r",
                    "metadata": {
                        "base16": "99f9d482abbdC5F05272A3C34a77E5933Bb1c615"
                    }
                },
                "hex_bytes": "088180b40a10bb011a144978075dd607933122f4355b220915efa51e84c722230a2102e44ef2c5c2031386faa6cafdf5f67318cc661871b0112a27458e65f37a35655e2a120a100000000000000000000001d1a94a200032120a10000000000000000000000000773594003801",
                "signature_type": "schnorr_1"
            },
            "public_key": {
                "hex_bytes": "02e44ef2c5c2031386faa6cafdf5f67318cc661871b0112a27458e65f37a35655e",
                "curve_type": "secp256k1"
            },
            "signature_type": "schnorr_1",
            "hex_bytes": "fcb93583d963a7c11f52f04b1ecbd129aa3df896e618b47ff163dc18c53b59afc4289851fd2d5a50eaa7d7ae0763eb912797b0b34e1cf1e6d3865a218e1066b7"
        }
    ]
}
```

Response:

Sample

```json
{
    "signed_transaction": "{\"amount\":2000000000000,\"code\":\"\",\"data\":\"\",\"gasLimit\":1,\"gasPrice\":2000000000,\"nonce\":187,\"pubKey\":\"02e44ef2c5c2031386faa6cafdf5f67318cc661871b0112a27458e65f37a35655e\",\"senderAddr\":\"zil1n8uafq4thhzlq5nj50p55al9jvamr3s45hm49r\",\"signature\":\"fcb93583d963a7c11f52f04b1ecbd129aa3df896e618b47ff163dc18c53b59afc4289851fd2d5a50eaa7d7ae0763eb912797b0b34e1cf1e6d3865a218e1066b7\",\"toAddr\":\"zil1f9uqwhwkq7fnzgh5x4djyzg4a7j3apx8dsnnc0\",\"version\":21823489}"
}
```

**/construction/metadata**

*Create a Request to Fetch Metadata*

Request:

`options` is from `/construction/preprocess`
```json
{
    "network_identifier": {
        "blockchain": "zilliqa",
        "network": "testnet"
    }, 
    "options": {
        "amount": "2000000000000",
        "gasLimit": "1",
        "gasPrice": "2000000000",
        "senderAddr": "zil1n8uafq4thhzlq5nj50p55al9jvamr3s45hm49r",
        "toAddr": "zil1f9uqwhwkq7fnzgh5x4djyzg4a7j3apx8dsnnc0"
    },
    "public_keys": [
        {
            "hex_bytes": "02e44ef2c5c2031386faa6cafdf5f67318cc661871b0112a27458e65f37a35655e",
            "curve_type": "secp256k1"
        }
    ]
}
```

Response:

Sample

```json
{
    "metadata": {
        "amount": "2000000000000",
        "gasLimit": "1",
        "gasPrice": "2000000000",
        "nonce": 187,
        "pubKey": "02e44ef2c5c2031386faa6cafdf5f67318cc661871b0112a27458e65f37a35655e",
        "senderAddr": "zil1n8uafq4thhzlq5nj50p55al9jvamr3s45hm49r",
        "toAddr": "zil1f9uqwhwkq7fnzgh5x4djyzg4a7j3apx8dsnnc0",
        "version": 21823489
    }
}
```

**/construction/derive**

*Derive an Address from a PublicKey*

Request:

```json
{
    "network_identifier": {
        "blockchain": "zilliqa",
        "network": "mainnet"
    },
    "public_key": {
        "hex_bytes": "026c7f3b8ac6f615c00c34186cbe4253a2c5acdc524b1cfae544c629d8e3564cfc",
        "curve_type": "secp256k1"
    },
    "metadata": {
    	"type": "bech32"
    }
}
```

Response:

Sample

```json
{
    "address": "zil1y9qmlzmdygfaf4eqfcka4wfx20wzghzl05xazc",
    "account_identifier": {
        "address": "zil1y9qmlzmdygfaf4eqfcka4wfx20wzghzl05xazc",
        "metadata": {
            "base16": "2141BF8B6D2213d4d7204E2DDAB92653dC245c5F"
        }
    }
}
```

**/construction/hash**

*Get the Hash of a Signed Transaction*

Request:

```json
{
    "network_identifier": {
        "blockchain": "zilliqa",
        "network": "testnet"
    },
	"signed_transaction": "{\"amount\":2000000000000,\"code\":\"\",\"data\":\"\",\"gasLimit\":1,\"gasPrice\":1000000000,\"nonce\":186,\"pubKey\":\"02e44ef2c5c2031386faa6cafdf5f67318cc661871b0112a27458e65f37a35655e\",\"signature\":\"51c69af638ad7afd39841a7abf937d5df99e20adedc4287f43c8070d497ba78136c951192b3920914feb83b9272ccb2ca7facd835dfad10eff2b848b13616daf\",\"toAddr\":\"zil1f9uqwhwkq7fnzgh5x4djyzg4a7j3apx8dsnnc0\",\"version\":21823489}"
}
```

Response:

Sample

```json
{
    "transaction_identifier": {
        "hash": "a17367c8bcd83cdc2d9ede4571c8e27ad74278ae195263f13e10ba84f12ab13c"
    }
}
```

**/construction/parse**

*Parse a Transaction*

__Note__: Set the `signed` flag accordingly if the transaction is signed or unsigned.

Request:

```json
{
    "network_identifier": {
        "blockchain": "zilliqa",
        "network": "testnet"
    },
    "signed": true,
    "transaction": "{\"amount\":2000000000000,\"code\":\"\",\"data\":\"\",\"gasLimit\":1,\"gasPrice\":1000000000,\"nonce\":186,\"pubKey\":\"02e44ef2c5c2031386faa6cafdf5f67318cc661871b0112a27458e65f37a35655e\",\"signature\":\"51c69af638ad7afd39841a7abf937d5df99e20adedc4287f43c8070d497ba78136c951192b3920914feb83b9272ccb2ca7facd835dfad10eff2b848b13616daf\",\"toAddr\":\"zil1f9uqwhwkq7fnzgh5x4djyzg4a7j3apx8dsnnc0\",\"version\":21823489}"
}
```

Response:

Sample

```json
{
    "signers": [
        "zil1n8uafq4thhzlq5nj50p55al9jvamr3s45hm49r"
    ],
    "operations": [
        {
            "operation_identifier": {
                "index": 0
            },
            "type": "transfer",
            "status": "",
            "account": {
                "address": "zil1n8uafq4thhzlq5nj50p55al9jvamr3s45hm49r",
                "metadata": {
                    "base16": "99f9d482abbdC5F05272A3C34a77E5933Bb1c615"
                }
            },
            "amount": {
                "value": "-2000000000000",
                "currency": {
                    "symbol": "ZIL",
                    "decimals": 12
                }
            }
        },
        {
            "operation_identifier": {
                "index": 1
            },
            "related_operations": [
                {
                    "index": 0
                }
            ],
            "type": "transfer",
            "status": "",
            "account": {
                "address": "zil1f9uqwhwkq7fnzgh5x4djyzg4a7j3apx8dsnnc0",
                "metadata": {
                    "base16": "4978075dd607933122f4355B220915EFa51E84c7"
                }
            },
            "amount": {
                "value": "2000000000000",
                "currency": {
                    "symbol": "ZIL",
                    "decimals": 12
                }
            }
        }
    ],
    "account_identifier_signers": [
        {
            "address": "zil1n8uafq4thhzlq5nj50p55al9jvamr3s45hm49r",
            "metadata": {
                "base16": "99f9d482abbdC5F05272A3C34a77E5933Bb1c615"
            }
        }
    ]
}
```

**/construction/payloads**

*Generate an Unsigned Transaction and Signing Payloads*

Request:

`metadata` for `operation_identifier 1` is from `/construction/metadata`
```json
{
    "network_identifier": {
        "blockchain": "zilliqa",
        "network": "testnet"
    },
	"operations": [
        {
            "operation_identifier": {
                "index": 0
            },
            "type": "transfer",
            "status": "",
            "account": {
                "address": "zil1n8uafq4thhzlq5nj50p55al9jvamr3s45hm49r",
                "metadata": {
                    "base16": "99f9d482abbdC5F05272A3C34a77E5933Bb1c615"
                }
            },
            "amount": {
                "value": "-2000000000000",
                "currency": {
                    "symbol": "ZIL",
                    "decimals": 12
                }
            }
        },
        {
            "operation_identifier": {
                "index": 1
            },
            "related_operations": [
                {
                    "index": 0
                }
            ],
            "type": "transfer",
            "status": "",
            "account": {
                "address": "zil1f9uqwhwkq7fnzgh5x4djyzg4a7j3apx8dsnnc0",
                "metadata": {
                    "base16": "4978075dd607933122f4355B220915EFa51E84c7"
                }
            },
            "amount": {
                "value": "2000000000000",
                "currency": {
                    "symbol": "ZIL",
                    "decimals": 12
                }
            }
        }
    ],
    "metadata": {       // from construction/metadata
        "amount": "2000000000000",
        "gasLimit": "1",
        "gasPrice": "2000000000",
        "nonce": 187,
        "pubKey": "02e44ef2c5c2031386faa6cafdf5f67318cc661871b0112a27458e65f37a35655e",
        "senderAddr": "zil1n8uafq4thhzlq5nj50p55al9jvamr3s45hm49r",
        "toAddr": "zil1f9uqwhwkq7fnzgh5x4djyzg4a7j3apx8dsnnc0",
        "version": 21823489
    },
    "public_keys": [
        {
            "hex_bytes": "02e44ef2c5c2031386faa6cafdf5f67318cc661871b0112a27458e65f37a35655e",
            "curve_type": "secp256k1"
        }
    ]
}
```

Response:

Sample

```json
{
    "unsigned_transaction": "{\"amount\":2000000000000,\"code\":\"\",\"data\":\"\",\"gasLimit\":1,\"gasPrice\":2000000000,\"nonce\":187,\"pubKey\":\"02e44ef2c5c2031386faa6cafdf5f67318cc661871b0112a27458e65f37a35655e\",\"senderAddr\":\"zil1n8uafq4thhzlq5nj50p55al9jvamr3s45hm49r\",\"toAddr\":\"zil1f9uqwhwkq7fnzgh5x4djyzg4a7j3apx8dsnnc0\",\"version\":21823489}",
    "payloads": [
        {
            "address": "zil1n8uafq4thhzlq5nj50p55al9jvamr3s45hm49r",
            "hex_bytes": "088180b40a10bb011a144978075dd607933122f4355b220915efa51e84c722230a2102e44ef2c5c2031386faa6cafdf5f67318cc661871b0112a27458e65f37a35655e2a120a100000000000000000000001d1a94a200032120a10000000000000000000000000773594003801",
            "account_identifier": {
                "address": "zil1n8uafq4thhzlq5nj50p55al9jvamr3s45hm49r",
                "metadata": {
                    "base16": "99f9d482abbdC5F05272A3C34a77E5933Bb1c615"
                }
            },
            "signature_type": "schnorr_1"
        }
    ]
}
```

**/construction/preprocess**

*Create a Request to Fetch Metadata*

Request:

```json
{
    "network_identifier": {
        "blockchain": "zilliqa",
        "network": "testnet"
    },
	"operations": [
        {
            "operation_identifier": {
                "index": 0
            },
            "type": "transfer",
            "status": "",
            "account": {
                "address": "zil1n8uafq4thhzlq5nj50p55al9jvamr3s45hm49r",
                "metadata": {
                    "base16": "99f9d482abbdC5F05272A3C34a77E5933Bb1c615"
                }
            },
            "amount": {
                "value": "-2000000000000",
                "currency": {
                    "symbol": "ZIL",
                    "decimals": 12
                }
            }
        },
        {
            "operation_identifier": {
                "index": 1
            },
            "related_operations": [
                {
                    "index": 0
                }
            ],
            "type": "transfer",
            "status": "",
            "account": {
                "address": "zil1f9uqwhwkq7fnzgh5x4djyzg4a7j3apx8dsnnc0",
                "metadata": {
                    "base16": "4978075dd607933122f4355B220915EFa51E84c7"
                }
            },
            "amount": {
                "value": "2000000000000",
                "currency": {
                    "symbol": "ZIL",
                    "decimals": 12
                }
            },
            "metadata": {
                "senderPubKey": "0x02e44ef2c5c2031386faa6cafdf5f67318cc661871b0112a27458e65f37a35655e"
            }
        }
    ],
    "metadata": {}
}
```

Response:

Sample

```json
{
    "options": {
        "amount": "2000000000000",
        "gasLimit": "1",
        "gasPrice": "2000000000",
        "senderAddr": "zil1n8uafq4thhzlq5nj50p55al9jvamr3s45hm49r",
        "toAddr": "zil1f9uqwhwkq7fnzgh5x4djyzg4a7j3apx8dsnnc0"
    },
    "required_public_keys": [
        {
            "address": "zil1n8uafq4thhzlq5nj50p55al9jvamr3s45hm49r",
            "metadata": {
                "base16": "99f9d482abbdC5F05272A3C34a77E5933Bb1c615"
            }
        }
    ]
}
```

**/construction/submit**

*Submit a Signed Transaction*

__Note__: Before calling `/submit`, please call `/combine` to obtain the `signed_transaction` required for the  request parameters.

Request:

```json
{
    "network_identifier": {
        "blockchain": "zilliqa",
        "network": "testnet"
    },
    "signed_transaction": "{\"amount\":2000000000000,\"code\":\"\",\"data\":\"\",\"gasLimit\":1,\"gasPrice\":2000000000,\"nonce\":187,\"pubKey\":\"02e44ef2c5c2031386faa6cafdf5f67318cc661871b0112a27458e65f37a35655e\",\"senderAddr\":\"zil1n8uafq4thhzlq5nj50p55al9jvamr3s45hm49r\",\"signature\":\"fcb93583d963a7c11f52f04b1ecbd129aa3df896e618b47ff163dc18c53b59afc4289851fd2d5a50eaa7d7ae0763eb912797b0b34e1cf1e6d3865a218e1066b7\",\"toAddr\":\"zil1f9uqwhwkq7fnzgh5x4djyzg4a7j3apx8dsnnc0\",\"version\":21823489}"
}
```

Response:

Sample

```json
{
    "transaction_identifier": {
        "hash": "963a984ee255cfd881b337a52caf699d4f05799c45cc0948d8a8ce72a6a12d8e"
    }
}
```

