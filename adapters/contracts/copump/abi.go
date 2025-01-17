package copump

var (
	ABI = `[
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "tokenAddress",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "buyer",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "totalPrice",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "fee",
				"type": "uint256"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "funding",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "supply",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "marketCap",
						"type": "uint256"
					}
				],
				"indexed": false,
				"internalType": "struct Copump.TokenState",
				"name": "tokenState",
				"type": "tuple"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "occuredAt",
				"type": "uint256"
			}
		],
		"name": "TokenPurchased",
		"type": "event"
	}
]`
)
