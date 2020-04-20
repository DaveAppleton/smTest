package main

import (
	"crypto/rand"
	"testing"
	"time"

	"github.com/spacemeshos/ed25519"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

func TestEd(t *testing.T) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(priv.Seed())
	addr := types.BytesToAddress(pub[:])
	t.Log(addr.Hex())
	t.Fail()
}

func TestTx(t *testing.T) {
	tx, err := getTransaction("510f943a25187657577565c7a75345226937e0c4b4b02cf5901a452a726f94eb")
	// "2688348143db9e33b18ae2a5d91d0d9d9bfda75b5ad2d28c0f4abf7a20573155")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("from", tx.Sender.Address)
	t.Log("to", tx.Receiver.Address)
	t.Log("amount", tx.Amount)
	t.Log("fee", tx.Fee)
	t.Log("status", tx.Status)
	tim := time.Unix(int64(tx.Timestamp), 0)
	t.Log("date", tim)
	t.Fail()
}

func TestGetTxs(t *testing.T) {
	for _, addr := range []string{"0x4c406e078f322a95940123f3d89a01978b32850d", "0x77d1f7b2554a9917e30e2dd9d3e5358a262d3f5c"} {
		t.Log(addr)
		txs, layer, err := getTransactions(addr)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(layer)
		for _, txid := range txs {
			t.Log(txid)
		}
	}
	t.Fail()
}
