package main

import (
	"context"
	"flag"
	"log"
	"github.com/smallnest/rpcx/client"
	"github.com/ar4mirez/go-learning/rpcx"
)

var (
	addr2 = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr2, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &service.Args{
		A: 10,
		B: 20,
	}

	reply := &service.Reply{}
	call, err := xclient.Go(context.Background(), "Mul", args, reply, nil)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	replyCall := <-call.Done
	if replyCall.Error != nil {
		log.Fatalf("failed to call: %v", replyCall.Error)
	} else {
		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
	}

}