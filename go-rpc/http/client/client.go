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
	client, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	defer client.Close()

	if err != nil {
		log.Fatal("dialing error:", err)
		return
	}

	args1 := &Args{2, 3}
	args2 := &Args{7, 2}
	args3 := &Args{7, 0}

	reply1 := 0
	reply2 := Quotient{}
	reply3 := Quotient{}

	err = client.Call("Arith.Multiply", args1, &reply1)
	if err != nil {
		log.Fatal("Arith error:", err)
		return
	}
	log.Info(reply1) //should be 6

	err = client.Call("Arith.Divide", args2, &reply2)
	if err != nil {
		log.Fatal("Arith error:", err)
		return
	}
	log.Info(reply2) //should be {3,1}

	err = client.Call("Arith.Divide", args3, &reply3)
	if err != nil {
		log.Fatal("Arith error:", err) //should be arith error:divide by zero
		return
	}
	log.Info(reply3)

}
