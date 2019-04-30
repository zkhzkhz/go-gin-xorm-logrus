package main

import (
	"../../../log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	client, err := rpc.Dial("tcp", ":1234")
	defer client.Close()

	if err != nil {
		log.Fatal("dialing error...", err)
		return
	}

	args1 := &Args{2, 3}
	args2 := &Args{7, 2}
	args3 := &Args{7, 0}

	reply1 := 0
	reply2 := Quotient{}
	reply3 := Quotient{}

	//同步的RPC
	err = client.Call("Arith.Multiply", args1, &reply1)
	if err != nil {
		log.Fatal("Arith error:", err)
		return
	}
	log.Info(reply1)

	//异步的RPC
	call2 := client.Go("Arith.Divide", args2, &reply2, nil)
	if call2 != nil {
		if replyCall, ok := <-call2.Done; ok {
			if replyCall.Error != nil {
				log.Fatal("Arith error:", replyCall.Error)
				return
			}
			log.Info(reply2)
		}
	}

	call3 := client.Go("Arith.Divide", args3, &reply3, nil)
	if call3 != nil {
		if replyCall, ok := <-call3.Done; ok {
			if replyCall.Error != nil {
				log.Fatal("Arith error:", replyCall.Error)
				return
			}
			log.Info(reply3)
		}
	}
}
