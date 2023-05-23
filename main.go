package main

import (
	"fmt"
	"log"
	"net"
)

func main(){
	listener,err := net.Listen("tcp", ":1729") // reserves port 1729 for listening to the client
	if err!=nil{
		log.Fatal(err)
	}
	conn,err := listener.Accept() // waits until the client gets connected
	if err!= nil{
		log.Fatal(err)
	}
	
	fmt.Println(conn)
}