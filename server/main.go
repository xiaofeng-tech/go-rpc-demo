package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Arg struct {}

type TimeServer struct {}

func (t TimeServer) GiveServerTime(args *Arg, reply *int64) error {
	fmt.Printf("Receive a new request")
	*reply = time.Now().Unix()
	return nil
}

func main() {
	timeserver := TimeServer{}
	rpc.Register(timeserver)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listener error: ", err)
	}

	http.Serve(listener, nil)
}