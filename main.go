package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

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

func (s *ContractConnectServer) ContractConnect(
	ctx context.Context,
	req *connect.Request[contract_connect_v1.ContractConnectRequest],
) (*connect.Response[contract_connect_v1.ContractConnectResponse], error) {
	log.Println("Request headers: ", req.Header())

	// Ethereum クライアントの初期化
	client, err := ethclient.Dial(fmt.Sprintf("http://127.0.0.1:%s", GANACHE_PORT))
	if err != nil {
		return nil, err
	}

	// スマートコントラクトとの接続
	conn, err := contracts.NewHelloworld(common.HexToAddress(CONTRACT_ADDRESS), client)
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

	res := connect.NewResponse(&contract_connect_v1.ContractConnectResponse{
		Result: reply,
	})
	res.Header().Set("Contract-Connect-Version", "v1")
	return res, nil
}

func server() http.Handler {
	mux := http.NewServeMux()
	mux.Handle(contract_connectorv1connect.NewContractConnectorServiceHandler(&ContractConnectServer{}))
	return mux
}

func main() {
	mux := server()
	http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
