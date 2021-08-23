package main

//usage: ./gScan -h <ip-address>

import (
	"flag"
	"fmt"
	// "log"
	"net"
	"strconv"
	"sync"
)

//port scan using go routines
func portScan(ip string, port string, wg *sync.WaitGroup) {
	defer wg.Done()
	//choose tcp or udp
	network := "tcp"
	address := ip + ":" + port
	connection, err := net.Dial(network, address)
	//error handling
	if err != nil {
		return
	}

	fmt.Printf("Port %s is open\n", port)
	connection.Close()
}

func main() {
	
}