package copump

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/copump"
)

const (
	// https://taikoscan.io/address/0x95e483Ce4acf1F24B6cBD8B369E0735a3e56f5BB
	TokenPurchasedAddress string = "0x95e483Ce4acf1F24B6cBD8B369E0735a3e56f5BB"

	logTokenPurchasedSignature string = "TokenPurchased(address,address,uint256,uint256,uint256,(uint256,uint256,uint256),uint256)"
)

type TokenStateInfo struct {
	Funding   *big.Int
	Supply    *big.Int
	MarketCap *big.Int
}

type TokenPurchasedEvent struct {
	TokenAddress common.Address
	Buyer        common.Address
	Amount       *big.Int
	TotalPrice   *big.Int
	Fee          *big.Int
	TokenState   TokenStateInfo
	OccuredAt    *big.Int
}

type TokenPurchasedIndexer struct {
	client    *ethclient.Client
	addresses []common.Address
}

func NewTokenPurchasedIndexer(client *ethclient.Client, addresses []common.Address) *TokenPurchasedIndexer {
	return &TokenPurchasedIndexer{
		client:    client,
		addresses: addresses,
	}
}

var _ adapters.LogIndexer[adapters.Whitelist] = &TokenPurchasedIndexer{}

func (indexer *TokenPurchasedIndexer) Addresses() []common.Address {
	return indexer.addresses
}

func (indexer *TokenPurchasedIndexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.Whitelist, error) {
	var whitelist []adapters.Whitelist

	for _, l := range logs {
		if !indexer.isTokenPurchasedLog(l) {
			continue
		}
		var tokenPurchasedEvent TokenPurchasedEvent

		tokenPurchasedABI, err := abi.JSON(strings.NewReader(copump.ABI))
		if err != nil {
			return nil, err
		}

		err = tokenPurchasedABI.UnpackIntoInterface(&tokenPurchasedEvent, "TokenPurchased", l.Data)
		if err != nil {
			return nil, err
		}

		block, err := indexer.client.BlockByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
		if err != nil {
			return nil, err
		}

		w := &adapters.Whitelist{
			User:        tokenPurchasedEvent.Buyer,
			Time:        block.Time(),
			BlockNumber: block.NumberU64(),
			TxHash:      l.TxHash,
		}

		whitelist = append(whitelist, *w)
	}

	return whitelist, nil
}

func (indexer *TokenPurchasedIndexer) isTokenPurchasedLog(l types.Log) bool {
	return l.Topics[0].Hex() == crypto.Keccak256Hash([]byte(logTokenPurchasedSignature)).Hex()
}
