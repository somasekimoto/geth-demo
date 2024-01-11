package lib

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const PRIVATE_KEY = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

func GetEthAuthConfig() (*bind.TransactOpts, *ethclient.Client, error) {

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	// auth.Value = big.NewInt(0) // in wei
	// // auth.GasLimit = uint64(3000000) // in units
	// // auth.GasPrice = big.NewInt(1000000)
	// auth.GasTipCap = big.NewInt(2000000000)   // 優先料金（例：2 Gwei）
	// auth.GasFeeCap = big.NewInt(100000000000) // 最大料金（例：100 Gwei）

	return auth, client, nil
}
