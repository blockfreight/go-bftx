package main

import (
	"flag"
	
	. "github.com/tendermint/go-common"
	"github.com/blockfreight/blockfreight-alpha/bft_dummy"
	"github.com/tendermint/abci/server"
	"github.com/tendermint/abci/types"
)

func main() {

	addrPtr := flag.String("addr", "tcp://127.0.0.0:46658", "Listen address")
	abciPtr := flag.String("bft", "socket", "socket | grpc")
	persistencePtr := flag.String("persist", "", "directory to use for a database")
	flag.Parse()
	
	// Create the application - in memory or persisted to disk
	var app types.Application
	if *persistencePtr == "" {
		app = bft_dummy.NewBftApplication()
	} else {
		app = bft_dummy.NewPersistentBftApplication(*persistencePtr)
	}

	// Start the listener
	srv, err := server.NewServer(*addrPtr, *abciPtr, app)
	if err != nil {
		Exit(err.Error())
	}

	// Wait forever
	TrapSignal(func() {
		// Cleanup
		srv.Stop()
	})

}