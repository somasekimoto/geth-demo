package maintest

import (
	"context"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
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
func TestSetData(t *testing.T) {
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

	want := true

	res, err := client.SetData(context.Background(), &connect.Request[contract_connectorv1.SetDataRequest]{
		Msg: &contract_connectorv1.SetDataRequest{
			Data: "foo",
		},
	})
	if err != nil {
		t.Error(err)
	}
	got := res.Msg.GetResult()
	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}
func TestWithdraw(t *testing.T) {

}
