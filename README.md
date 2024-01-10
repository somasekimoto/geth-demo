# geth-demo

.solファイルからABI, bin, goコードを作成する

```sh
solc --optimize --abi ./contracts/HelloWorld.sol -o build --overwrite
solc --optimize --bin ./contracts/HelloWorld.sol -o build --overwrite
abigen --abi ./build/HelloWorld.abi --bin build/HelloWorld.bin --pkg helloworld --out ./gocode/hello_world.go
```


## テストの実行
```sh
sh test.sh
```