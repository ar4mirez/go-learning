package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	service "github.com/ar4mirez/go-learning/rpcx"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main()  {
	flag.Parse()

	s := server.NewServer()
	s.Register(new(service.Arith), "")
	s.Serve("tcp", *addr)
}
