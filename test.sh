#!/bin/bash

# Anvilをバックグラウンドで実行
nohup anvil > /dev/null 2>&1 &
echo $! > anvil.pid
echo "Anvil is running with PID: $(cat anvil.pid)"

# Goのサーバーをビルドしてバックグラウンドで実行
go build -o go-server main.go
nohup ./go-server > /dev/null 2>&1 &
echo $! > go-server.pid
echo "Go server is running with PID: $(cat go-server.pid)"

# その他のGoコマンドを実行
go run ./cmd/create_contract_address.go
go test test/helloworld_test.go -v

# Anvilのプロセスを終了
if [ -f anvil.pid ]; then
    kill $(cat anvil.pid)
    rm anvil.pid
    echo "Anvil has been stopped."
else
    echo "Anvil is not running or pid file does not exist."
fi

# Goのサーバープロセスを終了
if [ -f go-server.pid ]; then
    kill $(cat go-server.pid)
    rm go-server.pid
    echo "Go server has been stopped."
else
    echo "Go server is not running or pid file does not exist."
fi
