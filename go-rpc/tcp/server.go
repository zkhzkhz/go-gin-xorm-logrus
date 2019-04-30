package main

import (
	"../../log"
	"errors"
	"net"
	"net/rpc"
	"time"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	server := rpc.NewServer()
	server.Register(arith)

	l, e := net.Listen("tcp", ":1234")
	defer l.Close()

	if e != nil {
		log.Fatal("listen error:", e)
		return
	}

	go server.Accept(l)
	log.Info("rpc server started!")

	for {
		time.Sleep(1 * time.Second)
	}
}
