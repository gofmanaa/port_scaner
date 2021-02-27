package main

import (
	"flag"
	"fmt"
	"net"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/gofmanaa/port_scanner/data"
)

var address string

const ttl = 2

func init() {
	flag.StringVar(&address, "addr", "localhost", "Host name")
}

func main() {
	fmt.Println("Start sacanning...")
	flag.Parse()
	ch := make(chan int, 1<<16)

	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	fmt.Println("CpuNum: ", cpus)

	scanPorts(address, ch)
	openPortsMap := getChanel(ch)
	openPorts := openPortsMap.GetInt()

	if len(openPorts) > 0 {
		for _, openPort := range openPorts {
			fmt.Printf("Port %d [open]\n", openPort)
		}
	} else {
		fmt.Printf("Host %s have no open ports.\n", address)
	}
	fmt.Println("Done.")
}

func isOpenPort(address string, port int) bool {
	address = address + ":" + strconv.Itoa(port)
	con, err := net.DialTimeout("tcp", address, time.Second*ttl)
	if err != nil {
		return false
	}

	con.Close()

	return true
}

func scanPorts(address string, out chan int) {
	var wg sync.WaitGroup

	wg.Add(1)
	go startScan(address, &wg, out)

	wg.Wait()
}

func startScan(address string, wg *sync.WaitGroup, out chan int) {
	for port := 1; port < 1<<16; port++ {
		if isOpenPort(address, port) {
			out <- port
		}
	}
	close(out)
	wg.Done()
}

func getChanel(in <-chan int) *data.Set {
	res := data.NewSet()

	for openPort := range in {
		res.Add(openPort)
	}

	return res
}
