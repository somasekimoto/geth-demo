package lib

import (
	"fmt"
	"math/big"

	// this would be your generated smart contract bindings

	contracts "github.com/somasekimoto/geth-demo/gocode"
)

// Ganacheのprivate keyの頭文字0xを外したキー

func DeployContract() {
	auth, client, err := GetEthAuthConfig()
	if err != nil {
		panic(err)
	}
	auth.Value = big.NewInt(0)                // in wei
	auth.GasTipCap = big.NewInt(2000000000)   // 優先料金（例：2 Gwei）
	auth.GasFeeCap = big.NewInt(100000000000) // 最大料金（例：100 Gwei）

	address, tx, instance, err := contracts.DeployHelloworld(auth, client)
	if err != nil {
		panic(err)
	}

	fmt.Println(address.Hex())

	_, _ = instance, tx
}
