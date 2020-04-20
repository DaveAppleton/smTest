package main

import (
	"context"
	"time"

	pb "github.com/smstuff/smTest/pb"
	"google.golang.org/grpc"
)

func getBalance(address string) (value string, err error) {
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
	//opts := grpc.CallOption{}
	res, err := c.GetBalance(ctx, &accountID)
	if err != nil {
		return
	}
	value = res.GetValue()
	if len(value) > 12 {
		dp := len(value) - 12
		value = value[0:dp] + "." + value[dp:]
	}
	return

}

func getNonce(address string) (value string, err error) {
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
	//opts := grpc.CallOption{}
	res, err := c.GetNonce(ctx, &accountID)
	if err != nil {
		return
	}
	value = res.GetValue()
	return
}

/*

0x4c406e078f322a95940123f3d89a01978b32850d 48.635559657472 1
0x77d1f7b2554a9917e30e2dd9d3e5358a262d3f5c    100000000001 0


*/
