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
	//get arg for ip
	ip := flag.string("h", "", "Select IP to Scan")
	//parse arg
	flag.Parse()
		//slice stores all 65536 ports
		var prt []int
		//slice sorts all 65536 ports and converts to string
		prtStr := []string{}
		//integer slices with 65536 slices
		allP := make([]int, 65536)
		//iterate all 65536 slices and append to prt
		for p := range allP {
			prt = append(prt, p)
		}
		//convert int slice into string slice
		for i := range prt {
			n := prt[i]
			text := strconv.Itoa(n)
			prtStr = append(prtStr, text)
		}
	
		var wg sync.WaitGroup
	
		for _, p := range prtStr {
			//if counter equals zero, all go routines blocked on Wait are released
			wg.Add(1)
			//call portScan function and iterate through every port on ip address concurrently
			go portScan(*ip, p, &wg)
		}
	
		wg.Wait()	
}