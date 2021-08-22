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

	port := [33]string{
		"21", "22", "23",
		"25", "53", "79",
		"80", "88", "110",
		"111", "135", "137",
		"139", "145", "161",
		"162", "389", "443",
		"587", "631", "636",
		"1433", "1521", "1723",
		"2049", "2100", "3306",
		"3389", "5900", "5985",
		"6379", "8080", "11211"}

	for _, p := range port {
		
	}

