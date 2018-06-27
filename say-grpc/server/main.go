package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os/exec"

	pb "github.com/ar4mirez/say-grpc/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		logrus.Fatalf("could not listen to port %d: %v", *port, err)
	}

	logrus.Infof("listening to port %d", *port)

	s := grpc.NewServer()
	pb.RegisterTextToSpeechServer(s, &server{})

	err = s.Serve(listener)

	if err != nil {
		logrus.Fatalf("could not serve: %v", err)
	}
}

type server struct{}

func (server) Say(ctx context.Context, text *pb.Text) (*pb.Speech, error) {

	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, fmt.Errorf("could not create tmp file: %v", err)
	}

	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("could not close %s: %v", f.Name(), err)
	}

	cmd := exec.Command("flite", "-t", text.Text, "-o", f.Name())

	if data, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("flite failed: %s", data)
	}

	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("could not read tmp file: %v", err)
	}

	return &pb.Speech{Audio: data}, nil
}
