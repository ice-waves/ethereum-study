package main

import (
	"fmt"
	"gin-greeting/ethereum/token"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	// 通过abi生成go版本智能合约接口，可以通过solcjs，也可以通过hardhat生成的json文件生成
	// abigen --abi=erc20_sol_ERC20.abi --pkg=token --out=erc20.go
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	tokenAddress := common.HexToAddress("0xa74476443119A942dE498590Fe1f2454d7D4aC0d")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	_ = instance
	//
	//address := common.HexToAddress("0x0536806df512d6cdde913cf95c9886f65b1d3462")
	//bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("wei: %s\n", bal)

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
}
