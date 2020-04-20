package main

import (
	"context"
	"time"

	"github.com/spacemeshos/ed25519"
	"github.com/spacemeshos/go-spacemesh/common/types"

	"github.com/spacemeshos/go-spacemesh/common/util"

	pb "github.com/smstuff/smTest/pb"
	"google.golang.org/grpc"
)

type InnerSerializableSignedTransaction struct {
	AccountNonce uint64
	Recipient    types.Address
	GasLimit     uint64
	Price        uint64
	Amount       uint64
}

// Once we support signed txs we should replace SerializableTransaction with this struct. Currently it is only used in the rpc server.
type SerializableSignedTransaction struct {
	InnerSerializableSignedTransaction
	Signature [64]byte
}

type Transaction struct {
	Amount    uint64
	Fee       uint64
	Receiver  *pb.AccountId
	Sender    *pb.AccountId
	Status    pb.TxStatus
	TxID      *pb.TransactionId
	LayerID   uint64
	Timestamp uint64
}

func getTransaction(txString string) (tx *Transaction, err error) {
	url := "localhost:9091"
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return
	}
	data := util.Hex2Bytes(txString)
	defer conn.Close()
	c := pb.NewSpacemeshServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	txID := pb.TransactionId{Id: data}

	res, err := c.GetTransaction(ctx, &txID)
	if err != nil {
		return
	}
	tx = new(Transaction)
	tx.Amount = res.GetAmount()
	tx.Fee = res.GetFee()
	tx.Receiver = res.GetReceiver()
	tx.Sender = res.GetSender()
	tx.TxID = res.GetTxId()
	tx.Timestamp = res.GetTimestamp()
	tx.Status = res.GetStatus()
	tx.LayerID = res.GetLayerId()
	return
}

func getTransactions(address string) (txs []string, layer uint64, err error) {
	url := "localhost:9091"
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return
	}

	defer conn.Close()
	c := pb.NewSpacemeshServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	accountID := pb.AccountId{Address: address}
	txsSinceLayerID := pb.GetTxsSinceLayer{Account: &accountID}

	resp, err := c.GetAccountTxs(ctx, &txsSinceLayerID)
	if err != nil {
		return
	}
	txs = resp.GetTxs()
	layer = resp.GetValidatedLayer()

	return
}

func submitTransaction(recipient string, amount uint64, fee uint64, nonce uint64, gasLimit int64, gasPrice uint64) (id uint64, value uint64, err error) {
	url := "localhost:9091"
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return
	}

	defer conn.Close()
	c := pb.NewSpacemeshServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	recipientAddress := types.HexToAddress(recipient)

	tx := SerializableSignedTransaction{AccountNonce: nonce, Amount: amount, Recipient: recipientAddress, GasLimit: gasLimit, Price: gasPrice}
	buf, _ := InterfaceToBytes(&tx.InnerSerializableSignedTransaction)
	copy(tx.Signature[:], ed25519.Sign2(key, buf))
	b, err := InterfaceToBytes(&tx)
	if err != nil {
		return "", err
	}
	conf, err := c.SubmitTransaction(ctx, b)

	if err != nil {
		return
	}
	id = conf.Id
	value = conf.Value
	return
}
