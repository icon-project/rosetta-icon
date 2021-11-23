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
          "BURN",
          "ISSUE",
          "WITHDRAWN",
          "ICXTRANSFER",
          "CLAIM",
          "MESSAGE",
          "DEPLOY",
          "CALL",
          "DEPOSIT"
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
      ...
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
		"network": "Mainnet"
	},
	"block_identifier": {
		"index": 42589147
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
      "index": 42589147,
      "hash": "0x4f54eb994dc9617dc75a7b67ea0762b8ad4ca168d9424fd91313d55320018038"
    },
    "parent_block_identifier": {
      "index": 42589146,
      "hash": "0x501c6190b95a3d66d76b551c37d94882df19c29c0337991998b89069c538acd7"
    },
    "timestamp": 1637680002937,
    "transactions": [
      {
        "transaction_identifier": {
          "hash": "0x0781ef62ccd47e488deaae84e7341dda2b278cc921f6de4b5428b8ccc92ca18b"
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
              "value": "3069013237069400562",
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
              "irep": "0x21e19e0c9bab2400000",
              "rrep": "0x149",
              "totalDelegation": "0x152a911f25a3d18549947f8",
              "value": "0x2a9753475e7e49f2"
            },
            "result": {
              "coveredByFee": "0x0",
              "coveredByOverIssuedICX": "0x0",
              "issue": "0x2a9753475e7e49f2"
            }
          },
          "dataType": "base",
          "nid": null,
          "nonce": null,
          "signature": null,
          "timestamp": "0x5d17616d4b384",
          "version": "0x3"
        }
      },
      {
        "transaction_identifier": {
          "hash": "0xe3f944bcbb6fddd1c7d67ff0d42748b972f93c41eae2d9b46b3fbdfadbf17a4f"
        },
        "operations": [
          {
            "operation_identifier": {
              "index": 0
            },
            "type": "CALL",
            "status": "SUCCESS",
            "account": {
              "address": "hxdacb2506200292bfe6ab2e3f3771b2d1786eaa77"
            },
            "amount": {
              "value": "-0",
              "currency": {
                "symbol": "ICX",
                "decimals": 18
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
            "type": "CALL",
            "status": "SUCCESS",
            "account": {
              "address": "cx0000000000000000000000000000000000000000"
            },
            "amount": {
              "value": "0",
              "currency": {
                "symbol": "ICX",
                "decimals": 18
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
            "type": "FEE",
            "status": "SUCCESS",
            "account": {
              "address": "hxdacb2506200292bfe6ab2e3f3771b2d1786eaa77"
            },
            "amount": {
              "value": "-1722500000000000",
              "currency": {
                "symbol": "ICX",
                "decimals": 18
              }
            }
          },
          {
            "operation_identifier": {
              "index": 3
            },
            "related_operations": [
              {
                "index": 2
              }
            ],
            "type": "FEE",
            "status": "SUCCESS",
            "account": {
              "address": "hx1000000000000000000000000000000000000000"
            },
            "amount": {
              "value": "1722500000000000",
              "currency": {
                "symbol": "ICX",
                "decimals": 18
              }
            }
          }
        ],
        "metadata": {
          "data": {
            "method": "setStake",
            "params": {
              "value": "0x1cf643ffa2f1183554f"
            }
          },
          "dataType": "call",
          "nid": "0x1",
          "nonce": null,
          "signature": "OyXkRB8oYE0GlyyYoxqiktOHdD4dOBrA3X6mugI15f0N0ktAFLR3NtJ+cpXgqaEjuPzjw1XT6vD84oIoBz6++AA=",
          "timestamp": "0x5d17616e685dc",
          "version": "0x3"
        }
      },
      {
        "transaction_identifier": {
          "hash": "0xd117502b236496ba580467481228524abd34b7ef393a5801a7f2ef9dc4e45ae7"
        },
        "operations": [
          {
            "operation_identifier": {
              "index": 0
            },
            "type": "CALL",
            "status": "SUCCESS",
            "account": {
              "address": "hxd48ab6a11220369368e8847fee7be4625f304b19"
            },
            "amount": {
              "value": "-5000000000000000000",
              "currency": {
                "symbol": "ICX",
                "decimals": 18
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
            "type": "CALL",
            "status": "SUCCESS",
            "account": {
              "address": "cx299d88908ab371d586c8dfe0ed42899a899e6e5b"
            },
            "amount": {
              "value": "5000000000000000000",
              "currency": {
                "symbol": "ICX",
                "decimals": 18
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
            "type": "FEE",
            "status": "SUCCESS",
            "account": {
              "address": "hxd48ab6a11220369368e8847fee7be4625f304b19"
            },
            "amount": {
              "value": "-7769000000000000",
              "currency": {
                "symbol": "ICX",
                "decimals": 18
              }
            }
          },
          {
            "operation_identifier": {
              "index": 3
            },
            "related_operations": [
              {
                "index": 2
              }
            ],
            "type": "FEE",
            "status": "SUCCESS",
            "account": {
              "address": "hx1000000000000000000000000000000000000000"
            },
            "amount": {
              "value": "7769000000000000",
              "currency": {
                "symbol": "ICX",
                "decimals": 18
              }
            }
          },
          {
            "operation_identifier": {
              "index": 4
            },
            "type": "ICXTRANSFER",
            "status": "SUCCESS",
            "account": {
              "address": "cx299d88908ab371d586c8dfe0ed42899a899e6e5b"
            },
            "amount": {
              "value": "-5000000000000000000",
              "currency": {
                "symbol": "ICX",
                "decimals": 18
              }
            }
          },
          {
            "operation_identifier": {
              "index": 5
            },
            "related_operations": [
              {
                "index": 4
              }
            ],
            "type": "ICXTRANSFER",
            "status": "SUCCESS",
            "account": {
              "address": "cx1b97c1abfd001d5cd0b5a3f93f22cccfea77e34e"
            },
            "amount": {
              "value": "5000000000000000000",
              "currency": {
                "symbol": "ICX",
                "decimals": 18
              }
            }
          },
          {
            "operation_identifier": {
              "index": 6
            },
            "type": "ICXTRANSFER",
            "status": "SUCCESS",
            "account": {
              "address": "cx1b97c1abfd001d5cd0b5a3f93f22cccfea77e34e"
            },
            "amount": {
              "value": "-5139000000000000000",
              "currency": {
                "symbol": "ICX",
                "decimals": 18
              }
            }
          },
          {
            "operation_identifier": {
              "index": 7
            },
            "related_operations": [
              {
                "index": 6
              }
            ],
            "type": "ICXTRANSFER",
            "status": "SUCCESS",
            "account": {
              "address": "hxd48ab6a11220369368e8847fee7be4625f304b19"
            },
            "amount": {
              "value": "5139000000000000000",
              "currency": {
                "symbol": "ICX",
                "decimals": 18
              }
            }
          }
        ],
        "metadata": {
          "data": {
            "method": "action",
            "params": {
              "model": "{\"name\":\"custom_bet\",\"params\":{\"number_of_tiles\":24,\"square_id\":1,\"user_seed\":\"\"}}"
            }
          },
          "dataType": "call",
          "nid": "0x1",
          "nonce": "0x1",
          "signature": "YP18Fa0efyVh//dUeH/MaG3B5jL5jN2YNdPADBydkOBSSTKOIoNsalA52/xK8GmYUmO8KOhu2gqL99LpFirjegA=",
          "timestamp": "0x5d176174e0ef8",
          "version": "0x3"
        }
      }
    ],
    "metadata": {
      "merkle_tree_root_hash": "0xda53bbbbb6cdb3a9250aabaf0a49151cd923053fbf59bd8429070f2785633238",
      "next_leader": "hx0000000000000000000000000000000000000000",
      "peer_id": "hx3d5c3ce7554f4d762f6396e53b2c5de07074ec39",
      "version": "2.0"
    }
  }
}
```

**/block/transaction**

*Get a Block Transaction - transfer*

Request:

```json
{
	"network_identifier": {
		"blockchain": "ICON",
		"network": "Mainnet"
	},
	"block_identifier": {
		"index": 7261768,
		"hash": "0xd46e773b23c36a29b05d18f2ff49e7a729511a6a9ff4b84320f0860067bb54ed"
	},
	"transaction_identifier": {
		"hash": "0x9f1b00faf95cce7e99d3616c51cec725b6cd6d870a80e30689fc724ba0f1368a"
	}
}
```

Response:

Sample
__Note__: The operation type is `base`

```json
{
    "transaction": {
        "transaction_identifier": {
            "hash": "0x9f1b00faf95cce7e99d3616c51cec725b6cd6d870a80e30689fc724ba0f1368a"
        },
        "operations": [
            {
                "operation_identifier": {
                    "index": 0
                },
                "type": "TRANSFER",
                "status": "SUCCESS",
                "account": {
                    "address": "hxe7af5fcfd8dfc67530a01a0e403882687528dfcb"
                },
                "amount": {
                    "value": "-1000000000000000000",
                    "currency": {
                        "symbol": "ICX",
                        "decimals": 18
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
                "type": "TRANSFER",
                "status": "SUCCESS",
                "account": {
                    "address": "hxe6080b42dd1c70dc3ce15ff73c19d9256ec2f76c"
                },
                "amount": {
                    "value": "1000000000000000000",
                    "currency": {
                        "symbol": "ICX",
                        "decimals": 18
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
                "type": "FEE",
                "status": "SUCCESS",
                "account": {
                    "address": "hxe7af5fcfd8dfc67530a01a0e403882687528dfcb"
                },
                "amount": {
                    "value": "-1000000000000000",
                    "currency": {
                        "symbol": "ICX",
                        "decimals": 18
                    }
                }
            },
            {
                "operation_identifier": {
                    "index": 3
                },
                "related_operations": [
                    {
                        "index": 2
                    }
                ],
                "type": "FEE",
                "status": "SUCCESS",
                "account": {
                    "address": "hx1000000000000000000000000000000000000000"
                },
                "amount": {
                    "value": "1000000000000000",
                    "currency": {
                        "symbol": "ICX",
                        "decimals": 18
                    }
                }
            }
        ],
        "metadata": {
            "data": null,
            "dataType": null,
            "nid": "0x50",
            "nonce": null,
            "signature": "/1XKWeEiJVnO6/q2Z3MqdKCpAwcvbbH5eDVMsQmSzjwSzQnXuPBjAtZTbZakCFRJLF9uA6HEJH4F2lP5HxB4VQA=",
            "timestamp": "0x5ba193198cc5f",
            "version": "0x3"
        }
    }
}
```



*Get a Block Transaction - Base*

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



*Get a Block Transaction - Contract Deployment*

Request:

```json
{
    "network_identifier": {
        "blockchain": "ICON",
        "network": "Testnet"
    },
    "block_identifier": {
        "index": 12337642,
        "hash": "0x16c8dc782c5636c98553b5f857933d2d843a90cd5cb311f95d05d2260e02fe26"
    },
    "transaction_identifier": {
    	"hash": "0x21ef4901281fa3a34533bcfc400a5fa5c25b25fe7a1780022ad15ab729afa8a7"
    }
}
```

Response:

Sample
__Note__: The operation type is `DEPLOY`.

```json
{
    "transaction": {
        "transaction_identifier": {
            "hash": "0x21ef4901281fa3a34533bcfc400a5fa5c25b25fe7a1780022ad15ab729afa8a7"
        },
        "operations": [
            {
                "operation_identifier": {
                    "index": 0
                },
                "type": "DEPLOY",
                "status": "SUCCESS",
                "account": {
                    "address": "hxe7af5fcfd8dfc67530a01a0e403882687528dfcb"
                },
                "amount": {
                    "value": "-0",
                    "currency": {
                        "symbol": "ICX",
                        "decimals": 18
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
                "type": "DEPLOY",
                "status": "SUCCESS",
                "account": {
                    "address": "cx0000000000000000000000000000000000000000"
                },
                "amount": {
                    "value": "0",
                    "currency": {
                        "symbol": "ICX",
                        "decimals": 18
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
                "type": "FEE",
                "status": "SUCCESS",
                "account": {
                    "address": "hxe7af5fcfd8dfc67530a01a0e403882687528dfcb"
                },
                "amount": {
                    "value": "-10227272000000000000",
                    "currency": {
                        "symbol": "ICX",
                        "decimals": 18
                    }
                }
            },
            {
                "operation_identifier": {
                    "index": 3
                },
                "related_operations": [
                    {
                        "index": 2
                    }
                ],
                "type": "FEE",
                "status": "SUCCESS",
                "account": {
                    "address": "hx1000000000000000000000000000000000000000"
                },
                "amount": {
                    "value": "10227272000000000000",
                    "currency": {
                        "symbol": "ICX",
                        "decimals": 18
                    }
                }
            }
        ],
        "metadata": {
            "data": {
                "content": "0x504b0304140000000800957ad0500000000002000000000000001700000073696d706c6553636f72652f5f5f696e69745f5f2e70790300504b0304140000000800957ad05004453b1b3f0000005d0000001800000073696d706c6553636f72652f7061636b6167652e6a736f6eabe6520002a5b2d4a2e2ccfc3c252b0525033d033d43251d88786e62665e7c6e7e4a694e2a48ae3833b72027353839bf28154545315804a820184901572d00504b0304140000000800957ad0509e3c3e9f1b010000c80200001a00000073696d706c6553636f72652f73696d706c6553636f72652e70798d50c16e842010bdfb15134ed06cfd00139b66b387eda5179bbd1a94716bc2820134ddbf2fa041dbd4b6731a98f7e6bd379dd137e85bad2c9aa96f11fadba08d83872ccb5ac9ad85caff48ac5a6d90be7860ec8edc222bb20c7c09eca0ae7bd5bbbaa616657700d11490b027ee7813f0f0f804af5a611169a1ec38a0a12c4f74d1b075e857e51397234209176e4e474ae293048103c4be76f7014beb0c9bcd3ce384ca497da52c99abd05d027631177905044e01838fb8c6d0ca1bb18e4b19b1bf38de20d917fe3808eef01ff405f853dedca2a3640688d92f49f93e1c1ac52535c88556f25ebe9911d7b0ef28a55ee57dca55dda01b8d02720e18f26d61da60f7cfb5e735768bc3344e67df4efff47fddaaef46d8e87b823fe227504b01021403140000000800957ad050000000000200000000000000170000000000000000000000a4810000000073696d706c6553636f72652f5f5f696e69745f5f2e7079504b01021403140000000800957ad05004453b1b3f0000005d000000180000000000000000000000a4813700000073696d706c6553636f72652f7061636b6167652e6a736f6e504b01021403140000000800957ad0509e3c3e9f1b010000c80200001a0000000000000000000000a481ac00000073696d706c6553636f72652f73696d706c6553636f72652e7079504b05060000000003000300d3000000ff0100000000",
                "contentType": "application/zip"
            },
            "dataType": "deploy",
            "nid": "0x3",
            "nonce": null,
            "signature": "fXiCHsxFY/zELt2xvRTCJNNUlAElum+MtD0e/75Rexx4o8CT2DgshS9s5pKRegDXKz4tpJ6i9FeUsj/SaoQ/YQA=",
            "timestamp": "0x5ba17d2e540f0",
            "version": "0x3"
        }
    }
}
```

### Displaying Contract Calls Information for Block Transactions
*Get a Block Transaction - Contract call*

Request:

```json
{
    "network_identifier": {
        "blockchain": "ICON",
        "network": "Testnet"
    },
    "block_identifier": {
        "index": 12338456,
        "hash": "0xe33458479ff9251f3c1e02e01ef74d656a8f3b90e27f6ab4cb46baaacf889453"
    },
    "transaction_identifier": {
    	"hash": "0xafb2049fc951a47659dcc0710fcdcb01caf553d58d3c1ef3da1a7267deeb1a5e"
    }
}
```

Response:

Sample
__Note__: The operation type is `CALL`.

```json
{
    "transaction": {
        "transaction_identifier": {
            "hash": "0xafb2049fc951a47659dcc0710fcdcb01caf553d58d3c1ef3da1a7267deeb1a5e"
        },
        "operations": [
            {
                "operation_identifier": {
                    "index": 0
                },
                "type": "CALL",
                "status": "SUCCESS",
                "account": {
                    "address": "hxe7af5fcfd8dfc67530a01a0e403882687528dfcb"
                },
                "amount": {
                    "value": "-0",
                    "currency": {
                        "symbol": "ICX",
                        "decimals": 18
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
                "type": "CALL",
                "status": "SUCCESS",
                "account": {
                    "address": "cx8f010b715aecb8ff7271176984dd2a25530af388"
                },
                "amount": {
                    "value": "0",
                    "currency": {
                        "symbol": "ICX",
                        "decimals": 18
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
                "type": "FEE",
                "status": "SUCCESS",
                "account": {
                    "address": "hxe7af5fcfd8dfc67530a01a0e403882687528dfcb"
                },
                "amount": {
                    "value": "-1380000000000000",
                    "currency": {
                        "symbol": "ICX",
                        "decimals": 18
                    }
                }
            },
            {
                "operation_identifier": {
                    "index": 3
                },
                "related_operations": [
                    {
                        "index": 2
                    }
                ],
                "type": "FEE",
                "status": "SUCCESS",
                "account": {
                    "address": "hx1000000000000000000000000000000000000000"
                },
                "amount": {
                    "value": "1380000000000000",
                    "currency": {
                        "symbol": "ICX",
                        "decimals": 18
                    }
                }
            }
        ],
        "metadata": {
            "data": {
                "method": "setValue",
                "params": {
                    "value": "Hello"
                }
            },
            "dataType": "call",
            "nid": "0x3",
            "nonce": null,
            "signature": "JobzyMm/1JPuWNTlJnPNDjo9GkyXNWHoR+jm5A8Xx7xaFBNaAxvJOfiAZ0v9FpXXmExK1uDGjDyepE/SXXSYLAA=",
            "timestamp": "0x5ba183435a411",
            "version": "0x3"
        }
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

__Note__: Before calling `/combine`, please call `/payloads` to have the `unsigned_transaction`. Next, use [ICON SDK](https://github.com/icon-project/icon-sdk-java) or other ICON's SDKs to craft a transaction object and sign the transaction object; print out the __*signature*__ and __*transaction object*__ in __hexadecimals__. 


Request:

```json
{
    "network_identifier": {
        "blockchain": "ICON",
        "network": "Mainnet"
    },
    "unsigned_transaction": "{\"version\":\"0x3\",\"from\":\"hxe7af5fcfd8dfc67530a01a0e403882687528dfcb\",\"to\":\"hx3a065000ab4183c6bf581dc1e55a605455fc6d61\",\"value\":\"0xde0b6b3a7640000\",\"stepLimit\":\"0x186a0\",\"timestamp\":\"0x5ba18a5764ee6\",\"nid\":\"0x1\",\"nonce\":\"0x1\"}",
    "signatures": [
        {
            "signing_payload": {
                "address": "hxe7af5fcfd8dfc67530a01a0e403882687528dfcb",
                "account_identifier": {
                    "address": "hxe7af5fcfd8dfc67530a01a0e403882687528dfcb",
                    "metadata": {}
                },
                "hex_bytes": "e7af5fcfd8dfc67530a01a0e403882687528dfcb",
                "signature_type": "ecdsa"
            },
            "public_key": {
                "hex_bytes": "04855312cefe3b6c7e86bbaec089f79376e9bb3c19c84b73492d0d5e4d2dbdd4bec1c8c9ddc55715b70ed920d08d1621f20ffaf0d09e5afc97fb1d327ea7070084",
                "curve_type": "secp256k1"
            },
            "signature_type": "ecdsa",
            "hex_bytes": "42ecd0ea707ac3aa69ac3a21d1676b1fd3cb85adc7279a7e28de86552442e2b763615340f7ae17f5a558f58030fb0cf6fbd84c84ed3544ddfc0d4ee86506d0bb00"
        }
    ]
}
```

Response:

Sample

```json
{
    "signed_transaction": "{\"version\":\"0x3\",\"from\":\"hxe7af5fcfd8dfc67530a01a0e403882687528dfcb\",\"to\":\"hx3a065000ab4183c6bf581dc1e55a605455fc6d61\",\"value\":\"0xde0b6b3a7640000\",\"stepLimit\":\"0x186a0\",\"timestamp\":\"0x5ba18a5764ee6\",\"nid\":\"0x1\",\"nonce\":\"0x1\",\"signature\":\"QuzQ6nB6w6pprDoh0WdrH9PLha3HJ5p+KN6GVSRC4rdjYVNA964X9aVY9YAw+wz2+9hMhO01RN38DU7oZQbQuwA=\"}"
}
```

**/construction/metadata**

*Create a Request to Fetch Metadata*

Request:

```json
{
  "network_identifier": {
    "blockchain": "ICON",
    "network": "Mainnet"
  },
  "options": {},
  "public_keys": [
    {
      "hex_bytes": "04855312cefe3b6c7e86bbaec089f79376e9bb3c19c84b73492d0d5e4d2dbdd4bec1c8c9ddc55715b70ed920d08d1621f20ffaf0d09e5afc97fb1d327ea7070084",
      "curve_type": "secp256k1"
    }
  ]
}
```

Response:

```json
{
  "metadata": {
    "default_step_cost": "0x186a0"
  }
}
```

**/construction/derive**

*Derive an Address from a PublicKey*

Request:

```json
{
    "network_identifier": {
        "blockchain": "ICON",
        "network": "Mainnet"
    },
    "public_key": {
        "hex_bytes": "04855312cefe3b6c7e86bbaec089f79376e9bb3c19c84b73492d0d5e4d2dbdd4bec1c8c9ddc55715b70ed920d08d1621f20ffaf0d09e5afc97fb1d327ea7070084",
        "curve_type": "secp256k1"
    },
    "metadata": {}
}
```

Response:

Sample

```json
{
    "address": "hxe7af5fcfd8dfc67530a01a0e403882687528dfcb",
    "account_identifier": {
        "address": "hxe7af5fcfd8dfc67530a01a0e403882687528dfcb"
    }
}
```

**/construction/hash**

*Get the Hash of a Signed Transaction*

Request:

```json
{
    "network_identifier": {
        "blockchain": "ICON",
        "network": "Mainnet"
    },
    "signed_transaction": "{\"version\":\"0x3\",\"from\":\"hxe7af5fcfd8dfc67530a01a0e403882687528dfcb\",\"to\":\"hx3a065000ab4183c6bf581dc1e55a605455fc6d61\",\"value\":\"0xde0b6b3a7640000\",\"stepLimit\":\"0x186a0\",\"timestamp\":\"0x5ba18a5764ee6\",\"nid\":\"0x1\",\"nonce\":\"0x1\",\"signature\":\"QuzQ6nB6w6pprDoh0WdrH9PLha3HJ5p+KN6GVSRC4rdjYVNA964X9aVY9YAw+wz2+9hMhO01RN38DU7oZQbQuwA=\"}"
}
```

Response:

Sample

```json
{
    "transaction_identifier": {
        "hash": "0x7a418a0fb98181a1701390163dd03f552a9ed843ab270d7fb53a374373337943"
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
        "blockchain": "ICON",
        "network": "Mainnet"
    },
    "signed": true,
    "transaction": "{\"version\":\"0x3\",\"from\":\"hxe7af5fcfd8dfc67530a01a0e403882687528dfcb\",\"to\":\"hx3a065000ab4183c6bf581dc1e55a605455fc6d61\",\"value\":\"0xde0b6b3a7640000\",\"stepLimit\":\"0x186a0\",\"timestamp\":\"0x5ba18a5764ee6\",\"nid\":\"0x50\",\"nonce\":\"0x1\",\"signature\":\"QuzQ6nB6w6pprDoh0WdrH9PLha3HJ5p+KN6GVSRC4rdjYVNA964X9aVY9YAw+wz2+9hMhO01RN38DU7oZQbQuwA=\"}"
}
```

Response:

Sample

```json
{
    "signers": [
        "hxe7af5fcfd8dfc67530a01a0e403882687528dfcb"
    ],
    "operations": [
        {
            "operation_identifier": {
                "index": 0
            },
            "type": "TRANSFER",
            "status": "",
            "account": {
                "address": "hxe7af5fcfd8dfc67530a01a0e403882687528dfcb"
            },
            "amount": {
                "value": "-1000000000000000000",
                "currency": {
                    "symbol": "ICX",
                    "decimals": 18
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
            "type": "TRANSFER",
            "status": "",
            "account": {
                "address": "hx3a065000ab4183c6bf581dc1e55a605455fc6d61"
            },
            "amount": {
                "value": "1000000000000000000",
                "currency": {
                    "symbol": "ICX",
                    "decimals": 18
                }
            }
        }
    ],
    "account_identifier_signers": [
        {
            "address": "hxe7af5fcfd8dfc67530a01a0e403882687528dfcb"
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
        "blockchain": "ICON",
        "network": "Mainnet"
    },
    "operations": [
        {
            "operation_identifier": {
                "index": 0,
                "network_index": 0
            },
            "type": "TRANSFER",
            "status": "",
            "account": {
                "address": "hx3a065000ab4183c6bf581dc1e55a605455fc6d61",
                "metadata": {}
            },
            "amount": {
                "value": "1000000000000000000",
                "currency": {
                    "symbol": "ICX",
                    "decimals": 18
                },
                "metadata": {
                    "currency": {
                        "symbol": "ICX",
                        "decimals": 18
                    }
                }
            },
            "metadata": {}
        },
        {
            "operation_identifier": {
                "index": 1,
                "network_index": 0
            },
            "type": "TRANSFER",
            "status": "",
            "account": {
                "address": "hxe7af5fcfd8dfc67530a01a0e403882687528dfcb",
                "metadata": {}
            },
            "amount": {
                "value": "-1000000000000000000",
                "currency": {
                    "symbol": "ICX",
                    "decimals": 18
                },
                "metadata": {
                    "currency": {
                        "symbol": "ICX",
                        "decimals": 18
                    }
                }
            },
            "metadata": {}
        }
    ],
    "metadata": {},
    "public_keys": [
        {
            "hex_bytes": "04855312cefe3b6c7e86bbaec089f79376e9bb3c19c84b73492d0d5e4d2dbdd4bec1c8c9ddc55715b70ed920d08d1621f20ffaf0d09e5afc97fb1d327ea7070084",
            "curve_type": "secp256k1"
        }
    ]
}
```

Response:

Sample

```json
{
    "unsigned_transaction": "{\"version\":\"0x3\",\"from\":\"hxe7af5fcfd8dfc67530a01a0e403882687528dfcb\",\"to\":\"hx3a065000ab4183c6bf581dc1e55a605455fc6d61\",\"value\":\"0xde0b6b3a7640000\",\"stepLimit\":\"0x186a0\",\"timestamp\":\"0x5ba18a5764ee6\",\"nid\":\"0x7\",\"nonce\":\"0x1\"}",
    "payloads": [
        {
            "address": "hxe7af5fcfd8dfc67530a01a0e403882687528dfcb",
            "hex_bytes": "7a418a0fb98181a1701390163dd03f552a9ed843ab270d7fb53a374373337943",
            "account_identifier": {
                "address": "hxe7af5fcfd8dfc67530a01a0e403882687528dfcb"
            },
            "signature_type": "ecdsa_recovery"
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
    "blockchain": "ICON",
    "network": "Mainnet"
  },
  "operations": [
    {
      "operation_identifier": {
        "index": 0,
        "network_index": 0
      },
      "type": "TRANSFER",
      "status": "",
      "account": {
        "address": "hx3a065000ab4183c6bf581dc1e55a605455fc6d61",
        "metadata": {}
      },
      "amount": {
        "value": "1000000000000000000",
        "currency": {
          "symbol": "ICX",
          "decimals": 18
        },
        "metadata": {
          "currency": {
            "symbol": "ICX",
            "decimals": 18
          }
        }
      },
      "metadata": {}
    },
    {
      "operation_identifier": {
        "index": 1,
        "network_index": 0
      },
      "type": "TRANSFER",
      "status": "",
      "account": {
        "address": "hxe7af5fcfd8dfc67530a01a0e403882687528dfcb",
        "metadata": {}
      },
      "amount": {
        "value": "-1000000000000000000",
        "currency": {
          "symbol": "ICX",
          "decimals": 18
        },
        "metadata": {
          "currency": {
            "symbol": "ICX",
            "decimals": 18
          }
        }
      },
      "metadata": {}
    }
  ]
}
```

Response:

Sample

```json
{
    "options": {
        "from": "hxe7af5fcfd8dfc67530a01a0e403882687528dfcb"
    }
}
```

**/construction/submit**

*Submit a Signed Transaction*

__Note__: Before calling `/submit`, please call `/combine` to obtain the `signed_transaction` required for the  request parameters.

Request:

```json
{
    "network_identifier": {
        "blockchain": "ICON",
        "network": "Mainnet"
    },
    "signed_transaction": "{\"version\":\"0x3\",\"from\":\"hxe7af5fcfd8dfc67530a01a0e403882687528dfcb\",\"to\":\"hx3a065000ab4183c6bf581dc1e55a605455fc6d61\",\"value\":\"0xde0b6b3a7640000\",\"stepLimit\":\"0x186a0\",\"timestamp\":\"0x5ba18a5764ee6\",\"nid\":\"0x1\",\"nonce\":\"0x1\",\"signature\":\"QuzQ6nB6w6pprDoh0WdrH9PLha3HJ5p+KN6GVSRC4rdjYVNA964X9aVY9YAw+wz2+9hMhO01RN38DU7oZQbQuwA=\"}"
}
```

Response:

Sample

```json
{
    "transaction_identifier": {
        "hash": "0x7a418a0fb98181a1701390163dd03f552a9ed843ab270d7fb53a374373337943"
    }
}
```

