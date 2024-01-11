package maintest

import (
	"context"
	"math/big"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	contract_connectorv1 "github.com/somasekimoto/geth-demo/gocode/protobuf/contract_connector/v1"
	"github.com/somasekimoto/geth-demo/lib"

	"github.com/somasekimoto/geth-demo/gocode/protobuf/contract_connector/v1/contract_connectorv1connect"
)

func TestHello(t *testing.T) {
	t.Parallel()
	mux := lib.Server()
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	t.Cleanup(server.Close)
	client := contract_connectorv1connect.NewContractConnectorServiceClient(
		server.Client(),
		server.URL,
	)

	res, err := client.Hello(context.Background(), &connect.Request[contract_connectorv1.HelloRequest]{})
	if err != nil {
		t.Error(err)
	}
	got := res.Msg.GetResult()
	want := "Hello World"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGreet(t *testing.T) {
	t.Parallel()
	mux := lib.Server()
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	t.Cleanup(server.Close)
	client := contract_connectorv1connect.NewContractConnectorServiceClient(
		server.Client(),
		server.URL,
	)

	want := "Gm World"

	res, err := client.Greet(context.Background(), &connect.Request[contract_connectorv1.GreetRequest]{
		Msg: &contract_connectorv1.GreetRequest{
			Str: want,
		},
	})
	if err != nil {
		t.Error(err)
	}
	got := res.Msg.GetStr()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGetBalance(t *testing.T) {
	t.Parallel()
	mux := lib.Server()
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	t.Cleanup(server.Close)
	client := contract_connectorv1connect.NewContractConnectorServiceClient(
		server.Client(),
		server.URL,
	)

	res, err := client.GetBalance(context.Background(), &connect.Request[contract_connectorv1.GetBalanceRequest]{})
	if err != nil {
		t.Error(err)
	}
	got := res.Msg.GetResult()
	want := uint64(0)
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func TestGetData(t *testing.T) {
	t.Parallel()
	mux := lib.Server()
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	t.Cleanup(server.Close)
	client := contract_connectorv1connect.NewContractConnectorServiceClient(
		server.Client(),
		server.URL,
	)

	res, err := client.GetData(context.Background(), &connect.Request[contract_connectorv1.GetDataRequest]{})
	if err != nil {
		t.Error(err)
	}
	got := res.Msg.GetData()
	want := ""
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func TestWriteTxDataFuncs(t *testing.T) {
	t.Parallel()
	mux := lib.Server()
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	t.Cleanup(server.Close)
	client := contract_connectorv1connect.NewContractConnectorServiceClient(
		server.Client(),
		server.URL,
	)

	t.Run("TestSetData", func(t *testing.T) {
		want := "foo"
		result, err := client.SetData(context.Background(), &connect.Request[contract_connectorv1.SetDataRequest]{
			Msg: &contract_connectorv1.SetDataRequest{
				Data: want,
			},
		})
		if err != nil {
			t.Error(err)
		}

		if !result.Msg.GetResult() {
			t.Error("SetData failed")
		}

		res, error := client.GetData(context.Background(), &connect.Request[contract_connectorv1.GetDataRequest]{})
		if error != nil {
			t.Error(error)
		}

		got := res.Msg.GetData()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("TestWithdraw", func(t *testing.T) {
		t.Run("0 balance", func(t *testing.T) {

			want := uint64(0)

			res, err := client.Withdraw(context.Background(), &connect.Request[contract_connectorv1.WithdrawRequest]{})
			if err != nil {
				t.Error(err)
			}
			got := res.Msg.GetBalance()
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})

		t.Run("send eth & withdraw", func(t *testing.T) {
			auth, ethclient, err := lib.GetEthAuthConfig()
			if err != nil {
				t.Error(err)
			}

			value := big.NewInt(100000000000000000)
			auth.Value = value // 0.1 eth

			gasTipCap, err := ethclient.SuggestGasTipCap(context.Background())
			if err != nil {
				t.Error(err)
			}
			auth.GasTipCap = gasTipCap

			// ベースフィーを超える最大料金を設定

			gasFeeCap := gasTipCap.Mul(gasTipCap, big.NewInt(2))
			auth.GasFeeCap = gasFeeCap

			chainID, err := ethclient.NetworkID(context.Background())
			if err != nil {
				t.Error(err)
			}

			privateKey, err := crypto.HexToECDSA(lib.PRIVATE_KEY)
			if err != nil {
				t.Error(err)
			}

			toAddress := common.HexToAddress(lib.CONTRACT_ADDRESS)

			tx := &types.DynamicFeeTx{
				ChainID:   chainID,
				Nonce:     auth.Nonce.Uint64(),
				GasFeeCap: gasFeeCap,
				GasTipCap: gasTipCap,
				Gas:       21000,
				To:        &toAddress,
				Value:     value, // 1 ETH
			}

			signedTx, err := types.SignNewTx(privateKey, types.NewCancunSigner(chainID), tx)
			if err != nil {
				t.Error(err)
			}

			error := ethclient.SendTransaction(context.Background(), signedTx)
			if error != nil {
				t.Error(error)
			}
		})
	})
}
