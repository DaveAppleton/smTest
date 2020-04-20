package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"log"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/common/util"
)

func makeAccount() {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(util.Bytes2Hex(priv.Seed()))
	addr := types.BytesToAddress(pub[:])
	fmt.Println(addr.Hex())
}

func info(address string) {

	bal, err := getBalance(address)
	if err != nil {
		log.Fatal(err)
	}
	nonce, err := getNonce(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address, bal, nonce)
}

func main() {
	info("0x4c406e078f322a95940123f3d89a01978b32850d")
	info("0x77d1f7b2554a9917e30e2dd9d3e5358a262d3f5c")
}
