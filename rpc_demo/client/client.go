package main

import (
	"fmt"
	"log"
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

func main() {
	client, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		fmt.Fprintf(os.Stdout, "rpc dial failed, err: %s", err.Error())
	}
	defer client.Close()
	args := &Args{
		A: 6,
		B: 8,
	}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	//call:= client.Go("Arith.Multiply", args, &reply, nil)
	//select {
	//case <-call.Done:
	//	if err != nil {
	//		fmt.Fprintf(os.Stdout, "Multiply call failed, err:%s", err.Error())
	//	}
	//	fmt.Printf("Multiply: %d*%d=%d\n", args.A, args.B, reply)
	//default:
	//	fmt.Printf("helle world")
	//}
	if err != nil {
		fmt.Fprintf(os.Stdout, "Multiply call failed, err:%s", err.Error())
	}
	fmt.Printf("Multiply: %d*%d=%d\n", args.A, args.B, reply)

	args = &Args{15, 6}
	var quo Quotient
	err = client.Call("Arith.Divide", args, &quo)
	if err != nil {
		log.Fatal("Divide error:", err)
	}
	fmt.Printf("Divide: %d/%d=%d...%d des: %s\n", args.A, args.B, quo.Quo, quo.Rem, quo.Des)
}
