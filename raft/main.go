package main

import (
	"flag"
	"fmt"
	"os"
	"raft/server"
	"strconv"
	"sync"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = 10000
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s nProcesses\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func StartRAFT(processes int) {
	var wg sync.WaitGroup
	sms := make([](*server.StateMachine), processes)
	for i := 0; i < processes; i++ {
		wg.Add(1)
		sms[i] = server.NewServerSm(i)
		go func(i int) {
			port := SERVER_PORT + i
			sms[i].Run(SERVER_HOST, port)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("raft: Missing number")
	}

	flag.Usage = usage
	flag.Parse()

	processes, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}

	StartRAFT(processes)
}
