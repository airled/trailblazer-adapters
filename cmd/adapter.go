package cmd

type adapter string

func (a adapter) String() string {
	return string(a)
}

const (
	IzumiLP                  adapter = "IzumiLP"
	RitsuLP                  adapter = "RitsuLP"
	TransactionSender        adapter = "TransactionSender"
	NftDeployed              adapter = "NftDeployed"
	GamingWhitelist          adapter = "GamingWhitelist"
	DotTaikoDomains          adapter = "DotTaikoDomains"
	OkxOrderFulfilled        adapter = "OkxOrderFulfilled"
	LoopexNewSale            adapter = "LoopexNewSale"
	OmnihubContractDeployed  adapter = "OmnihubContractDeployed"
	Nfts2meCollectionCreated adapter = "Nfts2meCollectionCreated"
	ConftTokenSold           adapter = "ConftTokenSold"
	CopumpTokenPurchased     adapter = "CopumpTokenPurchased"
	DripsLock                adapter = "DripsLock"
	SymmetricLock            adapter = "SymmetricLock"
	RobinosPrediction        adapter = "RobinosPrediction"
	LoopringLock             adapter = "LoopringLock"
	PolarisLP                adapter = "PolarisLP"
)

func adapterz() []adapter {
	return []adapter{
		IzumiLP,
		RitsuLP,
		TransactionSender,
		NftDeployed,
		GamingWhitelist,
		DotTaikoDomains,
		OkxOrderFulfilled,
		LoopexNewSale,
		OmnihubContractDeployed,
		Nfts2meCollectionCreated,
		ConftTokenSold,
		CopumpTokenPurchased,
		DripsLock,
		SymmetricLock,
		RobinosPrediction,
		LoopringLock,
		PolarisLP,
	}
}
