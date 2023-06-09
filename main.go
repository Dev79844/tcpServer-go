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
	time.Sleep(5 * time.Second) // trying to simulate a long process
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nWelcome to the server\r\n")) // writing the http headers
	conn.Close() // connection is closed
}

func worker(pool chan net.Conn){
	for conn := range pool{
		do(conn)
	}
}

func main(){
	listener,err := net.Listen("tcp", ":1729") // reserves port 1729 for listening to the client
	if err!=nil{
		log.Fatal(err)
	}

	pool := make(chan net.Conn,10)  // created a thread pool of 10 workers

	for i:=0; i<cap(pool);i++{
		go worker(pool) // creates worker threads and executes the program
	}

	for{
		conn,err := listener.Accept() // waits until the client gets connected, blocking call
		// err = conn.SetDeadline(time.Time{}.Add(time.Second * 5))
		if err != nil{
			log.Fatal(err)
		}
		if err!= nil{
			log.Fatal(err)
		}
		
		pool <- conn
	}
}