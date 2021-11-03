package main

import (
	"io"
	"log"
	"net"
	"time"
)

// f()
// go f()

//func f(){
//	for i := 0 ; ; i++ {
//		fmt.Println(i)
//		time.Sleep(1 * time.Second)
//	}
//}
//
//func fib(n int) int {
//	if n < 2{
//		return n
//	}
//	return fib(n - 1) + fib (n - 2)
//
//}

func handleConn(c net.Conn){
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().String() + "\n")
		if err != nil{
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
	}
}

func main(){
	//go f()
	//fmt.Println(fib(45))
	listener, err := net.Listen("tcp","localhost:8000")
	if err != nil{
		log.Fatal(err)
	}
	for{
		conn, err := listener.Accept()
		if err != nil{
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}