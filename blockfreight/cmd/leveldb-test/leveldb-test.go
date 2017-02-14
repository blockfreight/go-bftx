package main

import (
	"fmt"
	"strconv"

	"github.com/blockfreight/blockfreight-alpha/blockfreight/bft/leveldb"
)

func main() {
	//db_path := "bft-db/db"
	db_path := "bft-db"
	db, err := leveldb.OpenDB(db_path)
	defer leveldb.CloseDB(db)
	leveldb.HandleError(err, "Create or Open Database")
	//fmt.Println("Database created / open on "+db_path)
	
	for i := 1; i <= 50000; i++ {
		err = leveldb.InsertBFTX(strconv.Itoa(i), "Value for "+strconv.Itoa(i), db)
		//leveldb.HandleError(err, "Insert data for value "+strconv.Itoa(i))	
		//fmt.Println("Record saved!")	
	}

	//Iteration
	var n int
	n, err = leveldb.Iterate(db)
	leveldb.HandleError(err, "Iteration")
	fmt.Println("Total: "+strconv.Itoa(n))
}