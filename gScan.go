package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	ip := flag.String("h", "", "Select IP Address to scan")
	flag.Parse()

	network := "tcp"

	port := [33]string{"
		
		"}