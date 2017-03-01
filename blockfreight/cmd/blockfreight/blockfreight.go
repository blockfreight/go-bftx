package main

import (
	"flag"
	"fmt"
	
	. "github.com/tendermint/go-common"
	"github.com/blockfreight/blockfreight-alpha/blockfreight/bft"
	"github.com/tendermint/abci/server"
	"github.com/tendermint/abci/types"
)

func main() {

	fmt.Println("Blockfreight Go App")
	//Parameters
	addrPtr := flag.String("addr", "tcp://0.0.0.0:46658", "Listen address")
	abciPtr := flag.String("bft", "socket", "socket | grpc")
	//persistencePtr := flag.String("persist", "", "directory to use for a database")
	flag.Parse()
	
	// Create the application - in memory or persisted to disk
	var app types.Application
	app = bft.NewBftApplication() //if *persistencePtr != "" => NewPersistentBftApplication(*persistencePtr)
	
	// Start the listener
	srv, err := server.NewServer(*addrPtr, *abciPtr, app)
	fmt.Println(srv)
	if err != nil {
		Exit(err.Error())
	}
	fmt.Println("Service created by "+ *abciPtr +" server")

	// Wait forever
	TrapSignal(func() {
		// Cleanup
		fmt.Println("Stopping service")
		srv.Stop()
	})

}