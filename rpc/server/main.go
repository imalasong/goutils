package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {

	//planGo()
	jsonGo()

}

func jsonGo() {
	rpc.RegisterName("helloService", new(HelloService))

	listen, err := net.Listen("tcp", "localhost:8989")

	if err != nil {
		log.Fatalln("start server error :", err)
	}
	log.Println("server start success!")

	for {
		con, err := listen.Accept()
		if err != nil {
			log.Println("connector error,", err)
			continue
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(con))
	}
}

func planGo() {
	rpc.RegisterName("helloService", new(HelloService))

	listen, err := net.Listen("tcp", "localhost:8989")

	if err != nil {
		log.Fatalln("start server error :", err)
	}
	log.Println("server start success!")
	rpc.Accept(listen)
	//for {
	//	con, err := listen.Accept()
	//	if err != nil {
	//		log.Println("connector error,", err)
	//		continue
	//	}
	//	go rpc.ServeConn(con)
	//}
}

type HelloService struct {
}

func (h *HelloService) Hello(r string, w *string) error {
	log.Printf("接收请求:%v", r)
	//time.Sleep(time.Second * 5)
	*w = "hello," + r
	return nil
}
