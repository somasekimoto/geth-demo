package lib

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"connectrpc.com/connect"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	contracts "github.com/somasekimoto/geth-demo/gocode"
	contract_connect_v1 "github.com/somasekimoto/geth-demo/gocode/protobuf/contract_connector/v1"
	"github.com/somasekimoto/geth-demo/gocode/protobuf/contract_connector/v1/contract_connectorv1connect"
)

const (
	GANACHE_PORT     = "8545"
	CONTRACT_ADDRESS = "0x5FbDB2315678afecb367f032d93F642f64180aa3"
)

type ContractConnectServer struct{}

func getEthClient() (*ethclient.Client, error) {
	return ethclient.Dial(fmt.Sprintf("http://127.0.0.1:%s", GANACHE_PORT))
}

func connectContract() (*contracts.Helloworld, error) {
	// Ethereum クライアントの初期化
	client, err := getEthClient()
	if err != nil {
		return nil, err
	}
	return contracts.NewHelloworld(common.HexToAddress(CONTRACT_ADDRESS), client)
}

func (s *ContractConnectServer) Hello(
	ctx context.Context,
	req *connect.Request[contract_connect_v1.HelloRequest],
) (*connect.Response[contract_connect_v1.HelloResponse], error) {
	log.Println("Request headers: ", req.Header())

	// スマートコントラクトとの接続
	conn, err := connectContract()
	if err != nil {
		return nil, err
	}

	// スマートコントラクトのメソッドを呼び出す例
	reply, err := conn.Hello(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&contract_connect_v1.HelloResponse{
		Result: reply,
	})
	res.Header().Set("Contract-Connect-Version", "v1")
	return res, nil
}

func (s *ContractConnectServer) GetBalance(
	ctx context.Context,
	req *connect.Request[contract_connect_v1.GetBalanceRequest],
) (*connect.Response[contract_connect_v1.GetBalanceResponse], error) {
	log.Println("Request headers: ", req.Header())

	// スマートコントラクトとの接続
	conn, err := connectContract()
	if err != nil {
		return nil, err
	}

	// スマートコントラクトのメソッドを呼び出す例
	reply, err := conn.GetBalance(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&contract_connect_v1.GetBalanceResponse{
		Result: reply.Uint64(),
	})
	res.Header().Set("Contract-Connect-Version", "v1")
	return res, nil
}

func (s *ContractConnectServer) GetData(
	ctx context.Context,
	req *connect.Request[contract_connect_v1.GetDataRequest],
) (*connect.Response[contract_connect_v1.GetDataResponse], error) {
	log.Println("Request headers: ", req.Header())

	// スマートコントラクトとの接続
	conn, err := connectContract()
	if err != nil {
		return nil, err
	}

	// スマートコントラクトのメソッドを呼び出す例
	reply, err := conn.GetData(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&contract_connect_v1.GetDataResponse{
		Data: reply,
	})
	res.Header().Set("Contract-Connect-Version", "v1")
	return res, nil
}

func (s *ContractConnectServer) Greet(
	ctx context.Context,
	req *connect.Request[contract_connect_v1.GreetRequest],
) (*connect.Response[contract_connect_v1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())

	// スマートコントラクトとの接続
	conn, err := connectContract()
	if err != nil {
		return nil, err
	}

	// スマートコントラクトのメソッドを呼び出す例
	reply, err := conn.Greet(&bind.CallOpts{
		Context: ctx,
	}, req.Msg.GetStr())
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&contract_connect_v1.GreetResponse{
		Str: reply,
	})
	res.Header().Set("Contract-Connect-Version", "v1")
	return res, nil
}

func (s *ContractConnectServer) SetData(
	ctx context.Context,
	req *connect.Request[contract_connect_v1.SetDataRequest],
) (*connect.Response[contract_connect_v1.SetDataResponse], error) {
	log.Println("Request headers: ", req.Header())

	// スマートコントラクトとの接続
	conn, err := connectContract()
	if err != nil {
		return nil, err
	}

	auth, _, err := GetEthAuthConfig()
	if err != nil {
		panic(err)
	}
	auth.Value = big.NewInt(0)                // in wei
	auth.GasTipCap = big.NewInt(2000000000)   // 優先料金（例：2 Gwei）
	auth.GasFeeCap = big.NewInt(100000000000) // 最大料金（例：100 Gwei）

	// address, tx, instance, err := contracts.SetData(auth, ethclient)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(address.Hex())

	// スマートコントラクトのメソッドを呼び出す例
	reply, err := conn.SetData(&bind.TransactOpts{
		From:      auth.From,
		Signer:    auth.Signer,
		Value:     auth.Value,
		Nonce:     auth.Nonce,
		GasTipCap: auth.GasTipCap,
		GasFeeCap: auth.GasFeeCap,
		Context:   ctx,
	}, req.Msg.GetData())
	log.Println("reply is", reply)

	if err != nil {
		return nil, err
	}
	result := true

	res := connect.NewResponse(&contract_connect_v1.SetDataResponse{
		Result: result,
	})
	res.Header().Set("Contract-Connect-Version", "v1")
	return res, nil
}

func (s *ContractConnectServer) Withdraw(
	ctx context.Context,
	req *connect.Request[contract_connect_v1.WithdrawRequest],
) (*connect.Response[contract_connect_v1.WithdrawResponse], error) {
	return nil, nil
}

func Server() http.Handler {
	mux := http.NewServeMux()
	mux.Handle(contract_connectorv1connect.NewContractConnectorServiceHandler(&ContractConnectServer{}))
	return mux
}
