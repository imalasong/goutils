package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"sync"
)

func main() {

	//callPlanGo()
	callJsonGo()
}

func callJsonGo() {
	var wg sync.WaitGroup

	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			//dial, err := rpc.Dial("tcp", ":8989")
			conn, err := net.Dial("tcp", ":8989")
			if err != nil {
				log.Fatalln("dial error:", err)
			}
			client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
			defer wg.Done()
			var result string
			err = client.Call("helloService.Hello", "xiao肖", &result)
			defer client.Close()
			if err != nil {
				log.Fatalln("call error:", err)
			}
			log.Printf("call success: %v", result)
		}()
	}

	wg.Wait()
}

func callPlanGo() {
	var wg sync.WaitGroup

	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			dial, err := rpc.Dial("tcp", ":8989")

			if err != nil {
				log.Fatalln("dial error:", err)
			}
			defer wg.Done()
			var result string
			err = dial.Call("helloService.Hello", "xiao", &result)
			defer dial.Close()
			if err != nil {
				log.Fatalln("call error:", err)
			}
			log.Printf("call success: %v", result)
		}()
	}

	wg.Wait()
}
