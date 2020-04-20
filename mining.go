package main

import (
	"context"
	"time"

	pb "github.com/DaveAppleton/smTest/pb"
	"google.golang.org/grpc"
)

func setAwardAddress(address string) (value string, err error) {
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
	res, err := c.SetAwardsAddress(ctx, &accountID)
	//c.GetBalance(ctx, &accountID)
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
