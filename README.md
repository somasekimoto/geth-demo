# geth-demo

.solファイルからABI, bin, goコードを作成する

```sh
solc --optimize --abi ./contracts/HelloWorld.sol -o build
solc --optimize --bin ./contracts/HelloWorld.sol -o build
abigen --abi ./build/HelloWorld.abi --bin build/HelloWorld.bin --pkg gocode --out ./gocode/hello_world.go
```
