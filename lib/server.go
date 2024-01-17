package lib

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"

	"connectrpc.com/connect"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
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

func GetEthBalance(ownerAddress *common.Address, client *ethclient.Client, background context.Context) (*big.Int, error) {
	account := common.HexToAddress(ownerAddress.String())
	balance, err := client.BalanceAt(background, account, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
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
	log.Println("Request headers: ", req.Header())

	// スマートコントラクトとの接続
	conn, err := connectContract()
	if err != nil {
		return nil, err
	}

	auth, ethclient, err := GetEthAuthConfig()
	if err != nil {
		panic(err)
	}
	auth.Value = big.NewInt(0)                // in wei
	auth.GasTipCap = big.NewInt(2000000000)   // 優先料金（例：2 Gwei）
	auth.GasFeeCap = big.NewInt(100000000000) // 最大料金（例：100 Gwei）

	contractAddress := common.HexToAddress(CONTRACT_ADDRESS)

	contractBalance, err := GetEthBalance(&contractAddress, ethclient, context.Background())
	if err != nil {
		return nil, err
	}

	// スマートコントラクトのメソッドを呼び出す例
	reply, err := conn.Withdraw(&bind.TransactOpts{
		From:      auth.From,
		Signer:    auth.Signer,
		Value:     auth.Value,
		Nonce:     auth.Nonce,
		GasTipCap: auth.GasTipCap,
		GasFeeCap: auth.GasFeeCap,
		Context:   ctx,
	})
	log.Println("reply is", reply)

	if err != nil {
		return nil, err
	}

	afterBalance, err := GetEthBalance(&contractAddress, ethclient, context.Background())
	if err != nil {
		return nil, err
	}

	if afterBalance.Cmp(big.NewInt(0)) != 0 {
		err = errors.New("failed to withdraw")
		return nil, err
	}

	res := connect.NewResponse(&contract_connect_v1.WithdrawResponse{
		Balance: contractBalance.String(),
	})
	res.Header().Set("Contract-Connect-Version", "v1")
	return res, nil

}

type ReceivedEvent struct {
	Sender common.Address
	Amount *big.Int
}

func (s *ContractConnectServer) ReadReceiveEvent(
	ctx context.Context,
	req *connect.Request[contract_connect_v1.ReadReceiveEventRequest],
) (*connect.Response[contract_connect_v1.ReadReceiveEventResponse], error) {
	client, err := getEthClient()
	if err != nil {
		return nil, err
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(CONTRACT_ADDRESS)},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}

	contractAbi, err := abi.JSON(strings.NewReader(contracts.HelloworldABI)) // contracts.HelloWorldABIはあなたのコントラクトのABIです。
	if err != nil {
		return nil, err
	}

	// `Received`イベントのデータをデコード
	event := new(ReceivedEvent)
	log := logs[0]
	err = contractAbi.UnpackIntoInterface(event, "Received", log.Data)
	if err != nil {
		return nil, err
	}

	// `Topics`から送信者アドレスを取得
	if len(log.Topics) > 1 {
		event.Sender = common.HexToAddress(log.Topics[1].Hex())
		event.Amount = log.Topics[2].Big()
	}

	sender := event.Sender.String()
	amount := event.Amount.Uint64()

	res := connect.NewResponse(&contract_connect_v1.ReadReceiveEventResponse{
		Sender: sender,
		Amount: amount,
	})
	return res, nil
}

func Server() http.Handler {
	mux := http.NewServeMux()
	mux.Handle(contract_connectorv1connect.NewContractConnectorServiceHandler(&ContractConnectServer{}))
	return mux
}
