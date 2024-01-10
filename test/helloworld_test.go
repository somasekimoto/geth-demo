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

// type ContractConnectServer struct{}

// func server() http.Handler {
// 	mux := http.NewServeMux()
// 	mux.Handle(contract_connectorv1connect.NewContractConnectorServiceHandler(&ContractConnectServer{}))
// 	return mux
// }

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
