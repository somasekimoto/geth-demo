#!/bin/bash

#!/bin/bash

# Anvilをバックグラウンドで実行し、出力を無視する
nohup anvil > /dev/null 2>&1 &
# バックグラウンドプロセスのPIDを取得し、ファイルに書き込む
echo $! > anvil.pid
echo "Anvil is running with PID: $(cat anvil.pid)"

nohup go run main.go > /dev/null 2>&1 &
echo $! > go-server.pid
echo "Go server is running with PID: $(cat go-server.pid)"

go run ./cmd/create_contract_address.go
go test test/helloworld_test.go -v

# PIDファイルが存在するか確認
if [ -f anvil.pid ]; then
    # PIDを読み込み、プロセスを終了
    kill -9 $(cat anvil.pid)
    rm anvil.pid
    echo "Anvil has been stopped."
else
    echo "Anvil is not running or pid file does not exist."
fi

# Goのサーバープロセスを終了
if [ -f go-server.pid ]; then
    kill -9 $(cat go-server.pid)
    rm go-server.pid
    echo "Go server has been stopped."
else
    echo "Go server is not running or pid file does not exist."
fi