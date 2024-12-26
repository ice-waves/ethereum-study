package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"log"
)

func main() {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)

	password := "secret"
	account, err := ks.NewAccount(password)
	ks.Wallets()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex())
}
