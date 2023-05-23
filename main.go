package main

import (
	"log"
	"net"
	"time"
)

func do(conn net.Conn){
	buffer := make([]byte,1024)
	_,err := conn.Read(buffer) // blocking call
	if err!= nil{
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second) // trying to simulate a long process
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nWelcome to the server\r\n")) // writing the http headers
	conn.Close() // connection is closed
}

func main(){
	listener,err := net.Listen("tcp", ":1729") // reserves port 1729 for listening to the client
	if err!=nil{
		log.Fatal(err)
	}
	conn,err := listener.Accept() // waits until the client gets connected, blocking call
	if err!= nil{
		log.Fatal(err)
	}
	
	do(conn)
}