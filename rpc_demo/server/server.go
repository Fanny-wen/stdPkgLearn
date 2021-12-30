package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
	Des string
}

type Arith int

func (a *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (a *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by 0")
	}
	quo.Rem = args.A % args.B
	quo.Quo = args.A / args.B
	quo.Des = "你好, hello"
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Fprintf(os.Stdout, "net listen err: %s", err.Error())
	}
	defer listener.Close()
	conn, _ := listener.Accept()
	rpcServer := rpc.NewServer()
	//_ =rpcServer.Register(new(Arith))
	_ = rpcServer.RegisterName("Arith", new(Arith))
	rpcServer.ServeConn(conn)
	//jsonrpc.NewServerCodec(conn)
	//rpcServer.Accept(listener)
}
